package main

import (
	"fmt"

	"github.com/takahiroaoki/go-env/cmd"
	"github.com/takahiroaoki/go-env/util"
)

func main() {
	rootCmd := cmd.NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		util.FatalLog(fmt.Sprintf("Failed to execute the command. Error: %v", err))
	}
}
