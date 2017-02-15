// Code generated by go-bindata.
// sources:
// generator/tmpl/sql_script_table.tmpl
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

var _generatorTmplSql_script_tableTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x4b\xc4\x30\x10\x85\xcf\xe6\x57\x3c\xd6\x1e\x14\xb6\xfd\x01\x8a\x87\x55\x2b\x14\xaa\x82\xdb\x83\xb7\x25\x6b\xa6\x25\x58\xa3\xa6\x59\xa8\x84\xf9\xef\x92\x64\x6d\x44\x6f\x93\xc7\xf7\xde\xbc\x8c\xf7\x50\xd4\x6b\x43\x58\x4d\x9f\xe3\x6e\x7a\xb1\xfa\xc3\xed\x9c\xdc\x8f\xb4\x02\xf3\x29\xe4\xc1\xbd\x63\x20\x43\x56\x3a\x52\xd8\x7f\xc1\x1d\x5e\xcb\x81\x8c\xb8\x14\xc2\x7b\x58\x69\x06\x42\x11\x1d\xb8\xb8\x42\x85\x92\x59\xdc\x3c\xd5\x9b\xae\x46\xb7\xb9\x6e\x6b\x34\x77\x78\x78\xec\x50\x3f\x37\xdb\x6e\x0b\xef\x51\x75\x81\x6e\x14\x19\xa7\x7b\x4d\x16\xcc\x38\x13\x27\xde\xa3\x18\xb5\xa1\x29\xe4\xa4\xc4\xea\x96\xfa\x36\x4a\x21\x35\x12\x6f\x72\x6e\xd4\x1c\x91\x08\x57\xf7\x49\xf8\x01\x8e\x85\xb4\x9a\xd7\x89\xc8\x68\x0e\x89\x7a\x7a\x94\xd0\x3d\x4c\x72\x2c\xe9\xcc\x6b\x44\x92\x8c\x5a\x5c\x61\x66\x16\xe7\xe1\xdb\x05\xcd\xce\xca\xf6\x4f\xdb\x3a\x8b\xc1\x94\xdb\x2c\x3d\xe8\x1f\xb1\x74\xf9\xb5\x2d\x9d\xf6\xb8\x2f\x4f\xdf\x01\x00\x00\xff\xff\x1b\xf0\xb5\x96\xad\x01\x00\x00")

func generatorTmplSql_script_tableTmplBytes() ([]byte, error) {
	return bindataRead(
		_generatorTmplSql_script_tableTmpl,
		"generator/tmpl/sql_script_table.tmpl",
	)
}

func generatorTmplSql_script_tableTmpl() (*asset, error) {
	bytes, err := generatorTmplSql_script_tableTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "generator/tmpl/sql_script_table.tmpl", size: 429, mode: os.FileMode(420), modTime: time.Unix(1486109836, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generatorTmplSqlconnTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x90\x4d\x4e\x03\x31\x0c\x85\xd7\xf1\x29\xcc\xac\x5a\x54\x92\x3d\x52\x57\xec\x29\x82\x03\xa0\xfc\xb8\x21\x9a\xa9\x33\xcd\x0f\x15\x8a\x72\x77\x14\x68\x85\xc4\xee\xd9\xef\x7d\xb6\xe5\xd6\xd0\xd1\x31\x30\xe1\x94\xcf\x8b\x8d\xcc\x13\xf6\xae\x14\xea\x5a\x22\x7a\x62\x4a\xba\x90\x43\xf3\x85\xa5\xce\x0f\x9e\x18\x56\x6d\x67\xed\x09\x5b\x43\xf9\x72\xd5\xbd\x03\x84\xd3\x1a\x53\xc1\x0d\x88\xc9\x87\xf2\x51\x8d\xb4\xf1\xa4\x4a\x9d\x1d\xe5\x59\x5d\x69\x95\xcf\x8b\x72\x66\x82\x2d\xc0\xa7\x4e\x23\xfd\x8e\xce\xc8\xd7\x78\x79\xb3\x9a\x99\xd2\xcd\x09\xa3\x78\x8a\xcc\x78\xef\x8c\x1c\x02\xe0\x58\xd9\xe2\x61\x25\xde\x10\xfb\xc0\xb4\xc3\x1c\x6b\xb2\x84\xb9\xa4\xc0\x7e\x87\x71\x2d\x19\xa5\x94\xce\xc8\xc3\x5a\x42\xe4\x2d\x52\x4a\x31\x61\x03\x61\x77\x43\xe3\xe3\x7e\xec\x7b\xa6\xcb\xbf\x19\xbf\xb0\x94\x72\x0b\x22\x1c\x7f\xa2\x77\x7b\xe4\xb0\x0c\x56\x24\x2a\x35\xf1\xe8\x82\xe8\x00\xe2\xef\xba\x3d\x5a\xb8\xd9\x1c\x16\xe8\x00\xad\x21\xb1\x1b\x4f\xf9\x0e\x00\x00\xff\xff\xc0\xba\x5f\x28\x5f\x01\x00\x00")

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

	info := bindataFileInfo{name: "generator/tmpl/sqlconn.tmpl", size: 351, mode: os.FileMode(420), modTime: time.Unix(1487144781, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generatorTmplSqlmgrTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x59\x5b\x6f\xdb\xb8\x12\x7e\xf7\xaf\x98\x23\x9c\x53\x48\x85\xcb\xa4\xc0\xc1\x3e\x74\xe1\x05\xda\x5c\x8a\x60\xdb\xa4\x6d\xba\xbb\x0f\x41\x50\xe8\x32\x72\x19\xcb\x94\x43\x52\x75\x0d\x43\xff\x7d\xc1\x8b\x24\x4a\x96\x6c\x35\xdb\x00\xdd\x16\x45\x2d\x89\xc3\x99\x6f\x86\x33\x1f\x39\xec\x76\x0b\x09\xa6\x94\x21\x78\xe2\x3e\x5b\xce\xb9\x07\x65\x79\x74\x04\x61\x21\x73\x98\x23\x43\x1e\x4a\x4c\x20\xda\x80\x2c\x16\xcf\xe6\xc8\x26\xdb\x2d\xfc\x37\x8f\xee\x2e\xc3\x25\xc2\x8b\x19\x10\xfd\xf0\xac\x2c\xf5\x80\x90\xbc\x88\x65\x35\xb6\xe2\x94\xc9\x14\x3c\xf1\x3f\x71\xfd\xfe\xcd\x5b\xa5\xbc\x9e\xaa\x66\xac\xc2\x78\x11\xce\x11\xb6\x5b\x20\xef\xec\x73\x59\x4e\x26\x74\xb9\xca\xb9\x04\x7f\x02\x00\xe0\x25\xa1\x0c\xa3\x50\xe0\x91\xb8\xcf\xbc\x89\xf9\x36\xa7\xf2\x73\x11\x91\x38\x5f\x1e\xcd\xf3\x98\x87\xa9\x3c\x4a\x22\xee\xed\x0c\xca\x62\x91\xa0\x58\x1c\x59\xf0\x4a\xc5\x51\x12\x8d\x93\x93\x9b\x15\x0a\x6f\x12\x4c\x26\x5f\x42\x6e\xc1\x7c\x02\x71\x9f\x91\x0f\x28\x8a\x4c\xda\x0f\x49\x44\x4e\x72\xc6\xec\x9b\x9e\x44\x2e\x92\x7a\x90\x93\x53\x1a\x66\x18\x4b\x83\x9c\x32\xe1\x06\xb0\x2c\x4d\x60\x60\x06\x97\xb8\xee\x1b\xf1\x19\xcd\x02\x05\x22\x2d\x58\x3c\x28\x84\x5f\x31\x2e\x64\xce\x15\x9a\x33\xfb\x1c\xc0\xd3\xce\x92\x94\x25\x6c\x35\x0a\x8e\xb2\xe0\x0c\x9e\xec\x8c\x9b\x61\xf5\xa7\xd2\xf8\xa2\x7e\x9a\xea\xb1\x72\x52\x5a\x2c\xbd\x40\x0e\x1b\x1d\x88\x80\x52\xab\x82\x07\xbb\xd3\xcd\x9b\xd5\x72\x15\xdd\x61\x2c\x8d\x95\x46\xc7\xc4\x85\xec\x06\xa1\x46\xeb\xef\xe2\x0a\xe0\x63\x18\x65\xe8\x07\xca\x02\x65\xf3\x36\x4e\x4f\x65\xa5\x16\xb8\x48\x90\x49\x9a\x52\xe4\x50\x96\xde\x7e\x8d\x67\x6c\x4e\xd9\x8e\xca\xed\x16\x68\x0a\x0c\x81\x98\x71\xf0\x3c\x5d\x00\x5d\x73\x76\x54\x59\xb1\xd3\x30\x13\xd8\x15\x4d\x22\x72\x8a\x69\x58\x64\xb2\xb2\x56\x4b\xb3\x44\x0b\xd7\x10\xe5\x67\x2a\x7a\x56\x24\x00\x95\xb2\x7e\x00\xfe\x53\x9b\xbe\x53\x40\xce\x55\xce\x18\xc0\x34\x05\xca\x18\x72\x35\x04\xff\x99\x01\xa3\x19\x34\xb9\x51\xaf\xa4\x95\x98\xaa\x71\x9b\x1d\xfa\x27\xae\x34\x2a\x16\x48\x22\xf2\x1a\xa5\x86\x42\x2a\xc0\x41\x65\x45\xc9\x0c\xe9\x67\x34\xd3\x4a\x5c\xcd\x76\x28\xae\xad\x1e\x76\xf5\x6d\x21\xa4\x75\xb7\xf2\xf6\x41\x5e\xf6\xc0\x48\x22\xa2\xb4\xf7\xb8\x77\x18\x56\x95\xa1\x7e\xe0\xe6\x6b\x83\x4c\x2b\xac\x53\x7a\x08\x5d\x4b\xaa\x07\xa1\x1e\x6f\x22\x30\x02\xd7\xf5\xfb\x37\x36\x52\xaf\x0a\x9a\x25\xc8\xdb\x65\x61\xfc\xac\xb1\x13\x2d\x3e\x46\xad\x5c\xca\x0b\x26\x90\xcb\x0b\x26\x73\x3f\xce\xb3\x62\xc9\x80\x10\x62\x0a\x45\x5b\xe4\xc4\x48\x28\xd9\x1e\xab\xda\x14\x71\x94\xe8\xaf\xb6\x86\x03\x72\xa2\x55\x0a\xab\x9a\x10\x32\x16\xd6\x35\x2a\x7e\x1e\x82\x64\x46\xf7\x42\x6a\x29\x50\x86\xc9\x39\xcf\x97\x6d\x78\x23\xb1\xfc\xb1\x4a\x42\x89\xbe\xb5\x6d\xde\xf6\xda\xb6\x13\x1e\x62\xeb\x14\x33\x6c\x6c\x99\xb7\xbd\xb6\x8c\xc8\x83\x7c\xfb\xeb\x33\x72\xf4\xd7\xea\x5f\xa0\x4c\x22\x4f\xc3\x18\xb7\xe5\x54\x55\x72\x22\x54\xd0\x9d\xaf\xba\x22\xf4\x8c\x41\x28\x8e\x3e\xab\x63\xe4\x82\x9b\xf4\xf1\xf3\xe8\x4e\x5b\xed\xec\x23\x01\xf8\xcd\x06\xbf\xcb\x89\x19\x32\x3d\x33\x80\xd9\x0c\x8e\x07\x18\xab\xc3\x85\xaa\x38\x15\x0d\x76\x0a\xc7\x0c\x0a\x15\xee\x6a\xb0\x53\x22\xfa\x9b\xd9\xee\xc8\x39\xc5\x2c\x11\x7e\xa0\x9d\x54\x13\xd3\x9c\xc3\xa7\x29\xe4\xd1\x9d\x9a\xce\x43\x36\x47\xd0\x3e\x35\x98\xb4\xee\x99\xfe\x21\x7f\x86\x59\x81\x42\x61\xaf\x1e\x1b\x55\x6d\xc2\x50\x70\x35\x4c\x85\xc6\x57\x93\xc7\x44\xf5\x9c\xb2\xe4\x8a\xe1\xab\xcd\x3b\x4e\x97\x21\xdf\xfc\x8e\x1b\x3f\xd5\x98\xe1\xe6\xd6\x94\xd3\x14\x16\xb8\x31\x47\xbd\x5a\x86\x7c\xdc\xac\x90\xbc\xce\xd5\x8f\x89\x7e\x67\x3d\x3a\x4b\x60\xf2\xa7\x8a\x97\xc9\x01\x27\x29\xce\xee\x7d\xaf\x63\xc0\x90\x82\xda\x51\xb5\xfd\x20\xd8\x65\x47\x8b\xdd\xe2\x9d\x82\x4d\xaa\xe3\xb1\x7e\x9f\x17\x59\xd6\xf5\xfd\x07\x77\x55\xe7\xe9\xb7\xfa\xe9\x3a\x28\x7a\x57\x57\xbd\xee\xf5\x7a\x0a\x39\x4f\x90\xeb\xca\x4b\x22\x72\xa5\x5e\x02\xf0\x6f\x6e\x1f\x39\x16\x62\x28\x18\x3b\x8b\xae\xff\x1a\x90\x23\x09\xa5\x4a\x81\x76\x78\xfe\x9d\xd1\x68\xe5\xc5\xc3\x42\xd1\x54\x92\x93\x1b\x06\x72\x45\xea\x53\xc8\xd3\x54\xa0\x84\x82\x32\xf9\xcb\xff\x07\xc2\xb0\x3f\x08\x7b\x59\xd5\x32\xb5\x81\xb1\xc3\xd5\x16\x9d\x9d\xeb\x12\xec\xcb\x2c\xf3\x5b\x9c\xb8\x43\xcf\x76\xa7\x37\x2a\xf4\x4e\xff\x86\x2e\xa9\xf4\x9f\xd7\x07\x5a\xed\xaa\xdd\xbc\x76\x0e\x6d\x2e\x25\x3b\x3b\x58\x15\x16\xf3\x72\x52\xef\x66\x0e\x12\x9a\x56\x41\xfb\xad\xe5\x8c\xab\xf1\x4a\x0b\xf8\x46\xae\x35\xbb\xda\x2d\x54\x68\x9d\xfd\xc2\x84\xbd\x51\xa6\x8c\xa8\x6f\xe4\xf4\xec\xfa\xc4\xf9\xbe\x63\x48\x09\x9d\xa2\x88\x7d\x23\x6e\xf2\x2b\x68\xc9\xc7\x39\x93\x94\x15\x58\x7f\xb4\x50\x7a\x75\xbd\xec\x55\x55\x6d\x4c\xf9\x5a\x61\xd6\x1b\xd3\xfb\x02\xf9\xe6\x43\xbe\x76\x36\x27\x93\xc8\x42\x89\x3c\x69\xa7\xcc\xb6\x74\xbb\x0c\xe5\x36\x0a\x72\x1d\x87\xcc\xe7\xf9\x7a\x0a\xcd\x22\xfe\xfa\xd0\x36\x84\xa3\x18\xdb\x85\x38\x6c\x73\xb0\x30\xa6\x90\xa9\xac\xda\x5f\x1f\x87\x68\xe2\x47\xa9\x90\x1f\xba\x30\x54\x1c\x74\xa8\x87\x27\x9b\x02\xd7\x52\x3f\x51\x4d\x99\x82\x59\x86\x0b\xec\x4b\xa4\xe3\xa0\x2e\x3d\x51\xf7\xf1\x4d\x01\x76\xaa\x8f\xa6\x2a\x83\x2f\xc4\x19\xe7\x97\xf9\x87\x7c\x2d\x7c\xe4\x75\x1a\xf6\x55\x4b\x7b\x01\xbe\xb5\xf6\x54\xe0\x15\x30\x72\x89\x5f\xa5\xef\xda\xb1\xa7\xe1\x7e\x1a\x68\x53\x81\x3a\x09\x57\x54\x20\x46\x70\xc1\x10\xa6\xce\x2a\xa8\xb8\xce\x20\x5c\xad\x90\x25\xbe\xf6\x37\x8f\xee\xfa\x8e\xd9\xdf\xc0\x1c\x27\x79\xc1\xa4\xdf\x66\x8a\x00\x7c\xcb\x0d\xe3\x6b\x7e\xa8\x54\xbd\x58\x1b\x78\x1e\x78\x8f\x5d\xab\x23\x98\xfc\x4b\xc8\x41\xe3\x01\xed\x9f\xf1\xca\xb2\x77\xbe\x36\x4b\xf6\x44\x0b\x04\x9d\x0b\xa1\x82\x99\x86\x6d\x44\x44\x4d\xcf\x6c\x78\x4d\xf5\x44\xdd\xdb\xc4\x2a\x1f\xdc\xbb\x80\xde\x78\x7f\xd1\xbd\x94\xa9\xa4\xd5\x8d\x11\xbd\x75\x1a\x59\x9b\x7a\x96\x2a\xb4\xd2\x86\x2a\xac\x8d\x26\xb0\x46\xdb\x8d\xfe\x7e\x0b\x26\x47\x75\xb7\x66\xd8\xb4\x15\xca\xd5\xa2\xca\x62\xa7\xeb\xb0\x21\x39\x7c\x40\x5c\x2d\xa6\x8e\xf6\xd5\x22\xe8\x3b\x11\xda\x9b\x05\x7b\x1c\x34\xe0\xc6\x9c\x02\xdd\x79\xce\xd6\x66\x63\xd5\x1f\xa8\x81\xf8\xda\x6d\xca\xda\x1e\x68\xba\x8f\xff\x61\xcb\x5d\x5d\xb9\x90\x6b\x94\x6f\xc3\x55\x65\xed\xb1\x6b\x01\x3b\xcc\xda\xe9\xb9\x47\x90\xe3\xf1\x9e\x63\x09\x51\x3c\xfc\x32\x4d\x31\x96\x98\x8c\xba\xa0\x33\x97\x3b\xc3\x45\x31\xb0\x44\xdf\x3f\x11\xbb\x79\x68\xaf\xa9\xb4\xb2\xf1\x8e\xf4\xb4\x63\x84\x90\x43\x3d\xf9\xb0\x87\xee\x51\xe8\xa1\x9e\xda\xde\xeb\xfb\x38\xf8\x18\x3b\x42\x75\x23\xf8\xb3\xe5\xbe\xf3\x5f\x23\x7f\x07\x00\x00\xff\xff\x27\x7c\xb6\x7f\xf1\x1c\x00\x00")

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

	info := bindataFileInfo{name: "generator/tmpl/sqlmgr.tmpl", size: 7409, mode: os.FileMode(420), modTime: time.Unix(1487154472, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generatorTmplStructTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x5d\x6b\xf3\x36\x14\xbe\xd7\xaf\x38\x98\xae\xd8\x25\x51\xee\x0b\xb9\x18\x63\x2b\x63\x30\x46\x5b\x76\xb1\x52\x86\x62\x1f\xa7\x5a\x6c\x29\x93\xe5\x84\x60\xf4\xdf\x5f\x64\xc9\xb1\x13\x7f\xa4\x09\x2f\xe5\xbd\xaa\x2a\x4b\xe7\xf9\xd0\xa3\xa3\x54\x15\x24\x98\x72\x81\x10\x14\x5a\x95\xb1\x0e\xc0\x98\xc5\x02\x58\xa9\x25\xac\x51\xa0\x62\x1a\x13\x58\x1d\x40\x97\x9b\xf9\x1a\x05\xd9\xb2\x78\xc3\xd6\x08\x55\x05\xf4\x2f\x3f\x36\x86\x10\x9e\x6f\xa5\xd2\x10\x12\x00\x80\x20\xcd\x75\xe0\x46\x9a\xe7\x18\x10\x37\x5e\x73\xfd\x51\xae\x68\x2c\xf3\x85\x2e\x37\x09\x16\x9b\x85\x2f\xbb\x28\xfe\xcf\x16\xc9\x2a\xf8\xd4\x3a\x7d\xd8\x62\x11\x90\x88\x90\x1d\x53\x1e\xf1\x5f\x48\x73\x4d\x5f\xb4\xe2\x62\x8d\xca\x4f\x59\x6c\xfa\xca\x73\xf4\xff\x27\x2b\xfa\x2c\xf7\x2f\x31\x13\xa2\x5d\x63\x8b\xd1\xdf\x13\x5b\xae\xaa\xe0\x4e\xae\xfe\xfb\x93\xe5\x08\x8f\x4b\xa0\xf5\xc0\x18\x62\xd7\x40\xf7\xa3\x31\xe0\xec\x82\xaa\x2e\x53\x9b\xa1\x78\xce\xd4\xe1\x0f\x3c\x34\xfb\xce\xa7\x5f\x0f\x5b\xa4\x4f\xd2\xfe\xb1\x55\xdd\xc6\x39\x28\x26\xd6\x08\x77\x29\xc7\x2c\xa9\x61\x7f\xb3\xa3\xa2\x5d\xe2\xbf\x75\xcb\xf2\xf4\x38\x59\x66\x19\x5b\x65\xf6\xc3\x43\x55\x01\x8a\x04\x8c\x69\xf7\x8c\x60\xa2\x48\x8c\x21\x86\x90\xb4\x14\x31\x84\x0f\xa7\xe2\x22\x70\x14\xc2\x08\xde\xde\x8b\xda\x53\x2f\x54\xa1\x2e\x95\x38\xce\xba\x49\xcf\x72\x4c\xc7\xdc\xe3\xd6\x47\xdb\x32\xfb\x45\x66\x65\x2e\xc0\x98\x60\xd6\xad\xe2\x04\xd4\x33\x17\x09\xfe\x9c\x65\x57\x50\x0c\xce\x8e\x63\x84\xc0\x97\xca\x68\xe9\x84\x11\x0c\xc9\x18\x27\x7d\xc9\x9c\xe1\x4c\xde\x04\x33\xe1\x8a\x31\x97\x48\x9c\x86\x77\x02\xbf\x67\xa8\x87\x1e\xc8\xfa\xfc\x08\xab\x3f\x78\x01\x3d\xec\x27\xd4\x43\xc8\xe1\xd8\xbd\x98\xc1\x4a\xca\x2c\xf2\x9c\x78\x0a\xb6\x2c\xed\xdf\xbc\xe5\x12\x04\xcf\xa0\xcd\x94\xa7\xdf\xab\xfb\x0f\x2a\x59\xd7\x4d\x59\x56\xa0\x8f\x41\x67\xc3\xc3\x30\xc0\x0c\xb4\x2a\xb1\x3d\xd9\x61\x75\x2f\x03\xea\x5e\xa5\x35\x27\x6c\x24\x8c\xf1\xb7\xf4\x6f\x29\x1f\xee\xfa\x1a\xbd\x77\x17\x21\xc7\xda\xd5\xfd\xf1\x9a\xec\x88\x21\xbe\x2b\xd9\x40\xb5\xd7\x67\x8a\xe6\xdf\x2c\x2b\x31\x74\x81\x74\x91\x8a\x80\x0b\x8d\x2a\x65\x31\x56\xc6\xb3\x2a\xf6\x5c\xc7\x1f\xe0\x96\xb9\xa9\x98\x15\x38\x11\xf8\xc7\xf3\xc3\x6d\x84\x0d\xdc\x26\x42\x3e\xd9\x34\x8e\x98\xbd\x90\x8f\xc2\x9d\xfa\x78\x44\xea\x34\x96\x04\x53\x56\x66\xba\xad\xb0\x65\x82\xc7\xa1\x7d\x08\x7f\x55\x4a\xaa\x34\x0c\x4a\xb1\x11\x72\x2f\xbc\xfe\x9f\x8a\x60\xe6\x86\x51\x74\xde\x9a\x26\x4c\x76\xef\x40\xdf\xdb\x63\xa7\xed\x7f\xba\xa2\x97\x8e\xdc\x84\xcb\xcd\x74\x82\x71\xf3\x30\x5c\xcd\x79\xe2\xac\xaf\x7e\x24\xbe\xb3\xb0\x67\x8c\x91\xef\x50\x7d\x45\xe4\xef\xbf\x38\xf3\xf7\x3f\x54\xe8\x1b\xa7\x6f\x4e\xd1\x94\x7d\x57\xc7\x68\xc4\x9b\x9b\x73\x64\x7f\xfe\x86\x4a\xee\x4f\x7f\x0d\x7b\x8b\x0a\xa0\x94\x36\xd1\x42\xeb\x67\xa3\x38\xde\xa9\xc2\x52\xcc\xd9\x06\xc3\x13\xe1\x33\xc8\x50\xb8\x54\x16\xde\xe2\x54\x2a\xe0\xbe\xa4\xdd\xe4\x74\x7a\x84\xce\xeb\x69\x8b\xbe\xf1\x77\x58\xba\xcb\x72\x1a\xf1\xe6\xb4\xba\x8e\x2b\xb9\xa7\x4e\x80\xdd\x4a\x29\x8d\x3e\x23\xd7\x1e\x64\x4f\xf1\x88\xbe\x13\x22\x75\x02\x2e\xe2\x77\x9e\xae\x6f\x01\x00\x00\xff\xff\x2f\xa1\x45\x6c\x50\x0d\x00\x00")

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

	info := bindataFileInfo{name: "generator/tmpl/struct.tmpl", size: 3408, mode: os.FileMode(420), modTime: time.Unix(1487144780, 0)}
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
	"generator/tmpl/sql_script_table.tmpl": generatorTmplSql_script_tableTmpl,
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
			"sql_script_table.tmpl": &bintree{generatorTmplSql_script_tableTmpl, map[string]*bintree{}},
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

