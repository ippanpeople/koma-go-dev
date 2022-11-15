package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello Welcome to Go Server")

	// 實例化一個文件服務器, 且讓文件服務器使用本幣文件夾./static
	fileServer := http.FileServer(http.Dir("./static"))
	// 當訪問根（/）時響應這個文件服務器給使用者
	http.Handle("/", fileServer)

	fmt.Printf("服務器運行在端口 :8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
