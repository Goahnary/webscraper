package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
);

func main () {
	//get site from user
	fmt.Printf("Website to scrape: ");
	var siteUrl string
	fmt.Scan(&siteUrl)

	//get jquery selector from user
	fmt.Printf("\n");
	fmt.Printf("Jquery Selector:  ");
	var element string
	fmt.Scan(&element)

	//get search term from user
	fmt.Printf("\n");
	fmt.Printf("Filter by: ");
	var filter string
	fmt.Scan(&filter)

	//add http:// if user doesn't provide it
	var prepend string
	if !strings.Contains(siteUrl, "http") {
		prepend = "http://"
	}

	//make request
	res, err := http.Get(prepend + siteUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	//get filtered results
	filteredResults := doc.Find(element).FilterFunction(func(i int, s *goquery.Selection) bool {
		return strings.Contains(s.Text(), filter);
	});

	fmt.Printf("\n");
	fmt.Printf("Filtered Results:");

	filteredResults.Each(func(i int, s *goquery.Selection){
		fmt.Printf("\n");
		fmt.Printf("%v. %v\n", i+1, s.Text());
	});
}
