package classpath
/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/10 16:21
 */
import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	//remove the lastest char "*"
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(strings.ToLower(path), ".jar") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)

		}

		return nil
	}
	filepath.Walk(baseDir, walkFn)

	return compositeEntry

}
