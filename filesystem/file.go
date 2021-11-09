package filesystem

import (
	"bytes"
	"github.com/gobuffalo/packr/v2"
	"github.com/meshplus/hyperbench/common"
	"github.com/pingcap/failpoint"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

//go:generate rm -rf assets ; mkdir assets ; cp -r golua assets; cp -r config assets ; cp -r benchmark assets

//AssetsPath asset path
const AssetsPath = "assets"

//FileSystem the folder on a disk
var FileSystem = packr.New(AssetsPath, AssetsPath)
var logger = common.GetLogger("filesystem")

//Unpack unpack the path with prefix
func Unpack(prefix string) error {

	// to unpack some of files into the real file system
	var err error

	prefix = strings.TrimPrefix(prefix, "./")

	// unpack the path with prefix "lua"
	err = FileSystem.WalkPrefix(prefix, walk)
	if _, ok := failpoint.Eval(_curpkg_("unpack-err")); ok {
		err = errors.New("unpack-err")
	}
	if err != nil {
		logger.Errorf("can not unpack: %v", err)
	}
	return nil
}

func walk(name string, file packr.File) error {
	var err error
	buf := new(bytes.Buffer)

	_, err = buf.ReadFrom(file)

	if _, ok := failpoint.Eval(_curpkg_("walk-buf-err")); ok {
		err = errors.New("walk-buf-err")
	}

	if err != nil {
		logger.Errorf("can not read file(%v) from box: %v", name, err)
		return err
	}
	if _, err = EnsureFileExist(name, buf.Bytes()); err != nil {
		return err
	}
	return nil
}

//EnsureFileExist ensure file exist and file content is the specify content
func EnsureFileExist(path string, content []byte) (bool, error) {
	var f *os.File
	var err error
	var dir string

	if _, err = os.Stat(path); os.IsNotExist(err) {

		logger.Debugf("path (%v) does not exist, ready to create", path)
		// ensure path exit
		dir, _ = filepath.Split(path)

		if _, err = os.Stat(dir); os.IsNotExist(err) {
			// dir does not exist, then make it
			if err = os.MkdirAll(dir, os.ModePerm); err != nil {
				logger.Errorf("can not mkdir (%v): %v", dir, err)
				return false, err
			}
		}

		// create file
		f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
		if _, ok := failpoint.Eval(_curpkg_("openfile-err")); ok {
			err = errors.New("openfile-err")
		}
		if err != nil {
			logger.Errorf("can not open file (%v): %v", path, err)
			return false, err
		}

		// write file
		_, err = f.Write(content)
		if _, ok := failpoint.Eval(_curpkg_("writefile-err")); ok {
			err = errors.New("writefile-err")
		}
		if err != nil {
			logger.Errorf("can not write file (%v): %v", path, err)
			return false, err
		}

		// close file
		err = f.Close()
		if _, ok := failpoint.Eval(_curpkg_("closefile-err")); ok {
			err = errors.New("closefile-err")
		}
		if err != nil {
			logger.Errorf("can not close file (%v): %v", path, err)
			return false, err
		}
	}
	// todo: check file content
	return true, nil
}
