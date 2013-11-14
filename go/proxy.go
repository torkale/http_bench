package main

import (
    "fmt"
    "net/http"
    "net/url"
)

func main(){
    fmt.Printf("Starting http Server ... ")
    http.Handle("/", http.HandlerFunc(proxyHandlerFunc))
    err := http.ListenAndServe("0.0.0.0:3000", nil)
    if err != nil {
        fmt.Printf("ListenAndServe Error :" + err.Error())
    }
}

func proxyHandlerFunc(w http.ResponseWriter, r *http.Request) {
    // build proxy client
    proxyUrl, err := url.Parse("http://localhost:4000")
    tr := &http.Transport { Proxy: http.ProxyURL(proxyUrl) }
    defer tr.CloseIdleConnections()
    httpClient := &http.Client { Transport: tr }

    // override RequestURI
    r.RequestURI = ""
    r.URL.Scheme = "http"
    r.URL.Host= "localhost"
    res, err := httpClient.Do(r)

    if err != nil {
        fmt.Println(err.Error())
    }
    res.Body.Close()
}
