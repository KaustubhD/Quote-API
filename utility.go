package main

import(
  "fmt"
  "net/http"
  "log"
)

func WriteResponseOrPanic(writer http.ResponseWriter, message string){
  _, err := fmt.Fprintf(writer, message)
  if err != nil{
    log.Panic(err)
  }
}
