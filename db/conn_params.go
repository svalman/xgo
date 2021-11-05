package db

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync"
	"github.com/svalman/xgo/errors"
)

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
		DefaultSource string `yaml:"default_source"`
		Connections map[string]ConnectionParams `yaml:"connections"`
		cLock       sync.RWMutex
	}
)

/*
LoadConnectionsConfig
Формат файла

datasources:
  # служебная СУБД
  postgres:
    adapter: postgres
    dbname : postgres
    host: 127.0.0.1
    port: 5432
    username: postgres
    password: postgres

  bverify:
    adapter: postgres
    dbname: bverify
    host: 127.0.0.1
    port: 5432
    username: postgres
    password: post
 */
func LoadConnectionsConfig(configName string) (*ConnectionsConfig, error) {
	f, err := os.Open(configName)
	if err != nil {
		return nil, errors.Newf("Ошибка открытия файла конфигурации %v %w",
			configName, err)
	}
	defer f.Close()

	var cfg = new(ConnectionsConfig)
	cfg.cLock.Lock()
	defer cfg.cLock.Unlock()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(cfg); err != nil {
		return nil, errors.Newf("Ошибка разбора файла конфигурации %v %w",
			configName, err)
	}
	cfg.configName = configName
	return cfg, nil
}

func (cc *ConnectionsConfig) GetDatasource(dbname string) *ConnectionParams {
	cc.cLock.RLock()
	defer cc.cLock.RUnlock()

	if ds, ok := cc.Connections[dbname]; ok {
		return &ds
	} else {
		return nil
	}
}

func (cc *ConnectionsConfig) GetConfigName() string {
	return cc.configName
}

func (cc *ConnectionsConfig) GetDefaultDatasource() (*ConnectionParams, error) {
	if len(cc.DefaultSource)==0 {
		return nil, errors.New("В файле конфигурации не задан default_datasource")
	}
	cc.cLock.RLock()
	defer cc.cLock.RUnlock()

	c, ok := cc.Connections[cc.DefaultSource]
	if !ok {
		return nil, errors.Newf("В файле конфигурации задан неправильный default_datasource со значением %s. "+
									 "Этого параметра нет в datasources", cc.DefaultSource)
	}
	return &c, nil
}