package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"video_demo/src/config"
	"video_demo/src/model"
	"video_demo/src/router"
)

func main() {
	r := gin.Default()

	cfg, err := config.Load("./config.yml")
	if err != nil {
		fmt.Println("Loading config.yml failed.")
	}

	router.InitRouter(r)

	model.InitDB(cfg)

	port := ":" + strconv.Itoa(cfg.Net.Port)
	if err = r.Run(port); err != nil {
		log.Fatalf("Running on port %v failed", port)
	}
}
