package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://www.example.com")
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

	// Open a CSV file for writing
	file, err := os.Create("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)

	// Write the header row to the CSV file
	header := []string{"Header name"}
	if err := writer.Write(header); err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".ContentCard").Each(func(i int, selection *goquery.Selection) {
		// For each item found, get the band and title
		text := selection.Text()
		
		fmt.Println(text)

		// Write the extracted data to the CSV file
		if err := writer.Write([]string{text}); err != nil {
			log.Fatal(err)
		}
	})

	writer.Flush()

	// Check for errors
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
	
}