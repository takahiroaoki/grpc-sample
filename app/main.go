package main

import (
	"fmt"

	"github.com/takahiroaoki/go-env/app/cmd"
	"github.com/takahiroaoki/go-env/app/util"
)

func main() {
	rootCmd := cmd.NewCmdRoot()
	if err := rootCmd.Execute(); err != nil {
		util.FatalLog(fmt.Sprintf("Failed to execute the command. Error: %v", err))
	}
}
