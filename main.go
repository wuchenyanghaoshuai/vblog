package main

import "vblog/cmd"



func main() {
	
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}