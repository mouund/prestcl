package get
import (
  "github.com/spf13/cobra"
  "encoding/json"
  "prestctl/utils"
  "prestctl/controllers"
  "fmt"
)

// getOrdersCmd represents the orders command
var getOrdersCmd = &cobra.Command{
	Use:   "orders",
	Short: "Get a list of orders of the prestashop instance",
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
    apiSuffix, responseBytes = controllers.GetAllIdsByType(encodedToken, shopUrl, "orders",)
    var allorders controllers.Allorders
    json.Unmarshal(responseBytes, &allorders)
	  finalStringSlice = controllers.GetAllDetails(allorders, "orders", shopUrl, encodedToken)
    controllers.PrintSlice(finalStringSlice)
	},
}

func init() {
	GetCmd.AddCommand(getOrdersCmd)

}
