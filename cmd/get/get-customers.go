package get

import (
	"github.com/spf13/cobra"
	"encoding/json"
	"prestctl/utils"
	"prestctl/controllers"
	"fmt"
)

// getCustomersCmd represents the getcustomers command
var getCustomersCmd = &cobra.Command{
	Use:   "customers",
	Short: "Get a list of customers of the prestasop instance",
	Long: ``,
	PreRun: func(cmd *cobra.Command, args []string) {
		config, error := utils.LoadConfig("prestctl.yaml")
		if error != nil {
		  fmt.Println("Error Loading config" , error)
		}
		shopUrl = config.ShopURL
		token = config.Token
			if  shopUrl == "" || token == "" {
				//definition of shopUrl flag
				cmd.Flags().StringVarP(&shopUrl, "shopUrl", "u", "", "Url of the shop")
				cmd.MarkFlagRequired("shopUrl")
				// definition of token flag
				cmd.Flags().StringVarP(&token, "token", "t", "", "Token used to authentify to the shop")
				cmd.MarkFlagRequired("token")
	
			}
		},
	Run: func(cmd *cobra.Command, args []string) {

	//Iitialisation of variables
    //encoding token + :
    encodedToken = controllers.Encode64(token + ":")
    apiSuffix, responseBytes = controllers.GetAllIdsByType(encodedToken, shopUrl, "customers",)
    var allcustomers controllers.Allcustomers
    json.Unmarshal(responseBytes, &allcustomers)
	finalStringSlice = controllers.GetAllDetails(allcustomers, "customers", shopUrl, encodedToken)
    controllers.PrintSlice(finalStringSlice)

	},
}

func init() {
	GetCmd.AddCommand(getCustomersCmd)
	
}
