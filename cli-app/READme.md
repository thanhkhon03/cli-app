# CLI App
A simple command-line application built with Golang.

## Features
- Count words in a text (`count` command).

## Usage
```bash
# Count words
go run main.go count --text="Hello Golang CLI"


/*---

### **7. Đóng gói và phân phối**
#### **7.1. Biên dịch ứng dụng**
Biên dịch ứng dụng để chạy trên các hệ điều hành khác:
```bash
# Cho Windows
GOOS=windows GOARCH=amd64 go build -o cli-app.exe main.go

# Cho macOS
GOOS=darwin GOARCH=amd64 go build -o cli-app main.go

# Cho Linux
GOOS=linux GOARCH=amd64 go build -o cli-app main.go  **/

# How to run the application
go run C:\Users\Admin\Documents\Blockchain\cli-app\main.go -text="Hello World"
 