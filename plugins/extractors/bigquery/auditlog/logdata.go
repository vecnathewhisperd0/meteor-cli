package auditlog

import (
	"github.com/odpf/meteor/models"
	"github.com/pkg/errors"
	loggingpb "google.golang.org/genproto/googleapis/cloud/bigquery/logging/v1"
)

const serviceName = "bigquery"

type LogData struct {
	*loggingpb.AuditData
}

func (ld *LogData) validateAuditData() (err error) {
	if ld.GetJobCompletedEvent() == nil {
		err = errors.New("can't found jobCompletedEvent field")
		return
	}

	job := ld.GetJobCompletedEvent().GetJob()
	if job == nil {
		err = errors.New("can't found jobCompletedEvent.job field")
		return
	}

	jobStatus := job.GetJobStatus()
	if jobStatus == nil {
		err = errors.New("can't found jobCompletedEvent.job.jobStatus field")
		return
	}

	jobState := jobStatus.GetState()
	if jobState == "" {
		err = errors.New("jobCompletedEvent.job.jobStatus.state is empty")
		return
	}

	// ignoring the job that has not finished
	if jobState != "DONE" {
		err = errors.New("job status state is not DONE")
		return
	}

	if jobStatus.GetError() != nil {
		if jobErrMsg := jobStatus.GetError().GetMessage(); jobErrMsg != "" {
			err = errors.Errorf("job status has error: %s", jobErrMsg)
			return
		}
		err = errors.Errorf("job status error is not nil but cannot get the message")
		return
	}

	return
}

func (ld *LogData) GetReferencedTablesURN() (refTablesURN []string) {
	refTablesURN = []string{}
	var stats *loggingpb.JobStatistics
	if stats = ld.GetJobCompletedEvent().GetJob().GetJobStatistics(); stats == nil {
		return
	}
	for _, rt := range stats.ReferencedTables {
		tableURN := models.TableURN(serviceName, rt.ProjectId, rt.DatasetId, rt.TableId)
		refTablesURN = append(refTablesURN, tableURN)
	}
	return
}
