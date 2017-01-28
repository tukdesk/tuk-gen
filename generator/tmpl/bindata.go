// Code generated by go-bindata.
// sources:
// generator/tmpl/sqlconn.tmpl
// generator/tmpl/sqlmgr.tmpl
// generator/tmpl/struct.tmpl
// DO NOT EDIT!

package tmpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _generatorTmplSqlconnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xcb\x4e\x33\x31\x0c\x85\xd7\xf1\x53\xf8\x9f\x55\xfb\xab\x24\x7b\xa4\xae\xd8\x53\x04\x0f\x80\x72\xf1\x84\x68\xa6\xce\x34\x17\x2a\x14\xe5\xdd\x51\x80\x0a\x89\xdd\x67\x9f\x73\x6c\xcb\xad\xa1\xa3\x39\x30\xe1\x94\x2f\xab\x8d\xcc\x13\xf6\xae\x14\xea\x5a\x22\x7a\x62\x4a\xba\x90\x43\xf3\x81\xa5\x2e\x77\x9e\x18\x36\x6d\x17\xed\x09\x5b\x43\xf9\xf4\xc3\xbd\x03\x84\xf3\x16\x53\xc1\x1d\x88\xc9\x87\xf2\x56\x8d\xb4\xf1\xac\x4a\x5d\x1c\xe5\x45\x19\x9d\x49\xe5\xcb\xaa\x9c\x99\x60\x0f\xf0\xae\xd3\xb0\xbe\xa2\x33\xf2\x39\x5e\x5f\xac\x66\xa6\x74\x53\xc2\x28\x1e\x22\x33\xfe\x77\x46\x0e\x00\x98\x2b\x5b\x3c\x6d\xc4\x3b\x62\x1f\x98\x0e\x98\x63\x4d\x96\x30\x97\x14\xd8\x1f\x30\x6e\x25\xa3\x94\xd2\x19\x79\xda\x4a\x88\xbc\x47\x4a\x29\x26\x6c\x20\xec\x61\x30\xde\x1f\xc7\xbe\x47\xba\xfe\x99\xf1\x1d\x96\x52\xee\x41\x84\xf9\xcb\xfa\xef\x88\x1c\xd6\x91\x15\x89\x4a\x4d\x3c\xba\x20\x3a\x80\xf8\xbd\xee\x88\x16\x6e\x32\x87\x15\x3a\xb4\x86\xc4\x6e\x3c\xe4\x33\x00\x00\xff\xff\x9b\x4d\x4c\x34\x5b\x01\x00\x00")

func generatorTmplSqlconnTmplBytes() ([]byte, error) {
	return bindataRead(
		_generatorTmplSqlconnTmpl,
		"generator/tmpl/sqlconn.tmpl",
	)
}

func generatorTmplSqlconnTmpl() (*asset, error) {
	bytes, err := generatorTmplSqlconnTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generator/tmpl/sqlconn.tmpl", size: 347, mode: os.FileMode(420), modTime: time.Unix(1485581465, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generatorTmplSqlmgrTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xc4\x57\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x4c\x85\x36\xa0\x02\x87\xce\x39\x85\x0f\xc9\x7e\x04\x46\x93\x4d\x76\x1d\xb4\x87\x20\x08\x24\x71\xec\xa5\x2d\x93\x36\x49\xd5\x31\x04\xfd\xf7\x82\x1f\x92\x65\x5b\x76\xdc\x6d\x83\x9c\xd6\xe6\x70\xde\xbc\x99\xf7\x38\xc6\x56\x15\x30\x9c\x72\x81\x10\xeb\x75\xb1\x9c\xa9\x18\xea\x7a\x38\x84\xb4\x34\x12\x66\x28\x50\xa5\x06\x19\x64\x5b\x30\xe5\xe2\xc5\x0c\x45\x54\x55\xf0\xab\xcc\xe6\x77\xe9\x12\xe1\xd5\x08\xa8\xfb\xf0\xa2\xae\x5d\x40\x1b\x55\xe6\xa6\x89\xad\x14\x17\x66\x0a\xb1\xfe\x4d\x4f\xee\xdf\xbd\xb7\xe0\x6d\xaa\xcd\x58\xa5\xf9\x22\x9d\x21\x54\x15\xd0\x8f\xe1\x73\x5d\x47\x11\x5f\xae\xa4\x32\x40\x22\x00\x80\x98\xa5\x26\xcd\x52\x8d\x43\xbd\x2e\xe2\xc8\x9f\xcd\xb8\x79\x2c\x33\x9a\xcb\xe5\x70\x26\x73\x95\x4e\xcd\x90\x65\x2a\x3e\x0a\x9a\x72\xc1\x50\x2f\x86\x4d\xfe\x90\x65\x17\x5c\x32\xdb\x15\xea\x38\x4a\xa2\xe8\xef\x54\x05\x1a\x5f\x41\xaf\x0b\xfa\x80\xba\x2c\x4c\x38\x60\x19\xbd\x92\x42\x84\x6f\x2e\x89\x8e\x59\x1b\x54\xf4\x9a\xa7\x05\xe6\xc6\x73\xe6\x42\x77\x47\x57\xd7\x7e\x24\x30\x82\x67\x07\x93\xab\xeb\xaa\xb6\xc5\xa7\xa5\xc8\xa1\x2f\x87\x24\xf0\xfc\x28\x07\x2a\x57\x46\xa1\x29\x95\x38\x55\x2d\xaa\xa3\xc8\x12\x85\xe3\x74\xff\x2d\xa0\xc8\x6c\xee\x4b\xec\x00\x6c\xaa\x63\x44\x8e\x6b\x27\xf0\x29\xcd\x0a\x24\x89\x45\xe1\x62\xb6\xcf\x25\xb6\xfa\xba\x0b\x63\x86\xc2\xf0\x29\x47\x05\x75\x1d\x9f\x47\xbc\x11\x33\x2e\xce\x42\xfa\x1b\xfb\x48\xe6\x91\xeb\x9e\xe1\x24\x60\x95\x22\x09\x90\xe7\x41\xb5\x01\xa0\x52\x52\x25\x01\x97\x4f\x81\x0b\x81\xca\x86\xe0\x97\x11\x08\x5e\x84\xc8\xde\x50\xc3\x8d\x81\x8d\xbb\x68\xed\xc5\xcd\x1b\x44\x6b\x7b\x96\xd1\xb7\x68\x1c\x15\xda\x74\x91\x34\x55\xec\x9d\x53\xf8\x82\x17\x0e\xa4\x8b\x1c\x42\x79\x5b\xf5\xfb\xad\xbe\x2f\xb5\x09\xed\x36\xdd\x3e\xa9\xcb\x1e\x1a\x2c\xa3\x16\xbd\xa7\xbd\xef\xd3\x9a\x98\xa5\x19\x0b\x8d\xca\x8c\x85\x91\x24\x97\x45\xb9\x14\x40\x29\xf5\x02\x3b\xb2\x8a\xfa\x1b\xf6\x6e\x60\xa6\xed\x48\x5d\xb1\x5d\x5f\x74\x72\xff\xae\x4b\x4b\xd3\x0e\xb0\xbb\x1b\xfc\x98\xd0\x2b\x57\x46\x87\x72\x94\xd2\x4b\xa9\x4e\xd0\xbe\xdd\x53\x34\x7d\xf4\x5f\xd3\xdc\x03\xb5\x64\xe8\xad\x92\xcb\x7d\xca\x17\xf0\xf3\xdd\x12\x99\xcd\xb5\x65\x76\xf0\x52\x13\x20\xbb\x5d\x75\xec\xf3\x02\x85\xcb\x4c\x60\x34\x82\x97\x27\x5c\xd8\xe3\xef\x9e\x06\x7d\x50\xdb\x29\x34\xc1\x03\x95\xdd\x99\xcc\xe6\xf4\x96\x63\xc1\x34\x49\x9c\x02\x36\x6b\x2a\x15\x7c\x1d\xb8\x4d\xf3\x6a\x04\x2a\x15\x33\x04\xd7\xd0\x8e\x90\x03\x1e\xb9\x3f\xf4\xcf\xb4\x28\x51\x5b\xe2\xcd\xc7\x1d\x54\x20\xb9\x2e\x51\x6d\xdb\x57\x68\x29\x5b\x01\xe8\x9b\x92\x17\x8c\x58\x90\xff\xe3\x11\xd2\x9b\x6f\x98\x13\x57\xea\x12\xa5\x6e\xb9\x60\x1f\x04\xbe\xd9\x7e\x54\x7c\x99\xaa\xed\x1f\xb8\x25\x0b\xdc\xfa\x9f\xbc\xf6\x88\x7e\xda\xae\x90\xbe\x95\xf6\x8f\x17\xf0\x40\xd2\x03\x15\xcf\xca\xa1\xd0\x79\xf1\xd9\x3e\x42\x55\xfb\xa9\x3b\x19\xda\xdc\x56\x99\xd7\x45\x41\x92\x7e\x35\x83\x67\x7d\xa6\xf3\x6c\x3b\xb2\xbf\x1e\x51\x21\x69\x27\x7d\xb3\x26\xf1\x41\x63\xfe\xf9\xd9\x1d\x3d\x80\x05\x6e\x93\xe4\x87\x4b\x25\x37\x2d\xe4\xbd\xad\xf2\x20\x37\x8d\x5c\x5d\x4c\xeb\x39\xd4\x74\x92\xa7\xc2\xf6\xae\xe4\x26\xf9\xfd\xa9\xbe\x50\xa8\x2f\xdd\xcd\xd6\x0f\x5d\x33\x68\xeb\x06\x0d\x9f\xbf\x9c\x35\xc4\x00\xa4\x62\xa8\xdc\x63\x67\x19\xfd\x60\xbf\x24\x40\x3e\x7f\xf9\xef\x36\x59\xa6\x0b\xec\x03\x7a\x39\x70\x8b\xc2\xb2\x0b\x3f\x5e\x3f\xd9\x3a\xba\xf1\x4e\xb3\x38\xec\x10\x3a\xab\xc3\x0f\x68\x27\x1b\x9f\xfa\x33\x7a\x7d\x33\xb9\xea\x9c\x1f\x2e\x16\x37\xcc\x6b\xd4\x39\xf1\xd7\x7d\xd9\x64\xef\x7e\x2e\x85\xe1\xa2\xc4\xf6\x30\xe8\xdf\x8b\xf5\xba\x17\xea\xc7\xef\x28\xb9\xd1\xfb\xc0\xce\xfe\x8d\xf7\x03\x2a\xcb\xe8\x58\xdf\x28\x75\x27\x1f\xe4\x46\x13\x54\xad\x59\xfa\xfc\xdc\x81\x7f\x02\x23\x2b\x94\x65\x45\xef\xf0\x9b\x21\xdd\x3a\x61\xe9\xf7\xaf\xa8\xfd\x47\x6a\x6d\xd6\x79\xa4\xfa\xd4\x2b\x3d\xc5\xe5\x40\x2d\xeb\xfa\x11\xa4\xab\x15\x0a\x46\x5c\x9f\x32\x9b\x27\xe7\xdf\x74\x55\x01\x0a\xe6\xfe\x63\xf9\x27\x00\x00\xff\xff\x66\x1e\xf1\x94\x2f\x0d\x00\x00")

func generatorTmplSqlmgrTmplBytes() ([]byte, error) {
	return bindataRead(
		_generatorTmplSqlmgrTmpl,
		"generator/tmpl/sqlmgr.tmpl",
	)
}

func generatorTmplSqlmgrTmpl() (*asset, error) {
	bytes, err := generatorTmplSqlmgrTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generator/tmpl/sqlmgr.tmpl", size: 3375, mode: os.FileMode(420), modTime: time.Unix(1485583432, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generatorTmplStructTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x5f\x6b\xeb\x36\x14\x7f\xd7\xa7\x38\x98\xbb\x62\x97\x44\x79\xbf\x90\x87\x31\xb6\xcb\x18\x8c\xd1\x96\x3d\xac\x94\x21\xdb\xc7\xa9\x16\x5b\xca\x64\x39\x21\x18\x7d\xf7\x21\x4b\x8e\x9d\xf8\x4f\x9a\x30\xca\x7d\xaa\x2a\x4b\xe7\xf7\x47\x3f\x1d\xa5\xae\x21\xc5\x8c\x0b\x84\xa0\xd4\xaa\x4a\x74\x00\xc6\xac\x56\xc0\x2a\x2d\x61\x83\x02\x15\xd3\x98\x42\x7c\x04\x5d\x6d\x97\x1b\x14\x64\xc7\x92\x2d\xdb\x20\xd4\x35\xd0\x3f\xfc\xd8\x18\x42\x78\xb1\x93\x4a\x43\x48\x00\x00\x82\xac\xd0\x81\x1b\x69\x5e\x60\x40\xdc\x78\xc3\xf5\x7b\x15\xd3\x44\x16\x2b\x5d\x6d\x53\x2c\xb7\xab\x98\x95\xb8\x2a\xff\xcd\x57\x69\x1c\x5c\x5f\xa4\x8f\x3b\x2c\x03\x12\x11\xb2\x67\xca\x63\xfd\x0d\x59\xa1\xe9\xb3\x56\x5c\x6c\x50\xf9\x29\x8b\x4a\x5f\x78\x81\xfe\xff\x34\xa6\x4f\xf2\xf0\x9c\x30\x21\xba\x35\xb6\x18\xfd\x35\xb5\xe5\xea\x1a\xbe\xc8\xf8\x9f\xdf\x59\x81\xf0\x75\x0d\xb4\x19\x18\x43\xec\x1a\xe8\x7f\x34\x06\x9c\x51\x50\x37\x65\x1a\x1b\x14\x2f\x98\x3a\xfe\x86\xc7\x76\xdf\xe5\xf4\xcb\x71\x87\xf4\x9b\xb4\x7f\x6c\x55\xb7\x71\x09\x8a\x89\x0d\xc2\x97\x8c\x63\x9e\x36\xb0\xbf\xd8\x51\xd9\x2d\xf1\xdf\xfa\x65\x79\x76\x9a\xac\xf2\x9c\xc5\xb9\xfd\xf0\x58\xd7\x80\x22\x05\x63\xba\x3d\x13\x98\x28\x52\x63\x88\x21\x24\xab\x44\x02\xe1\xe3\xb9\xb8\x08\x1c\x85\x30\x82\xd7\xb7\xb2\xf1\xd4\x0b\x55\xa8\x2b\x25\x4e\xb3\x6e\xd2\xb3\x9c\xd2\xb1\xf4\xb8\xcd\xb9\x76\xcc\x7e\x92\x79\x55\x08\x30\x26\x58\xf4\xab\x38\x01\xcd\xcc\x55\x82\x3f\xe6\xf9\x0d\x14\x83\x8b\xe3\x98\x20\xf0\xa9\x32\x3a\x3a\x61\x04\x63\x32\xa6\x49\x5f\x33\x67\x3c\x93\x77\xc1\xcc\xb8\x62\xcc\x35\x12\xe7\xe1\x9d\xc1\x1f\x18\xea\xa1\x47\xb2\xbe\x3c\xc1\xea\x77\x5e\xc2\x00\xfb\x1b\xea\x31\xe4\x70\xea\x5e\x2c\x20\x96\x32\x8f\x3c\x27\x9e\x81\x2d\x4b\x87\x37\x6f\xbd\x06\xc1\x73\xe8\x32\xe5\xe9\x0f\xea\xfe\x85\x4a\x36\x75\x33\x96\x97\xe8\x63\xd0\xdb\xf0\x38\x0e\xb0\x00\xad\x2a\xec\x4e\x76\x5c\xdd\xf3\x88\xba\x17\x69\xcd\x09\x5b\x09\x53\xfc\x2d\xfd\x7b\xca\x87\xfb\xa1\x46\xef\xdd\x55\xc8\xa9\x76\xf5\x70\xba\x26\x7b\x62\x88\xef\x4a\x36\x50\xdd\xf5\x99\xa3\xf9\x27\xcb\x2b\x0c\x5d\x20\x5d\xa4\x22\xe0\x42\xa3\xca\x58\x82\xb5\xf1\xac\xca\x03\xd7\xc9\x3b\xb8\x65\x6e\x2a\x61\x25\xce\x04\xfe\xeb\xe5\xe1\xb6\xc2\x46\x6e\x13\x21\x1f\x6c\x1a\x27\xcc\x41\xc8\x27\xe1\xce\x7d\x3c\x21\xf5\x1a\x4b\x8a\x19\xab\x72\xdd\x55\xd8\x31\xc1\x93\xd0\x3e\x84\x3f\x2b\x25\x55\x16\x06\x95\xd8\x0a\x79\x10\x5e\xff\x0f\x65\xb0\x70\xc3\x28\xba\x6c\x4d\x33\x26\xbb\x77\x60\xe8\xed\xa9\xd3\x0e\x3f\xdd\xd0\x4b\x27\x6e\xc2\xf5\x66\x3a\xc3\xb8\x7d\x18\x6e\xe6\x3c\x73\xd6\x37\x3f\x12\xff\xb3\xb0\x27\x4c\x90\xef\x51\x7d\x46\xe4\x1f\x3e\x39\xf3\x0f\xdf\x55\xe8\x5b\xa7\xef\x4e\xd1\x9c\x7d\x37\xc7\x68\xc2\x9b\xbb\x73\x64\x7f\xfe\x86\x4a\x1e\xce\x7f\x0d\x7b\x8b\x4a\xa0\x94\xb6\xd1\x42\xeb\x67\xab\x38\xd9\xab\xd2\x52\x2c\xd8\x16\xc3\x33\xe1\x0b\xc8\x51\xb8\x54\x96\xde\xe2\x4c\x2a\xe0\xbe\xa4\xdd\xe4\x74\x7a\x84\xde\xeb\x69\x8b\xbe\xf2\x37\x58\xbb\xcb\x72\x1e\xf1\xf6\xb4\xfa\x8e\x2b\x79\xa0\x4e\x80\xdd\x4a\x29\x8d\x3e\x22\xd7\x1e\xe4\x40\xf1\x84\xbe\x33\x22\x4d\x02\xae\xe2\xf7\x9e\xae\xff\x02\x00\x00\xff\xff\x68\xf2\x51\x67\x4a\x0d\x00\x00")

func generatorTmplStructTmplBytes() ([]byte, error) {
	return bindataRead(
		_generatorTmplStructTmpl,
		"generator/tmpl/struct.tmpl",
	)
}

func generatorTmplStructTmpl() (*asset, error) {
	bytes, err := generatorTmplStructTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generator/tmpl/struct.tmpl", size: 3402, mode: os.FileMode(420), modTime: time.Unix(1485582894, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"generator/tmpl/sqlconn.tmpl": generatorTmplSqlconnTmpl,
	"generator/tmpl/sqlmgr.tmpl": generatorTmplSqlmgrTmpl,
	"generator/tmpl/struct.tmpl": generatorTmplStructTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"generator": &bintree{nil, map[string]*bintree{
		"tmpl": &bintree{nil, map[string]*bintree{
			"sqlconn.tmpl": &bintree{generatorTmplSqlconnTmpl, map[string]*bintree{}},
			"sqlmgr.tmpl": &bintree{generatorTmplSqlmgrTmpl, map[string]*bintree{}},
			"struct.tmpl": &bintree{generatorTmplStructTmpl, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

