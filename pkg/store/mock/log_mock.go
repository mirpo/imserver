package storeMock

import (
	"github.com/codenotary/immudb/pkg/errors"
	"imserver/pkg/model"
	"imserver/pkg/store"
)

type LogStoreMock struct {
	store.LogStore
}

func (l *LogStoreMock) GetLogs(sourceID int64, count int) (*[]model.Log, error) {
	if sourceID == 1 {
		logs := []model.Log{
			{
				SourceID:  sourceID,
				CreatedAT: 12345,
				Metrics:   "metric1",
			},
			{
				SourceID:  sourceID,
				CreatedAT: 23456,
				Metrics:   "metric2",
			},
		}
		return &logs, nil
	}
	return nil, errors.New("error")
}
