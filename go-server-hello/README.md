# 利用Go建立一個HTTP服務器 並寫一個URL的映射函數

## 步驟
1. 初始化項目工程
2. 導入net http包 https://pkg.go.dev/net/http
3. 編寫項目程序
4. 編寫URL映射函數
5. 運行調適

## 實作
- 建立工程目錄
```bash
mkdir go-server
```
- 進入目錄
```bash
cd go-server
```
- 初始化工程模塊
```bash
go mod github.com/ippanpeople/go-server
```
- 編寫服務器程序
```bash
vim server.go
```
- 運行程序
```bash
go run server.go
```
- 編譯程序
```bash
go build
```
-運行程序
```bash
./go-server
```