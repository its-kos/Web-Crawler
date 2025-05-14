package utils

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func ValidateInternalURL(givenUrl string, domainURL string) (string, bool) {
	parsedDomainURL, err := url.Parse(domainURL)
	if err != nil {
		return "", false
	}

	parsedURL, err := url.Parse(givenUrl)
	if err != nil {
		return "", false
	}

	resolvedURL := parsedDomainURL.ResolveReference(parsedURL)
	if resolvedURL.Host != parsedDomainURL.Host {
		return "", false
	}

	return resolvedURL.String(), true
}

func FetchHTML(givenUrl string) (*goquery.Document, error) {
	res, err := http.Get(givenUrl)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func ExtractPageURLs(currUrl string) ([]string, error) {
	doc, err := FetchHTML(currUrl)
	if err != nil {
		return nil, err
	}

	var hrefs []string

	doc.Find("a").Each(func(i int, s *goquery.Selection) { // Apparently this is the goquery way ??
		// Safer cause it checks first if the attribute exists
		if href, exists := s.Attr("href"); exists {
			hrefs = append(hrefs, href)
		}
	})

	return hrefs, nil
}

// BFS instead if DFS because links on the same level might be more relevant (Not that it matters that much).
// This also assumes that the first url is valid and in the same domain.
// For the rest of the urls we check.
func BFS(start string) ([]string, error) {

	finalUrls := []string{start}
	queue := []string{start}
	visited := make(map[string]bool)

	for len(queue) > 0 {
		currUrl := queue[0]
		queue = queue[1:]

		if visited[currUrl] {
			continue
		}

		visited[currUrl] = true

		urls, err := ExtractPageURLs(currUrl)
		if err != nil {
			return nil, err
		}

		for _, url := range urls {
			if resolved, valid := ValidateInternalURL(url, currUrl); valid {
				if _, ok := visited[resolved]; !ok {
					finalUrls = append(finalUrls, resolved)
					queue = append(queue, resolved)
				}
			}
		}
	}
	return finalUrls, nil
}
