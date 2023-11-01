package main

import "C"
import (
	"micoService/src/algorithm/analysis"

	"github.com/gin-gonic/gin"
)

func main() {
	//发布服务  http://localhost:8080/pointbuffer?req=POINT(1%202)&distance=0.5&n=16
	r := gin.Default()
	r.GET("/pointbuffer", analysis.PointBuffer)
	r.Run(":8080")
}
