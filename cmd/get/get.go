package get

import (
  "prestctl/cmd"
  "github.com/spf13/cobra"
)

var (
    carrierId int
    customerId int
    categoryId int
    orderId int
    shopUrl string
    token string
    method string
    apiSuffix string
    productId int 
    jsonSuffix string
    completeUrl string
    encodedToken string
	  objectType string
    responseBytes []byte
	  finalStringSlice []string
)

// getCmd represents the get command
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a list of a ressource from the prestahsop website",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
       cmd.Help()
	},
}

func init() {
  cmd.RootCmd.AddCommand(GetCmd)
}
