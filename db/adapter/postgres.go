package adapter

import (
	"context"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4/stdlib"
	"xml-diff/xgo/xconfig"
)

type (
	PgAdapter struct {
		TAdapter
		connConfig *pgx.ConnConfig
		conn       *pgx.Conn
		pool       *pgxpool.Pool
	}
)

func NewPostgresAdapter(ds *xconfig.DbConnectionParams) (*PgAdapter, error) {
	if ds == nil {
		return nil, errors.New("Нет сведений о соединении")
	}

	sqlxDsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		ds.Username,
		ds.Password,
		ds.Host,
		ds.Port,
		ds.Dbname)

	db := &PgAdapter{
		TAdapter: TAdapter{
			Params:   *ds,
			HaveConn: false,
			Dsn:      sqlxDsn,
		},
	}

	cc, err := pgx.ParseConfig(sqlxDsn) // ParseConfig(db.dsn)
	if err != nil {
		return nil, err
	}

	cc.PreferSimpleProtocol = true // включить бинарный протокол
	cc.BuildStatementCache = nil   // отключить prepared statements
	//TODO отключить отмену запросов
	//CustomCancel: func(_ *pgx.Conn) error { return nil },
	db.connConfig = cc

	return db, nil
}

//Open - соединение с базой
func (db *PgAdapter) Open(ctx context.Context) error {

	var err error

	if ctx == nil {
		ctx = context.Background()
	}
	db.Ctx = ctx

	timedCtx, cancelFunc := context.WithTimeout(ctx, DbConnectTimeout)
	defer cancelFunc()

	conn, err := pgx.ConnectConfig(timedCtx, db.connConfig)
	if err != nil {
		return err
	}
	db.conn = conn
	db.HaveConn = true

	stdlib.RegisterConnConfig(db.connConfig)
	if db.pool, err = pgxpool.Connect(db.Ctx, db.Dsn); err != nil {
		return errors.New("Не установить соединение с СУБД: " + err.Error())
	}

	return nil
	// ВАЖНО! сначала настраивается драйвер PGX для соединения с СУБД
	// и только потом подключается SQLX
	//db.xconn, err = sqlx.Connect("pgx", db.Dsn)
	//if err != nil {
	//	return fmt.Errorf("Ошибка соединения с СУБД %s %w",
	//		db.Params.Dbname, err)
	//}
	//db.xconn.MapperFunc(strings.ToUpper)
	//connStr := stdlib.RegisterConnConfig(db.connConfig)
	//connPool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
	//	ConnConfig:     connConfig,
	//	AfterConnect:   nil,
	//	MaxConnections: 20,
	//	AcquireTimeout: 30 * time.Second,
	//})
	//if err != nil {
	//	return nil, error.New(err, "Call to pgx.NewConnPool failed")
	//}

}

func (db *PgAdapter) getPgx() *pgx.Conn {
	return db.conn
}

func (db *PgAdapter) _closePgxConnection(ctx context.Context) error {
	db.HaveConn = false
	return db.conn.Close(ctx)
}

//Close - закрыть все соединения с базой
func (db *PgAdapter) Close(ctx context.Context) error {
	return db._closePgxConnection(ctx)
}

func (db *PgAdapter) Status() string {
	if db.HaveConn {
		return "connected"
	} else {
		return "closed"
	}
}

func (dba *PgAdapter) Select(dst interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(dba.Ctx, dba.pool, dst, query, args)
}
