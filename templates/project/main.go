package main

//go:generate go generate ./components/...
//go:generate go generate ./pages/...
//go:generate go generate ./service/...

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/influx6/npkg/ndaemon"
	"github.com/influx6/npkg/nenv"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/sabuhp"
	"github.com/influx6/sabuhp/actions"
	"github.com/influx6/sabuhp/ochestrator"

	"github.com/influx6/groundlayer/pkg/peji"
	"github.com/influx6/groundlayer/templates/project/pages"
	"github.com/influx6/groundlayer/templates/project/service"
)

func runApp(envFileReader io.Reader) error {
	var envConfig = nenv.New(service.ProjectNamespace)

	var environmentLoader nenv.EnvironmentLoader
	var loadErr = environmentLoader.Register(envConfig)
	if loadErr != nil {
		return nerror.WrapOnly(loadErr)
	}

	if envFileReader != nil {
		var envReaderLoader = &nenv.EnvReaderLoader{File: envFileReader}
		var readEnvFileErr = envReaderLoader.Register(envConfig)
		if readEnvFileErr != nil {
			return nerror.WrapOnly(readEnvFileErr)
		}
	}

	var serviceId, hasServiceId = envConfig.GetString("SERVICE_ID")
	if !hasServiceId {
		return nerror.New("serviceId for app is required, please use %q", envConfig.KeyFor("REDIS_ADDR"))
	}

	var serviceAddr, hasServiceAddr = envConfig.GetString("ADDR")
	if !hasServiceAddr {
		return nerror.New("serviceId for app is required, please use %q", envConfig.KeyFor("ADDR"))
	}

	var pageMaxIdlenessDuration, hasPageMaxIdlenessDuration = envConfig.GetDuration("PAGE_MAX_IDLENESS")
	if !hasPageMaxIdlenessDuration {
		pageMaxIdlenessDuration = peji.DefaultMaxPageIdleness
	}

	var pageIdleCheckDuration, hasPageIdleCheckDuration = envConfig.GetDuration("PAGE_CHECK_IDLENESS")
	if !hasPageIdleCheckDuration {
		pageIdleCheckDuration = peji.DefaultPageIdlenessChecksInterval
	}

	var actionRegistry = actions.NewWorkerTemplateRegistry()
	service.AddActions(actionRegistry)

	var cliCtx, cliCtxCanceler = context.WithCancel(context.Background())
	ndaemon.WaiterForKillWithSignal(ndaemon.WaitForKillChan(), cliCtxCanceler)

	var station = ochestrator.DefaultStation(
		cliCtx,
		serviceId,
		serviceAddr,
		service.Logger,
		actionRegistry,
	)

	station.CreateEnvConfig = func(_ context.Context, _ sabuhp.Logger) (*nenv.EnvStore, error) {
		return envConfig, nil
	}

	if beforeInitErr := service.BeforeInit(station, envConfig); beforeInitErr != nil {
		return nerror.WrapOnly(beforeInitErr)
	}

	if stationInitErr := station.Init(); stationInitErr != nil {
		var wrapErr = nerror.WrapOnly(stationInitErr)
		return nerror.New("Closing application due to station initialization:\n %+s", wrapErr)
	}

	if afterInitErr := service.AfterInit(station, envConfig); afterInitErr != nil {
		return nerror.WrapOnly(afterInitErr)
	}

	var appPages, addPagesErr = pages.CreatePages(
		cliCtx,
		service.Logger,
		station,
		envConfig,
		pageMaxIdlenessDuration,
		pageIdleCheckDuration,
	)

	if addPagesErr != nil {
		cliCtxCanceler()
		return nerror.WrapOnly(addPagesErr)
	}

	if err := station.Wait(); err != nil {
		var wrapErr = nerror.WrapOnly(err)
		return nerror.New("Closing application: %s", wrapErr)
	}

	// ensure all page managers have closed
	appPages.Wait()
	return nil
}

func main() {
	var app = cli.App{
		Name:     "void",
		HelpName: "Void CLI",
		Usage:    "void provides the CLI tool for the void library",
		Commands: []*cli.Command{
			{
				Name:        "start",
				Description: "start service server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "env-file",
						Usage: "Sets path to an optional dotenv file to load environment values from",
					},
				},
				Action: func(ctx *cli.Context) error {
					var envFileReader io.Reader
					var envFile = strings.TrimSpace(ctx.String("env-file"))
					if len(envFile) != 0 {
						var content, fileErr = ioutil.ReadFile(envFile)
						if fileErr != nil {
							return nerror.WrapOnly(fileErr)
						}

						fmt.Printf("Content: %q\n", content)
						envFileReader = bytes.NewBuffer(content)
					}

					return runApp(envFileReader)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
