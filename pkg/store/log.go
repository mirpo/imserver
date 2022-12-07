package store

import (
	"context"
	"fmt"
	immudb "github.com/codenotary/immudb/pkg/client"
	"imserver/pkg/model"
)

type LogStore struct {
	dbClient immudb.ImmuClient
}

func NewLogStore(dbClient immudb.ImmuClient) LogStoreType {
	return &LogStore{
		dbClient: dbClient,
	}
}

const insertLogQuery = "INSERT INTO logs (source_id, created_at, metrics) VALUES (@source_id, NOW(), @metrics)"

func (lg *LogStore) CreateLog(log model.Log) error {
	_, err := lg.dbClient.SQLExec(
		context.TODO(),
		insertLogQuery,
		map[string]interface{}{
			"source_id": log.SourceID,
			"metrics":   log.Metrics,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (lg *LogStore) CreateLogs(logs []model.Log) error {
	tx, err := lg.dbClient.NewTx(context.TODO())
	if err != nil {
		return err
	}

	for _, log := range logs {
		err = tx.SQLExec(
			context.TODO(),
			insertLogQuery,
			map[string]interface{}{
				"source_id": log.SourceID,
				"metrics":   log.Metrics,
			},
		)
		if err != nil {
			return err
		}
	}

	_, err = tx.Commit(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func (lg *LogStore) GetTotal(sourceID int64) (int64, error) {
	res, err := lg.dbClient.SQLQuery(
		context.TODO(),
		"SELECT COUNT(*) FROM logs WHERE source_id = @source",
		map[string]interface{}{
			"source": sourceID,
		},
		true,
	)

	if err != nil {
		return 0, err
	}

	return res.Rows[0].Values[0].GetN(), nil
}

func (lg *LogStore) GetLogs(sourceID int64, count int) (*[]model.Log, error) {
	query := "SELECT source_id, created_at, metrics FROM logs WHERE source_id = @source"

	if count > 0 {
		query += fmt.Sprintf(" LIMIT %d", count)
	}

	res, err := lg.dbClient.SQLQuery(
		context.TODO(),
		query,
		map[string]interface{}{
			"source": sourceID,
		},
		true,
	)

	if err != nil {
		return nil, err
	}

	logs := []model.Log{}
	for _, row := range res.Rows {
		log := model.Log{
			SourceID:  row.Values[0].GetN(),
			CreatedAT: row.Values[1].GetTs(),
			Metrics:   row.Values[2].GetS(),
		}
		logs = append(logs, log)
	}

	return &logs, nil
}
