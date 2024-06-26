# Multi-Backend Caching Library in Go

## Steps to Run the Project

1. **Create a new folder and initialize the module:**
   ```bash
   go mod init app

2. **Get the package**
   ```bash
   go get github.com/devisettymahidhar315/cache
3. **Create a new file in the directory with the following content**
   ```bash
   package main
   import "github.com/devisettymahidhar315/cache"
   func main() {
    r := cache.Hello()
    r.Run()
   }
4. **Run the Program**
   ```bash
      go run main.go length
   ```
   length means number of elements are storing
   ```bash
      go run main.go
   ```
   without length if we run the program, it takes the default length has 2

# Functions Present in the Project
### `get`  
### `post`
### `delete`

# Accessing the Functions
## Get Function
### you can access on web broswer.
### for pritinf the data 
### redis data ```http://localhost:8080/redis/print```
### inmemory data ```http://localhost:8080/inmemory/print```
### particular data ```http://localhost:8080/key```

## Delete Function
### delete the data
### open the terminal and type the following commands for
### delete partcular command ```curl -X DELETE http://localhost:8080/key```
### delete entire data ```curl -X DELETE http://localhost:8080/all```

## Post Function
### store the data
### Implement the "time to live" (TTL) feature. If the time is -1, it means no time is set. If the time is greater than 1, the key-value pair will be deleted after that many seconds
### open the terminal and type the following command ```curl -X POST http://localhost:8080/key/value/time```
