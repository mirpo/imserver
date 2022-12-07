package db

import (
	"context"
	immudb "github.com/codenotary/immudb/pkg/client"
	"log"
	"os"
)

func executeSql(dbClient immudb.ImmuClient, sqlString string) {
	_, err := dbClient.SQLExec(context.Background(), sqlString, map[string]interface{}{})
	if err != nil {
		log.Fatal(err)
	}
}

func initDb(dbClient immudb.ImmuClient) {
	sqlString := `CREATE TABLE IF NOT EXISTS logs (
    				id          INTEGER AUTO_INCREMENT, 
    				source_id   INTEGER,
    				created_at  TIMESTAMP,
    				metrics     VARCHAR NOT NULL,
    				PRIMARY KEY (id)
			     )`
	executeSql(dbClient, sqlString)

	sqlString = `CREATE TABLE IF NOT EXISTS sources (
    				id          INTEGER AUTO_INCREMENT, 
    				name        VARCHAR,
					roles       VARCHAR,
    				PRIMARY KEY (id)
			     )`
	executeSql(dbClient, sqlString)

	sqlString = `UPSERT INTO sources (id, name, roles) 
				 VALUES				 (1, 'serviceA', 'ROLE_READ,ROLE_WRITE'),
									 (2, 'serviceB', 'ROLE_WRITE'),
									 (3, 'serviceC', '');
				`
	executeSql(dbClient, sqlString)
}

func NewClient() immudb.ImmuClient {
	opts := immudb.
		DefaultOptions().
		WithAddress(os.Getenv("IMMUDB_HOST")).
		WithPort(3322)

	client := immudb.NewClient().WithOptions(opts)

	err := client.OpenSession(
		context.Background(),
		[]byte(os.Getenv("IMMUDB_USERNAME")),
		[]byte(os.Getenv("IMMUDB_PASSWORD")),
		os.Getenv("IMMUDB_DATABASE"),
	)
	if err != nil {
		log.Fatal(err)
	}

	initDb(client)

	return client
}
