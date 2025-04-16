package config

import (
	"errors"
	"fmt"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

type Config struct {
	LogLevel LogLevel `name:"log-level" description:"Logging level for the application. One of debug, info, warn, or error" default:"info"`
	Storage  Storage  `name:"storage" description:"Storage configuration"`
	HTTP     HTTP     `name:"http" description:"HTTP server configuration"`
	Metrics  Metrics  `name:"metrics" description:"Metrics server configuration"`
	PProf    PProf    `name:"pprof" description:"PProf server configuration"`
}

type HTTP struct {
	Address string `name:"address" description:"Address to listen on"`
	Port    int    `name:"port" description:"Port to listen on" default:"8080"`
}

type Metrics struct {
	Enabled bool   `name:"enabled" description:"Enable metrics server"`
	Address string `name:"address" description:"Address to listen on"`
	Port    int    `name:"port" description:"Port to listen on" default:"9000"`
}

type PProf struct {
	Enabled bool   `name:"enabled" description:"Enable pprof server"`
	Address string `name:"address" description:"Address to listen on"`
	Port    int    `name:"port" description:"Port to listen on" default:"9999"`
}

type StorageType string

const (
	StorageTypeRedis    StorageType = "redis"
	StorageTypeMySQL    StorageType = "mysql"
	StorageTypePostgres StorageType = "postgres"
	StorageTypeSQLite   StorageType = "sqlite"
)

type Storage struct {
	Type StorageType `name:"type" description:"Storage type. One of redis, mysql, postgres, sqlite" default:"sqlite"`
	DSN  string      `name:"dsn" description:"Data source name for the storage"`
}

var (
	ErrBadLogLevel    = errors.New("invalid log level provided")
	ErrBadStorageType = errors.New("invalid storage type provided")
)

func (c Config) Validate() error {
	if c.LogLevel != LogLevelDebug &&
		c.LogLevel != LogLevelInfo &&
		c.LogLevel != LogLevelWarn &&
		c.LogLevel != LogLevelError {
		return fmt.Errorf("%w: %s", ErrBadLogLevel, c.LogLevel)
	}

	if c.Storage.Type != StorageTypeRedis &&
		c.Storage.Type != StorageTypeMySQL &&
		c.Storage.Type != StorageTypePostgres &&
		c.Storage.Type != StorageTypeSQLite {
		return fmt.Errorf("%w: %s", ErrBadStorageType, c.Storage.Type)
	}

	// TODO: validate HTTP addresses and ports
	// TODO: validate storage DSN

	return nil
}
