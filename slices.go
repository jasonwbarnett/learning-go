package main

import "fmt"

func main() {
  jason := "Jason Barnett"
  fmt.Println(len(jason))
  for i:=0;i<len(jason);i++ {
  fmt.Printf("%X ", jason[i])}
}
