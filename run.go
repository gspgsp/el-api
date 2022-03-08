package main

import (
	"fmt"
	"github.com/gspgsp/el-api/services"
)

func main()  {
	num := services.GetNumber()
	fmt.Println("just a test", num)
}


