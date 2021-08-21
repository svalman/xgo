package adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/nakagami/firebirdsql"
	"github.com/svalman/xgo/db"
)

type (
	FbAdapter struct {
		params   db.DbConnectionParams
		dsn      string
		haveConn bool
		conn     *sql.DB
		ctx      context.Context
	}
)

func NewFirebirdAdapter(ds *db.DbConnectionParams) (*FbAdapter, error) {
	if ds == nil {
		return nil, errors.New("Нет сведений о соединении: NewAdapter ds nil.")
	}
	fbdsn := fmt.Sprintf(
		"%s:%s@%s:%s/%s?charset=utf8",
		ds.Username,
		ds.Password,
		ds.Host,
		ds.Port,
		ds.Dbname)

	db := &FbAdapter{
		params:   *ds,
		haveConn: false,
		dsn:      fbdsn,
	}

	return db, nil
}

func (fa *FbAdapter) Context() context.Context {
	return fa.ctx
}

func (fa *FbAdapter) Open(ctx context.Context) error {

	if ctx == nil {
		ctx = context.Background()
	}
	fa.ctx = ctx

	conn, err := sql.Open("firebirdsql", fa.dsn)
	if err != nil {
		return err
	}
	fa.haveConn = true
	fa.conn = conn
	return nil
}

func (fa *FbAdapter) Close(ctx context.Context) error {
	if fa.haveConn == false {
		return nil
	}
	if err := fa.conn.Close(); err != nil {
		return err
	}
	return nil
}

func (fa *FbAdapter) Status() string {
	if fa.haveConn {
		return "connected"
	} else {
		return "closed"
	}
}
