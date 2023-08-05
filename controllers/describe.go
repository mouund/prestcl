package controllers

import (
     "fmt"
     "net/http" 
     "strconv" 
	 "os"
	 "io"
  )

func DescribeObjectFromId (encodedToken, shopUrl, objectType string, id int) []byte {

    method ="GET"
    if string(shopUrl[len(shopUrl)-1]) == "/" {
        apiSuffix = "api/" + objectType +"/"
    } else {
        apiSuffix = "/api/" + objectType + "/"
    }
    //Initialisation json suffix
    jsonSuffix = "?output_format=JSON"
    //Generation complete URL
    completeUrl = shopUrl+apiSuffix+strconv.Itoa(id)+jsonSuffix
    request, error := http.NewRequest(method, completeUrl, nil)
    request.Header.Set("Authorization", "Basic " + encodedToken)
    client := &http.Client{}
    response, error := client.Do(request)

    if response.Status != "200 OK" {
        fmt.Printf("error during request %s ", response.Status )
		os.Exit(1)
    }
	if error != nil {
		fmt.Printf("error during request %s ", error)
		os.Exit(1)
	 }
	 responseBytes, error := io.ReadAll(response.Body)
	 if error != nil {
		fmt.Printf("error during Body read %s ", error)
		os.Exit(1)
	 }
     return responseBytes
}

