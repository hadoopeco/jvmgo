package classpath

/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/10 17:13
 */
import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absDir}
}

// iterate jar or zip to find classfile
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {

	r, err := zip.OpenReader(self.absDir)

	// return if read zip error
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	//find class file from jar or zip
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()

			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil

		}
	}

	return nil, nil, errors.New("class not found = " + className)
}

func (self *ZipEntry) String() string {
	return self.absDir
}
