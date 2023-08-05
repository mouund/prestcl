package controllers

import (
    "fmt"
     "net/http" 
     "strconv"
      "os"
  )

//Variables for delete package


func DeleteObjectFromId (encodedToken, shopUrl, objectType string, id int) string {
    method ="DELETE"
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
	if error != nil {
		fmt.Println("error during request")
		fmt.Println(error)
		os.Exit(1)
	 }
     return response.Status



}