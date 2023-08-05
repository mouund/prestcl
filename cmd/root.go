package cmd

import (
	"os"
	"github.com/spf13/cobra"
)



// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "prestctl",
	Short: "A simple prestashop CLI to administer your instance",
	Long: ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.prestctl.yaml)")

}


