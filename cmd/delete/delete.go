package delete

import (
  "prestctl/cmd"
	"github.com/spf13/cobra"
)

var (
    responseStatus string
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
    responseBytes []byte
	  finalStringSlice []string
)

// DeleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a ressource from the prestashop website",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
  cmd.RootCmd.AddCommand(DeleteCmd)
}
