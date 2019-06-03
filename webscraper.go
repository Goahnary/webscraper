package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
);

func main () {
	fmt.Printf("What website should I scrape? ");
	var siteUrl string
	fmt.Scan(&siteUrl)

	fmt.Printf("\n");
	fmt.Printf("What kind of element? ");
	var element string
	fmt.Scan(&element)

	fmt.Printf("\n");
	fmt.Printf("Filter by? ");
	var filter string
	fmt.Scan(&filter)

	fmt.Printf("\n");
	fmt.Printf("Site: %v", siteUrl);
	fmt.Printf("\n");
	fmt.Printf("Element: %v", element);
	fmt.Printf("\n");
	fmt.Printf("Filter: %v", filter);
	fmt.Printf("\n");

	var prepend string
	if !strings.Contains(siteUrl, "http") {
		prepend = "http://"
	}

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

    // fmt.Printf("\n");

	// fmt.Printf("Results:\n");
	// doc.Find(element).Each(func(i int, s *goquery.Selection){
	// 	fmt.Printf("%v\n", s.Text());
	// });

	filteredResults := doc.Find(element).FilterFunction(func(i int, s *goquery.Selection) bool {
		return strings.Contains(s.Text(), filter);
	});

	fmt.Printf("\n");
	
	fmt.Printf("Filtered:\n");

	filteredResults.Each(func(i int, s *goquery.Selection){
		fmt.Printf("%v\n", s.Text());
	});
}
