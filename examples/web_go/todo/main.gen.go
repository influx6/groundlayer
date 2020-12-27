//+build ignore

package main

import (
	"log"
	"os"

	"github.com/influx6/groundlayer/pkg/miru"
	"github.com/influx6/npkg/nunsafe"
)

func main() {
	var dir = miru.NewVDir("./")
	var indexFile, err = dir.GetFile("main.html")
	if err != nil {
		log.Fatal("Failed read file 'todo':", err)
	}

	var parsedData, parseErr = indexFile.Parse(miru.DefaultOption, nil)
	if parseErr != nil {
		log.Fatal("Failed to parse template files: ", parseErr)
	}

	var parsedString, parsedStrErr = parsedData.Format("todo.go")
	if parsedStrErr != nil {
		log.Fatal("Failed to format code properly: ", parsedStrErr)
	}

	var newFile, newFileErr = os.Create("todo.go")
	if newFileErr != nil {
		log.Fatal("Failed to create go file('todo.go'): ", newFileErr)
	}

	defer newFile.Close()

	var _, writeErr = newFile.Write(nunsafe.String2Bytes(parsedString))
	if writeErr != nil {
		log.Fatal("Failed to write to go file('todo.go'): ", writeErr)
	}
}
