package db

import (
	"github.com/svalman/xgo/db/adapter"
	"github.com/svalman/xgo/db/types"
)

type (
	Connection struct {
		connParams *types.DbConnectionParams
		db         *adapter.IAdapter
	}
)

func NewConnection(connParams *types.DbConnectionParams) (*Connection, error) {

	c := &Connection{connParams: connParams}

	if connParams.Adapter == "" {

	}

	return c, nil

}
