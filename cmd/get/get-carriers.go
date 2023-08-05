package get



import (
  "github.com/spf13/cobra"
  "encoding/json"
  "prestctl/utils"
  "prestctl/controllers"
  "fmt"
)

// getCarriersCmd represents the getCarriers command
var getCarriersCmd = &cobra.Command{
	Use:   "carriers",
	Short: "Get a list of carriers of the prestashop instance",
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
    apiSuffix, responseBytes = controllers.GetAllIdsByType(encodedToken, shopUrl, "carriers")
    var allcarriers controllers.Allcarriers
    json.Unmarshal(responseBytes, &allcarriers)
	  finalStringSlice = controllers.GetAllDetails(allcarriers, "carriers", shopUrl,  encodedToken)
    controllers.PrintSlice(finalStringSlice)
	},
}

func init() {
	GetCmd.AddCommand(getCarriersCmd)
}
