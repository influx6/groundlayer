package miru

import (
	"go/format"
	"io"
	"os"
	"path"
	"strings"

	"github.com/influx6/groundlayer/pkg/miru/parse"
	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/nunsafe"
)

var _ SourceFileSystem = (*VDir)(nil)
var _ File = (*VFile)(nil)

// VDir implements SourceFileSystem provided a
// caching system for files and directories.
type VDir struct {
	targetDir string
	cache     map[string]*VFile
}

// NewVDir returns a new VDir for the giving target path
func NewVDir(targetDir string) *VDir {
	var vd VDir
	vd.targetDir = targetDir
	vd.cache = map[string]*VFile{}
	return &vd
}

// LoadDir loads a giving directory files as cache VFiles for quicker
// access.
func (v *VDir) LoadDir() error {
	var fileHandle, fileErr = os.Open(v.targetDir)
	if fileErr != nil {
		return nerror.WrapOnly(fileErr)
	}
	var dirFiles, readDirErr = fileHandle.Readdir(-1)
	if readDirErr != nil {
		return nerror.WrapOnly(readDirErr)
	}

	v.cache = make(map[string]*VFile, len(dirFiles))
	for _, dirFile := range dirFiles {
		if dirFile.IsDir() {
			continue
		}
		v.cache[dirFile.Name()] = NewVFileWithFS(path.Join(v.targetDir, dirFile.Name()), v)
	}
	return nil
}

// GetFile returns a file within the directory which provides a
// VFile.
func (v *VDir) GetFile(subPath string) (File, error) {
	var targetFile = path.Join(v.targetDir, subPath)

	// check if file exists
	if _, err := os.Stat(targetFile); err != nil {
		return nil, nerror.WrapOnly(err)
	}

	var nf = NewVFileWithFS(targetFile, v)
	v.cache[targetFile] = nf
	return nf, nil
}

type VFile struct {
	targetFile string
	fs         SourceFileSystem
	builder    *strings.Builder
}

// NewVFile returns a new file.
func NewVFile(targetFile string) *VFile {
	return &VFile{
		targetFile: targetFile,
	}
}

// NewVFileWithFS returns a new VFile with associated fs.
func NewVFileWithFS(targetFile string, fs SourceFileSystem) *VFile {
	return &VFile{
		fs:         fs,
		targetFile: targetFile,
	}
}

// Parse parses the content of the File using the miru.TextWriter
// returning a strings.Builder containing the parsed content.
func (vt *VFile) Parse(ops Options, fns map[string]string) (*ParsedData, error) {
	if vt.builder == nil {
		if err := vt.Refresh(); err != nil {
			return nil, err
		}
	}
	return ParseTree(vt.targetFile, vt.builder.String(), ops, fns, vt.fs, nil)
}

// ParseFor parses the content of the File using the miru.TextWriter
// returning a strings.Builder containing the parsed content.
func (vt *VFile) ParseFor(ops Options, fns parse.FuncMaps, treeset Treeset) (*ParsedData, error) {
	if vt.builder == nil {
		if err := vt.Refresh(); err != nil {
			return nil, err
		}
	}
	return ParseTreeWithFuncMaps(vt.targetFile, vt.builder.String(), ops, fns, vt.fs, treeset)
}

// ParseBlock parses the content of the File using the miru.TextWriter
// returning a strings.Builder containing the parsed content.
func (vt *VFile) ParseBlock(blockName string, ops Options, fns map[string]string) (*ParsedData, error) {
	if vt.builder == nil {
		if err := vt.Refresh(); err != nil {
			return nil, err
		}
	}
	return ParseTree(blockName, vt.builder.String(), ops, fns, vt.fs, nil)
}

// ParseBlockFor parses the content of the File using the miru.TextWriter
// returning a strings.Builder containing the parsed content.
func (vt *VFile) ParseBlockFor(blockName string, ops Options, fns parse.FuncMaps, treeset Treeset) (*ParsedData, error) {
	if vt.builder == nil {
		if err := vt.Refresh(); err != nil {
			return nil, err
		}
	}
	return ParseTreeWithFuncMaps(blockName, vt.builder.String(), ops, fns, vt.fs, treeset)
}

// Read returns the string from it's cache (a strings.Builder)
// which uses unsafe to convert a byteslice into a string when
// calling strings.Builder.String, so ensure to not write to the
// returned string memory location as it will cause a memory corruption.
func (vt *VFile) Read() (string, error) {
	if vt.builder.Len() == 0 {
		if err := vt.Refresh(); err != nil {
			return vt.builder.String(), nerror.WrapOnly(err)
		}
	}
	return vt.builder.String(), nil
}

// Refresh reloads the internal cache of the content of the file.
func (vt *VFile) Refresh() error {
	if vt.builder == nil {
		var bs strings.Builder
		vt.builder = &bs
	}
	var file, fileErr = os.Open(vt.targetFile)
	if fileErr != nil {
		return nerror.WrapOnly(fileErr)
	}
	if _, copyErr := io.Copy(vt.builder, file); copyErr != nil && copyErr != io.EOF {
		return nerror.WrapOnly(copyErr)
	}
	return nil
}

type ParsedData struct {
	typeName          string
	root              *parse.RootNode
	functions         *strings.Builder
	types             *strings.Builder
	body              *strings.Builder
	options           Options
	komponentTypeName string
}

func newParsedData(typeName string, types, funcs, body *strings.Builder, root *parse.RootNode, op Options) *ParsedData {
	return &ParsedData{
		options:   op,
		types:     types,
		functions: funcs,
		body:      body,
		root:      root,
		typeName:  typeName,
	}
}

func (pd *ParsedData) Format(packageName string) (string, error) {
	if len(packageName) == 0 {
		return "", nerror.New("package name argument is required")
	}

	var combined strings.Builder

	combined.WriteString("\n")
	combined.WriteString("// Source is auto-generated by miru. DO NOT EDIT")
	combined.WriteString("\n")

	combined.WriteString("\n")
	combined.WriteString("package ")
	combined.WriteString(packageName)
	combined.WriteString("\n")

	combined.WriteString("\n")
	combined.WriteString("import (\n")
	combined.WriteString(`"github.com/influx6/groundlayer/pkg/domu"`)
	combined.WriteString("\n")

	for alias, pkgPath := range pd.options.Packages {
		combined.WriteString("\n")
		combined.WriteString(alias)
		combined.WriteString(` `)
		combined.WriteString(`"`)
		combined.WriteString(pkgPath)
		combined.WriteString(`"`)
	}

	combined.WriteString("\n")
	combined.WriteString(")")
	combined.WriteString("\n")

	if pd.types.Len() != 0 {
		combined.WriteString("\n")
		combined.WriteString("\n// Generated go-based types from template")
		combined.WriteString("\n")
		combined.WriteString("\n")
		combined.WriteString("\n")
		combined.WriteString(pd.types.String())
		combined.WriteString("\n")
	}

	if pd.functions.Len() != 0 {
		combined.WriteString("\n")
		combined.WriteString("\n// Generated go-based functions from template")
		combined.WriteString("\n")
		combined.WriteString("\n")
		combined.WriteString("\n")
		combined.WriteString(pd.functions.String())
		combined.WriteString("\n")
	}

	if pd.body.Len() != 0 {
		combined.WriteString("\n")
		combined.WriteString("\n// Generated go-based rendering of output template")
		combined.WriteString("\n")
		combined.WriteString("\n")
		combined.WriteString("\n")

		combined.WriteString("func RenderLayout(")

		combined.WriteString("page *peji.Page")
		combined.WriteString(", ")

		combined.WriteString("ctx ")
		combined.WriteString(`peji.Data`)
		combined.WriteString(", ")

		combined.WriteString("rootDoc *domu.Node")
		combined.WriteString(")")
		combined.WriteString("{")
		combined.WriteString("\n")
		combined.WriteString(pd.body.String())
		combined.WriteString("\n")
		combined.WriteString("}")
		combined.WriteString("\n")
	}

	var parsedContent = combined.String()
	var formatted, err = format.Source(nunsafe.String2Bytes(parsedContent))
	if err != nil {
		return parsedContent, nerror.Wrap(err, "Content: %s", parsedContent)
	}
	return nunsafe.Bytes2String(formatted), nil
}
