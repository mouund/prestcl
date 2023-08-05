package get

import (

  "github.com/spf13/cobra"
  "encoding/json"
  "prestctl/utils"
  "prestctl/controllers"
  "fmt"
)



// getProductsCmd represents the products command
var getProductsCmd = &cobra.Command{
	Use:   "products",
	Short: "Get a list of products from your prestashop instance",
	Long: `Get a list of products on the shop

Required flags 
--token
--shopUrl `,
  Example: `- prestctl get object --shopUrl <shop URL> --token <API token>
- prestctl get object --shopUrl http://my-shop.com --token fazfazflpfl`,
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
    //encoding token + :
    encodedToken = controllers.Encode64(token + ":")
    apiSuffix, responseBytes = controllers.GetAllIdsByType(encodedToken, shopUrl, "products")
    var allproducts controllers.Allproducts
    json.Unmarshal(responseBytes, &allproducts)
    finalStringSlice = controllers.GetAllDetails (allproducts,"products",shopUrl, encodedToken)
    controllers.PrintSlice(finalStringSlice)
	},
}


func init() {
	GetCmd.AddCommand(getProductsCmd)

}
