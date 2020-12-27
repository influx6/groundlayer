// Code generated by bindata. DO NOT EDIT.
// sources:
// templates/project/.env.dev
// templates/project/.env.prod
// templates/project/.env.qa
// templates/project/Makefile
// templates/project/components/dropdown/blocks.html
// templates/project/components/dropdown/index.go
// templates/project/doc.go
// templates/project/docker-compose.yml
// templates/project/go.mod.template
// templates/project/main.go
// templates/project/pages/hello/index.go
// templates/project/pages/home/about/index.go
// templates/project/pages/index.go
// templates/project/pages/layout/index.go
// templates/project/pages/not_found/index.go
// templates/project/pages/themes/dark.go
// templates/project/pages/themes/light.go
// templates/project/service/hello.go
// templates/project/service/index.go

package templates

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type gzipAsset struct {
	bytes []byte
	info  gzipFileInfoEx
}

type gzipFileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type gzipBindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi gzipBindataFileInfo) Name() string {
	return fi.name
}
func (fi gzipBindataFileInfo) Size() int64 {
	return fi.size
}
func (fi gzipBindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi gzipBindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi gzipBindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi gzipBindataFileInfo) IsDir() bool {
	return false
}
func (fi gzipBindataFileInfo) Sys() interface{} {
	return nil
}

var _gzipBindataTemplatesProjectEnvdev = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x08\xf2\xf7\x72\x75\x0e\x89\x0f\x72\x75\xf1\x0c\x8e\x77\x74\x71\x09" +
		"\xb2\x55\xb2\x32\x33\x36\xb7\x54\xe2\x82\x49\x05\xbb\x06\x85\x79\x3a\xbb\xc6\x7b\xba\xd8\x16\x14\xe5\x67\xa5\x26" +
		"\x97\xe8\x1a\x72\xc1\x65\xa1\x5a\xcc\x0d\x2c\x0d\x94\xb8\x00\x01\x00\x00\xff\xff\xec\xd2\x4e\xea\x4e\x00\x00\x00" +
		"")

func gzipBindataTemplatesProjectEnvdev() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectEnvdev
	info := gzipBindataFileInfo{
		name:        "templates/project/.env.dev",
		size:        78,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608994304, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectEnvprod = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x08\xf2\xf7\x72\x75\x0e\x89\x0f\x72\x75\xf1\x0c\x8e\x77\x74\x71\x09" +
		"\xb2\x55\xb2\x32\x33\x36\xb7\x54\xe2\x82\x49\x05\xbb\x06\x85\x79\x3a\xbb\xc6\x7b\xba\xd8\x16\x14\xe5\x67\xa5\x26" +
		"\x97\xe8\x1a\x72\xc1\x65\xa1\x5a\xcc\x0d\x2c\x0d\x94\xb8\x00\x01\x00\x00\xff\xff\xec\xd2\x4e\xea\x4e\x00\x00\x00" +
		"")

func gzipBindataTemplatesProjectEnvprod() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectEnvprod
	info := gzipBindataFileInfo{
		name:        "templates/project/.env.prod",
		size:        78,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608994310, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectEnvqa = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x08\xf2\xf7\x72\x75\x0e\x89\x0f\x72\x75\xf1\x0c\x8e\x77\x74\x71\x09" +
		"\xb2\x55\xb2\x32\x33\x36\xb7\x54\xe2\x82\x49\x05\xbb\x06\x85\x79\x3a\xbb\xc6\x7b\xba\xd8\x16\x14\xe5\x67\xa5\x26" +
		"\x97\xe8\x1a\x72\xc1\x65\xa1\x5a\xcc\x0d\x2c\x0d\x94\xb8\x00\x01\x00\x00\xff\xff\xec\xd2\x4e\xea\x4e\x00\x00\x00" +
		"")

func gzipBindataTemplatesProjectEnvqa() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectEnvqa
	info := gzipBindataFileInfo{
		name:        "templates/project/.env.qa",
		size:        78,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608994313, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectMakefile = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\x4d\xaa\xc3\x30\x0c\x84\xd7\x99\x53\xe8\x02\xd2\xdb\x07\xde\x59" +
		"\x82\xb1\x95\x34\x34\xb1\x8c\xff\x72\xfd\x62\x02\xdd\x84\xee\x86\xf9\xa4\x8f\x69\x69\x09\xe6\xdf\x9a\x67\x4c\x77" +
		"\x60\x6f\x67\xb2\xa2\xd4\x12\x71\x00\x82\x5d\xf1\x09\x47\x0b\x1c\xb6\x95\x27\x1b\x2d\x70\xb9\xea\x5f\xfc\xfb\x82" +
		"\x78\x05\x5a\x5a\x5c\x4a\x33\xa6\xcd\x28\xb7\x48\xa7\xdb\xa3\x6c\x46\xa5\xba\x5c\x89\x59\x63\xe7\x75\x3f\xf4\x5f" +
		"\x34\x76\x09\xda\xc7\xcb\x4c\xdf\xd5\x74\x0b\x80\xaa\xa5\xde\x9a\x91\x88\xb3\xf3\x4a\xdc\x49\xfe\x44\x04\x9f\x00" +
		"\x00\x00\xff\xff\xef\x66\xb7\xef\xe6\x00\x00\x00")

func gzipBindataTemplatesProjectMakefile() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectMakefile
	info := gzipBindataFileInfo{
		name:        "templates/project/Makefile",
		size:        230,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1609083233, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectComponentsDropdownBlockshtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\x51\xd4\xd5\x55\x48\xce\xcf\x2d\xc8\xcf\x4b\xcd\x2b\xb1\x52\xc8\x48" +
		"\x4d\x4c\x49\x2d\x52\xd0\xd5\xb5\xe3\x02\x04\x00\x00\xff\xff\x82\x0b\x00\xc4\x1b\x00\x00\x00")

func gzipBindataTemplatesProjectComponentsDropdownBlockshtml() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectComponentsDropdownBlockshtml
	info := gzipBindataFileInfo{
		name:        "templates/project/components/dropdown/blocks.html",
		size:        27,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608986138, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectComponentsDropdownIndexgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x29\xca\x2f\x48\xc9\x2f\xcf\xe3\x02" +
		"\x04\x00\x00\xff\xff\xc2\x97\x61\x97\x11\x00\x00\x00")

func gzipBindataTemplatesProjectComponentsDropdownIndexgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectComponentsDropdownIndexgo
	info := gzipBindataFileInfo{
		name:        "templates/project/components/dropdown/index.go",
		size:        17,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608986084, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectDocgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\x02\x04\x00\x00\xff" +
		"\xff\xec\x33\xc7\x17\x0d\x00\x00\x00")

func gzipBindataTemplatesProjectDocgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectDocgo
	info := gzipBindataFileInfo{
		name:        "templates/project/doc.go",
		size:        13,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608900301, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectDockercomposeyml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4b\x2d\x2a\xce\xcc\xcf\xb3\x52\x50\x32\xd6\xb3\x50\xe2\xe2\x2a\x4e" +
		"\x2d\x2a\xcb\x4c\x4e\x2d\xb6\xe2\x52\x50\x28\x4a\x4d\xc9\x04\x33\x14\x14\x32\x73\x13\xd3\x53\xad\xa0\x22\x39\x89" +
		"\x25\xa9\xc5\x25\x60\xf1\x82\xfc\xa2\x12\xa8\x12\x5d\x05\x33\x63\x73\x4b\x2b\x10\x01\x08\x00\x00\xff\xff\x65\xcf" +
		"\x35\x53\x55\x00\x00\x00")

func gzipBindataTemplatesProjectDockercomposeyml() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectDockercomposeyml
	info := gzipBindataFileInfo{
		name:        "templates/project/docker-compose.yml",
		size:        85,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608994571, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectGomodtemplate = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x51\x0e\x80\x20\x08\x00\xd0\x7f\x4f\xe1\x09\x64\x7e\xd4\x7d\x4c" +
		"\xc9\x6c\x20\x8e\x60\xab\xdb\xf7\x58\x9a\x13\xc6\x3e\xec\xf2\x23\x55\x61\x18\xf3\x24\x7f\x77\xe8\x2a\x3e\x1b\x95" +
		"\x0f\x15\x0c\x79\x51\x31\x7c\x60\xa9\xdc\x58\x2d\x84\x2e\x31\xa7\xbc\x85\x3f\x00\x00\xff\xff\x2e\xe8\x97\xce\x41" +
		"\x00\x00\x00")

func gzipBindataTemplatesProjectGomodtemplate() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectGomodtemplate
	info := gzipBindataFileInfo{
		name:        "templates/project/go.mod.template",
		size:        65,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1609057382, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectMaingo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x57\x4f\x6f\xdb\xb8\x12\x3f\x4b\x9f\x82\x15\x90\x42\xea\xf3\x93\x1f" +
		"\xde\x61\x0f\x2e\x7c\x70\x6d\xa7\x35\x9a\x76\x83\xb8\xdb\x2c\xb0\x5d\x04\x8c\x34\x92\xd9\x50\xa4\x4a\x52\x6e\xb2" +
		"\x45\xbe\xfb\x62\x48\x4a\x96\xfc\x27\x41\x17\x58\x1f\x2c\x69\x38\x7f\x7e\x33\x1c\xce\x0c\x6b\x9a\xdd\xd1\x12\x48" +
		"\x45\x99\x08\xc3\xf1\xb8\x94\x93\x12\x04\x28\x6a\x80\x94\x92\x74\xef\xe9\x38\x93\x55\x2d\x05\x08\xa3\xc7\x69\x9a" +
		"\x3e\xc9\x5a\xd3\x12\x9e\xe7\xd2\xa0\xb6\x2c\x03\xcb\x17\xb2\xaa\x96\xca\x90\x38\x0c\xa2\xdb\x07\x03\x3a\x0a\x83" +
		"\x28\x93\xc2\xc0\xbd\xc1\xd7\xa2\xb2\x0f\x26\xdd\xff\x98\xc9\xc6\x30\x8e\x1f\x5c\x96\xf8\x90\x56\x42\x1b\xc5\x44" +
		"\xa9\xa3\x30\x0c\xa2\x92\x99\x4d\x73\x9b\x66\xb2\x1a\x37\xaa\xa0\x5b\x18\x67\x9c\x8d\xb7\xff\xdf\x5f\x64\xa2\xe0" +
		"\xcd\xfd\x2f\x63\x51\xdf\x95\x63\x91\x53\xa8\xa4\x88\x9e\x62\x01\xb1\x7d\x7a\x5d\x29\xa9\x4e\x70\x68\x7a\xdb\x6c" +
		"\xea\x27\x17\xc7\x34\x33\x4c\x0a\xfd\x34\x93\xcc\x36\xa0\x8d\xa2\x06\x4d\x1d\xe7\x2c\x95\x6c\x44\xce\xe9\x03\xa8" +
		"\x31\x02\xab\xe1\x2b\x3b\xa1\xb4\xcf\x6a\xa0\xaa\x39\x35\xa0\xc7\xb5\x92\x5f\x21\x33\x6e\x3b\xff\x91\xa4\xdf\xe2" +
		"\x28\x4c\xc2\xb0\x68\x44\x46\x54\x23\x66\x75\x1d\x83\xd8\x9e\x33\x0e\x57\x40\x73\x50\x84\xc9\xd4\xbd\x25\xc4\xc6" +
		"\x8e\xfc\x08\x83\x2d\x55\x04\xc4\x76\x2e\x45\xc1\x4a\x32\x25\x18\xf5\xf4\x23\x7c\x8f\xbd\xca\xf4\xd2\x99\xf8\x48" +
		"\x2b\xd0\x35\xcd\x20\x09\x3b\x21\xa6\xa4\xa8\x40\x98\x0b\x69\xd5\x5b\xd1\xe5\x3e\xd9\x71\x73\x49\xf3\xa5\x52\x64" +
		"\x7a\x28\x97\x5e\x41\xc9\xb4\x01\x15\x77\x38\x92\x30\x60\x45\x27\xf3\x62\x4a\x04\xe3\x08\x36\x50\x60\x1a\x25\x88" +
		"\xdb\xfa\xf4\x5a\xd1\xfa\x57\xc1\x1f\x62\xcf\x99\x84\xc1\x63\x68\x45\x87\x7e\xf7\x14\x78\xe4\x6e\xc1\xe3\x9e\x92" +
		"\x97\x2d\xf4\x3e\xfd\x07\x6a\x98\x0c\x55\x3d\x7a\x15\x0a\x68\xbe\x74\x0b\x9d\x57\x7d\xd9\xe3\x3e\x21\xb2\x3d\xc9" +
		"\x1e\xb4\x53\xce\x0d\x25\x50\xcd\xa3\xf3\x13\x81\xf8\x5d\x5a\xe5\x23\xb2\xa1\x7a\xdd\x7e\x39\x44\xce\x70\xfa\x16" +
		"\xcc\xda\x9e\xd7\x38\x5a\x2f\xaf\x3e\xaf\xe6\xcb\x9b\xd5\x22\x72\x31\x7e\x31\x90\x3a\x0c\x31\x66\x42\xd4\x19\x21" +
		"\x85\x54\x84\xd6\x35\x61\x9a\x28\xf8\xd6\x30\x05\xf9\x88\xd4\x1c\xa8\x06\xd2\x68\x20\x67\xdf\xa2\x51\xcf\xf2\x7b" +
		"\x78\x38\x97\x2a\x8e\xae\x96\x8b\xd5\xfa\x66\xb6\x58\x5c\x45\x49\x72\x00\x7e\x96\xe7\xaa\x0f\x1f\xbf\x4f\x39\xe0" +
		"\x74\xec\x43\xb7\x12\xff\x16\xf8\x03\xd8\x78\x4c\x3f\xd0\xfb\x55\xce\x41\x80\xd6\x8b\x46\x51\x2c\x24\xd6\x85\xcb" +
		"\xe3\x6b\xfb\xee\xb4\xf4\x38\xba\x9c\xbd\x5d\xde\x7c\x98\xfd\x7e\xb3\x5a\x5c\x2c\x3f\x2e\xd7\xeb\x9e\x77\xa7\x94" +
		"\xa1\xa7\x27\x40\x90\x29\xc1\xf2\x93\x2e\xa0\xa0\x0d\x37\x1f\xe8\x3d\x2a\x69\xb9\x86\x3e\x20\x75\xbe\x81\xec\xee" +
		"\xc0\x83\x83\x95\x67\xf0\xcf\xdf\x2d\xe7\xef\x4f\x79\x70\xa8\xac\xc5\x7f\xcc\x4c\x1f\x7d\x1f\xba\xe5\xd3\x2b\x61" +
		"\x40\x6d\x29\xdf\x39\xe2\x8a\xb8\x3b\x6e\xea\x81\x4c\x3d\x41\xe3\xe6\x5f\x4b\x75\x07\xea\x93\xaf\x96\x2d\x4f\x9c" +
		"\x84\x41\x5b\xdc\x66\x79\x3e\x73\xfc\xf1\x50\x51\x5b\xe5\x32\xce\xe6\xe6\x7e\xe4\x9f\x73\x2a\x32\xe0\xb6\x64\xf8" +
		"\x9e\x99\x5e\x33\xb3\x71\xe4\xb8\x25\xbd\xa1\xd9\x9d\x2b\xd7\x31\xa6\x8d\x6f\x75\xe9\x35\x65\x06\xd4\xb9\x54\xef" +
		"\x19\xe7\x28\xb6\x66\xa5\xa0\x3c\xee\xaf\xfb\xd5\xf9\x86\x8a\x38\xd9\xb7\xda\x62\xd2\xa6\x0d\x56\xaf\x39\xb5\x31" +
		"\x5b\xbb\xc5\x38\x0c\x02\x8f\x3d\x0c\x82\x5d\x95\xd8\x7d\xd8\x53\xb7\xfb\x4c\x2f\x64\x59\x82\xa5\x0c\x23\x31\x0a" +
		"\x03\x34\xec\x8d\xa6\x73\x05\xd4\xc0\xb2\xd7\x31\xb0\xdf\xc4\x37\x5d\x40\xe6\xee\x39\x22\x37\xc4\xb5\x50\xaf\x39" +
		"\x21\xf1\xab\xb6\xcc\xae\x8d\x54\x30\x72\x4d\x28\xe9\x1f\xdc\x2e\xc9\x46\x58\x16\xbb\x72\x7e\x0b\x85\x54\xb0\x12" +
		"\xcc\x60\xcd\x9c\x4c\xdb\xca\x91\xbe\xe9\x16\x62\x0f\xb0\x77\x82\x93\xd7\x7b\x82\xcf\x37\x92\x01\xff\xae\x9d\x78" +
		"\xdd\x7d\x00\x3e\x1c\xd6\x74\xf2\x7a\x9f\x63\xaf\xe3\x7c\x57\xb4\x76\x7d\x62\xdf\xe0\x50\x2e\x39\x5e\xc2\xe6\x5c" +
		"\x6a\x26\x4a\x2c\x5e\x9c\x65\x6e\xf3\xf3\x06\x88\x91\x5d\x2e\x30\xc1\x0c\xa3\x9c\xfd\x65\x3f\x27\x5f\x04\x39\xfb" +
		"\x8f\x8e\x46\xad\xe5\x9d\x2b\xb4\x30\xa0\x8e\x44\x72\xd6\xd2\x8f\x07\x72\x20\xf6\x7c\x1c\xfb\xec\xbd\xca\x49\xeb" +
		"\x1a\xcf\xb4\x1e\x11\x9a\xe7\xf6\xcd\x45\xc5\x0e\x3e\x3e\xb5\x2c\xf9\x68\xfa\xf6\x52\xb4\xc5\x18\x06\xc1\x2e\x65" +
		"\x4e\x57\xc5\xd1\xa9\x82\xe3\x93\x1b\xe3\xd2\x03\xd4\xf3\x6f\x78\x02\xe3\xc3\x0d\xda\xb9\xbc\x53\xd0\x9b\x43\x86" +
		"\xd9\x82\x47\x1c\xb3\x05\x7e\x26\x45\xe0\x67\xf2\x62\x42\xce\x0e\xb7\x7d\x3c\x26\x20\x74\xa3\x80\x50\xce\x6d\xb0" +
		"\x49\x45\x05\x2d\x41\x69\xb2\xa1\x5b\x20\x19\x97\x1a\xf2\x30\x68\xf7\xc7\x03\x0d\x3b\xa3\x8c\x87\x8f\x7e\xb2\xc4" +
		"\x0b\x4c\x9c\xb4\xa3\x23\xf6\xd3\x29\xd6\xa9\x74\x56\xd7\xe8\x0b\xce\x89\x13\x82\xbf\x68\x2b\x59\x1e\x61\xe8\xdf" +
		"\x01\xaf\x1d\x3d\xfa\x2c\x59\x4e\xe6\x17\x2b\x4b\xff\x4d\xd3\xd2\x31\x5b\x5e\x52\x2b\xb9\x65\x39\x68\x62\x36\x80" +
		"\x4c\xc4\x48\xc9\x6d\xdb\x46\x82\x65\xe1\xec\x56\x51\xf5\x60\xc5\xe7\xb2\xaa\xa8\xc8\xf5\x84\xfc\xf1\xe7\x2b\x84" +
		"\xe0\x09\x76\x98\xb2\x7f\x3d\x34\x68\x44\x1b\xaa\x8c\x15\x0d\x82\x05\xe8\x4c\xb1\xda\xc5\xcc\xad\xb4\x67\xc1\x3e" +
		"\x41\x79\xc6\x73\x4e\x4b\x6b\x02\x2d\xe0\x87\xd3\x1c\xbc\xc4\x6f\x37\x96\xf4\xa8\xad\xc5\x08\xc4\xf6\xbf\x05\xe3" +
		"\xe0\xb5\x74\xce\x46\x6b\x30\x9a\xd4\xd4\x6c\xf0\xf8\x52\x41\xa4\xc5\x40\x39\xc9\xa5\x01\xb1\x25\x28\x84\x4b\x38" +
		"\xd3\xf6\x87\x65\xb2\xa5\xbc\x01\x4d\x0a\x25\xab\x56\xe9\xa3\x7b\xfa\x87\xeb\x64\x13\x57\x90\x33\x73\x4f\x7c\x50" +
		"\x6c\x3d\xee\x8d\xfc\xf8\xf3\x73\xf0\xb1\xcb\xc1\x01\x03\xc1\x04\xb6\xf7\xbd\xf4\x93\x62\xd5\x1a\x6f\x01\xa8\x3f" +
		"\x6d\x87\xb2\xce\x57\x6c\x77\xf8\xc3\xe1\x1d\x44\x7b\xfb\x48\x30\xd7\xff\xd7\x5a\x76\x6d\x15\x41\x09\x33\xb2\xde" +
		"\xba\xb4\x77\x37\x4d\x8b\x01\x85\x3a\x61\x2f\xc5\x8a\x8e\xb7\x3f\x33\x07\xa7\xe7\xe6\x62\x37\x30\xdb\x20\x85\xfe" +
		"\xa5\xa8\x4c\x7a\xa9\x98\x30\x45\x1c\xcd\x1d\x8e\x09\x39\xfb\xf6\x45\x44\xa3\x16\x57\x2b\x33\x8c\xd0\x94\xd8\x1b" +
		"\x33\x9e\xbd\x37\x4d\x51\x80\x8a\x87\xec\xad\x05\x8f\xe7\xd8\x15\x2c\xe9\xed\x97\xfd\xc7\xbf\x61\xad\xa0\x75\x9d" +
		"\x5e\x35\x22\x96\x3a\x9d\xa9\x52\x1f\x56\x0b\x2e\xcb\xf4\x9c\x1a\xca\x7d\x65\x78\x0c\x1f\xc3\xbf\x03\x00\x00\xff" +
		"\xff\xeb\xa8\x3e\x45\x5e\x10\x00\x00")

func gzipBindataTemplatesProjectMaingo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectMaingo
	info := gzipBindataFileInfo{
		name:        "templates/project/main.go",
		size:        4190,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1609057382, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectPagesHelloIndexgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\x41\x6b\xdc\x30\x10\x85\xcf\x3b\xbf\x62\x6a\x68\x91\x82\xb1\x6f" +
		"\xbd\xf5\x50\x52\x42\x0e\x25\x84\x34\x90\xf3\xc4\x1a\x5b\x6a\x6c\x49\x95\x47\x69\x96\xc5\xff\xbd\xc8\x4e\x9b\x6d" +
		"\xd3\xc0\x5e\x0c\x9e\x79\xcc\xfb\x9e\x9f\x23\x75\x0f\x34\x30\x5a\x1e\xc7\x00\xe0\xa6\x18\x92\xa0\x82\x5d\xd5\x4f" +
		"\x52\x01\xec\xaa\xc1\x89\xcd\xf7\x4d\x17\xa6\xd6\xf9\x7e\xcc\x4f\x1f\xdb\x99\xee\xb3\x8d\x6f\x6d\x87\x14\xb2\x37" +
		"\x23\xed\x39\xb5\xf1\x61\x68\x4d\x98\x72\x75\x9a\x34\xf2\x77\x57\x81\x06\x78\xa4\x84\x97\x05\xe9\x2e\xa4\xd1\x5c" +
		"\x05\xb9\x28\x42\xfc\x84\x45\xd1\x5c\x92\x37\x23\xa7\x8b\xec\x3b\xd5\x97\x87\x21\xa1\x6d\xf5\x85\x84\x34\x9e\x15" +
		"\xcf\xe6\x2a\x18\xc6\x03\xec\x12\x4b\x4e\x1e\xd7\xd9\x2d\x3f\x89\xea\x27\x69\xbe\xc5\xe4\xbc\xf4\xaa\x5a\x6d\x70" +
		"\xf5\x41\x13\x78\x46\x1f\x04\x2d\x3d\x32\x8a\x75\x33\xbe\xff\xf1\xae\xaa\xb1\x18\x34\xd7\x24\x56\x6b\x58\x34\x80" +
		"\xec\x23\x1f\x01\x7e\xa5\x7d\xc8\x82\xb3\xa4\xdc\xc9\x61\x01\x28\x54\xa8\xec\x2b\x89\xc6\x1b\xf6\x86\x93\x8a\x78" +
		"\xb6\xf2\x5e\xd3\xc0\xdb\xf9\x17\xfe\x1a\x23\x25\xf6\x72\x14\x43\x97\x1c\xdb\xb4\xf9\x1c\x23\x7b\x73\x6e\xdd\x68" +
		"\xd4\x09\x99\x5e\x05\xd0\xf0\x1b\xf0\x3c\x31\x09\xbf\x30\x16\x18\x15\x69\xe0\x9b\x90\x85\x4b\x1c\xe7\x87\x1a\x25" +
		"\xe1\xd6\x78\x73\x9b\xc8\xcf\xe5\x0f\xd1\x47\xf8\x05\xad\xf4\xe5\xf9\xe7\xfa\xfa\x5c\xd2\x9d\x13\xfb\xf7\xc1\x1a" +
		"\x3f\xfc\xfb\x3d\x0e\x4b\xfd\x9f\x9e\xf5\x9f\xce\x9e\x6f\xc2\x02\xbf\x02\x00\x00\xff\xff\x7c\x7f\x69\xbf\xa9\x02" +
		"\x00\x00")

func gzipBindataTemplatesProjectPagesHelloIndexgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectPagesHelloIndexgo
	info := gzipBindataFileInfo{
		name:        "templates/project/pages/hello/index.go",
		size:        681,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1609057381, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectPagesHomeAboutIndexgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\x4c\xca\x2f\x2d\xe1\x02\x04\x00\x00" +
		"\xff\xff\x6d\x92\xba\x2f\x0e\x00\x00\x00")

func gzipBindataTemplatesProjectPagesHomeAboutIndexgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectPagesHomeAboutIndexgo
	info := gzipBindataFileInfo{
		name:        "templates/project/pages/home/about/index.go",
		size:        14,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608900307, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectPagesIndexgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\xc1\x6e\xdc\x38\x0c\x3d\x4b\x5f\xc1\xd5\xc9\x1e\x08\xf6\x6d\x0f" +
		"\xbb\x98\x43\x30\x99\x36\x05\x8a\x24\x68\x02\xe4\x18\x28\x36\x6d\x2b\xd1\x48\x06\x4d\x4f\x27\x28\xf2\xef\x85\x24" +
		"\xcf\x24\x68\x9b\x22\x27\x53\xd4\x23\xf9\x9e\x1e\x3d\x9a\xe6\xc9\xf4\x08\xa3\xe9\x71\x92\xd2\xee\xc6\x40\x0c\x85" +
		"\x14\xaa\x09\x9e\xf1\xc0\x4a\x0a\xc5\x76\x87\x4a\x4a\xa1\x7a\xcb\xc3\xfc\x50\x35\x61\x57\x5b\xdf\xb9\xf9\xf0\x6f" +
		"\xed\xc7\xa7\xbe\xf6\xe8\xf7\xea\xaf\xf7\x44\x81\xde\x41\x4c\xe6\x61\x1e\xc6\x3a\x34\x03\x4e\x4c\x86\x23\xf0\xcf" +
		"\xc8\x9e\xc2\xec\x5b\x67\x9e\x91\xea\xd8\x76\xc4\x47\xfb\x11\x2c\xe3\x6e\x74\x86\x71\xaa\x47\x0a\x8f\xd8\x70\x9d" +
		"\xe4\xd6\x03\x3a\x17\xde\x61\xf5\x91\x7a\x1f\xf8\xbe\x8b\x30\x25\x4b\x29\x9b\xe0\xa7\xf4\x74\xd7\xa6\xc7\x6b\xc2" +
		"\xce\x1e\x60\x0d\x2a\x63\x55\x4e\x7f\x0b\x33\x23\xbc\xa6\xeb\xd5\x68\x78\x50\x50\xd7\x90\x22\x60\x74\x6e\x02\x1e" +
		"\x10\x28\x42\x09\x38\xc0\xce\x70\x33\x80\xf1\xcf\x3c\x58\xdf\x83\xe9\x52\x7e\x40\xc8\x3d\xe2\xec\x6e\xf6\x0d\x6c" +
		"\x08\x0d\x63\x1c\x33\x15\x52\x34\x7c\x80\xc5\xc3\x6a\x93\xbf\x5a\x0a\x17\xfa\x1e\x09\xe2\xcb\x55\x5f\x53\xac\xa5" +
		"\x98\xd8\xb0\x0d\x1e\x56\x6f\x3c\xa8\x6e\x72\x52\x4b\x81\x7e\xbf\x09\xbe\xb3\x3d\xac\xa2\xd1\xd5\xd6\xef\x6f\x38" +
		"\x10\x6a\x29\x76\xe6\xf0\xa5\x75\xe8\x71\x9a\x20\x6e\x49\x75\x3e\xd3\xb1\xcc\xb6\x0e\x37\x03\x36\x4f\xc7\xdc\xaf" +
		"\x88\x12\x8a\x55\x22\x92\x28\x6b\x48\x4b\x52\xc2\x0f\x29\xf6\x86\xf2\x46\xc2\x3a\x53\xbd\xc4\xef\x47\x5d\x51\x98" +
		"\x96\x62\x51\x12\xa3\xd7\xf7\x8e\xa7\x37\x94\xe2\xf1\x37\x16\x31\xb9\x08\xae\x6e\xc9\xf8\x29\x6e\x7c\x51\xc6\xf4" +
		"\xc9\xcf\xea\x32\xf0\xa7\x18\x68\x29\x4a\x99\xf9\x2c\x7e\xac\xe1\x58\x9c\xbc\xa4\xa2\x94\x22\x51\xad\xce\xda\xf6" +
		"\xca\x9f\x4c\x2e\xa2\x25\x45\x2a\x82\x89\xc9\xfa\x5e\xc3\x3d\xbc\xd1\x9b\x84\x8a\x7b\x58\x2f\x9d\xab\xed\x1e\x3d" +
		"\xe7\x0a\x9d\xd5\x97\x52\xbc\xc4\xf1\xb6\x03\xd3\xb6\x5b\x22\xf8\x6f\x0d\xa7\x61\x85\xca\xfb\xab\x21\x7d\xab\x6c" +
		"\xff\x45\x8c\xef\x02\xb9\x36\x4e\x29\xff\x3f\x56\xfe\xb3\x06\x6f\x5d\x9a\x49\xc8\x33\xf9\x78\xd4\x90\xff\xcc\xea" +
		"\x8e\xcc\x78\xe5\xdd\x73\x91\xd1\x71\xb0\x94\x62\x21\x76\xc1\x3c\xde\x20\xed\x6d\x83\xc5\x49\xdf\x42\x51\x83\xba" +
		"\xd8\x9e\x9d\x2b\x0d\xea\xf3\xf6\x56\x95\xf2\xd8\x7d\xb9\xf5\xd6\xc9\x17\xf9\x33\x00\x00\xff\xff\xfe\x5a\xfe\x9a" +
		"\x66\x04\x00\x00")

func gzipBindataTemplatesProjectPagesIndexgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectPagesIndexgo
	info := gzipBindataFileInfo{
		name:        "templates/project/pages/index.go",
		size:        1126,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1609087815, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectPagesLayoutIndexgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x49\xac\xcc\x2f\x2d\xe1\x02\x04\x00" +
		"\x00\xff\xff\x1c\x59\xc2\x86\x0f\x00\x00\x00")

func gzipBindataTemplatesProjectPagesLayoutIndexgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectPagesLayoutIndexgo
	info := gzipBindataFileInfo{
		name:        "templates/project/pages/layout/index.go",
		size:        15,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608901259, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectPagesNotfoundIndexgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\xbd\x4a\x44\x31\x10\x85\xeb\x3b\x4f\x31\x06\x84\x44\x24\xe9\xec" +
		"\xec\x64\xb1\x5a\x04\xed\x65\xbc\xf9\xd9\xb8\x37\x3f\x8c\x13\x59\x11\xdf\x5d\x72\x7d\x01\x9b\x29\xce\xf9\x98\xc3" +
		"\xd7\x69\x3d\x53\x0a\x58\x9b\xbc\xc6\x36\xaa\x07\xc8\xa5\x37\x16\xd4\xb0\xa8\x58\x44\x01\x2c\x2a\x65\x39\x8d\x37" +
		"\xbb\xb6\xe2\x72\x8d\xdb\xb8\xdc\xb9\xc4\x93\xde\xe8\x2b\xb0\xeb\xe7\xe4\x7c\x2b\x43\xfd\x0f\xed\xe1\x3d\x2b\x30" +
		"\x00\x9f\xc4\x78\x6c\x72\x98\x35\xde\xe3\xcc\xed\x23\x55\xbf\x05\x3e\x8c\xba\xea\x38\x8f\x27\xa1\xbf\xea\x81\x84" +
		"\x0c\xde\xcc\x25\x7b\x6c\x3e\xe0\x37\x2c\x1c\x64\x70\xc5\x3d\x7b\x09\x17\xd1\xb1\x88\x7d\xee\x9c\xab\x44\xad\xae" +
		"\x3f\xa6\x18\xee\x62\x57\xea\x16\xe7\x2f\xfb\x44\x72\x32\x06\x7e\x0c\xfc\x06\x00\x00\xff\xff\x0d\xdb\x4b\xfb\xfd" +
		"\x00\x00\x00")

func gzipBindataTemplatesProjectPagesNotfoundIndexgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectPagesNotfoundIndexgo
	info := gzipBindataFileInfo{
		name:        "templates/project/pages/not_found/index.go",
		size:        253,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1609057382, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectPagesThemesDarkgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\xc9\x48\xcd\x4d\x2d\x06\x04\x00\x00" +
		"\xff\xff\xc0\xe1\x99\x31\x0e\x00\x00\x00")

func gzipBindataTemplatesProjectPagesThemesDarkgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectPagesThemesDarkgo
	info := gzipBindataFileInfo{
		name:        "templates/project/pages/themes/dark.go",
		size:        14,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608901236, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectPagesThemesLightgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\xc9\x48\xcd\x4d\x2d\xe6\x02\x04\x00" +
		"\x00\xff\xff\xc2\x5d\x82\xa9\x0f\x00\x00\x00")

func gzipBindataTemplatesProjectPagesThemesLightgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectPagesThemesLightgo
	info := gzipBindataFileInfo{
		name:        "templates/project/pages/themes/light.go",
		size:        15,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608901244, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectServiceHellogo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x4f\x8b\xdb\x30\x10\xc5\xcf\x9a\x4f\x31\xe8\x64\x15\xaf\x42\x2f" +
		"\x3d\x04\xf6\x10\x02\xdb\x52\xd8\x50\xba\x81\x3d\x94\x52\x64\x65\xe2\xb8\x6b\x6b\x5c\xfd\x59\xa7\x2c\xf9\xee\x45" +
		"\x96\x9b\xb6\x7b\xe8\x49\xf0\x9e\xe6\xcd\x6f\xde\x68\xec\x93\x69\x09\x03\xf9\xe7\xce\x12\x40\x37\x8c\xec\x23\x56" +
		"\x20\xa4\x65\x17\xe9\x1c\x25\x80\x90\x6d\x17\x4f\xa9\xd1\x96\x87\x55\xe7\x8e\x7d\x3a\xbf\x5b\x05\xd3\xa4\xd3\x28" +
		"\xff\x67\xae\x8c\x8d\x1d\xbb\x20\x41\x01\x1c\x93\xb3\x78\xa2\xbe\xe7\x47\xf6\xfd\x61\x33\x5b\x95\x8d\x67\x5c\x16" +
		"\xe9\x6d\x79\x6b\xfc\xce\x0d\x2e\xa3\xfa\x23\x37\x0a\x5f\x40\x7c\xc3\xdb\xac\xeb\xbd\x37\x2e\x64\x46\xfd\x40\xee" +
		"\xb0\xe7\x4d\xdf\x57\x65\x9b\xde\xd1\x74\x4f\x21\x98\x96\xaa\xfc\xf3\x3e\xb4\xfa\xce\xf3\xb0\x39\x1c\x7c\x5d\x66" +
		"\xb9\xc6\x2f\x5f\x9b\x9f\x91\x2a\x39\x93\xe0\x94\x51\xa4\x52\x35\xde\xbc\x55\x70\x01\x78\x36\x1e\x3f\x5c\x21\x3f" +
		"\xd3\x8f\x44\x21\xe2\xed\x15\xe7\x91\xfd\x13\xf9\x45\x7f\x01\x51\xce\xd8\x99\x81\xd6\x88\xf2\xcf\x79\xb2\x06\xf1" +
		"\x29\x35\x0f\xa9\xd9\xf3\xd8\xd9\xf5\xe2\xdd\x4c\xbf\xcd\x12\xb4\xf5\x64\x22\xfb\x35\xe6\x76\x2a\xcb\xee\xd8\xb5" +
		"\xaf\x76\x6d\x67\x51\xe1\x9b\x7f\xe5\xf7\x9e\xd3\x98\x8b\x11\x65\x4a\x17\x92\xbf\x50\x8b\x70\x97\x83\x5f\xd7\xae" +
		"\x40\x08\x4f\x31\x79\x77\xfd\xbd\xa3\x29\x07\xcf\xb1\x0b\x88\x02\x71\xa9\xe1\x02\xbf\x02\x00\x00\xff\xff\xce\xae" +
		"\xec\x71\x25\x02\x00\x00")

func gzipBindataTemplatesProjectServiceHellogo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectServiceHellogo
	info := gzipBindataFileInfo{
		name:        "templates/project/service/hello.go",
		size:        549,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1608986761, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

var _gzipBindataTemplatesProjectServiceIndexgo = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\xcd\x6e\xe3\x36\x10\xc7\xcf\xe2\x53\xcc\x0a\xd8\x85\x64\xa8\xd4" +
		"\xad\x28\x52\xf8\xe0\xda\x6e\xeb\xb4\x75\x02\x3b\x6d\x0e\x6d\x51\x30\xd4\x58\xe6\x46\x26\x95\x21\xe5\xc4\x08\xf2" +
		"\xee\x05\xa9\x0f\xdb\x41\xeb\xed\xa1\x39\xc4\x22\x67\x38\x5f\xff\x1f\x59\x0b\xf9\x28\x4a\x04\x8b\xb4\x57\x12\x19" +
		"\x53\xbb\xda\x90\x83\x84\x45\xb1\x34\xda\xe1\x8b\x8b\x59\x14\x57\xa6\x8c\x19\x8b\xe2\x52\xb9\x6d\xf3\xc0\xa5\xd9" +
		"\xe5\xa5\xf9\x8a\xb0\x50\x36\x6f\xff\xef\xbf\x79\xef\xa0\xf4\xa6\x6a\x5e\xbe\xce\x75\xfd\x58\xe6\x1a\xf5\x3e\xbe" +
		"\x68\x27\x32\x74\xd1\xe3\xb3\x35\xfa\x5f\x1c\xac\x78\x68\xb6\xf5\x45\x63\x2e\xa4\x53\x46\xdb\xcb\x4e\xd2\x14\x28" +
		"\xbf\xe0\xa3\xf4\x67\x94\xce\xd0\x17\xdc\x8c\xdc\xa2\x75\x24\x9c\x6f\x2b\x65\x4c\x1a\x6d\xc3\x5c\x6f\xc9\xf8\x00" +
		"\x4b\xb1\x43\x08\x7f\x63\x88\x7f\xaf\xdb\xcd\x3f\xe3\x33\xbb\xad\x85\xc4\x60\xbf\x5d\xdd\x5c\xcf\xa7\x77\xde\x3e" +
		"\xd7\xfb\xa9\xd1\x1b\x55\x4e\x2a\x25\x6c\x7b\x7e\xe4\x47\xc4\xe7\x7a\xbf\x76\x86\x30\xe4\xdb\x0b\x82\x9f\x4d\x59" +
		"\x22\xc1\x18\x3e\x55\xe1\x6b\xb1\xab\xab\xd7\x37\xc6\xdc\xa1\x46\x38\x6e\x81\x75\xd4\x48\xe7\x2d\x9b\x46\x4b\x48" +
		"\xaa\x13\x63\xea\xa3\x24\xf2\x01\x46\x41\x01\x7e\xbd\xbe\x59\xa6\xf0\xca\xa2\xca\x94\xfc\x96\x94\x76\x9b\x24\xfe" +
		"\x68\xff\xd0\x71\x06\xf2\x81\xff\x82\xd6\x8a\x12\x93\x34\x65\x7d\xb8\xef\x70\x63\x08\x17\x5a\xb9\xc4\x3a\xe1\x55" +
		"\x80\xd1\xc9\x74\xf8\xba\xdd\xcc\x00\xfb\xce\x60\xe4\x79\x19\xfa\x49\x21\xd0\xe1\x93\x76\x01\xf8\x94\x50\x38\x5c" +
		"\x74\x52\xc0\x18\x7e\xb5\xc3\xea\xbd\xd7\xd4\x8b\xda\xba\x84\x4f\xc6\x22\x3f\x9c\xc0\xed\xa4\x28\x28\x83\xad\xb0" +
		"\xab\x7e\x05\xe3\x63\x1d\xfc\x07\x74\x6b\x47\x4a\x97\x49\xbc\x9a\xcf\x16\xeb\xbf\x26\xb3\xd9\x2a\x4e\x59\xa4\x36" +
		"\xf0\xe1\xec\xd4\x2b\x8b\x22\x42\xd7\x90\x86\x96\x65\xbe\xc4\xe7\x24\x16\x45\x41\x68\x2d\x6c\x4c\x97\x10\x94\x05" +
		"\xc2\xa7\x46\x11\x16\xf0\xf1\x29\x3e\xe9\x9a\xff\x84\x87\xef\x0d\x9d\xa7\x4a\x59\xf4\xc6\xde\x77\x74\x47\x42\xdb" +
		"\x70\x4f\xc7\x70\x3a\xc9\x50\xce\x60\xbc\x57\x6e\x7b\x53\x07\xea\x93\x90\x9b\x77\x2b\x5f\xeb\x12\xdd\xb3\xa1\xc7" +
		"\x2b\x88\x9d\xac\xe3\x8c\x45\x91\x6f\xe3\xca\xf3\x78\x1c\x0c\x8b\xde\x52\xc6\x86\xbe\x54\x35\x68\x3a\xd9\x38\xa4" +
		"\xff\x45\xd2\x7f\x0c\xdf\x4b\x95\x48\xf7\x02\xdd\x3b\xc4\xa7\xed\x6f\xd6\xc1\x09\xed\x45\xe3\x2d\xe4\x29\x24\xdd" +
		"\x3a\x1c\xcc\xda\xf8\x01\xd4\x2e\xfe\xa7\xf6\x72\xf3\x6b\x6b\x74\xf0\x79\x7d\xcb\xde\x27\xed\x11\xfa\xef\x79\x2f" +
		"\xb5\x98\x8c\x86\xb7\x82\xf7\x91\x4f\x0b\xf3\x14\xaa\x23\xc2\x47\xe7\x25\x3e\x0f\x95\x78\x01\xf2\xbc\x33\xfa\x64" +
		"\xa0\xb4\x33\x30\x5b\x04\x08\xdb\xed\x39\x11\x5c\x1d\x03\xf0\x49\x51\xfc\x26\xaa\x06\x93\xa1\xb6\x0c\xce\x9f\x8d" +
		"\xf4\xdb\x93\xa3\x1f\xc6\x7e\x0e\x67\x08\xab\x2a\xeb\x39\xbe\x27\x51\xdf\xe8\xea\x90\x0c\x07\x3a\x26\xf3\x1c\x26" +
		"\x45\x01\x3b\x43\x08\xb3\x05\x28\x87\xbb\x16\xf4\x02\x6b\xd4\x05\x6a\x79\xe8\x92\x78\xe8\x8e\x42\xab\x61\x14\x67" +
		"\x44\x15\xc5\x44\xf6\xb0\x96\xca\x3a\x3a\xc0\xa8\x7b\xb4\xf9\xbd\xa1\x47\xa4\x3b\xdc\xd5\x95\x70\xb8\xea\xec\x9d" +
		"\xba\xed\x82\xb7\xbb\x48\xc9\x8f\x58\x55\xe6\xde\x50\x55\xac\xf0\xa9\x41\xeb\xfc\x53\xf4\x77\x00\x00\x00\xff\xff" +
		"\xe7\x92\xc0\x6b\xe9\x06\x00\x00")

func gzipBindataTemplatesProjectServiceIndexgo() (*gzipAsset, error) {
	bytes := _gzipBindataTemplatesProjectServiceIndexgo
	info := gzipBindataFileInfo{
		name:        "templates/project/service/index.go",
		size:        1769,
		md5checksum: "",
		mode:        os.FileMode(436),
		modTime:     time.Unix(1609057641, 0),
	}

	a := &gzipAsset{bytes: bytes, info: info}

	return a, nil
}

// GzipAsset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func GzipAsset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _gzipbindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("GzipAsset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// MustGzipAsset is like GzipAsset but panics when GzipAsset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
func MustGzipAsset(name string) []byte {
	a, err := GzipAsset(name)
	if err != nil {
		panic("asset: GzipAsset(" + name + "): " + err.Error())
	}

	return a
}

// GzipAssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
func GzipAssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _gzipbindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("GzipAssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

// GzipAssetNames returns the names of the assets.
// nolint: deadcode
func GzipAssetNames() []string {
	names := make([]string, 0, len(_gzipbindata))
	for name := range _gzipbindata {
		names = append(names, name)
	}
	return names
}

//
// _gzipbindata is a table, holding each asset generator, mapped to its name.
//
var _gzipbindata = map[string]func() (*gzipAsset, error){
	"templates/project/.env.dev":                        gzipBindataTemplatesProjectEnvdev,
	"templates/project/.env.prod":                       gzipBindataTemplatesProjectEnvprod,
	"templates/project/.env.qa":                         gzipBindataTemplatesProjectEnvqa,
	"templates/project/Makefile":                        gzipBindataTemplatesProjectMakefile,
	"templates/project/components/dropdown/blocks.html": gzipBindataTemplatesProjectComponentsDropdownBlockshtml,
	"templates/project/components/dropdown/index.go":    gzipBindataTemplatesProjectComponentsDropdownIndexgo,
	"templates/project/doc.go":                          gzipBindataTemplatesProjectDocgo,
	"templates/project/docker-compose.yml":              gzipBindataTemplatesProjectDockercomposeyml,
	"templates/project/go.mod.template":                 gzipBindataTemplatesProjectGomodtemplate,
	"templates/project/main.go":                         gzipBindataTemplatesProjectMaingo,
	"templates/project/pages/hello/index.go":            gzipBindataTemplatesProjectPagesHelloIndexgo,
	"templates/project/pages/home/about/index.go":       gzipBindataTemplatesProjectPagesHomeAboutIndexgo,
	"templates/project/pages/index.go":                  gzipBindataTemplatesProjectPagesIndexgo,
	"templates/project/pages/layout/index.go":           gzipBindataTemplatesProjectPagesLayoutIndexgo,
	"templates/project/pages/not_found/index.go":        gzipBindataTemplatesProjectPagesNotfoundIndexgo,
	"templates/project/pages/themes/dark.go":            gzipBindataTemplatesProjectPagesThemesDarkgo,
	"templates/project/pages/themes/light.go":           gzipBindataTemplatesProjectPagesThemesLightgo,
	"templates/project/service/hello.go":                gzipBindataTemplatesProjectServiceHellogo,
	"templates/project/service/index.go":                gzipBindataTemplatesProjectServiceIndexgo,
}

// GzipAssetDir returns the file names below a certain
// directory embedded in the file by bindata.
// For example if you run bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then GzipAssetDir("data") would return []string{"foo.txt", "img"}
// GzipAssetDir("data/img") would return []string{"a.png", "b.png"}
// GzipAssetDir("foo.txt") and GzipAssetDir("notexist") would return an error
// GzipAssetDir("") will return []string{"data"}.
func GzipAssetDir(name string) ([]string, error) {
	node := _gzipbintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type gzipBintree struct {
	Func     func() (*gzipAsset, error)
	Children map[string]*gzipBintree
}

var _gzipbintree = &gzipBintree{Func: nil, Children: map[string]*gzipBintree{
	"templates": {Func: nil, Children: map[string]*gzipBintree{
		"project": {Func: nil, Children: map[string]*gzipBintree{
			".env.dev":  {Func: gzipBindataTemplatesProjectEnvdev, Children: map[string]*gzipBintree{}},
			".env.prod": {Func: gzipBindataTemplatesProjectEnvprod, Children: map[string]*gzipBintree{}},
			".env.qa":   {Func: gzipBindataTemplatesProjectEnvqa, Children: map[string]*gzipBintree{}},
			"Makefile":  {Func: gzipBindataTemplatesProjectMakefile, Children: map[string]*gzipBintree{}},
			"components": {Func: nil, Children: map[string]*gzipBintree{
				"dropdown": {Func: nil, Children: map[string]*gzipBintree{
					"blocks.html": {Func: gzipBindataTemplatesProjectComponentsDropdownBlockshtml, Children: map[string]*gzipBintree{}},
					"index.go":    {Func: gzipBindataTemplatesProjectComponentsDropdownIndexgo, Children: map[string]*gzipBintree{}},
				}},
			}},
			"doc.go":             {Func: gzipBindataTemplatesProjectDocgo, Children: map[string]*gzipBintree{}},
			"docker-compose.yml": {Func: gzipBindataTemplatesProjectDockercomposeyml, Children: map[string]*gzipBintree{}},
			"go.mod.template":    {Func: gzipBindataTemplatesProjectGomodtemplate, Children: map[string]*gzipBintree{}},
			"main.go":            {Func: gzipBindataTemplatesProjectMaingo, Children: map[string]*gzipBintree{}},
			"pages": {Func: nil, Children: map[string]*gzipBintree{
				"hello": {Func: nil, Children: map[string]*gzipBintree{
					"index.go": {Func: gzipBindataTemplatesProjectPagesHelloIndexgo, Children: map[string]*gzipBintree{}},
				}},
				"home": {Func: nil, Children: map[string]*gzipBintree{
					"about": {Func: nil, Children: map[string]*gzipBintree{
						"index.go": {Func: gzipBindataTemplatesProjectPagesHomeAboutIndexgo, Children: map[string]*gzipBintree{}},
					}},
				}},
				"index.go": {Func: gzipBindataTemplatesProjectPagesIndexgo, Children: map[string]*gzipBintree{}},
				"layout": {Func: nil, Children: map[string]*gzipBintree{
					"index.go": {Func: gzipBindataTemplatesProjectPagesLayoutIndexgo, Children: map[string]*gzipBintree{}},
				}},
				"not_found": {Func: nil, Children: map[string]*gzipBintree{
					"index.go": {Func: gzipBindataTemplatesProjectPagesNotfoundIndexgo, Children: map[string]*gzipBintree{}},
				}},
				"themes": {Func: nil, Children: map[string]*gzipBintree{
					"dark.go":  {Func: gzipBindataTemplatesProjectPagesThemesDarkgo, Children: map[string]*gzipBintree{}},
					"light.go": {Func: gzipBindataTemplatesProjectPagesThemesLightgo, Children: map[string]*gzipBintree{}},
				}},
			}},
			"service": {Func: nil, Children: map[string]*gzipBintree{
				"hello.go": {Func: gzipBindataTemplatesProjectServiceHellogo, Children: map[string]*gzipBintree{}},
				"index.go": {Func: gzipBindataTemplatesProjectServiceIndexgo, Children: map[string]*gzipBintree{}},
			}},
		}},
	}},
}}
