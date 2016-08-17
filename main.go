package main

import (
    "fmt"
    "net/http"
    "io"
   	"os"
    "math/rand"
    "strconv"
)

var client = http.Client{}

var http_codes = [...]int{101, 200, 201, 202, 204, 206, 207, 300, 301, 303, 304, 305, 307, 400, 401, 402, 403, 404, 405, 406, 408, 409, 410, 411, 413, 414, 416, 417, 418, 422, 423, 424, 425, 426, 429, 431, 444, 450, 500, 502, 503, 506, 507, 508, 509, 599}
var roots = [...]string{"https://httpstatusdogs.com/img/", "https://http.cat/"}

func statusHandler(res http.ResponseWriter, req *http.Request) {
    var code = strconv.Itoa(http_codes[rand.Intn(len(http_codes))])
    var root = roots[rand.Intn(2)]
    reqImg, err := client.Get(root + code + ".jpg")
    if err != nil {
        fmt.Fprintf(res, "Error %d", err)
        return
    }
    res.Header().Set("Content-Length", fmt.Sprint(reqImg.ContentLength))
    res.Header().Set("Content-Type", reqImg.Header.Get("Content-Type"))
    if _, err = io.Copy(res, reqImg.Body); err != nil {
        // handle error
    }
    reqImg.Body.Close()
}

func main() {
    http.HandleFunc("/status", statusHandler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
