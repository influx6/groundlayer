package pages

import (
	"context"
	"time"

	"github.com/influx6/groundlayer/pkg/peji"
	"github.com/influx6/npkg/nenv"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/sabuhp/ochestrator"

	"github.com/influx6/groundlayer/templates/project/pages/hello"
	"github.com/influx6/groundlayer/templates/project/pages/not_found"
)

const (
	PagePrefix = "/pages"
	PageRoute  = "/pages/*path"
)

func CreatePages(
	ctx context.Context,
	logger peji.Logger,
	station *ochestrator.Station,
	envConfig *nenv.EnvStore,
	maxIdleness time.Duration,
	idleCheckDuration time.Duration,
) (*peji.Pages, error) {
	var pages = peji.NewPages(
		ctx,
		logger,
		PagePrefix,
		maxIdleness,
		idleCheckDuration,
		station.Transport(),
		not_found.NotFound,
	)

	if addErr := pages.Add("hello", hello.CreateHelloWorldPage); addErr != nil {
		return nil, nerror.WrapOnly(addErr)
	}

	station.Router().HttpService(PageRoute, pages, "HEAD", "GET")
	return pages, nil
}
