package delete

import (
	"github.com/spf13/cobra"
	"fmt"
	"prestctl/utils"
	"prestctl/controllers"
)

// deleteOrderCmd represents the deleteOrder command
var deleteOrderCmd = &cobra.Command{
	Use:   "order",
	Short: "Delete an order from the prestashop instance",
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
		 responseStatus = controllers.DeleteObjectFromId (encodedToken, shopUrl, "orders", id)
		 if responseStatus == "200 OK" {
			fmt.Println("Order deleted successfully")
		 }else{
			fmt.Println("Eror: " + responseStatus )
		 }
	},
}

func init() {
	DeleteCmd.AddCommand(deleteOrderCmd)

	 // definition of id flag
	 deleteOrderCmd.Flags().IntVarP(&id, "id", "i",0, "id of the resource to be deleted")
	 deleteOrderCmd.MarkFlagRequired("id")
}
