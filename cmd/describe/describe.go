package describe

import (
  "prestctl/cmd"
  "github.com/spf13/cobra"
)

var (
	id int
    shopUrl string
    token string
    method string
    apiSuffix string
    productId int 
    jsonSuffix string
    completeUrl string
    encodedToken string
	  objectType string
	  finalStringSlice []string
)

// DescribeCmd represents the describe command
var DescribeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe a ressource from the prestahsop website",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
       cmd.Help()
	},
}

func init() {
  cmd.RootCmd.AddCommand(DescribeCmd)
}
