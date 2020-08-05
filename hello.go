package main

import "fmt"

//return hello gretting
func hello(user string) string {

	if len(user) == 0 {
		return "Hello BA!"
	} else {
		return fmt.Sprintf("Hello %v!", user)
	}

}

func main() {
	fmt.Println(hello(""))
}
