package main

import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "os"
    "io"
)

func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/upload", upload)
    log.Fatal(http.ListenAndServe(":8080", nil))
    fmt.Println("Listen on port: 8080")
}

func index(w http.ResponseWriter, r *http.Request) {
    file, _ := ioutil.ReadFile("form.html")

	w.Header().Set("Content-Type", "text/html")
	w.Write(file)
}

func upload(w http.ResponseWriter, r *http.Request) {
    r.ParseMultipartForm(32 << 20)

    if r.Method != "POST" {
        fmt.Fprintf(w, "Method not allowed")
        return
    }

    file, handler, err := r.FormFile("uploadfile")
    if err != nil {
        fmt.Print("Err: ")
        fmt.Println(err)
        return
    }
    defer file.Close()

    f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
       fmt.Println(err)
       return
    }
    defer f.Close()

    io.Copy(f, file)

    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "Upload <br />")

}
