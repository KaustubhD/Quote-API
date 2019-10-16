package main

import(
  "fmt"
  "log"
  "net/http"
)

func main(){

  http.HandleFunc("/", homePage)

  fmt.Println("Server listening on port 8000")
  log.Panic(http.ListenAndServe(":8000", nil))

}

func homePage(writer http.ResponseWriter, request *http.Request){
  _, err := fmt.Fprintf(writer, "Welcome to the Motivational Quote API")
  if err != nil{
    log.Panic(err)
  }

}
