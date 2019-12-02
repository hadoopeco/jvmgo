package classpath
/**
 * Copyright (C) 2018
 * All rights reserved
 *
 * @author: mark.wei
 * @mail: marks@126.com
 * Date: 2018/3/10 21:08
 */
import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath  Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data, entry, err := self.bootClasspath.readClass(className); err == nil{
		return data, entry, err
	}

	if data, entry, err := self.extClasspath.readClass(className); err == nil{
		return data, entry, err
	}

	return self.userClasspath.readClass(className)

}


func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// /jre/lib/*
	jrelibDir := filepath.Join(jreDir, "lib","*")
	self.bootClasspath = newWildcardEntry(jrelibDir)

	//  /jre/lib/ext/*
	jreExtDir := filepath.Join(jreDir,"lib", "ext","*")
	self.extClasspath = newWildcardEntry(jreExtDir)
}

func (self *Classpath) parseUserClasspath(cpOption string)  {
	if cpOption  == ""{
		cpOption ="."
	}

	self.userClasspath = newEntry(cpOption)
}

// find the jre directory address
func getJreDir(jreOption string) string {

	if jreOption != "" && exists(jreOption) {
		return jreOption + "/jre"
	}

	if exists("./jre") {
		return "./jre"
	}

	// get env JAVA_HOME
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("can't find jre")
}

// file exists or not
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}

	}
	return true
}


func (self *Classpath) String() string {
	return self.userClasspath.String()
}
