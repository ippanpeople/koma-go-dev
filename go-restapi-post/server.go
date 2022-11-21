package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 聲明標準返回類型結構體
type ResponseResult struct {
	Result string `json:"result"`
}

// 聲明書籍結構體
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Price  int     `json:"price"`
	Author *Author `json:"author"`
}

// 聲明作者結構體
type Author struct {
	Name string `json:"name"`
}

// 聲明書籍數組
var books []Book

// 聲明一過初始化函數
func initBooks() {
	books = append(books, Book{
		ID:     "1",
		Title:  "GO hello 1",
		Price:  51,
		Author: &Author{Name: "rin"},
	})
	books = append(books, Book{
		ID:     "2",
		Title:  "GO hello 2",
		Price:  52,
		Author: &Author{Name: "rinrin"},
	})
	books = append(books, Book{
		ID:     "3",
		Title:  "GO hello 3",
		Price:  53,
		Author: &Author{Name: "rinrinrin"},
	})
}

// 定義設置共同的Http response Header 訊息的函數
func setSameHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// 定義用來處理跟請求的函數
func index(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	fmt.Fprintf(w, "Hello Go REST API Server")
}

// 定義基礎的CRUD函數, 並調用setSameHeader()函數使CRUD函數能返回JSON類型的資料
func listBooks(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	// 調用mux路由組建的Vars方法解析r請求流
	params := mux.Vars(r) // 解析成功的結果（數組）
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "getBook:" + params["id"],
	})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	var book Book
	// 獲取 body 劉轉換為結構體 json -> struct
	_ = json.NewDecoder(r.Body).Decode(&book)
	// 生成新書 ID（刪除後會有重複ＩＤ問提, 此方法只限測試環境使用）
	book.ID = strconv.Itoa(len(books) + 1)
	// 添加進數組
	books = append(books, book)
	// 返回客戶端
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "updateBook",
	})
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	setSameHeader(w)
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			// 取得從頭到要刪除前的書, 跟要刪除後一本的書到結束的所有資料, 再利用append拼成一個新的數組
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(ResponseResult{
		Result: "deleteBook:" + params["id"],
	})
}

func main() {
	fmt.Println(">>Hello Welcome to Go REST API Server<<")

	initBooks()

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/books", listBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("服務器運行在端口 :8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}

}
