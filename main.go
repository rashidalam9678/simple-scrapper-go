package main
 import ( 
	"fmt"
	"net/http"
	"os"
	"strings"
	"golang.org/x/net/html"

 )
 func main(){
	foundUrls := make(map[string]bool)
	seedUrls := os.Args[1:]

	chUrls := make(chan string)
	chFinished := make(chan bool)

	for _, url := range seedUrls {
		go crawl(url, chUrls, chFinished);
	}
	for c:=0; c<len(seedUrls); {
		select {
		case url := <-chUrls:
			foundUrls[url] = true
		case <-chFinished:
			c++
		}
	}
	fmt.Println("found", len(foundUrls), "unique urls")
	for url, _ := range foundUrls {
		fmt.Println("-",url,"")
	}	
 }
