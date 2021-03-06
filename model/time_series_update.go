package model

import (
	"context"
	"fmt"
	"time"

	"github.com/evergreen-ci/cedar"
	"github.com/mongodb/anser/bsonutil"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

// PerfAnalysis contains information about when the associated PerformanceResult was analyzed for change points.
type PerfAnalysis struct {
	ProcessedAt time.Time `bson:"processed_at" json:"processed_at" yaml:"processed_at"`
}

var (
	perfAnalysisProcessedAtKey = bsonutil.MustHaveTag(PerfAnalysis{}, "ProcessedAt")
)

// PerformanceResultSeriesID is the id of a group of PerformanceResults. It is used to create a job to analyze the group
// of PerfResults for change points.
type PerformanceResultSeriesID struct {
	Project     string           `bson:"project"`
	Variant     string           `bson:"variant"`
	Task        string           `bson:"task"`
	Test        string           `bson:"test"`
	Measurement string           `bson:"measurement"`
	Arguments   map[string]int32 `bson:"args"`
}

// String creates a string representation of a PerformanceResultSeriesID.
func (p PerformanceResultSeriesID) String() string {
	return fmt.Sprintf("%s %s %s %s %v", p.Project, p.Variant, p.Task, p.Test, p.Arguments)
}

// TimeSeriesEntry is information about a single PerfResult.
type TimeSeriesEntry struct {
	PerfResultID string  `bson:"perf_result_id"`
	Value        float64 `bson:"value"`
	Order        int     `bson:"order"`
	Version      string  `bson:"version"`
}

// PerformanceData contains information about PerformanceResults and their associated measurement data.
type PerformanceData struct {
	PerformanceResultId PerformanceResultSeriesID `bson:"_id"`
	TimeSeries          []TimeSeriesEntry         `bson:"time_series"`
}

// MarkPerformanceResultsAsAnalyzed marks all the PerformanceResults associated a PerformanceResultSeriesID as analyzed
// for change points.
func MarkPerformanceResultsAsAnalyzed(ctx context.Context, env cedar.Environment, performanceResultId PerformanceResultSeriesID) error {
	filter := bson.M{
		bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoProjectKey):   performanceResultId.Project,
		bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoVariantKey):   performanceResultId.Variant,
		bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTaskNameKey):  performanceResultId.Task,
		bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTestNameKey):  performanceResultId.Test,
		bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoArgumentsKey): performanceResultId.Arguments,
	}

	update := bson.M{
		"$currentDate": bson.M{
			bsonutil.GetDottedKeyName(perfAnalysisKey, perfAnalysisProcessedAtKey): true,
		},
	}
	_, err := env.GetDB().Collection(perfResultCollection).UpdateMany(ctx, filter, update)
	if err != nil {
		return errors.Wrapf(err, "Unable to mark performance results as analyzed for change points")
	}
	return nil
}

// GetPerformanceResultSeriesIdsNeedingTimeSeriesUpdate queries the database and gets all the PerformanceResultSeriesIDs
// for PerformanceResults that haven't been analyzed for change points yet.
func GetPerformanceResultSeriesIdsNeedingTimeSeriesUpdate(ctx context.Context, env cedar.Environment) ([]PerformanceResultSeriesID, error) {
	cur, err := env.GetDB().Collection(perfResultCollection).Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoOrderKey):    bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoMainlineKey): true,
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoProjectKey):  bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoVariantKey):  bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTaskNameKey): bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTestNameKey): bson.M{"$exists": true},
			},
		},
		{
			"$match": bson.M{
				"$expr": bson.M{
					"$lt": []string{
						"$" + bsonutil.GetDottedKeyName(perfAnalysisKey, perfAnalysisProcessedAtKey),
						"$" + bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsProcessedAtKey),
					},
				},
			},
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"project": "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoProjectKey),
					"variant": "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoVariantKey),
					"task":    "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTaskNameKey),
					"test":    "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTestNameKey),
				},
			},
		},
		{
			"$replaceRoot": bson.M{
				"newRoot": "$_id",
			},
		},
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to get metrics needing change point detection")
	}
	defer cur.Close(ctx)
	var res []PerformanceResultSeriesID
	err = cur.All(ctx, &res)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to decode metrics needing change results")
	}
	return res, nil
}

// GetPerformanceResultSeriesIDs finds the PerformanceResultSeriesIDs for PerformanceResults that need to be processed
// to find change points in them.
func GetPerformanceResultSeriesIDs(ctx context.Context, env cedar.Environment) ([]PerformanceResultSeriesID, error) {
	cur, err := env.GetDB().Collection(perfResultCollection).Aggregate(ctx, []bson.M{
		{
			"$match": bson.M{
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoProjectKey):  bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoVariantKey):  bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoOrderKey):    bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTaskNameKey): bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTestNameKey): bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoMainlineKey): true,
			},
		},
		{
			"$unwind": "$" + bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsStatsKey),
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"project": "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoProjectKey),
					"variant": "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoVariantKey),
					"task":    "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTaskNameKey),
					"test":    "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTestNameKey),
				},
			},
		},
		{
			"$replaceRoot": bson.M{
				"newRoot": "$_id",
			},
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "Cannot aggregate time series ids")
	}
	defer cur.Close(ctx)
	var res []PerformanceResultSeriesID
	err = cur.All(ctx, &res)
	if err != nil {
		return nil, errors.Wrap(err, "Could not decode time series ids")
	}
	return res, nil
}

// GetPerformanceData gets the PerfResult time series data associated with the given PerformanceResultSeriesID.
func GetPerformanceData(ctx context.Context, env cedar.Environment, performanceResultId PerformanceResultSeriesID) ([]PerformanceData, error) {
	pipe := []bson.M{
		{
			"$match": bson.M{
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoProjectKey):  performanceResultId.Project,
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoVariantKey):  performanceResultId.Variant,
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoOrderKey):    bson.M{"$exists": true},
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTaskNameKey): performanceResultId.Task,
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTestNameKey): performanceResultId.Test,
				bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoMainlineKey): true,
			},
		},
		{
			"$unwind": "$" + bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsStatsKey),
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"project":     "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoProjectKey),
					"variant":     "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoVariantKey),
					"task":        "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTaskNameKey),
					"test":        "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoTestNameKey),
					"measurement": "$" + bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsStatsKey, perfRollupValueNameKey),
					"args":        "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoArgumentsKey),
				},
				"time_series": bson.M{
					"$push": bson.M{
						"value": bson.M{
							"$ifNull": bson.A{
								"$" + bsonutil.GetDottedKeyName(perfRollupsKey, perfRollupsStatsKey, perfRollupValueValueKey),
								0,
							},
						},
						"order":          "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoOrderKey),
						"perf_result_id": "$_id",
						"version":        "$" + bsonutil.GetDottedKeyName(perfInfoKey, perfResultInfoVersionKey),
					},
				},
			},
		},
	}
	cur, err := env.GetDB().Collection(perfResultCollection).Aggregate(ctx, pipe)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot aggregate time series")
	}
	defer cur.Close(ctx)
	var res []PerformanceData
	err = cur.All(ctx, &res)
	if err != nil {
		return nil, errors.Wrap(err, "Could not decode time series")
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res, nil
}
