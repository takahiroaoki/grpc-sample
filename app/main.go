package main

import (
	"fmt"

	"github.com/takahiroaoki/grpc-sample/app/cmd"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

func main() {
	if err := cmd.NewBundle().Execute(); err != nil {
		util.FatalLog(fmt.Sprintf("Failed to execute the command. Error: %v", err))
	}
}
