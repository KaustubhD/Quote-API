package main

import(
  "fmt"
  "net/http"
  "log"
  "encoding/json"
)

func WriteResponseOrPanic(writer http.ResponseWriter, message string){
  _, err := fmt.Fprintf(writer, message)
  if err != nil{
    log.Panic(err)
  }
}
func WriteJSON(writer http.ResponseWriter, data interface{}, status int){
  writer.Header().Set("Content-Type", "application/json")
  writer.WriteHeader(status)

  actualJSON, err := json.Marshal(data)

  if err != nil{
    log.Panic(err)
  }

  WriteResponseOrPanic(writer, string(actualJSON))
}
