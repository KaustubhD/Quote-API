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

func LogOnServerOrPanic(data interface{}){
  _, err := fmt.Printf("%+v\n", data)
  if err != nil{
    log.Panic(err)
  }
}

func RespondAndLog(writer http.ResponseWriter, data interface{}, status int){
  writer.Header().Set("Content-Type", "application/json")
  writer.WriteHeader(status)

  JSON, err := json.Marshal(data)

  if err != nil{
    log.Panic(err)
  }
  JSONInString := string(JSON)

  LogOnServerOrPanic(JSONInString)
  WriteResponseOrPanic(writer, JSONInString)
}
