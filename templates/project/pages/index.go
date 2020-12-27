package pages

import (
	"context"
	"time"

	"github.com/influx6/npkg/nenv"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/njson"
	"github.com/influx6/sabuhp/ochestrator"

	"github.com/influx6/groundlayer/pkg/peji"

	"github.com/influx6/groundlayer/templates/project/pages/hello"
	"github.com/influx6/groundlayer/templates/project/pages/not_found"
)

const (
	PagePrefix     = "/pages"
	PageRoute      = "/pages/*path" // *path tells the router to match anything after the /pages
	PageEventRoute = "/pages/*"
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

	var router = station.Router()
	pages.AddOnPageRoute(func(route string, _ *peji.Page) {
		var stack = njson.Log(logger)
		stack.Message("Received new page route").String("route", route).End()
		njson.ReleaseLogStack(stack)

		// add event for route, so that event pipelines can also request page
		// based routes through pubsub.
		_ = router.Event(route, pages)
	})

	router.Service(PageEventRoute, PageRoute, pages, "HEAD", "GET")
	return pages, nil
}
