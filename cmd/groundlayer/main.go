package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/njobs"
	"github.com/influx6/npkg/nunsafe"
	"github.com/urfave/cli/v2"

	"github.com/influx6/groundlayer/templates"
)

func main() {
	var app = cli.App{
		Name:     "groundlayer",
		HelpName: "Void CLI",
		Usage:    "groundlayer provides the CLI tool for the groundlayer library",
		Commands: []*cli.Command{
			{
				Name:        "new",
				Description: "create a new groundlayer project",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "dir",
						Usage: "--dir=./apps; sets directory to be used to store project",
					},
					&cli.StringFlag{
						Name:     "package",
						Required: true,
						Usage:    "--package=github.com/bigproject/apps/newsletter, set package name for project",
					},
				},
				Action: func(ctx *cli.Context) error {
					var currentDir, err = os.Getwd()
					if err != nil {
						return err
					}

					var packagePath = ctx.String("package")
					if len(packagePath) == 0 {
						return nerror.New("name is not provided")
					}

					var packageName = path.Base(packagePath)
					log.Println("Creating project for: ", packagePath)
					log.Println("Using package name: ", packageName)

					var stringSwitchMap = [][]string{
						{"github.com/influx6/groundlayer/templates/project", packagePath},
						{"project", packageName},
						{"[project]", packageName},
						{"[PROJECT]", strings.ToUpper(packageName)},
						{"PROJECT_", strings.ToUpper(packageName) + "_"},
					}

					var nameSwitchMap = map[string]string{
						"go.mod.template": "go.mod",
					}

					var rootDirectory = "templates/project"
					var directoryMapping = templates.GzipAssetNames()

					_ = currentDir
					var targetDir = ctx.String("dir")
					var projectRoot = path.Join(targetDir, packageName)

					var jobs njobs.Jobs

					var buildJob = func(nextFile string) error {
						var assetData, assetErr = templates.GzipAsset(nextFile)
						if assetErr != nil {
							return nerror.WrapOnly(assetErr)
						}

						var nextFilePart = strings.Replace(nextFile, rootDirectory, "", 1)
						var targetFile = path.Join(projectRoot, nextFilePart)

						for fromName, toName := range nameSwitchMap {
							targetFile = strings.ReplaceAll(targetFile, fromName, toName)
						}

						var reader, readerErr = gzip.NewReader(bytes.NewReader(assetData))
						if readerErr != nil {
							return nerror.WrapOnly(readerErr)
						}

						var contentBytes, contentReadErr = ioutil.ReadAll(reader)
						if contentReadErr != nil {
							return nerror.WrapOnly(contentReadErr)
						}

						var contentString = nunsafe.Bytes2String(contentBytes)
						for _, opList := range stringSwitchMap {
							contentString = strings.ReplaceAll(contentString, opList[0], opList[1])
						}

						var partJob njobs.Jobs
						partJob.Add(njobs.Mkdir(path.Dir(targetFile), 0777))
						partJob.Add(njobs.Println("Created: %s", os.Stdout))
						partJob.Add(njobs.File(path.Base(targetFile), 0777, strings.NewReader(contentString)))
						partJob.Add(njobs.Println("Created: %s", os.Stdout))
						partJob.Add(njobs.JobFunction(func(v interface{}) (interface{}, error) {
							return v, reader.Close()
						}))

						jobs.Add(njobs.MoveLastForward(partJob.Do))
						return nil
					}

					var buildErr error
					for _, nextFile := range directoryMapping {
						if buildErr = buildJob(nextFile); buildErr != nil {
							break
						}
					}

					if buildErr != nil {
						return nerror.WrapOnly(buildErr)
					}

					var goModFunction = njobs.ExecuteCommand(ctx.Context, projectRoot, "go mod tidy", nil, nil, map[string]string{})
					jobs.Add(njobs.JobFunction(func(i interface{}) (interface{}, error) {
						log.Println("Initializing go modules using: go mod tidy ...")
						var res, err = goModFunction.Do(i)
						if err != nil {
							log.Println("go mod tidy failed: You will need to manually do this yourself it seems! :)")
						}
						return res, nil
					}))
					_, err = jobs.Do(currentDir)
					return err
				},
			},
			{
				Name:        "component",
				Description: "create a .render.html file for existing package or directory",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "dir",
					},
					&cli.StringFlag{
						Name: "name",
					},
				},
				Action: func(ctx *cli.Context) error {
					var currentDir, err = os.Getwd()
					if err != nil {
						return err
					}

					var componentName = ctx.String("name")
					if len(componentName) == 0 {
						return nerror.New("name is not provided")
					}

					var targetDir = ctx.String("dir")
					var componentFile = fmt.Sprintf("%s.render.html", componentName)

					var viewsJob njobs.Jobs

					if len(targetDir) != 0 {
						viewsJob.Add(njobs.Mkdir(targetDir, 0777))
						viewsJob.Add(njobs.Println("Created: %s", os.Stdout))
					}

					var formattedContent = fmt.Sprintf(componentContent, componentName)
					viewsJob.Add(njobs.File(componentFile, 0777, bytes.NewBufferString(formattedContent)))
					viewsJob.Add(njobs.Println("Created: %s", os.Stdout))

					_, err = viewsJob.Do(currentDir)
					return err
				},
			},
			{
				Name:        "view",
				Description: "creates a groundlayer dom package generated from a set of go/html templates",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "dir",
					},
					&cli.StringFlag{
						Name: "name",
					},
				},
				Action: func(ctx *cli.Context) error {
					var currentDir, err = os.Getwd()
					if err != nil {
						return err
					}

					var viewName = ctx.String("name")
					if len(viewName) == 0 {
						return nerror.New("name is not provided")
					}

					var rootFile = "main.html"
					var viewDocFile = "doc.go"
					var viewGoGenFile = "main.gen.go"

					var targetDir = ctx.String("dir")
					var targetViewDir = path.Join(targetDir, viewName)
					var viewsDir = path.Join(targetDir, viewName)
					var viewGoFile = fmt.Sprintf("%s.go", viewName)

					var generatedCode = fmt.Sprintf(
						generatorContent,
						"./",
						rootFile,
						viewGoFile,
						viewName,
						viewGoFile,
						viewGoFile,
						viewGoFile,
					)
					var data, dataErr = format.Source(nunsafe.String2Bytes(generatedCode))

					if dataErr != nil {
						return dataErr
					}

					var viewsJob njobs.Jobs

					viewsJob.Add(njobs.Mkdir(viewsDir, 0777))
					viewsJob.Add(njobs.Println("Created: %s", os.Stdout))
					viewsJob.Add(njobs.File("blocks.html", 0777, blockContent))
					viewsJob.Add(njobs.Println("Created: %s", os.Stdout))
					viewsJob.Add(njobs.BackupPath())
					viewsJob.Add(njobs.File("main.html", 0777, indexContent))
					viewsJob.Add(njobs.Println("Created: %s", os.Stdout))

					var goJob njobs.Jobs
					goJob.Add(njobs.Mkdir(targetViewDir, 0777))
					goJob.Add(njobs.Println("Created: %s", os.Stdout))
					goJob.Add(njobs.File(viewGoGenFile, 0777, bytes.NewBuffer(data)))
					goJob.Add(njobs.Println("Created: %s", os.Stdout))
					goJob.Add(njobs.BackupPath())
					goJob.Add(njobs.File(viewDocFile, 0777, bytes.NewBufferString(fmt.Sprintf(docContent, viewName))))
					goJob.Add(njobs.Println("Created: %s", os.Stdout))

					var jobs njobs.Jobs
					jobs.Add(njobs.MoveLastForward(viewsJob.Do))
					jobs.Add(njobs.MoveLastForward(goJob.Do))
					_, err = jobs.Do(currentDir)
					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var componentContent = `<!-- component: %s -->

`

var indexContent = bytes.NewBufferString(`

{{/* you can reference a specific 'define' block from an external file using the # prefix*/}}
{{ template "blocks.html#header" .Path }}
`)

var blockContent = bytes.NewBufferString(`{{/* blocks: This file can only contain define blocks */}}

{{ define "header" string }}
	<h1>{{ . }}</h1>
{{ end }}
`)

var docContent = `package %s

//go:generate go run main.gen.go

`

var generatorContent = `
//+build ignore

package main

import (
	"os"
	"log"
	"github.com/influx6/groundlayer/pkg/miru"
	"github.com/influx6/npkg/nunsafe"
)

func main() {
	var dir = miru.NewVDir(%q)
	var indexFile, err = dir.GetFile(%q)
	if err != nil {
		log.Fatal("Failed read file '%s':", err)
	}

	var parsedData, parseErr = indexFile.Parse(miru.DefaultOption, nil)
	if parseErr != nil {
		log.Fatal("Failed to parse template files: ", parseErr)
	}

	var parsedString, parsedStrErr = parsedData.Format(%q)
	if parsedStrErr != nil {
		log.Fatal("Failed to format code properly: ", parsedStrErr)
	}

	var newFile, newFileErr = os.Create(%q)
	if newFileErr != nil {
		log.Fatal("Failed to create go file('%s'): ", newFileErr)
	}

	defer newFile.Close()

	var _, writeErr = newFile.Write(nunsafe.String2Bytes(parsedString))
	if writeErr != nil {
		log.Fatal("Failed to write to go file('%s'): ", writeErr)
	}
}
`
