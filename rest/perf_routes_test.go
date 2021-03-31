package rest

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/evergreen-ci/cedar/model"
	"github.com/evergreen-ci/cedar/rest/data"
	datamodel "github.com/evergreen-ci/cedar/rest/model"
	"github.com/evergreen-ci/gimlet"
	"github.com/stretchr/testify/suite"
)

type PerfHandlerSuite struct {
	sc         data.MockConnector
	rh         map[string]gimlet.RouteHandler
	apiResults map[string]datamodel.APIPerformanceResult

	suite.Suite
}

func (s *PerfHandlerSuite) setup() {
	s.sc = data.MockConnector{
		CachedPerformanceResults: map[string]model.PerformanceResult{
			"abc": model.PerformanceResult{
				ID: "abc",
				Info: model.PerformanceResultInfo{
					Version:  "1",
					Order:    1,
					TaskID:   "123",
					TaskName: "taskname0",
					Tags:     []string{"a", "b"},
					Mainline: true,
				},
			},
			"def": model.PerformanceResult{
				ID:          "def",
				CreatedAt:   time.Date(2018, time.December, 1, 1, 1, 0, 0, time.UTC),
				CompletedAt: time.Date(2018, time.December, 1, 2, 1, 0, 0, time.UTC),
				Info: model.PerformanceResultInfo{
					Version:  "1",
					Order:    2,
					TaskID:   "123",
					TaskName: "taskname0",
					Tags:     []string{"a"},
					Mainline: true,
				},
			},
			"ghi": model.PerformanceResult{
				ID:          "ghi",
				CreatedAt:   time.Date(2018, time.December, 1, 1, 1, 0, 0, time.UTC),
				CompletedAt: time.Date(2018, time.December, 1, 2, 1, 0, 0, time.UTC),
				Info: model.PerformanceResultInfo{
					Version:  "1",
					Order:    3,
					TaskID:   "123",
					TaskName: "taskname0",
					Tags:     []string{"b"},
					Mainline: true,
				},
			},
			"jkl": model.PerformanceResult{
				ID:          "jkl",
				CreatedAt:   time.Date(2018, time.December, 1, 1, 1, 0, 0, time.UTC),
				CompletedAt: time.Date(2018, time.December, 1, 2, 1, 0, 0, time.UTC),
				Info: model.PerformanceResultInfo{
					Version:  "1",
					Order:    4,
					TaskID:   "123",
					TaskName: "taskname0",
					Tags:     []string{"a", "b", "c", "d"},
					Mainline: true,
				},
			},
			"lmn": model.PerformanceResult{
				ID:          "lmn",
				CreatedAt:   time.Date(2018, time.December, 5, 1, 1, 0, 0, time.UTC),
				CompletedAt: time.Date(2018, time.December, 6, 2, 1, 0, 0, time.UTC),
				Info: model.PerformanceResultInfo{
					Version:  "2",
					Order:    1,
					TaskID:   "456",
					TaskName: "taskname1",
					Mainline: true,
				},
			},
			"delete": model.PerformanceResult{
				ID:          "delete",
				CreatedAt:   time.Date(2018, time.December, 5, 1, 1, 0, 0, time.UTC),
				CompletedAt: time.Date(2018, time.December, 6, 2, 1, 0, 0, time.UTC),
				Info: model.PerformanceResultInfo{
					Version:  "1",
					Order:    2,
					TaskID:   "456",
					TaskName: "taskname1",
					Mainline: true,
				},
			},
		},
	}
	s.sc.ChildMap = map[string][]string{
		"abc": []string{"def"},
		"def": []string{"jkl"},
	}
	s.rh = map[string]gimlet.RouteHandler{
		"id":            makeGetPerfById(&s.sc),
		"remove":        makeRemovePerfById(&s.sc),
		"task_id":       makeGetPerfByTaskId(&s.sc),
		"task_name":     makeGetPerfByTaskName(&s.sc),
		"version":       makeGetPerfByVersion(&s.sc),
		"children":      makeGetPerfChildren(&s.sc),
		"change_points": makePerfSignalProcessingRecalculate(&s.sc),
	}
	s.apiResults = map[string]datamodel.APIPerformanceResult{}
	for key, val := range s.sc.CachedPerformanceResults {
		apiResult := datamodel.APIPerformanceResult{}
		s.Require().NoError(apiResult.Import(val))
		s.apiResults[key] = apiResult
	}
}

func TestPerfHandlerSuite(t *testing.T) {
	s := new(PerfHandlerSuite)
	s.setup()
	suite.Run(t, s)
}

func (s *PerfHandlerSuite) TestPerfGetByIdHandlerFound() {
	rh := s.rh["id"]
	rh.(*perfGetByIdHandler).id = "abc"
	expected := s.apiResults["abc"]

	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	s.Require().NotNil(resp.Data())
	s.Equal(&expected, resp.Data())
}

func (s *PerfHandlerSuite) TestPerfGetByIdHandlerNotFound() {
	rh := s.rh["id"]
	rh.(*perfGetByIdHandler).id = "DNE"

	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.NotEqual(http.StatusOK, resp.Status())
}

func (s *PerfHandlerSuite) TestPerfRemoveByIdHandler() {
	rh := s.rh["remove"]
	rh.(*perfRemoveByIdHandler).id = "delete"

	_, ok := s.sc.CachedPerformanceResults["delete"]
	s.True(ok)
	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	_, ok = s.sc.CachedPerformanceResults["delete"]
	s.False(ok)

	// should not fail on non-existent id
	resp = rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
}

func (s *PerfHandlerSuite) TestPerfGetByTaskIdHandlerFound() {
	rh := s.rh["task_id"]
	rh.(*perfGetByTaskIdHandler).taskId = "123"
	rh.(*perfGetByTaskIdHandler).tags = []string{"d"}
	expected := []datamodel.APIPerformanceResult{s.apiResults["jkl"]}

	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	s.Require().NotNil(resp.Data())
	s.Equal(expected, resp.Data())
}

func (s *PerfHandlerSuite) TestPerfGetByTaskIdHandlerNotFound() {
	rh := s.rh["task_id"]
	rh.(*perfGetByTaskIdHandler).taskId = "555"

	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.NotEqual(http.StatusOK, resp.Status())
}

func (s *PerfHandlerSuite) TestPerfGetByTaskNameHandlerFound() {
	rh := s.rh["task_name"]
	rh.(*perfGetByTaskNameHandler).taskName = "taskname0"
	rh.(*perfGetByTaskNameHandler).interval = model.TimeRange{
		StartAt: time.Date(2018, time.November, 5, 0, 0, 0, 0, time.UTC),
		EndAt:   time.Now(),
	}
	rh.(*perfGetByTaskNameHandler).tags = []string{"b"}
	expected := []datamodel.APIPerformanceResult{
		s.apiResults["jkl"],
		s.apiResults["ghi"],
	}
	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	s.Require().NotNil(resp.Data())
	s.Equal(expected, resp.Data())

	rh.(*perfGetByTaskNameHandler).interval = model.TimeRange{
		StartAt: time.Time{},
		EndAt:   time.Now(),
	}
	rh.(*perfGetByTaskNameHandler).tags = []string{}
	expected = []datamodel.APIPerformanceResult{
		s.apiResults["jkl"],
		s.apiResults["ghi"],
		s.apiResults["def"],
		s.apiResults["abc"],
	}
	resp = rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	s.Require().NotNil(resp.Data())
	s.Equal(expected, resp.Data())

	rh.(*perfGetByTaskNameHandler).limit = 3
	expected = []datamodel.APIPerformanceResult{
		s.apiResults["jkl"],
		s.apiResults["ghi"],
		s.apiResults["def"],
	}
	resp = rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	s.Require().NotNil(resp.Data())
	s.Equal(expected, resp.Data())

}

func (s *PerfHandlerSuite) TestPerfGetByTaskNameHandlerNotFound() {
	rh := s.rh["task_name"]
	rh.(*perfGetByTaskNameHandler).taskName = "taskname2"
	rh.(*perfGetByTaskNameHandler).interval = model.TimeRange{
		StartAt: time.Date(2018, time.November, 5, 0, 0, 0, 0, time.UTC),
		EndAt:   time.Now(),
	}

	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.NotEqual(http.StatusOK, resp.Status())
}

func (s *PerfHandlerSuite) TestPerfGetByVersionHandlerFound() {
	rh := s.rh["version"]
	rh.(*perfGetByVersionHandler).version = "1"
	rh.(*perfGetByVersionHandler).tags = []string{"d"}
	expected := []datamodel.APIPerformanceResult{s.apiResults["jkl"]}

	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	s.Require().NotNil(resp.Data())
	s.Equal(expected, resp.Data())
}

func (s *PerfHandlerSuite) TestPerfGetByVersionHandlerNotFound() {
	rh := s.rh["version"]
	rh.(*perfGetByVersionHandler).version = "2"

	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.NotEqual(http.StatusOK, resp.Status())
}

func (s *PerfHandlerSuite) TestPerfGetChildrenHandlerFound() {
	rh := s.rh["children"]
	rh.(*perfGetChildrenHandler).id = "abc"
	rh.(*perfGetChildrenHandler).maxDepth = 1
	rh.(*perfGetChildrenHandler).tags = []string{"a"}
	expected := []datamodel.APIPerformanceResult{
		s.apiResults["abc"],
		s.apiResults["def"],
	}

	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	s.Require().NotNil(resp.Data())
	s.Equal(expected, resp.Data())
}

func (s *PerfHandlerSuite) TestPerfGetChildrenHandlerNotFound() {
	rh := s.rh["children"]
	rh.(*perfGetChildrenHandler).id = "DNE"
	rh.(*perfGetChildrenHandler).maxDepth = 5

	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.NotEqual(http.StatusOK, resp.Status())
}

func (s *PerfHandlerSuite) TestPerfSignalProcessingRecalculateHandlerFound() {
	rh := s.rh["change_points"]
	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	response := resp.Data().(struct{})
	s.Require().NotNil(response)
}

func (s *PerfHandlerSuite) TestPerfChangePointTriageMarkHandlerFound() {
	rh := s.rh["change_points"]
	resp := rh.Run(context.TODO())
	s.Require().NotNil(resp)
	s.Equal(http.StatusOK, resp.Status())
	response := resp.Data().(struct{})
	s.Require().NotNil(response)
}

func (s *PerfHandlerSuite) TestParse() {
	for _, test := range []struct {
		urlString string
		query     string
		handler   string
		limit     bool
	}{
		{
			handler:   "task_id",
			urlString: "http://example.com/perf/task_id/task_id0",
			query:     "?",
		},
		{
			handler:   "task_name",
			urlString: "http://example.com/perf/task_name/task_name0",
			query:     "?started_after=2020-03-15&finished_before=2021-09-01",
			limit:     true,
		},
		{
			handler:   "version",
			query:     "?",
			urlString: "http://example.com/perf/version/verison0",
		},
	} {
		s.T().Run(test.handler, func(t *testing.T) {
			s.testParseValid(test.handler, test.urlString, test.query, test.limit)
			s.testParseDefaults(test.handler, test.urlString, test.query, test.limit)
		})
	}
}

func (s *PerfHandlerSuite) testParseValid(handler, urlString, query string, limit bool) {
	ctx := context.Background()
	query = strings.Join([]string{query, "tags=hello", "tags=world", "limit=5"}, "&")
	req := &http.Request{Method: "GET"}
	url, err := url.Parse(urlString + query)
	s.Require().NoError(err)
	req.URL = url
	expectedTags := []string{"hello", "world"}
	rh := s.rh[handler]

	s.Require().NoError(rh.Parse(ctx, req))
	s.Equal(expectedTags, getPerfTags(rh, handler))
	if limit {
		s.Equal(5, getPerfLimit(rh, handler))
	}
}

func (s *PerfHandlerSuite) testParseDefaults(handler, urlString, query string, limit bool) {
	ctx := context.Background()
	req := &http.Request{Method: "GET"}
	url, err := url.Parse(urlString + query)
	s.Require().NoError(err)
	req.URL = url
	rh := s.rh[handler]

	s.NoError(rh.Parse(ctx, req))
	s.Nil(getPerfTags(rh, handler))
	if limit {
		s.Zero(getPerfLimit(rh, handler))
	}
}

func getPerfTags(rh gimlet.RouteHandler, handler string) []string {
	switch handler {
	case "task_id":
		return rh.(*perfGetByTaskIdHandler).tags
	case "task_name":
		return rh.(*perfGetByTaskNameHandler).tags
	case "version":
		return rh.(*perfGetByVersionHandler).tags
	default:
		return []string{}
	}
}

func getPerfLimit(rh gimlet.RouteHandler, handler string) int {
	switch handler {
	case "task_name":
		return rh.(*perfGetByTaskNameHandler).limit
	default:
		return 0
	}
}
