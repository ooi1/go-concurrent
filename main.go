//Beginning of the main package
package main

//These imports are auto added once u use them
import (
	"fmt"
	"net/http"
	"time"
)

//The main function is where program starts running
func main() {
	//These are link definitions to various websites
	// links is a array of type string
	// the links are seperated by commas and a last comma is required
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
		"https://www.panzerrush.com",
	}

	// This creates a channel to receive values of type string from goroutines
	c := make(chan string)

	// This recurse through range links and runs each as goroutines
	for _, link := range links {
		go checkLink(link, c)
	}

	// here we recurse through the range of channels to wait for incoming goroutines
	// this is done thru a function literal
	// when a goroutine returns, the channel is activated and goroutine runs to completion
	for l := range c {
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

//This is the function that "gets" each link via HTTP
// it returns a string to indicate success or failure
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
