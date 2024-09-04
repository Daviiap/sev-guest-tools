package main

import (
	"fmt"
	"os"
	"sev-guest/src/commands"
	"strconv"
)

func printUsage() {
	fmt.Println("sevtool --[command] [command opts]")
	fmt.Println("commands:")
	fmt.Println("  --get_report")
	fmt.Println("    Input params:")
	fmt.Println("      data file name - name of the file containing the data to be included into the report")
	fmt.Println("  --read_report")
	fmt.Println("    Input params:")
	fmt.Println("      filename - name of the report binary file to read")
}

type commandsOpts struct {
	PrintUsage        bool
	GetReport         bool
	GetExtendedReport bool
	GetReportOptions  commands.GetReportOptions
	ReadReport        bool
	ReadReportOptions commands.ReadReportOptions
}

func isValidIndex(index int, length int) bool {
	return length > index
}

func parseOptions(cmdOpts *commandsOpts) {
	args := os.Args[1:]

	i := 0

	for _, s := range args {
		switch s {
		case "--get_report":
			cmdOpts.GetReport = true
			cmdOpts.GetReportOptions.Filename = "report.bin"

			if isValidIndex(i+1, len(args)) {
				cmdOpts.GetReportOptions.DataFileName = args[i+1]
			}
		case "--vmpl":
			if isValidIndex(i+1, len(args)) {
				vmpl, err := strconv.Atoi(args[i+1])
				if err != nil {
					panic(err)
				}
				cmdOpts.GetReportOptions.VMPL = vmpl
			}
		case "--get_extended_report":
			cmdOpts.GetExtendedReport = true
			cmdOpts.GetReportOptions.Filename = "report.bin"

			if isValidIndex(i+1, len(args)) {
				cmdOpts.GetReportOptions.DataFileName = args[i+1]
			}
		case "--read_report":
			cmdOpts.ReadReport = true

			if isValidIndex(i+1, len(args)) {
				cmdOpts.ReadReportOptions.Filename = args[i+1]
			} else {
				panic("Invalid argument")
			}
		case "--help":
			cmdOpts.PrintUsage = true
		}

		i++
	}
}

func main() {
	cmds := commandsOpts{}
	parseOptions(&cmds)

	if cmds.PrintUsage {
		printUsage()
	} else if cmds.GetReport {
		commands.GetReportCommand(cmds.GetReportOptions)
	} else if cmds.GetExtendedReport {
		commands.GetExtendedReportCommand(cmds.GetReportOptions)
	} else if cmds.ReadReport {
		commands.ReadReportCommand(cmds.ReadReportOptions)
	}
}
