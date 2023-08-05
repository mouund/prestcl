package get

import (
  "github.com/spf13/cobra"
  "encoding/json"
  "prestctl/utils"
  "prestctl/controllers"
  "fmt"
)

// getcategoriescmd represents the getCategories command
var getCategoriesCmd = &cobra.Command{
	Use:   "categories",
	Short: "Get a list of categories of the prestashop instance",
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
    apiSuffix, responseBytes = controllers.GetAllIdsByType(encodedToken, shopUrl, "categories",)
    var allcategories controllers.Allcategories
    json.Unmarshal(responseBytes, &allcategories)
	  finalStringSlice = controllers.GetAllDetails(allcategories, "categories", shopUrl, encodedToken)
    controllers.PrintSlice(finalStringSlice)
	},
}

func init() {
	GetCmd.AddCommand(getCategoriesCmd)
}  