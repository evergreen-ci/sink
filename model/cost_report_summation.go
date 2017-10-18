package model

import (
	"encoding/json"

	"github.com/evergreen-ci/sink"
	"github.com/mongodb/anser/bsonutil"
	"github.com/mongodb/anser/model"
	"github.com/mongodb/grip"
	"github.com/mongodb/grip/message"
	"github.com/pkg/errors"
)

const costReportSummaryCollection = "spending.summaries"

type CostReportSummary struct {
	ID                string                    `bson:"_id" json:"-" yaml:"-"`
	Metadata          CostReportMetadata        `bson:"metadata" json:"metadata" yaml:"metadata"`
	EvergreenProjects []EvergreenProjectSummary `bson:"projects" json:"projects" yaml:"projects"`
	ProviderSummary   []CloudProviderSummary    `bson:"providers" json:"providers" yaml:"providers"`
	TotalCost         float64                   `bson:"total_cost" json:"total_cost" yaml:"total_cost"`

	env       sink.Environment
	populated bool
}

var (
	costReportSummaryIDKey                = bsonutil.MustHaveTag(CostReportSummary{}, "ID")
	costReportSummaryMetadataKey          = bsonutil.MustHaveTag(CostReportSummary{}, "Metadata")
	costReportSummaryEvergreenProjectsKey = bsonutil.MustHaveTag(CostReportSummary{}, "EvergreenProjects")
	costReportSummaryProviderSummaryKey   = bsonutil.MustHaveTag(CostReportSummary{}, "ProviderSummary")
	costReportSummaryTotalCostKey         = bsonutil.MustHaveTag(CostReportSummary{}, "TotalCost")
)

type EvergreenProjectSummary struct {
	Name      string                         `bson:"name" json:"name" yaml:"name"`
	Versions  int                            `bson:"versions" json:"versions" yaml:"versions"`
	Resources []EvergreenResourceCostSummary `bson:"resources" json:"resources" yaml:"resources"`
}

type EvergreenResourceCostSummary struct {
	Name    string  `bson:"name" json:"name" yaml:"name"`
	Seconds int64   `bson:"seconds" json:"seconds" yaml:"seconds"`
	Cost    float64 `bson:"cost" json:"cost" yaml:"cost"`
}

var (
	costReportEvergrenProjectSummaryNameKey        = bsonutil.MustHaveTag(EvergreenProjectSummary{}, "Name")
	costReportEvergrenProjectSummaryVersionsKey    = bsonutil.MustHaveTag(EvergreenProjectSummary{}, "Versions")
	costReportEvergrenProjectSummaryResourceUseKey = bsonutil.MustHaveTag(EvergreenProjectSummary{}, "ResourceUse")
)

type CloudProviderSummary struct {
	Name     string             `bson:"name" json:"name" yaml:"name"`
	Services map[string]float64 `bson:"services" json:"services" yaml:"services"`
}

var (
	costReportCloudProviderSummaryNameKey     = bsonutil.MustHaveTag(CloudProviderSummary{}, "Name")
	costReportCloudProviderSummaryServicesKey = bsonutil.MustHaveTag(CloudProviderSummary{}, "Services")
)

func NewCostReportSummary(r *CostReport) *CostReportSummary {
	r.refresh()

	out := CostReportSummary{
		Metadata: r.Report,
		env:      r.env,
	}

	for _, p := range r.Evergreen.Projects {
		psum := EvergreenProjectSummary{Name: p.Name}
		commitSet := map[string]struct{}{}
		resources := map[string]EvergreenResourceCostSummary{}

		for _, t := range p.Tasks {
			commitSet[t.Githash] = struct{}{}
			distro := r.Evergreen.distro[t.Distro]
			if distro.InstanceType == "" {
				r := resources[distro.Name]
				r.Name = distro.Name
				r.Seconds += t.TaskSeconds
				r.Cost += t.EstimatedCost
				resources[distro.Name] = r
			} else {
				r := resources[distro.InstanceType]
				r.Name = distro.Name
				r.Seconds += t.TaskSeconds
				r.Cost += t.EstimatedCost
				resources[distro.InstanceType] = r
			}
		}

		psum.Versions = len(commitSet)
		for _, rsum := range resources {
			psum.Resources = append(psum.Resources, rsum)
		}
		out.EvergreenProjects = append(out.EvergreenProjects, psum)
	}

	for _, p := range r.Providers {
		psum := CloudProviderSummary{Name: p.Name, Services: map[string]float64{}}
		for _, account := range p.Accounts {
			for _, service := range account.Services {
				psum.Services[service.Name] += float64(service.Cost)
				out.TotalCost += float64(service.Cost)
			}
		}

		out.ProviderSummary = append(out.ProviderSummary, psum)
	}
	out.populated = true
	out.ID = r.ID

	return &out
}

func (r *CostReportSummary) Setup(e sink.Environment) { r.env = e }
func (r *CostReportSummary) IsNil() bool              { return r.populated }
func (r *CostReportSummary) Save() error {
	if !r.populated {
		return errors.New("cannot save unpopulated report")
	}

	conf, session, err := sink.GetSessionWithConfig(r.env)
	if err != nil {
		return errors.WithStack(err)
	}
	defer session.Close()

	changeInfo, err := session.DB(conf.DatabaseName).C(costReportSummaryCollection).UpsertId(r.ID, r)
	grip.Debug(message.Fields{
		"ns":          model.Namespace{DB: conf.DatabaseName, Collection: costReportSummaryCollection},
		"id":          r.ID,
		"operation":   "save build cost report",
		"change-info": changeInfo,
	})

	return errors.Wrap(err, "problem saving cost report summaryt")
}

func (r *CostReportSummary) String() string {
	jsonReport, err := json.MarshalIndent(r, "", "    ")
	if err != nil {
		return ""
	}

	return string(jsonReport)
}
