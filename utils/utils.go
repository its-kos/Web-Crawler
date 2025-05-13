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

func ExtractPageURLs(doc goquery.Document) ([]string, error) {
	var hrefs []string
	urls := doc.Find("a")

	//log.Println("Found ", urls.Size(), " links")

	for _, url := range urls.Nodes {
		log.Println("Link: ", url.Attr[0].Val)
		hrefs = append(hrefs, url.Attr[0].Val)
	}

	return hrefs, nil
}

func BFS(start string) ([]string, error) { // BFS because links on the same level might be more relevant
	finalUrls := []string{start}
	queue := []string{start}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		currUrl := queue[0]
		log.Println("Visiting: ", currUrl)
		queue = queue[1:]
		if visited[currUrl] {
			continue
		}

		visited[currUrl] = true

		doc, err := FetchHTML(currUrl)
		if err != nil {
			return nil, err
		}

		urls, err := ExtractPageURLs(*doc)
		if err != nil {
			return nil, err
		}

		for _, u := range urls {
			if _, ok := visited[u]; !ok {
				finalUrls = append(finalUrls, u)
				queue = append(queue, u)
			}
		}
	}
	return finalUrls, nil
}
