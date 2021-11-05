package adapter

import (
	"context"
	"errors"
	"github.com/svalman/xgo/config"
	"time"
)

type (
	IAdapter interface {
		Context() context.Context
		Open(ctx context.Context) error
		Close(ctx context.Context) error
		Status() string
	}

	TAdapter struct {
		Params   config.DbConnectionParams
		Dsn      string
		HaveConn bool
		Ctx      context.Context
	}
)

func (ta *TAdapter) Context() context.Context {
	return ta.Ctx
}

var (
	DbConnectTimeout = 5 * time.Second
)

const (
	Firebird = "firebird"
	Postgres = "postgres"
)

func GetAdapter(params *config.DbConnectionParams) (IAdapter, error) {
	if params.Adapter == Firebird {
		return NewFirebirdAdapter(params)
	}
	if params.Adapter == Postgres {
		return NewPostgresAdapter(params)
	}

	return nil, errors.New("Тип " + params.Adapter + " не реализован")
}
