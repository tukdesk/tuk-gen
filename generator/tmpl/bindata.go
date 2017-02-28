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

var _generatorTmplSql_script_tableTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\x90\x41\x4b\xc4\x30\x10\x85\xcf\xe6\x57\x3c\xd6\x1e\x14\xb6\xfd\x01\x8a\x87\x55\x2b\x14\xaa\x82\xdb\x83\xb7\x25\x6b\xa6\x25\x58\xa3\xa6\x59\xa8\x84\xf9\xef\x92\x64\x6d\x44\x6f\x93\xc7\xf7\xde\xbc\x8c\xf7\x50\xd4\x6b\x43\x58\x4d\x9f\xe3\x6e\x7a\xb1\xfa\xc3\xed\x9c\xdc\x8f\xb4\x02\xf3\x29\xe4\xc1\xbd\x63\x20\x43\x56\x3a\x52\xd8\x7f\xc1\x1d\x5e\xcb\x81\x8c\xb8\x14\xc2\x7b\x58\x69\x06\x42\x11\x1d\xb8\xb8\x42\x85\x92\x59\xdc\x3c\xd5\x9b\xae\x46\xb7\xb9\x6e\x6b\x34\x77\x78\x78\xec\x50\x3f\x37\xdb\x6e\x0b\xef\x51\x75\x81\x6e\x14\x19\xa7\x7b\x4d\x16\xcc\x38\x13\x27\xde\xa3\x18\xb5\xa1\x29\xe4\xa4\xc4\xea\x96\xfa\x36\x4a\x21\x35\x12\x6f\x72\x6e\xd4\x1c\x91\x08\x57\xf7\x49\xf8\x01\x8e\x85\xb4\x9a\xd7\x89\xc8\x68\x0e\x89\x7a\x7a\x94\xd0\x3d\x4c\x72\x2c\xe9\xcc\xeb\x08\x92\x51\x8b\x29\xcc\xcc\xe2\x3c\xfc\xba\xa0\xd9\x59\xd9\xfe\x29\x5b\x67\x31\x98\x72\x99\xa5\x06\xfd\x23\x96\x2a\xbf\xb6\xa5\xcb\x1e\xf7\xe5\xe9\x3b\x00\x00\xff\xff\xe2\xda\x6e\x23\xac\x01\x00\x00")

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

	info := bindataFileInfo{name: "generator/tmpl/sql_script_table.tmpl", size: 428, mode: os.FileMode(420), modTime: time.Unix(1488296500, 0)}
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

	info := bindataFileInfo{name: "generator/tmpl/sqlconn.tmpl", size: 351, mode: os.FileMode(420), modTime: time.Unix(1488291231, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _generatorTmplSqlmgrTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x59\x5b\x6f\xdb\xb8\x12\x7e\xf7\xaf\x98\x23\x9c\x53\x48\x85\xcb\xa4\xc0\xc1\x3e\x74\xe1\x05\xda\x5c\x8a\x60\xdb\xa4\x6d\xba\xbb\x0f\x41\x50\xe8\x32\x72\x19\xcb\x94\x43\x52\x75\x0d\x43\xff\x7d\xc1\x8b\x24\x4a\x96\x6c\x35\xdb\x00\xdd\x16\x45\x2d\x91\xc3\x99\x6f\x86\x33\x1f\x39\xea\x76\x0b\x09\xa6\x94\x21\x78\xe2\x3e\x5b\xce\xb9\x07\x65\x79\x74\x04\x61\x21\x73\x98\x23\x43\x1e\x4a\x4c\x20\xda\x80\x2c\x16\xcf\xe6\xc8\x26\xdb\x2d\xfc\x37\x8f\xee\x2e\xc3\x25\xc2\x8b\x19\x10\xfd\xf0\xac\x2c\xf5\x84\x90\xbc\x88\x65\x35\xb7\xe2\x94\xc9\x14\x3c\xf1\x3f\x71\xfd\xfe\xcd\x5b\xa5\xbc\x5e\xaa\x56\xac\xc2\x78\x11\xce\x11\xb6\x5b\x20\xef\xec\x73\x59\x4e\x26\x74\xb9\xca\xb9\x04\x7f\x02\x00\xe0\x25\xa1\x0c\xa3\x50\xe0\x91\xb8\xcf\xbc\x89\x19\x9b\x53\xf9\xb9\x88\x48\x9c\x2f\x8f\xe6\x79\xcc\xc3\x54\x1e\x25\x11\xf7\x76\x26\x65\xb1\x48\x50\x2c\x8e\x2c\x78\xa5\xe2\x28\x89\xc6\xc9\xc9\xcd\x0a\x85\x37\x09\x26\x93\x2f\x21\xb7\x60\x3e\x81\xb8\xcf\xc8\x07\x14\x45\x26\xed\x40\x12\x91\x93\x9c\x31\xfb\xa6\x17\x91\x8b\xa4\x9e\xe4\xe4\x94\x86\x19\xc6\xd2\x20\xa7\x4c\xb8\x01\x2c\x4b\x13\x18\x98\xc1\x25\xae\xfb\x66\x7c\x46\xb3\x40\x81\x48\x0b\x16\x0f\x0a\xe1\x57\x8c\x0b\x99\x73\x85\xe6\xcc\x3e\x07\xf0\xb4\xb3\x25\x65\x09\x5b\x8d\x82\xa3\x2c\x38\x83\x27\x3b\xf3\x66\x5a\xfd\xa9\x34\xbe\xa8\x9f\xa6\x7a\xae\x9c\x94\x16\x4b\x2f\x90\xc3\x46\x07\x22\xa0\xd4\xaa\xe0\xc1\xee\x72\xf3\x66\xb5\x5c\x45\x77\x18\x4b\x63\xa5\xd1\x31\x71\x21\xbb\x41\xa8\xd1\xfa\xbb\xb8\x02\xf8\x18\x46\x19\xfa\x81\xb2\x40\xd9\xbc\x8d\xd3\x53\x59\xa9\x05\x2e\x12\x64\x92\xa6\x14\x39\x94\xa5\xb7\x5f\xe3\x19\x9b\x53\xb6\xa3\x72\xbb\x05\x9a\x02\x43\x20\x66\x1e\x3c\x4f\x17\x40\xd7\x9c\x9d\x55\x56\xec\x32\xcc\x04\x76\x45\x93\x88\x9c\x62\x1a\x16\x99\xac\xac\xd5\xd2\x2c\xd1\xc2\x35\x44\xf9\x99\x8a\x9e\x1d\x09\x40\xa5\xac\x1f\x80\xff\xd4\xa6\xef\x14\x90\x73\x95\x33\x06\x30\x4d\x81\x32\x86\x5c\x4d\xc1\x7f\x66\xc0\x68\x06\x4d\x6e\xd4\x3b\x69\x25\xa6\x6a\xde\x66\x87\xfe\x89\x2b\x8d\x8a\x05\x92\x88\xbc\x46\xa9\xa1\x90\x0a\x70\x50\x59\x51\x32\x43\xfa\x19\xcd\xb4\x12\x57\xb3\x9d\x8a\x6b\xab\x87\x5d\x7d\x5b\x08\x69\xdd\xad\xbc\x7d\x90\x97\x3d\x30\x92\x88\x28\xed\x3d\xee\x1d\x86\x55\x65\xa8\x1f\xb8\xf9\xda\x20\xd3\x0a\xeb\x94\x1e\x42\xd7\x92\xea\x41\xa8\xe7\x9b\x08\x8c\xc0\x75\xfd\xfe\x8d\x8d\xd4\xab\x82\x66\x09\xf2\x76\x59\x18\x3f\x6b\xec\x44\x8b\x8f\x51\x2b\x97\xf2\x82\x09\xe4\xf2\x82\xc9\xdc\x8f\xf3\xac\x58\x32\x20\x84\x98\x42\xd1\x16\x39\x31\x12\x4a\xb6\xc7\xaa\x36\x45\x1c\x25\x7a\xd4\xd6\x70\x40\x4e\xb4\x4a\x61\x55\x13\x42\xc6\xc2\xba\x46\xc5\xcf\x43\x90\xcc\xec\x5e\x48\x2d\x05\xca\x30\x39\xe7\xf9\xb2\x0d\x6f\x24\x96\x3f\x56\x49\x28\xd1\xb7\xb6\xcd\xdb\x5e\xdb\x76\xc1\x43\x6c\x9d\x62\x86\x8d\x2d\xf3\xb6\xd7\x96\x11\x79\x90\x6f\x7f\x7d\x46\x8e\xfe\x5a\xfd\x0b\x94\x49\xe4\x69\x18\xe3\xb6\x9c\xaa\x4a\x4e\x84\x0a\xba\x33\xaa\x2b\x42\xaf\x18\x84\xe2\xe8\xb3\x3a\x46\x6e\xb8\x49\x1f\x3f\x8f\xee\xb4\xd5\xce\x39\x12\x80\xdf\x1c\xf0\xbb\x9c\x98\x21\xd3\x2b\x03\x98\xcd\xe0\x78\x80\xb1\x3a\x5c\xa8\x8a\x53\xd1\x60\xa7\x70\xcc\xa4\x50\xe1\xae\x26\x3b\x25\xa2\xc7\xcc\x71\x47\xce\x29\x66\x89\xf0\x03\xed\xa4\x5a\x98\xe6\x1c\x3e\x4d\x21\x8f\xee\xd4\x72\x1e\xb2\x39\x82\xf6\xa9\xc1\xa4\x75\xcf\xf4\x0f\xf9\x33\xcc\x0a\x14\x0a\x7b\xf5\xd8\xa8\x6a\x13\x86\x82\xab\x61\x2a\x34\xbe\x5a\x3c\x26\xaa\xe7\x94\x25\x57\x0c\x5f\x6d\xde\x71\xba\x0c\xf9\xe6\x77\xdc\xf8\xa9\xc6\x0c\x37\xb7\xa6\x9c\xa6\xb0\xc0\x8d\xb9\xea\xd5\x32\xe4\xe3\x66\x85\xe4\x75\xae\x7e\x4c\xf4\x3b\xfb\xd1\xd9\x02\x93\x3f\x55\xbc\x4c\x0e\x38\x49\x71\x76\xef\x7b\x1d\x03\x86\x14\xd4\x89\xaa\xed\x07\xc1\x2e\x3b\x5a\xec\x16\xef\x14\x6c\x52\x1d\x8f\xf5\xfb\xbc\xc8\xb2\xae\xef\x3f\xb8\xab\x3a\x4f\xbf\xd5\x4f\xd7\x41\xd1\xbb\xbb\xea\x75\xaf\xd7\x53\xc8\x79\x82\x5c\x57\x5e\x12\x91\x2b\xf5\x12\x80\x7f\x73\xfb\xc8\xb1\x10\x43\xc1\xd8\xd9\x74\xfd\xd7\x80\x1c\x49\x28\x55\x0a\xb4\xc3\xf3\xef\x8c\x46\x2b\x2f\x1e\x16\x8a\xa6\x92\x9c\xdc\x30\x90\x2b\x52\x9f\x42\x9e\xa6\x02\x25\x14\x94\xc9\x5f\xfe\x3f\x10\x86\xfd\x41\xd8\xcb\xaa\x96\xa9\x0d\x8c\x1d\xae\xb6\xe8\xec\x5a\x97\x60\x5f\x66\x99\xdf\xe2\xc4\x1d\x7a\xb6\x27\xbd\x51\xa1\x4f\xfa\x37\x74\x49\xa5\xff\xbc\xbe\xd0\x6a\x57\xed\xe1\xb5\x73\x69\x73\x29\xd9\x39\xc1\xaa\xb0\x98\x97\x93\xfa\x34\x73\x90\xd0\xb4\x0a\xda\x6f\x2d\x67\x5c\x8d\x57\x5a\xc0\x37\x72\xad\xd5\xd5\x69\xa1\x42\xeb\x9c\x17\x26\xec\x8d\x32\x65\x44\x8d\x91\xd3\xb3\xeb\x13\x67\x7c\xc7\x90\x12\x3a\x45\x11\xfb\x46\xdc\xe4\x57\xd0\x92\x8f\x73\x26\x29\x2b\xb0\x1e\xb4\x50\x7a\x75\xbd\xec\x55\x55\x1d\x4c\xf9\x5a\x61\xd6\x07\xd3\xfb\x02\xf9\xe6\x43\xbe\x76\x0e\x27\x93\xc8\x42\x89\x3c\x69\xa7\xcc\xb6\x74\xbb\x0c\xe5\x36\x0a\x72\x1d\x87\xcc\xe7\xf9\x7a\x0a\xcd\x26\xfe\xfa\xd0\x36\x84\xa3\x18\xdb\x85\x38\x6c\x73\xb0\x30\xa6\x90\xa9\xac\xda\x5f\x1f\x87\x68\xe2\x47\xa9\x90\x1f\xba\x30\x54\x1c\x74\xa8\x87\x17\x9b\x02\xd7\x52\x3f\x51\x4d\x99\x82\x59\x86\x0b\xec\x4b\xa4\xe3\xa0\x2e\x3d\x51\xf7\xf1\x4d\x01\x76\xaa\x8f\xa6\x2a\x83\x2f\xc4\x19\xe7\x97\xf9\x87\x7c\x2d\x7c\xe4\x75\x1a\xf6\x55\x4b\x7b\x03\xbe\xb5\xf6\x12\x4c\x91\x6b\x68\xe4\x24\xcb\x05\x56\x39\xad\x36\x44\x8f\x5e\xe2\x57\xe9\xbb\xf6\xed\x2d\xb9\x9f\x1e\xda\x14\xa1\x6e\xc8\x15\x45\x88\x11\x1c\x31\x84\xb5\xb3\x3b\x2a\xde\x33\x08\x57\x2b\x64\x89\xaf\xe3\x90\x47\x77\x7d\xd7\xef\x6f\x60\x94\x93\xbc\x60\xd2\x6f\x33\x48\x00\xbe\xe5\x8c\xf1\x5c\x30\x54\xc2\x5e\xac\x0d\x3c\x0f\xbc\xc7\xae\xe1\x11\x0c\xff\x25\xe4\xa0\xf1\x80\xf6\xcf\x78\x65\x59\x3d\x5f\x9b\x2d\x7b\xa2\x05\x82\xce\x87\xa2\x82\x99\x46\x6e\x44\x44\x4d\x2f\x6d\xf8\x4e\xf5\x4a\xdd\xaf\x8c\x55\x3e\xb8\xdf\x08\x7a\xe3\xfd\x45\xf7\x58\xa6\xc2\x56\x37\x46\xf4\xd6\x69\x70\x6d\xea\x59\x0a\xd1\x4a\x1b\x0a\xb1\x36\x9a\xc0\x1a\x6d\x37\x7a\xfc\x16\x4c\x8e\xea\x2e\xce\xb0\x6c\x2b\x94\xab\x45\x95\xc5\x4e\x37\x62\x43\x72\xf8\xe2\xb8\x5a\x4c\x1d\xed\xab\x45\xd0\x77\x53\xb4\x5f\x1c\xec\x35\xd1\x80\x1b\x73\x3b\x74\xd7\x39\x47\x9e\x8d\x55\x7f\xa0\x06\xe2\x6b\x8f\x2f\x6b\x7b\xa0\x19\x3f\xfe\x87\xad\x78\xf5\x29\x86\x5c\xa3\x7c\x1b\xae\x2a\x6b\x8f\x5d\x0b\xd8\x61\xdc\x4e\x2f\x3e\x82\x34\x8f\xf7\x5c\x57\x88\xe2\xe7\x97\x69\x8a\xb1\xc4\x64\xd4\x87\x3b\xf3\xd1\x67\xb8\x28\x06\xb6\xe8\xfb\x27\x62\x37\x0f\xed\xe7\x2b\xad\x6c\xbc\x23\x3d\x6d\x1a\x21\xe4\x50\xaf\x3e\xec\xa1\x7b\x45\x7a\xa8\xa7\xb6\x27\xfb\x3e\x0e\x3e\xc6\x89\x50\x7d\x29\xfc\xd9\x72\xdf\xf9\x2f\x93\xbf\x03\x00\x00\xff\xff\x42\xf6\x54\x5c\x09\x1d\x00\x00")

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

	info := bindataFileInfo{name: "generator/tmpl/sqlmgr.tmpl", size: 7433, mode: os.FileMode(420), modTime: time.Unix(1488291231, 0)}
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

	info := bindataFileInfo{name: "generator/tmpl/struct.tmpl", size: 3408, mode: os.FileMode(420), modTime: time.Unix(1488291231, 0)}
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

