package main

import (
	"../cmd"
	"fmt"
)

func startJvm(command *cmd.Cmd)  {
	fmt.Printf("classpath:%s class:%s args:%v\n", command.CpOption,command.Class, command.Args)
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
