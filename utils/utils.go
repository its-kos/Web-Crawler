package utils

import (
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func ValidateURL(givenUrl string) (*url.URL, error) {
	parsedURL, err := url.Parse(givenUrl)
	// This might be too strict for this use case. We can modify it later.
	// Only allows http(s)://www.domain.com type of URLs
	if err == nil && parsedURL.Scheme != "" && parsedURL.Host != "" {
		return parsedURL, nil
	} else {
		return nil, err
	}
}

func FetchHTML(givenUrl string) (*goquery.Document, error) {
	res, err := http.Get(givenUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// For now this only find the top level links
// This must be expanded upon to do a full "tree" search.
// Same as BFS/DFS with a stack/queue
func ExtractURLs(doc goquery.Document) ([]string, error) {
	var hrefs []string
	urls := doc.Find("a")
	// Images can also link but they point to a file, not a page
	//imgs := doc.Find("img") 

	log.Println("Found ", urls.Size(), " links")
	//log.Println("Found", imgs.Size(), "images")

	for _, url := range urls.Nodes {
		log.Println("Link: ", url.Attr[0].Val)
		hrefs = append(hrefs, url.Attr[0].Val)
	}

	return hrefs, nil
}
