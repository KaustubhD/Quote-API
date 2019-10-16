package main

import(
  "fmt"
  "log"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type quoteString struct{
  Quote string `json:"quote"`
}
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
  body, err := ioutil.ReadAll(request.Body)
  if err != nil{
    _, err := fmt.Fprintf(writer, "Error read request body")
    if err != nil{
      log.Panic(err)
    }
    return
  }

  var quote quoteString
  err = json.Unmarshal(body, &quote)
  if err != nil{
    _, err := fmt.Fprintf(writer, "Error parsing JSON data")
    if err != nil{
      log.Panic(err)
    }
    return
  }

  if quote.Quote == ""{
    _, err := fmt.Fprintf(writer, "At least enter a quote")
    if err != nil{
      log.Panic(err)
    }
    return
  }

  _, err = fmt.Fprintf(writer, "Quote received: \"%s\"\n", quote.Quote)
  if err != nil{
    log.Panic(err)
  }
}
