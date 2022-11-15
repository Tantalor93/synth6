package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use: "synth6",
}

func init() {
}

func Execute() {
	RootCmd.Execute()
}
