package main

import (
	_ "ZShortUrl/client"
	_ "ZShortUrl/config"
	"ZShortUrl/router"
)

func main() {
	r := router.Router()
	_ = r.Run()
}
