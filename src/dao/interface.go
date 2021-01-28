package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DAO interface {
}

type dao struct {
	connectionString string
	connection       *sql.DB
}

type PostgresBuilder struct {
	dao
}

func (pb PostgresBuilder) SetConnectionString(connectionString string) PostgresBuilder {
	a := pb
	a.connectionString = connectionString
	return a
}

func (pb PostgresBuilder) Build() *dao {
	var err error
	psqlInfo := fmt.Sprintf("%s", pb.connectionString)

	pb.connection, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error opening connection: ", err)
		panic(err) //TODO: either import logger into DAO as well, or do something better than panic!
	}

	return &pb.dao
}
