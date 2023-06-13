package configuration

import (
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/keepalive"
)

type Config struct {
	Logger          zap.Config             `yaml:"logger"`
	HTTPServer      ServerConfiguration    `yaml:"httpServer"`
	Storage         StorageConfiguration   `yaml:"storage"`
	GrpcServer      GrpcConfiguration      `yaml:"grpcServer"`
	Scheduler       SchedulerConfiguration `yaml:"scheduler"`
	ShutdownTimeout time.Duration          `yaml:"shutdownTimeout"`
}

type ServerConfiguration struct {
	Host         string        `yaml:"host"`
	Port         int           `yaml:"port"`
	IdleTimeout  time.Duration `yaml:"idleTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
}

type StorageConfiguration struct {
	UsePostgresStorage bool   `yaml:"usePostgres"`
	PostgresConnection string `yaml:"postgresConnection"`
}

type GrpcConfiguration struct {
	Host              string                      `yaml:"host"`
	Port              int                         `yaml:"port"`
	EnforcementPolicy keepalive.EnforcementPolicy `yaml:"enforcementPolicy"`
	ServerParameters  keepalive.ServerParameters  `yaml:"serverParameters"`
}

type SchedulerConfiguration struct {
	ScanDelay      time.Duration `yaml:"scanDelay"`
	CleanDelay     time.Duration `yaml:"cleanDelay"`
	CleanOlderThan time.Duration `yaml:"cleanOlderThan"`
}

func NewConfig(configFile string) (Config, error) {
	config := Config{}

	err := Configure(&config, configFile)
	if err != nil {
		return config, err
	}

	return config, nil
}
