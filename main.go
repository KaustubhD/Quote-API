package main

import(
  "fmt"
  "log"
  "net/http"
)

func main(){

  http.HandleFunc("/", homePage)
  http.HandleFunc("/quotes", handleQuotes)

  fmt.Println("Server listening on port 8000")
  log.Panic(http.ListenAndServe(":8000", nil))

}

func homePage(writer http.ResponseWriter, request *http.Request){
  _, err := fmt.Fprintf(writer, "Welcome to the Motivational Quote API")
  if err != nil{
    log.Panic(err)
  }

}

func handleQuotes(writer http.ResponseWriter, request *http.Request){
  if request.Method == http.MethodPost{
    postNewQuote(writer, request)
  }else{
    _, err := fmt.Fprintf(writer, "Only POST request available right now !")
    if err != nil{
      log.Panic(err)
    }
  }
}

func postNewQuote(writer http.ResponseWriter, request *http.Request){
  quote, err := GetQuoteFromRequest(request)
  if err != nil{
    WriteResponseOrPanic(writer, fmt.Sprintf("Error: Cannot parse a quote from JSON.\n%s\n", err.Error()))
    return
  }

  err = quote.StoreInDB()
  if err != nil{
    WriteResponseOrPanic(writer, fmt.Sprintf("Error: Cannot save quote.\n%s\n", err.Error()))
    return
  }

  WriteResponseOrPanic(writer, fmt.Sprintf("Quote received: \"%s\"\n", quote.Quote))
}
