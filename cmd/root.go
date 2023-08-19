package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "musoon",
	Short: "musoon is a tool for opening youtube to listen to music from the command line according to the setting.",
}

func exitError(msg interface{}) {
	fmt.Print(os.Stderr, msg)
	os.Exit(1)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exitError(err)
	}
}
