package miru_test

import (
	"testing"

	"github.com/influx6/groundlayer/pkg/miru"
	"github.com/stretchr/testify/require"
)

func TestSourceFileSystem(t *testing.T) {
	var dir = miru.NewVDir("./testdata")
	require.NotNil(t, dir)
	require.NoError(t, dir.LoadDir())

	var file1, file1Err = dir.GetFile("index.html")
	require.NoError(t, file1Err)
	require.NotNil(t, file1)
}
