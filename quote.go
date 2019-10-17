package main

import(
  "net/http"
  "io/ioutil"
  "encoding/json"
  "errors"
)

type QuoteString struct{
  Quote string `json:"quote"`
}

func GetQuoteFromRequest(request *http.Request) (*QuoteString, error){
  body, err := ioutil.ReadAll(request.Body)
  if err != nil{
    return nil, err
  }

  var quote QuoteString
  err = json.Unmarshal(body, &quote)
  if err != nil{
    return nil, err
  }

  if quote.Quote == ""{
    return nil, errors.New("At least enter a quote")
  }

  return &quote, nil
}
