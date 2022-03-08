package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gspgsp/el-api/services"
)

func main()  {
	num := services.GetNumber()
	gin.Default()
	fmt.Println("just a test", num)
}


