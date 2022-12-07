package store

import (
	"context"
	immudb "github.com/codenotary/immudb/pkg/client"
	"github.com/codenotary/immudb/pkg/errors"
	"imserver/pkg/model"
)

type SourceStoreType interface {
	Get(sourceID int64) (*model.Source, error)
}

type SourceStore struct {
	dbClient immudb.ImmuClient
}

func NewSourceStore(dbClient immudb.ImmuClient) SourceStoreType {
	return &SourceStore{
		dbClient: dbClient,
	}
}

func (s *SourceStore) Get(sourceID int64) (*model.Source, error) {
	query := "SELECT id, name, roles FROM sources WHERE id = @source_id LIMIT 1"

	res, err := s.dbClient.SQLQuery(
		context.TODO(),
		query,
		map[string]interface{}{
			"source_id": sourceID,
		},
		true,
	)

	if err != nil {
		return nil, err
	}

	if len(res.Rows) < 1 {
		return nil, errors.New("not found")
	}

	row := res.Rows[0]
	return &model.Source{
		SourceID: row.Values[0].GetN(),
		Name:     row.Values[1].GetS(),
		Roles:    row.Values[2].GetS(),
	}, nil
}
