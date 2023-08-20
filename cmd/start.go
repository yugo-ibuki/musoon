package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yugo-ibuki/musoon/internal/browser"
	"github.com/yugo-ibuki/musoon/internal/config"
	"os"
)

type flag struct {
	id         *string
	configPath *string
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start is a command for a quick start for a music you want to listen to.",
	Long: `start is a command for a quick start for a music you want to listen to.
You don't need to open the browser and search for the music you want to listen to.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := start(cmd, args); err != nil {
			fmt.Println("can not parse command line argument (--id)")
			os.Exit(1)
		}
	},
}

func init() {
	startCmd.Flags().StringP("id", "i", "", "Specify the id of the music you want to listen to")
	startCmd.Flags().StringP("configPath", "c", "", "If you have the config file(toml), you can specify it.")
	rootCmd.AddCommand(startCmd)
}

func parseArgs(cmd *cobra.Command, _ ...[]string) (*flag, error) {
	flg := flag{}

	// if id is specified, it doesn't read the config file.
	id, err := cmd.Flags().GetString("id")
	if err != nil {
		return nil, fmt.Errorf("can not parse command line argument (--id)")
	}
	flg.id = &id

	// if the both of id and configPath are not specified, you may need to specify id or configPath.
	configPath, err := cmd.Flags().GetString("configPath")
	if err != nil {
		return nil, fmt.Errorf("can not parse command line argument (--config)")
	}
	flg.configPath = &configPath

	if *flg.id == "" && *flg.configPath == "" {
		return nil, fmt.Errorf("can not parse command line argument")
	}

	return &flg, nil
}

func start(cmd *cobra.Command, args []string) error {
	flag, err := parseArgs(cmd, args)
	if err != nil {
		return err
	}

	brws := browser.NewBrowser()

	if len(*flag.id) != 0 {
		// if id is specified directly, it sets id to the browser.
		if err := brws.Open(*flag.id); err != nil {
			fmt.Print(err)
			return err
		}
	} else {
		// if configPath is specified, it reads the config file and sets id to the browser.
		conf := config.NewConfig()
		content, err := conf.Read(*flag.configPath)
		if err != nil {
			return err
		}

		if err := brws.Open(content.ID); err != nil {
			fmt.Print(err)
			return err
		}
	}

	return nil
}
