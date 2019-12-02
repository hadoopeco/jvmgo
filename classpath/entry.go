package classpath
/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/10 14:13
 */
import (
	"os"
	"strings"
	"fmt"
)

//
const pathListSeparater = string(os.PathListSeparator)

type Entry interface {
	// className: fully/qualified/ClassName.class
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparater) {
		return newCompositeEntry(path)
	}

	// handle the path end with *
	if strings.HasSuffix(path, "*") {
		fmt.Println("return  wildcard entry")
		return newWildcardEntry(path)
	}
	// handle the path end with .zip and .jar
	if strings.HasSuffix(strings.ToLower(path), ".zip") || strings.HasSuffix(strings.ToLower(path), ".jar") {
		fmt.Println("return  zip entry")
		return newZipEntry(path)
	}

	return newDirEntry(path)

}
