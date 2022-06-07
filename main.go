package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func login(w http.ResponseWriter, r *http.Request) {
  var fileName = "login.html"
  t, err := template.ParseFiles(fileName)
  if err != nil {
    fmt.Println("Error when parsing file", err)
    return
  }

  err = t.ExecuteTemplate(w, fileName, data:nil) {
    if err != nil {
      fmt.Println("Error when executing template", err)
      return
    }
  }
}

// Creating a mock DB
var mockDB = map[string]string {
  "Wallace": "goodPassword"
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
  username := r.FormValue("username")
  password := r.FormValue("password")
  fmt.Println(username, password)
}

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/login":
      login(w, r)
    case "/login-submit":
      loginSubmit(w, r)
    }
    default: 
      fmt.Fprintf(w, format: "Sup Ninjas")
} 
