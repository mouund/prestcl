package describe

import (

	"github.com/spf13/cobra"
	"prestctl/utils"
	"prestctl/controllers"
	"fmt"
	"encoding/json"
	"gopkg.in/yaml.v2"
  )
  
  // describeOrderCmd represents the order command
  var describeOrderCmd = &cobra.Command{
	  Use:   "order",
	  Short: "Describe a order from the prestashop instance",
	  Long: `Describe a order from the prestashop instance
  
  Required flags 
  --token
  --shopUrl `,
	Example: `- prestctl describe object --shopUrl <shop URL> --token <API token>
  - prestctl describe object --shopUrl http://my-shop.com --token fazfazflpfl`,
	PreRun: func(cmd *cobra.Command, args []string) {
	config, error := utils.LoadConfig("prestctl.yaml")
	if error != nil {
	  fmt.Println("Error Loading config, either use config file or use flags --shopUrl and --token" , error)
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
		   responseBytes := controllers.DescribeObjectFromId (encodedToken, shopUrl, "orders", id)
	   var order controllers.Order
		   json.Unmarshal(responseBytes, &order)
	   describedObject, error := yaml.Marshal(order)
		if error != nil {
		  fmt.Println(error)
		}
  
	  fmt.Printf("%s", describedObject)
  
	  },
  }
  
  
  func init() {
	  DescribeCmd.AddCommand(describeOrderCmd)
  
		 // definition of id flag
	   describeOrderCmd.Flags().IntVarP(&id, "id", "i",0, "id of the resource to be described")
	   describeOrderCmd.MarkFlagRequired("id")
  
  }
  