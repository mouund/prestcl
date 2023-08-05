/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
       _ "prestctl/cmd/delete" 
       _ "prestctl/cmd/get"
       _ "prestctl/cmd/describe"
       "prestctl/cmd"
)

func main() {
       
       cmd.Execute()
       
}
