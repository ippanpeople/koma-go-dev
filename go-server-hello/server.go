package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(">>Hello Welcome to Go Server<<")

	// 實例化一個文件服務器, 且讓文件服務器使用本幣文件夾./static
	fileServer := http.FileServer(http.Dir("./static"))

}
