package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    _ "database/sql"
    "io/ioutil"
)

type Get_json_sign struct{
    Id        string    `json:"id"`
    Password  string    `json:"password"`
}

type Get_json_score struct{
    Id        string    `json:"id"`
    Score     int       `json:"score"`
}

type Post_json_signup struct{
    Msg       string    `json:"msg"`
}

func signup (w http.ResponseWriter,r *http.Request){
    _len:=r.ContentLength
    body:=make([]byte,_len)
    r.Body.Read(body)
    decoder:=json.NewDecoder(body)
    for{
        get_json_signup:=Get_json_sign{}
        err:=decoder.Decode(&get_json_signup)
        if err==io.EOF{
            break
        }
        if err!=nil{
            fmt.Println("Error decoding json",err)
            break
        }
    }
    user_id:=get_json_signup.id
    pasword:=get_json_signup.password
/*
    此处为数据库操作。。。
*/
    w.Header().Set("Content-Type","application/json")
    post_json_signup:=&Post_json_signup{
        Msg: "successful",
    }
//    json, _ :=json.Marshal(post_json_signup)
    w.Write(json)
}

func main(){
    server:=http.Server{
        Addr: "127.0.0.1:8080",
    }
    http.HandleFunc("/signup",signup)

    server.ListenAndServe()
}
