package main

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
func main() {

}
