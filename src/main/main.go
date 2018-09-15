package main

import (
	"../cmd"
	"fmt"
	"../classpath"
	"strings"
)

func startJvm(command *cmd.Cmd)  {
	cp := classpath.Parse(command.XjreOption, command.CpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, command.Class, command.Args)
	className := strings.Replace(command.Class, ".", "/", -1)
	classDate, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", command.Class)
		return
	}
	fmt.Printf("class data:%v\n", classDate)
}

func main()  {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if command.HelpFlag || command.Class == "" {
		cmd.PrintUsag();
	} else {
		startJvm(command)
	}
}
