package store

import "imserver/pkg/model"

type LogStoreType interface {
	CreateLog(log model.Log) error
	CreateLogs(logs []model.Log) error
	GetTotal(sourceID int64) (int64, error)
	GetLogs(sourceID int64, count int) (*[]model.Log, error)
}
