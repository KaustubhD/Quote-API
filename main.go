package main

import(
  "fmt"
  "log"
  "net/http"
)

type JSONMessage struct{
  Message string
}

func main(){

  http.HandleFunc("/", homePage)
  http.HandleFunc("/quotes", handleQuotes)

  fmt.Println("Server listening on port 8000")
  log.Panic(http.ListenAndServe(":8000", nil))

}

func homePage(writer http.ResponseWriter, request *http.Request){
  WriteJSON(writer, &JSONMessage{"Welcome to the Motivational Quote API"}, 200)

}

func handleQuotes(writer http.ResponseWriter, request *http.Request){
  if request.Method == http.MethodPost{
    postNewQuote(writer, request)
  }else if request.Method == http.MethodGet{
    getQuote(writer)
  }else{
    WriteJSON(writer, &JSONMessage{"Invalid request method."}, 405)
  }
}

func postNewQuote(writer http.ResponseWriter, request *http.Request){
  quote, err := GetQuoteFromRequest(request)
  if err != nil{
    WriteJSON(writer, &JSONMessage{err.Error()}, 422)
    return
  }

  err = quote.StoreInDB()
  if err != nil{
    WriteJSON(writer, &JSONMessage{"Database Error"}, 503)
    return
  }

  WriteJSON(writer, &JSONMessage{"Quote received"}, 200)
}

func getQuote(writer http.ResponseWriter){
  quoteString, err := GetQuoteFromDB()
  if err != nil{
    WriteJSON(writer, &JSONMessage{err.Error()}, 422)
    return
  }

  WriteJSON(writer, quoteString, 200)

}
