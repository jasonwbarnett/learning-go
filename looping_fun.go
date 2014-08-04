package main
import "fmt"

func printNum(number int, eORo string) (n int, err error) {
  var word string
  switch number {
    case 0: word = "Zero"
    case 1: word = "One"
    case 2: word = "Two"
    case 4: word = "Four"
    case 8: word = "Eight"
    default: word = "Unknown number"
  }
  return fmt.Printf("%d (%s) %s\n", number, word, eORo)
}

func main() {
  for i:=1;i<=10;i++ {
    if i % 2 == 0 {
      printNum(i, "even")
    } else {
      printNum(i, "odd")
    }
  }

}
