package main

import (
	"fmt"

	"github.com/takahiroaoki/go-env/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Failed to execute the command. Error: %v", err)
	}
}
