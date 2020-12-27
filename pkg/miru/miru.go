package miru

import "github.com/influx6/groundlayer/pkg/miru/parse"

// File represents a actually file (be it memory or local)
// which contains contents which can be refreshed
// to get the latest content.
type File interface {
	Refresh() error
	Read() (string, error)
	Parse(ops Options, ms map[string]string) (*ParsedData, error)
	ParseFor(ops Options, fm parse.FuncMaps, t Treeset) (*ParsedData, error)
	ParseBlock(blockName string, ops Options, ms map[string]string) (*ParsedData, error)
	ParseBlockFor(blockName string, ops Options, fm parse.FuncMaps, t Treeset) (*ParsedData, error)
}

// SourceFileSystem embodies a provider which allows retrieval of
// specific sources based on giving unique name.
type SourceFileSystem interface {
	GetFile(name string) (File, error)
}
