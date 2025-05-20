package main

import "vblog/cmd"



func main() {
	
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}