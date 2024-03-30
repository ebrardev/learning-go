package main

import "fmt"

func main() {
	websites := map[string]string{
		"Facebook": "https://facebook.com",
		"Google":   "https://google.com",
		"Twitter":  "https://twitter.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Facebook"])
	websites["AkrepNalan"] = "https://akrepnalan.com"
	fmt.Println(websites)
	delete(websites, "Facebook")
	fmt.Println(websites)
}
