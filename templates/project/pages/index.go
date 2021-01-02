package pages

import (
	"context"
	"time"

	"github.com/influx6/npkg/nenv"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/sabuhp/ochestrator"

	"github.com/influx6/groundlayer/pkg/peji"

	"github.com/influx6/groundlayer/templates/project/pages/hello"
	"github.com/influx6/groundlayer/templates/project/pages/not_found"
)

const (
	PagePrefix = "/pages"
	PageRoute  = "/pages/*path" // *path tells the router to match anything after the /pages
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
		Theme,
		station.Transport(),
		not_found.NotFound,
	)

	var router = station.Router()
	pages.AddOnPageRoute(func(route string, _ *peji.Pages) {
		_ = router.Event(route, pages)
	})

	if addErr := pages.Add("hello", hello.CreateHelloWorldPage); addErr != nil {
		return nil, nerror.WrapOnly(addErr)
	}

	router.HttpService(PageRoute, pages, "HEAD", "GET")
	return pages, nil
}
