package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"xml-diff/xgo/db/adapter"
	"xml-diff/xgo/xconfig"
)

/*
Использование модуля:
1. Прочитать файл настроек подключений  к БД LoadConnectionsConfig
2. Открыть подключение NewConnection для базы, которая есть в конфигурации

*/

type (
	Connection struct {
		connParams *xconfig.DbConnectionParams
		adapter    adapter.IAdapter
		logger     *logrus.Logger
	}
)

func NewConnection(connParams *xconfig.DbConnectionParams, logger *logrus.Logger) (*Connection, error) {

	c := &Connection{
		connParams: connParams,
		logger:     logger,
	}

	if connParams.Adapter != adapter.Postgres {
		return nil, errors.New("Адаптер " + connParams.Adapter + " не поддерживается")
	}
	return c, nil
}

func (c *Connection) Connect(ctx context.Context) error {
	if len(c.connParams.Dbname) == 0 {
		return errors.New("Не указана база, с которой следует соединиться\n")
	}

	//ds := c.connParams. xconfig.GetDatasource(c.Catalog.DbName)
	//if ds == nil {
	//	msg := fmt.Sprintf("В файле настроек %s не указаны параметры соединения с СУБД %s\n",
	//		c.xconfig.GetConfigName(),
	//		c.Catalog.DbName)
	//	return errors.New(msg)
	//}

	conn, err := adapter.GetAdapter(c.connParams)
	if err != nil {
		msg := fmt.Sprintf("Не настроить параметры соединения с СУБД %s. Ошибка: %s\n",
			c.connParams.Dbname, err.Error())
		return errors.New(msg)
	} else {
		c.adapter = conn
	}

	//open db
	if err = conn.Open(ctx); err != nil {
		e := fmt.Errorf("Невозможно соединиться с СУБД %s. Ошибка: %s",
			c.connParams.Dbname, err.Error())
		return e
	}
	return nil
}

func (c *Connection) GetDbAdapter() adapter.IAdapter {
	return c.adapter
}
