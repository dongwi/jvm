package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag	bool
	VersionFlag	bool
	CpOption	string
	Class		string
	Args		[]string
}

func PrintUsag()  {
	fmt.Printf("Usage: %s [-optioins] class [args...]\n", os.Args[0])
}

func ParseCmd()  *Cmd{
	cmd := &Cmd{}
	flag.Usage = PrintUsag
	flag.BoolVar(&cmd.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.CpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.CpOption, "cp", "", "classpath")
	flag.Parse()

	if args := flag.Args(); len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}
	return cmd
}
