package main

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
	"jvmgo/rtdata/heap"
	"strings"
)

func main() {

	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	classLoader := heap.NewClassLoader(cp, false)

	className := strings.Replace(cmd.class, ".", "/", -1)
	class := classLoader.LoadClass(className)
	mainMethod := class.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, true)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

//func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
//	classData, _, err := cp.ReadClass(className)
//	if err != nil {
//		panic(err)
//	}
//
//	cf, err := classfile.Parse(classData)
//	if err != nil {
//		panic(err)
//	}
//
//	return cf
//}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("  %s\n", cf.ConstantPool())

	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SupperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterFaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
