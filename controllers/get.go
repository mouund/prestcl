package controllers

import (
  "fmt"
  "encoding/json"
   "net/http"
   "io"
   "strconv"
   "github.com/microcosm-cc/bluemonday"
	"os"
)
//Variables for get package
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
//Structs for get package
type Allcarriers struct {
	Carriers []struct {
		ID int `json:"id"`
	} `json:"carriers"`
}
type Carrier struct {
	Details struct {
		ID                 int    `json:"id"`
		Deleted            string `json:"deleted"`
		IsModule           string `json:"is_module"`
		IDTaxRulesGroup    string `json:"id_tax_rules_group"`
		IDReference        string `json:"id_reference"`
		Name               string `json:"name"`
		Active             string `json:"active"`
		IsFree             string `json:"is_free"`
		URL                string `json:"url"`
		ShippingHandling   string `json:"shipping_handling"`
		ShippingExternal   string `json:"shipping_external"`
		RangeBehavior      string `json:"range_behavior"`
		ShippingMethod     int    `json:"shipping_method"`
		MaxWidth           string `json:"max_width"`
		MaxHeight          string `json:"max_height"`
		MaxDepth           string `json:"max_depth"`
		MaxWeight          string `json:"max_weight"`
		Grade              string `json:"grade"`
		ExternalModuleName string `json:"external_module_name"`
		NeedRange          string `json:"need_range"`
		Position           string `json:"position"`
		Delay              string `json:"delay"`
	} `json:"carrier"`
}

type Allcustomers struct {
	Customers []struct {
		ID int `json:"id"`
	} `json:"customers"`
}

type Customer struct {
	Details struct {
		ID                       int    `json:"id"`
		IDDefaultGroup           string `json:"id_default_group"`
		IDLang                   string `json:"id_lang"`
		NewsletterDateAdd        string `json:"newsletter_date_add"`
		IPRegistrationNewsletter string `json:"ip_registration_newsletter"`
		LastPasswdGen            string `json:"last_passwd_gen"`
		SecureKey                string `json:"secure_key"`
		Deleted                  string `json:"deleted"`
		Passwd                   string `json:"passwd"`
		Lastname                 string `json:"lastname"`
		Firstname                string `json:"firstname"`
		Email                    string `json:"email"`
		IDGender                 string `json:"id_gender"`
		Birthday                 string `json:"birthday"`
		Newsletter               string `json:"newsletter"`
		Optin                    string `json:"optin"`
		Website                  string `json:"website"`
		Company                  string `json:"company"`
		Siret                    string `json:"siret"`
		Ape                      string `json:"ape"`
		OutstandingAllowAmount   string `json:"outstanding_allow_amount"`
		ShowPublicPrices         string `json:"show_public_prices"`
		IDRisk                   string `json:"id_risk"`
		MaxPaymentDays           string `json:"max_payment_days"`
		Active                   string `json:"active"`
		Note                     string `json:"note"`
		IsGuest                  string `json:"is_guest"`
		IDShop                   string `json:"id_shop"`
		IDShopGroup              string `json:"id_shop_group"`
		DateAdd                  string `json:"date_add"`
		DateUpd                  string `json:"date_upd"`
		ResetPasswordToken       string `json:"reset_password_token"`
		ResetPasswordValidity    string `json:"reset_password_validity"`
		Associations             struct {
			Groups []struct {
				ID string `json:"id"`
			} `json:"groups"`
		} `json:"associations"`
	} `json:"customer"`
}

type Order struct {
	Details struct {
		ID                    int    `json:"id"`
		IDAddressDelivery     string `json:"id_address_delivery"`
		IDAddressInvoice      string `json:"id_address_invoice"`
		IDCart                string `json:"id_cart"`
		IDCurrency            string `json:"id_currency"`
		IDLang                string `json:"id_lang"`
		IDCustomer            string `json:"id_customer"`
		IDCarrier             string `json:"id_carrier"`
		CurrentState          string `json:"current_state"`
		Module                string `json:"module"`
		InvoiceNumber         string `json:"invoice_number"`
		InvoiceDate           string `json:"invoice_date"`
		DeliveryNumber        string `json:"delivery_number"`
		DeliveryDate          string `json:"delivery_date"`
		Valid                 string `json:"valid"`
		DateAdd               string `json:"date_add"`
		DateUpd               string `json:"date_upd"`
		ShippingNumber        string `json:"shipping_number"`
		IDShopGroup           string `json:"id_shop_group"`
		IDShop                string `json:"id_shop"`
		SecureKey             string `json:"secure_key"`
		Payment               string `json:"payment"`
		Recyclable            string `json:"recyclable"`
		Gift                  string `json:"gift"`
		GiftMessage           string `json:"gift_message"`
		MobileTheme           string `json:"mobile_theme"`
		TotalDiscounts        string `json:"total_discounts"`
		TotalDiscountsTaxIncl string `json:"total_discounts_tax_incl"`
		TotalDiscountsTaxExcl string `json:"total_discounts_tax_excl"`
		TotalPaid             string `json:"total_paid"`
		TotalPaidTaxIncl      string `json:"total_paid_tax_incl"`
		TotalPaidTaxExcl      string `json:"total_paid_tax_excl"`
		TotalPaidReal         string `json:"total_paid_real"`
		TotalProducts         string `json:"total_products"`
		TotalProductsWt       string `json:"total_products_wt"`
		TotalShipping         string `json:"total_shipping"`
		TotalShippingTaxIncl  string `json:"total_shipping_tax_incl"`
		TotalShippingTaxExcl  string `json:"total_shipping_tax_excl"`
		CarrierTaxRate        string `json:"carrier_tax_rate"`
		TotalWrapping         string `json:"total_wrapping"`
		TotalWrappingTaxIncl  string `json:"total_wrapping_tax_incl"`
		TotalWrappingTaxExcl  string `json:"total_wrapping_tax_excl"`
		RoundMode             string `json:"round_mode"`
		RoundType             string `json:"round_type"`
		ConversionRate        string `json:"conversion_rate"`
		Reference             string `json:"reference"`
		Associations          struct {
			OrderRows []struct {
				ID                 string `json:"id"`
				ProductID          string `json:"product_id"`
				ProductAttributeID string `json:"product_attribute_id"`
				ProductQuantity    string `json:"product_quantity"`
				ProductName        string `json:"product_name"`
				ProductReference   string `json:"product_reference"`
				ProductEan13       string `json:"product_ean13"`
				ProductIsbn        string `json:"product_isbn"`
				ProductUpc         string `json:"product_upc"`
				ProductPrice       string `json:"product_price"`
				IDCustomization    string `json:"id_customization"`
				UnitPriceTaxIncl   string `json:"unit_price_tax_incl"`
				UnitPriceTaxExcl   string `json:"unit_price_tax_excl"`
			} `json:"order_rows"`
		} `json:"associations"`
	} `json:"order"`
}


type Allorders struct {
	Orders []struct {
		ID int `json:"id"`
	} `json:"orders"`
}

type Allproducts struct {
	Products []struct {
		ID int `json:"id"`
	} `json:"products"`
}

type Product struct {
	Details struct {
		ID                      int    `json:"id"`
		IDManufacturer          string `json:"id_manufacturer"`
		IDSupplier              string `json:"id_supplier"`
		IDCategoryDefault       string `json:"id_category_default"`
		New                     any    `json:"new"`
		CacheDefaultAttribute   string `json:"cache_default_attribute"`
		IDDefaultImage          string `json:"id_default_image"`
		IDDefaultCombination    string `json:"id_default_combination"`
		IDTaxRulesGroup         string `json:"id_tax_rules_group"`
		PositionInCategory      string `json:"position_in_category"`
		ManufacturerName        string `json:"manufacturer_name"`
		Quantity                string `json:"quantity"`
		Type                    string `json:"type"`
		IDShopDefault           string `json:"id_shop_default"`
		Reference               string `json:"reference"`
		SupplierReference       string `json:"supplier_reference"`
		Location                string `json:"location"`
		Width                   string `json:"width"`
		Height                  string `json:"height"`
		Depth                   string `json:"depth"`
		Weight                  string `json:"weight"`
		QuantityDiscount        string `json:"quantity_discount"`
		Ean13                   string `json:"ean13"`
		Isbn                    string `json:"isbn"`
		Upc                     string `json:"upc"`
		Mpn                     string `json:"mpn"`
		CacheIsPack             string `json:"cache_is_pack"`
		CacheHasAttachments     string `json:"cache_has_attachments"`
		IsVirtual               string `json:"is_virtual"`
		State                   string `json:"state"`
		AdditionalDeliveryTimes string `json:"additional_delivery_times"`
		DeliveryInStock         string `json:"delivery_in_stock"`
		DeliveryOutStock        string `json:"delivery_out_stock"`
		OnSale                  string `json:"on_sale"`
		OnlineOnly              string `json:"online_only"`
		Ecotax                  string `json:"ecotax"`
		MinimalQuantity         string `json:"minimal_quantity"`
		LowStockThreshold       any    `json:"low_stock_threshold"`
		LowStockAlert           string `json:"low_stock_alert"`
		Price                   string `json:"price"`
		WholesalePrice          string `json:"wholesale_price"`
		Unity                   string `json:"unity"`
		UnitPriceRatio          string `json:"unit_price_ratio"`
		AdditionalShippingCost  string `json:"additional_shipping_cost"`
		Customizable            string `json:"customizable"`
		TextFields              string `json:"text_fields"`
		UploadableFiles         string `json:"uploadable_files"`
		Active                  string `json:"active"`
		RedirectType            string `json:"redirect_type"`
		IDTypeRedirected        string `json:"id_type_redirected"`
		AvailableForOrder       string `json:"available_for_order"`
		AvailableDate           string `json:"available_date"`
		ShowCondition           string `json:"show_condition"`
		Condition               string `json:"condition"`
		ShowPrice               string `json:"show_price"`
		Indexed                 string `json:"indexed"`
		Visibility              string `json:"visibility"`
		AdvancedStockManagement string `json:"advanced_stock_management"`
		DateAdd                 string `json:"date_add"`
		DateUpd                 string `json:"date_upd"`
		PackStockType           string `json:"pack_stock_type"`
		MetaDescription         string `json:"meta_description"`
		MetaKeywords            string `json:"meta_keywords"`
		MetaTitle               string `json:"meta_title"`
		LinkRewrite             string `json:"link_rewrite"`
		Name                    string `json:"name"`
		Description             string `json:"description"`
		DescriptionShort        string `json:"description_short"`
		AvailableNow            string `json:"available_now"`
		AvailableLater          string `json:"available_later"`
		Associations            struct {
			Categories []struct {
				ID string `json:"id"`
			} `json:"categories"`
			Images []struct {
				ID string `json:"id"`
			} `json:"images"`
			Combinations []struct {
				ID string `json:"id"`
			} `json:"combinations"`
			ProductOptionValues []struct {
				ID string `json:"id"`
			} `json:"product_option_values"`
			ProductFeatures []struct {
				ID             string `json:"id"`
				IDFeatureValue string `json:"id_feature_value"`
			} `json:"product_features"`
			StockAvailables []struct {
				ID                 string `json:"id"`
				IDProductAttribute string `json:"id_product_attribute"`
			} `json:"stock_availables"`
		} `json:"associations"`
	} `json:"product"`
}

type Category struct {
	Details struct {
		ID                  int    `json:"id"`
		IDParent            string `json:"id_parent"`
		LevelDepth          string `json:"level_depth"`
		NbProductsRecursive string `json:"nb_products_recursive"`
		Active              string `json:"active"`
		IDShopDefault       string `json:"id_shop_default"`
		IsRootCategory      string `json:"is_root_category"`
		Position            string `json:"position"`
		DateAdd             string `json:"date_add"`
		DateUpd             string `json:"date_upd"`
		Name                string `json:"name"`
		LinkRewrite         string `json:"link_rewrite"`
		Description         string `json:"description"`
		MetaTitle           string `json:"meta_title"`
		MetaDescription     string `json:"meta_description"`
		MetaKeywords        string `json:"meta_keywords"`
		Associations        struct {
			Categories []struct {
				ID string `json:"id"`
			} `json:"categories"`
		} `json:"associations"`
	} `json:"category"`
}

type Allcategories struct {
	Categories []struct {
		ID int `json:"id"`
	} `json:"categories"`
}

//func for get packages


func makeRequest(method, url, encodedToken string) []byte{
    // create new http request
    request, error := http.NewRequest(method, url, nil)
    request.Header.Set("Authorization", "Basic " + encodedToken)	
    // sending the request
    client := &http.Client{}
    response, error := client.Do(request)
	if error != nil {
		fmt.Println("error during request")
		fmt.Println(error)
		os.Exit(1)
	 }
    responseBody, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println("error during body read")
		fmt.Println(error)
		os.Exit(1)
	 }
    defer response.Body.Close()
    return responseBody
}

func GetAllIdsByType( encodedToken, shopUrl, objectType string) (string,[]byte) {

    method = "GET"
    //initializing api_suffix
    if string(shopUrl[len(shopUrl)-1]) == "/" {
		  apiSuffix = "api/" + objectType +"/"
	  } else {
		  apiSuffix = "/api/" + objectType + "/"
	  }
    //Initialisation json suffix
    jsonSuffix = "?output_format=JSON"
    //Generation complete URL
    completeUrl = shopUrl+apiSuffix+jsonSuffix
    responseBytes = makeRequest(method,completeUrl, encodedToken)
    return apiSuffix, responseBytes
}

func GetDetailsById (Id int, shopUrl, apiSuffix, encodedToken string) []byte {

		completeUrl = shopUrl + apiSuffix + strconv.Itoa(Id) + jsonSuffix
        responseBytes = makeRequest(method,completeUrl, encodedToken)
		return responseBytes

}
func GetAllDetails (objectStruct  interface{}, reqType, shopUrl, encodedToken string) []string{
    finalStringSlice = nil
	switch reqType {
		case "products" :
			allproducts := objectStruct.(Allproducts)
			fmt.Printf("%-10s %-40s %-10s\n", "ID", "Name", "Prix")
			for i := 0; i < len(allproducts.Products); i++ {
				//HTML stripping method
				p := bluemonday.StripTagsPolicy()
				productId = allproducts.Products[i].ID
				var product Product
				responseBytes = GetDetailsById (productId, shopUrl, apiSuffix, encodedToken)
				json.Unmarshal(responseBytes, &product)
			    finalStringSlice = append(finalStringSlice, fmt.Sprintf("%-10s %-40s %-10s", strconv.Itoa(product.Details.ID) ,product.Details.Name , p.Sanitize(product.Details.Price)))
				
			}
				
		case "orders" :
			allorders := objectStruct.(Allorders)
			fmt.Printf("%-10s %-40s %-10s\n", "ID", "Name", "Prix")
    		// iterate through all items id
			for i := 0; i < len(allorders.Orders); i++ {
				orderId = allorders.Orders[i].ID
				var order Order
				responseBytes = GetDetailsById (orderId, shopUrl, apiSuffix, encodedToken)
				json.Unmarshal(responseBytes, &order)
				finalStringSlice = append(finalStringSlice, fmt.Sprintf("%-10s %-40s %-10s", strconv.Itoa(order.Details.ID) ,order.Details.TotalPaid , order.Details.Reference))
			}
			
		case "categories" :
			allcategories := objectStruct.(Allcategories)
			fmt.Printf("%-10s %-40s %-10s\n", "ID", "Creation Date", "Name")
			// iterate through all items id
			for i := 0; i < len(allcategories.Categories); i++ {
				categoryId = allcategories.Categories[i].ID
				var category Category
				responseBytes = GetDetailsById (categoryId, shopUrl, apiSuffix, encodedToken)
				json.Unmarshal(responseBytes, &category)
				finalStringSlice = append(finalStringSlice, fmt.Sprintf("%-10s %-40s %-10s", strconv.Itoa(category.Details.ID) ,category.Details.DateAdd , category.Details.Name))
			}
			
		case "customers" :
			allcustomers := objectStruct.(Allcustomers)
			fmt.Printf("%-10s %-40s %-10s\n", "ID", "First Name", "Last Name")
			// iterate through all items id
			for i := 0; i < len(allcustomers.Customers); i++ {
				customerId = allcustomers.Customers[i].ID
				var customer Customer
				responseBytes = GetDetailsById (customerId, shopUrl, apiSuffix, encodedToken)
				json.Unmarshal(responseBytes, &customer)
				finalStringSlice = append(finalStringSlice, fmt.Sprintf("%-10s %-40s %-10s", strconv.Itoa(customer.Details.ID) ,customer.Details.Lastname , customer.Details.Firstname))
	        }
		case "carriers" :
			allcarriers := objectStruct.(Allcarriers)
			fmt.Printf("%-10s %-40s %-10s\n", "ID", "Name", "Delay")
			// iterate through all items id
			for i := 0; i < len(allcarriers.Carriers); i++ {
				carrierId = allcarriers.Carriers[i].ID
				var carrier Carrier
				responseBytes = GetDetailsById (carrierId, shopUrl, apiSuffix, encodedToken)
				json.Unmarshal(responseBytes, &carrier)
				finalStringSlice = append(finalStringSlice, fmt.Sprintf("%-10s %-40s %-10s", strconv.Itoa(carrier.Details.ID) ,carrier.Details.Name , carrier.Details.Delay))
	        }
	}

	 return finalStringSlice
	}

func PrintSlice(s []string) {
	if len(s) == 0 {
		return
	}
	fmt.Println(s[0])
	PrintSlice(s[1:])
}
