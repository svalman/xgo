package db

import (
	"github.com/svalman/xgo/db/adapter"
)

type (
	DbConnectionParams struct {
		Adapter  string `yaml:"adapter"`
		Dbname   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	Connection struct {
		connParams *DbConnectionParams
		db         *adapter.IAdapter
	}
)

func NewConnection(connParams *DbConnectionParams) (*Connection, error) {

	c := &Connection{connParams: connParams}

	if connParams.Adapter == "" {

	}

	return c, nil

}
