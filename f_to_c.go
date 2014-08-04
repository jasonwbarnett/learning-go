package main

import(
  "fmt"
)

func main() {
  fmt.Print("Enter F to convert to C: ")
  var input float64
  fmt.Scanf("%f", &input)
  fmt.Println(input)
  output := (input - 32) * 5/9
  fmt.Printf("Celsius: %f\n", output)
}
