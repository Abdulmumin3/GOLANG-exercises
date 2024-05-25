package main

import (
	"fmt"

	"time"

	"rsc.io/quote"
)

func main(){

  switch(time.Now().Weekday()){
  case time.Saturday, time.Sunday:
    fmt.Println("It's weekend")
  default:
    fmt.Println("It's weekday")
  }

  fmt.Println(quote.Go())



}