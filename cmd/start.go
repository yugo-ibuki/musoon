package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/musoon/internal/browser"
	"os"
)

var id *string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start is a command for a quick start for a music you want to listen to.",
	Long: `start is a command for a quick start for a music you want to listen to.
You don't need to open the browser and search for the music you want to listen to.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(*id) == 0 {
			fmt.Println("can not parse command line argument (--id)")
			os.Exit(1)
		}
		os.Exit(start(cmd, args))
	},
}

func init() {
	id = startCmd.Flags().StringP("id", "i", "", "Specify the id of the music you want to listen to")
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) int {
	id, err := cmd.Flags().GetString("id")
	if err != nil {
		fmt.Print("no id specified")
		return 1
	}

	// open the browser
	brws := browser.NewBrowser()
	if err := brws.Open(id); err != nil {
		fmt.Print(err)
		return 1
	}

	return 0
}
