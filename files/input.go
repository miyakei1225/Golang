package main

import (
    "database/sql"
    "fmt"
    //"net/http"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "miyano246:dyqatyz3tc@tcp(localhost:3306)/procir_miyano246?charset=utf8")

    defer db.Close()

    if err != nil {
        fmt.Println(err.Error())
    }

    //Select
    rows, err := db.Query("select * from user")
    defer rows.Close()


    if err != nil {
        fmt.Println("データベース接続失敗")
        return
    } else {
        fmt.Println("データベース接続成功")
    }
}
