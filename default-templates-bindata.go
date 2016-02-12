// Code generated by go-bindata.
// sources:
// templates/default.html
// templates/detailed.html
// DO NOT EDIT!

package main

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

var _templatesDefaultHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\x8e\xc1\x0e\x82\x30\x10\x44\xcf\xf8\x15\xeb\x0f\x40\xb8\x9a\x86\x8b\x57\x63\xfc\x05\x68\x57\xa9\x81\x5d\x52\xb6\x07\x42\xf8\x77\xb7\xad\x5e\x3c\x4d\x3b\xd3\x79\x1d\x73\x76\x6c\x65\x5b\x10\x46\x99\xa7\xee\x64\x8a\x54\x66\xc4\xde\x25\x6d\x7e\x87\x81\xdd\xa6\xaa\x49\xdb\xdd\xfc\x2a\xc0\x4f\xe0\x05\x09\x56\x8e\xc1\x22\x58\x26\x09\x7e\x88\xe2\x99\x56\xad\xb5\xf9\x71\x4c\xb0\x6a\xdf\x21\xf4\xf4\x42\xa8\xe1\x38\x92\x3d\xf9\x4e\xbd\xfa\x11\xf8\x8d\x56\xea\x7b\x3f\xa3\x26\x17\x48\xe6\x95\x23\x89\xde\xfe\x89\xda\x29\x28\x24\xf7\xc5\x34\x19\x6f\x9a\xb2\x4d\x3f\xcd\xe3\x3f\x01\x00\x00\xff\xff\x50\x3e\x7e\x46\xd4\x00\x00\x00")

func templatesDefaultHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDefaultHtml,
		"templates/default.html",
	)
}

func templatesDefaultHtml() (*asset, error) {
	bytes, err := templatesDefaultHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/default.html", size: 212, mode: os.FileMode(420), modTime: time.Unix(1455236630, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesDetailedHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x6c\x90\x4d\x6e\x84\x30\x0c\x85\xd7\xed\x29\xdc\x39\x00\xd1\xec\xd3\x6c\xda\xe5\xa8\xaa\x2a\xf5\x00\xf9\x31\x25\xd5\x4c\x8c\x92\xb0\x18\x21\xee\x5e\x87\x14\xc4\x20\x56\x7e\x3c\x7f\xc6\xf1\x93\x2f\x8e\x6c\xbe\xf7\x08\x5d\xbe\x5d\xd5\xb3\xac\xe5\x49\x76\xa8\x5d\xa9\x62\x11\x86\xdc\x9d\x2b\x77\xce\xea\xe2\x53\x06\x6a\x81\x7a\x0c\x90\x68\x88\x16\xc1\x52\xc8\xd1\x9b\x21\x7b\x0a\x89\xc7\xce\x05\x1e\x47\x88\x3a\xfc\x20\x34\x30\x4d\xf5\xdb\xb7\xd0\x7c\x46\xfa\x45\x9b\x9b\xef\xaf\x4b\xf5\xa5\x51\x52\x43\x17\xb1\x7d\x3d\x31\xb3\x03\x4e\x6a\xeb\x7d\xe8\x1b\xb2\x29\x85\x56\x52\x98\xff\x2d\x78\x4d\xb8\xfe\xea\x90\x5e\xc9\xe0\x2a\x58\xd0\x28\xd4\xac\xb6\x13\xef\x98\x6c\xf4\x7d\xb9\xe3\x18\x7c\xa3\x21\x64\x6e\x3d\x9e\xbc\x03\x17\xb5\xd9\xc8\x6f\x98\x33\xe4\x70\xe6\x90\xff\x02\x00\x00\xff\xff\x7d\x2c\x3b\xfc\x7c\x01\x00\x00")

func templatesDetailedHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDetailedHtml,
		"templates/detailed.html",
	)
}

func templatesDetailedHtml() (*asset, error) {
	bytes, err := templatesDetailedHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/detailed.html", size: 380, mode: os.FileMode(420), modTime: time.Unix(1455236623, 0)}
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
	"templates/default.html": templatesDefaultHtml,
	"templates/detailed.html": templatesDetailedHtml,
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
	"templates": &bintree{nil, map[string]*bintree{
		"default.html": &bintree{templatesDefaultHtml, map[string]*bintree{}},
		"detailed.html": &bintree{templatesDetailedHtml, map[string]*bintree{}},
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

