package main

import (
	"lilith/job"
	"lilith/router"
)

func init() {
	go job.InitSpider()
}

func main() {
	r := router.InitRouter()
	_ = r.Run(":8080")
}
