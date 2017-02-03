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

var _generatorTmplSqlconnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\xcb\x4e\x33\x31\x0c\x85\xd7\xf1\x53\xf8\x9f\x55\xfb\xab\x24\x7b\xa4\xae\xd8\x53\x04\x0f\x80\x72\xf1\x84\x68\xa6\xce\x34\x17\x2a\x14\xe5\xdd\x51\x80\x0a\x89\xdd\x67\x9f\x73\x6c\xcb\xad\xa1\xa3\x39\x30\xe1\x94\x2f\xab\x8d\xcc\x13\xf6\xae\x14\xea\x5a\x22\x7a\x62\x4a\xba\x90\x43\xf3\x81\xa5\x2e\x77\x9e\x18\x36\x6d\x17\xed\x09\x5b\x43\xf9\xf4\xc3\xbd\x03\x84\xf3\x16\x53\xc1\x1d\x88\xc9\x87\xf2\x56\x8d\xb4\xf1\xac\x4a\x5d\x1c\xe5\x45\x19\x9d\x49\xe5\xcb\xaa\x9c\x99\x60\x0f\xf0\xae\xd3\xb0\xbe\xa2\x33\xf2\x39\x5e\x5f\xac\x66\xa6\x74\x53\xc2\x28\x1e\x22\x33\xfe\x77\x46\x0e\x00\x98\x2b\x5b\x3c\x6d\xc4\x3b\x62\x1f\x98\x0e\x98\x63\x4d\x96\x30\x97\x14\xd8\x1f\x30\x6e\x25\xa3\x94\xd2\x19\x79\xda\x4a\x88\xbc\x47\x4a\x29\x26\x6c\x20\xec\x61\x30\xde\x1f\xc7\xbe\x47\xba\xfe\x99\xf1\x1d\x96\x52\xee\x41\x84\xf9\xcb\xfa\xef\x88\x1c\xd6\x91\x15\x89\x4a\x4d\x3c\xba\x20\x3a\x80\xf8\xbd\xee\x88\x16\x6e\x32\x87\x15\x3a\x40\x6b\x48\xec\xc6\x47\x3e\x03\x00\x00\xff\xff\x58\x2f\x1e\xac\x5c\x01\x00\x00")

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

	info := bindataFileInfo{name: "generator/tmpl/sqlconn.tmpl", size: 348, mode: os.FileMode(420), modTime: time.Unix(1485606521, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generatorTmplSqlmgrTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x58\xdf\x6f\xdb\xb6\x13\x7f\xf7\x5f\x71\x5f\xe3\xbb\x42\x2a\x5c\x3a\x0f\xc3\x1e\x3a\x78\x40\xda\xa4\x45\xb0\x36\x69\xeb\x6e\x7b\x08\x8a\x42\x3f\x4e\x0e\x63\x99\x74\x48\xaa\xa9\x61\xe8\x7f\x1f\x8e\xa4\x7e\xd9\x72\xa2\xb8\xdb\x8a\x02\x96\x78\xc7\xe3\xe7\xee\x3e\x77\x3c\x65\xbb\x85\x14\x33\x2e\x10\xc6\xfa\x2e\x5f\x2d\xd4\x18\xca\x72\x3a\x85\xa8\x30\x12\x16\x28\x50\x45\x06\x53\x88\x37\x60\x8a\xe5\x8b\x05\x8a\xd1\x76\x0b\xff\x97\xf1\xed\x65\xb4\x42\x78\x39\x03\x66\x1f\x5e\x94\xa5\x15\x68\xa3\x8a\xc4\x54\xb2\xb5\xe2\xc2\x64\x30\xd6\x3f\xe9\xf9\xc7\x77\xef\xc9\x78\xbd\x95\x76\xac\xa3\x64\x19\x2d\x10\xb6\x5b\x60\x1f\xfc\x73\x59\x8e\x46\x7c\xb5\x96\xca\x40\x30\x02\x00\x18\xa7\x91\x89\xe2\x48\xe3\x54\xdf\xe5\xe3\x91\x5b\x5b\x70\x73\x53\xc4\x2c\x91\xab\xe9\x42\x26\x2a\xca\xcc\x34\x8d\xd5\x78\x4f\x68\x8a\x65\x8a\x7a\x39\xad\xf6\x4f\xd3\x78\x80\x92\xd9\xac\x51\x8f\x47\xe1\x68\xf4\x2d\x52\x1e\xc6\x57\xd0\x77\x39\xfb\x84\xba\xc8\x8d\x5f\x48\x63\xf6\x5a\x0a\xe1\xdf\xec\x26\x76\x91\xd6\x42\xc5\xce\x78\x94\x63\x62\x1c\x66\x2e\x74\x3b\x74\x65\xe9\x42\x02\x33\xb8\xc4\xfb\x3e\x49\x20\x78\x1e\x12\x88\xac\x10\xc9\x41\x25\xfc\x8e\x49\x61\xa4\x22\x34\xe7\xfe\x39\x84\xe7\x3b\xc9\x28\x4b\xd8\x5a\x14\x0a\x4d\xa1\x04\x3c\xdb\x93\x3b\x31\xfd\xab\x2c\xbe\xac\x9f\x26\x56\x56\x8e\x4a\x8f\xa5\x17\xc8\xe3\x87\x1e\x88\x00\x99\xa5\xe0\xc1\xfe\x76\xf7\xe6\xad\xc8\xf8\xd6\x1d\xd1\x18\x18\xb5\xf1\xb6\x23\x50\x43\x0d\xf6\x41\x85\xf0\x39\x8a\x73\x0c\x42\x32\xcf\xc5\xa2\x0b\x72\x4c\x64\xb4\x0a\x17\x29\x0a\xc3\x33\x8e\x0a\xca\x72\xfc\xb0\xc5\x73\xb1\xe0\x62\xcf\xe4\x76\x0b\x3c\x03\x81\xc0\x9c\x1c\xc6\x63\xcb\xfb\xdd\xe3\xbc\x94\x4e\xf1\xdb\x30\xd7\xb8\xab\x9a\xc6\xec\x0c\xb3\xa8\xc8\x4d\x75\x5a\xad\x2d\x52\xab\x5c\x43\x34\x37\x5c\xf7\xa4\x23\x04\xe2\x6b\x10\x42\xf0\xdc\x73\x77\x02\xa8\x14\x11\xc6\x01\xe6\x19\x70\x21\x50\x91\x08\xfe\x37\x03\xc1\x73\x68\x88\x51\xa7\xd1\x6b\x4c\x48\xee\xa9\x61\x7f\x92\xca\x22\x15\x7f\x1a\xb3\xb7\x68\x2c\x14\x56\x01\x0e\xab\x53\x48\xe7\x90\x7d\xc1\x73\x6b\xa4\x6d\xd9\x8b\x92\xfa\xd4\xc7\x5d\x7d\x5f\x68\xe3\xdd\xad\xbc\x3d\xca\xcb\x1e\x18\x69\xcc\xc8\x7a\x8f\x7b\x8f\xc3\xaa\x18\x1a\x84\x6d\xbe\x36\xc8\xac\xc1\x9a\xd2\x87\xd0\x75\xb4\x7a\x10\x5a\x79\x13\x81\x01\xb8\xe6\x66\x65\x2e\x84\x46\x65\x2e\x84\x91\x41\x22\xf3\x62\x25\x80\x31\xe6\x18\x6d\x83\xa8\x98\xd3\x20\x5d\x8f\x49\x53\xaa\x5d\x10\x6a\xc7\xd8\xfc\xe3\x3b\x4f\x4e\x0f\x47\xb3\x96\x69\xab\xed\x4b\x30\x64\xaf\xed\x41\xda\x1f\xc8\x18\x1b\x0a\x76\x8e\xd4\x5b\x0f\x01\x75\xd2\x23\x80\x76\xcc\x12\x1c\xf6\x46\xc9\x55\x17\xf4\x40\x84\x7f\xac\xd3\xc8\x60\xe0\x11\xb9\xb7\x23\x10\x79\x33\xc7\x20\x38\xc3\x1c\x1b\x04\xee\xed\x08\x04\x6e\xe3\x51\x71\x70\x79\x0f\x64\x7c\xab\x29\x47\x3b\xfd\x3b\x84\xa0\xb9\x55\xf7\x7b\x51\x8e\xc2\xee\x0c\x61\x36\x83\x93\x03\x9d\x62\xa7\x07\x51\x51\xf4\xb8\xe5\x84\x9a\x7c\xaf\x84\x3b\x8c\xb7\x6b\x32\xbe\x65\x6f\x38\xe6\xa9\x0e\x42\xcb\x45\xda\x95\x49\x05\x5f\x27\xf6\xfe\x79\x39\x03\x15\x89\x05\x82\x75\xa8\x01\x64\x0d\xcf\xec\x0f\xfb\x33\xca\x0b\xd4\x04\xbc\x7a\x6c\x4c\x75\xab\x94\xb0\x5a\x8c\x04\x25\xa0\xcd\x43\x42\xfa\x86\x8b\xf4\x4a\xe0\xab\xcd\x07\xc5\x57\x91\xda\xfc\x8e\x9b\x60\x89\x1b\x37\x45\xd5\x4b\xec\xf3\x66\x8d\xec\xad\xa4\x1f\x17\xe9\x9d\xd8\xef\x84\xfb\xc1\xb8\xb5\xbb\x8a\x3f\xde\x0e\x1e\x8e\x2c\xec\xfc\x2e\x18\xef\x9c\xee\xca\x9a\xae\xb4\x09\x2c\x71\x13\xda\x3c\x0d\xf5\xae\xed\x9a\x26\xdf\x34\x5c\x7f\x79\xd0\xbd\x09\x48\x95\xa2\x22\xbd\x34\x66\x57\xf4\x1c\x42\x70\xfd\xe5\x1f\xf4\xf9\x69\x0e\x6b\xe7\xf1\x04\x4e\xec\x7f\x87\xee\x09\xd9\x0d\xee\x6f\x50\xa1\x9d\x22\x5f\x15\x3c\x4f\x51\x4d\xe8\xf6\x4b\xc9\x43\x2e\x0c\xaa\x2c\x4a\x70\xfb\x83\x69\xad\xca\xa1\xe5\x99\x6f\x7f\xdd\x62\x38\xcd\x73\xc7\xe1\x9e\x5e\xc8\xfe\x22\xa0\x0e\xae\x87\x68\xd9\xee\x62\x28\xef\x6b\xfb\x1f\x0b\x54\x9b\x4f\xf2\xbe\xc5\x75\x17\x66\xdb\x85\x9e\x75\xbd\xd8\x96\xed\x49\x81\xaa\x0e\x35\x9b\x27\x91\x20\x28\x4a\xde\x87\xbf\x1e\x3b\x42\x28\xd4\x43\x27\x08\x9b\xf5\x61\x79\x98\x80\xcc\x32\x8d\x66\x02\x39\x5f\x71\x03\x05\x17\xe6\x97\x9f\xff\x05\x56\x3e\x3d\x65\x75\x6c\x7a\x72\x57\xcb\x1e\x4a\x22\xcf\xbc\x73\xf0\x5b\xa7\x09\xb7\x7b\xde\x95\x55\x08\x9c\x5e\xa7\xd5\x51\x1b\xb7\x21\x39\xbc\xf9\x1d\xc9\x03\xab\xd5\xd9\x5a\xf5\x5d\x8a\x5b\xab\xf3\xba\x90\x36\xa6\x08\x1f\xad\xb1\xb3\xf3\xf9\xeb\xd6\xfa\x1e\x46\x52\x3a\x43\x9d\x04\x4e\xdd\x55\x6c\xd8\xd1\x4f\xa4\x30\x5c\x14\x58\x2f\x7a\x28\xbd\xb6\x4e\x7b\x4d\xd5\x74\xb3\xbc\x5e\x45\x4b\xec\xcb\xf8\x49\x53\x21\xba\x1e\x99\x9b\x3a\xd9\x29\x12\x9e\xd1\xb0\x78\xa1\xcf\x95\xba\x94\x9f\xe4\xbd\x0e\x50\xd5\x7c\xe9\x23\x77\x37\x01\x4f\x2d\x15\x0a\x3c\x01\x63\x97\xf8\xdd\x04\xed\x73\xfc\x1d\xd8\x5f\xad\xdd\x8a\x25\x2a\xb6\x2a\x56\x1f\x2a\xd9\x43\x58\x76\xa2\x4f\xf1\x9c\x41\xb4\x5e\xa3\x48\x03\xeb\xa7\x8c\x6f\xfb\x2e\xd5\x27\x14\xb8\x9b\xac\xae\xe2\x5b\xaa\x1f\xfb\x79\x79\x9a\x24\xb2\x10\x66\x02\x99\x1d\x02\xda\x63\x65\xe0\x4b\x7a\x78\xad\x7e\xb3\xb7\xbf\x63\xc1\xfa\xda\xd9\x69\x37\x0d\x1f\x36\x4f\x73\x7b\x62\x43\x73\x0f\xa0\x89\x93\xb3\x76\x6d\xd7\xbf\x80\x8b\xaf\x9d\x2f\x02\xbb\xd4\x09\xc5\x7a\x59\x65\xa0\x35\x23\xf4\xdc\x6b\x7e\xb4\xec\xde\x6c\xeb\xe5\xa4\x65\x7c\xbd\x0c\xab\xab\xcc\x21\x18\x72\x89\x79\xbb\x43\x7b\xa7\x0f\x54\x7f\x94\x0e\x44\xde\x8f\x87\x1e\xd3\x81\x01\xf1\xe4\x07\xc7\xc3\x6a\x82\x67\x73\x34\xef\xa3\x75\x75\x5a\x05\xc0\xf9\xb7\xc7\xe8\x76\xa3\x38\xd0\x58\xbb\x8d\xa2\xdb\x00\x76\x06\xc2\x01\x35\x7c\xf2\xc0\x65\xc7\xa8\x5d\x9c\x66\x19\x26\x06\xd3\x41\x5f\x82\x6e\xda\xef\xa9\x8a\x63\x6a\x60\x28\x13\xfd\x97\xca\x20\x26\x0e\x77\xa2\x67\x90\x64\x8c\x3d\x36\x28\x1f\xe9\x64\x7d\x01\x1f\xe9\x69\x33\x35\x0e\x77\x70\xf8\x98\xf8\x64\xa7\xf6\x8a\xa1\xfa\x98\x7c\x78\xe2\xfb\xaf\xd9\xdc\xfa\xf3\xd7\xdf\x01\x00\x00\xff\xff\x0c\x2f\xcf\x13\xcc\x16\x00\x00")

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

	info := bindataFileInfo{name: "generator/tmpl/sqlmgr.tmpl", size: 5836, mode: os.FileMode(420), modTime: time.Unix(1486094912, 0)}
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

