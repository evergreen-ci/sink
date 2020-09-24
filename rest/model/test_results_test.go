package model

import (
	"fmt"
	"testing"
	"time"

	dbmodel "github.com/evergreen-ci/cedar/model"
	"github.com/stretchr/testify/assert"
)

func TestTestResultImport(t *testing.T) {
	t.Run("InvalidType", func(t *testing.T) {
		apiTestResult := &APITestResult{}
		assert.Error(t, apiTestResult.Import(dbmodel.TestResults{}))
	})
	t.Run("ValidTestResult", func(t *testing.T) {
		tr := dbmodel.TestResult{
			TaskID:         "task_id",
			Execution:      2,
			TestName:       "test_name",
			Trial:          3,
			Status:         "pass",
			LineNum:        102,
			TaskCreateTime: time.Now().Add(-time.Hour),
			TestStartTime:  time.Now().Add(-30 * time.Minute),
			TestEndTime:    time.Now(),
		}
		expected := &APITestResult{
			TaskID:         ToAPIString(tr.TaskID),
			Execution:      tr.Execution,
			TestName:       ToAPIString(tr.TestName),
			Trial:          tr.Trial,
			Status:         ToAPIString(tr.Status),
			LogURL:         ToAPIString(fmt.Sprintf("test_log/%s/%d/%s", tr.TaskID, tr.Execution, tr.TestName)),
			LineNum:        tr.LineNum,
			TaskCreateTime: NewTime(tr.TaskCreateTime),
			TestStartTime:  NewTime(tr.TestStartTime),
			TestEndTime:    NewTime(tr.TestEndTime),
		}
		apiTestResult := &APITestResult{}
		assert.NoError(t, apiTestResult.Import(tr))
		assert.Equal(t, expected, apiTestResult)
	})
}
