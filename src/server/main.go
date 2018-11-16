package main

import (
  "fmt"
  "net/http"
  // "html/template"
)

func ping(w http.ResponseWriter, r *http.Request) {
      d1 := []byte("pong")
  w.Write([]byte(d1))
}

func element(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>this shit is confusing but I'll figure it out</h1><p>testing multiples</p><input type='email' /><input tupe='submit'/>")
}

// type NewOut struct {
//   Name string
//   Age string
// }

// func objectFunc(w http.ResponseWriter, r *http.Request) {

//   out := NewOut{Name:"Lee", Age: "Never"}
//   data := NewOut{Name:"Lee", Age: "Never"}
//   tmpl, err := template.New("name").Parse("<h1>NOPE</h1>")
//   // Error checking elided
//   err = tmpl.Execute(out, data)
// }

func main() {
  // http.HandleFunc("/template", objectFunc)
  http.Handle("/", http.FileServer(http.Dir("./src/server")))
  http.HandleFunc("/ping", ping)
  http.HandleFunc("/element", element)
  http.ListenAndServe(":8081", nil)
}