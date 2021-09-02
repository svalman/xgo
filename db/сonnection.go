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
