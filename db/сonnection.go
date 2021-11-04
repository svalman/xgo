package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/svalman/xgo/db/adapter"
	"gopkg.in/yaml.v3"
	"os"
)

/*
Использование модуля:
1. Прочитать файл настроек подключений  к БД LoadConnectionsConfig
2. Открыть подключение NewConnection для из баз, которая есть в конфигурации

 */

type (
	ConnectionParams struct {
		Adapter  string `yaml:"adapter"`
		Dbname   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	ConnectionsConfig struct {
		configName  string
		Connections map[string]ConnectionParams `yaml:"connections"`
	}

	Connection struct {
		connParams *ConnectionParams
		adapter    *adapter.IAdapter
		logger     *logrus.Logger
		ctx 		context.Context
	}
)

func LoadConnectionsConfig(configName string) (*ConnectionsConfig, error) {
	f, err := os.Open(configName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия файла конфигурации %v %w",
			configName, err)
	}
	defer f.Close()

	var cfg = new(ConnectionsConfig)
	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(cfg); err != nil {
		return nil, fmt.Errorf("Ошибка разбора файла конфигурации %v %w",
			configName, err)
	}
	cfg.configName = configName
	return cfg, nil
}

func (cc *ConnectionsConfig) GetDatasource(dbname string) *ConnectionParams {
	if ds, ok := cc.Connections[dbname]; ok {
		return &ds
	} else {
		return nil
	}
}

func (cc *ConnectionsConfig) GetConfigName() string {
	return cc.configName
}

func NewConnection(connParams *ConnectionParams, logger *logrus.Logger) (*Connection, error) {

	c := &Connection{
		connParams: connParams,
		logger: logger,
		ctx : context.Background(),
	}

	if connParams.Adapter != adapter.Postgres {
		return nil,  errors.New("Адаптер "+connParams.Adapter+" не поддерживается")
	}
	return c, nil
}

func (c *Connection) Connect() error {
	if len(c.connParams.Dbname) == 0 {
		return errors.New("Не указана база, с которой следует соединиться\n")
	}

	//ds := c.connParams. config.GetDatasource(c.Catalog.DbName)
	//if ds == nil {
	//	msg := fmt.Sprintf("В файле настроек %s не указаны параметры соединения с СУБД %s\n",
	//		c.config.GetConfigName(),
	//		c.Catalog.DbName)
	//	return errors.New(msg)
	//}

	conn, err := adapter.GetAdapter(c.connParams)
	if err != nil {
		msg := fmt.Sprintf("Не настроить параметры соединения с СУБД %s. Ошибка: %s\n",
			c.connParams.Dbname, err.Error())
		return errors.New(msg)
	} else {
		c.adapter = &conn
	}

	//open db
	if err = conn.Open(c.ctx); err != nil {
		e := fmt.Errorf("Невозможно соединиться с СУБД %s. Ошибка: %s",
			c.connParams.Dbname, err.Error())
		return e
	}
	return nil
}

func (c *Connection) GetDbAdapter() *adapter.IAdapter {
	return c.adapter
}