package main

import (
	"fmt"
	"reflect"
)
import "strings"
import "jvmgo/ch02/classpath"

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
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, enty, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Println("加载类的类型:" + reflect.TypeOf(enty).String())
	fmt.Println("类的文件路径:" + enty.String())
	fmt.Printf("class data:%v\n", classData)
}

//go run . -Xjre="/Users/zxj/go/src/Learn/jvmgo-book/jre" java.lang.Object
