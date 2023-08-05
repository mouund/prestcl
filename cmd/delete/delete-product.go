package delete

import (
	"github.com/spf13/cobra"
	"fmt"
	"prestctl/utils"
	"prestctl/controllers"
)

// deleteProductCmd represents the deleteProduct command
var deleteProductCmd = &cobra.Command{
	Use:   "product",
	Short: "Delete a product from the prestashop instance",
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
		 //encoding token + :
		 encodedToken = controllers.Encode64(token + ":")
		 responseStatus = controllers.DeleteObjectFromId (encodedToken, shopUrl, "products", id)
		 if responseStatus == "200 OK" {
			fmt.Println("Product deleted successfully")
		 }else{
			fmt.Println("Eror: " + responseStatus )
		 }
	},
}

func init() {
	DeleteCmd.AddCommand(deleteProductCmd)

	 // definition of id flag
	 deleteProductCmd.Flags().IntVarP(&id, "id", "i",0, "id of the resource to be deleted")
	 deleteProductCmd.MarkFlagRequired("id")
}
