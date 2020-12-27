package service

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/influx6/npkg/nenv"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/njson"
	"github.com/influx6/sabuhp"
	"github.com/influx6/sabuhp/actions"
	"github.com/influx6/sabuhp/codecs"
	"github.com/influx6/sabuhp/injectors"
	"github.com/influx6/sabuhp/ochestrator"
)

const (
	ProjectName      = "[project]"
	ProjectNamespace = "[PROJECT]"
	EnvConfigAlias   = "*npkg.EnvStore"
)

var Logger = &loggerImpl{}

type loggerImpl struct{}

func (l loggerImpl) Log(cb *njson.JSON) {
	log.Printf("%s\n", cb.Message())
}

func BeforeInit(station *ochestrator.Station, envConfig *nenv.EnvStore) error {
	station.CreateInjector = UseInjector
	station.CreateCodec = UseCodec

	var redisAddr, hasRedisAddr = envConfig.GetString("REDIS_ADDR")
	if !hasRedisAddr {
		return nerror.New("address for redis is required %q", envConfig.KeyFor("REDIS_ADDR"))
	}

	station.CreateTransport = ochestrator.RedisTransportWithOptions(redis.Options{
		Network: "tcp",
		Addr:    redisAddr,
	})

	return nil
}

func AfterInit(station *ochestrator.Station, envConfig *nenv.EnvStore) error {

	return nil
}

func UseCodec(ctx context.Context, logger sabuhp.Logger) (sabuhp.Codec, error) {
	return &codecs.JsonCodec{}, nil
}

func UseInjector(ctx context.Context, logger sabuhp.Logger, envConfig *nenv.EnvStore) (*injectors.Injector, error) {
	var injector = injectors.NewInjector()

	// inject env into DI
	if injectErr := injector.AddValue(envConfig, EnvConfigAlias); injectErr != nil {
		return nil, nerror.WrapOnly(injectErr)
	}

	// Add more DI items for dependency injections

	return injector, nil
}

func AddActions(registry *actions.WorkerTemplateRegistry) {
	registry.Register(HelloWorldRequest)
}
