package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "miyano246:dyqatyz3tc@tcp(localhost:3306)/procir_miyano246?charset=utf8")

    if err != nil {
        fmt.Println(err.Error())
    }

    defer db.Close()

    err = db.Ping()

    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }
}
