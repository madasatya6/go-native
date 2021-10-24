package utils

import (
    "fmt"
    "net/http"
    "time"
)

func SetFlashCookie(rw http.ResponseWriter, r *http.Request, name, value string) {
    c := &http.Cookie{
        Name:     name,
        Value:    value,
        Path:     "/",
        MaxAge:   1,
        HttpOnly: true,
    }

    http.SetCookie(rw, c)
}

func DeleteCookieHandler(rw http.ResponseWriter, r *http.Request, name, value string) {

    c, err := r.Cookie("storage")
    if err != nil {
        panic(err.Error())
    }
    c.Name = name
    c.Value = value
    c.Expires = time.Unix(1414414788, 1414414788000)
}

func ReadCookieHandler(rw http.ResponseWriter, r *http.Request, name string) {

    c, err := r.Cookie(name)
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(c.Expires)
}

func EvaluateCookieHandler(rw http.ResponseWriter, r *http.Request) {

    c, err := r.Cookie("storage")
    if err != nil {
        panic(err.Error())
    }

    if time.Now().After(c.Expires) {
        fmt.Println("Cookie is expired.")
    }
}


