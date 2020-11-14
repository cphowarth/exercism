package main
import "fmt"
import "strings"

func main() {
	words := [4]string{
		"cat", 
		"Dog", 
		"b4t", 
		"gO",
	}
	for _, word := range words {
		upperword := strings.ToUpper(word)
		fmt.Println(upperword)
    }
}

