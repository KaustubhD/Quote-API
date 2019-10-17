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

func (quote *QuoteString) StoreInDB() error{
  query := "INSERT into quotes (quote) VALUES (?)"
  _, err := ExecDB(query, quote.Quote)

  return err
}

func GetQuoteFromDB() (*QuoteString, error){
  query := "SELECT quote from quotes order by RANDOM() limit 1"
  row, err := QueryDB(query)

  if err != nil {
    return nil, err
  }

  if !row.Next(){
    return nil, errors.New("Cannot find a quote")
  }

  var q string
  err = row.Scan(&q)
  if err != nil{
    return nil, err
  }

  actualQuote := &QuoteString{
    Quote: q,
  }

  return actualQuote, nil
}
