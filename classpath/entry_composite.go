package classpath
/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/10 19:41
 */
import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparater) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)

	}
	return compositeEntry

}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)

		if err == nil {
			return data, from, err
		}
	}

	return nil, nil, errors.New("class not founed:" + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparater)
}
