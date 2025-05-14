
# Web Crawler

A simple web crawler running as an HTTP server that takes a URL as input and finds all unique pages for the given website. Only crawls pages within the same domain.

## Getting Started

### Requirements
Built using Go version 1.23+.

### Installation & Run
```bash
  git clone https://github.com/its-kos/Web-Crawler.git
  cd Web-Crawler/cmd
  go build .
  ./main
```
    
## API Reference

#### Get all unique pages from given URL

```http
  GET /pages
```

| Parameter | Type     | Required  | Description                |
| :-------- | :------- | :-------- |:------------------------- |
| `target`  | `string`  | Yes      | The URL to analyzew. |

## Features

- Modular implementation.
- The web crawler runs as an HTTP server (localhost for now).
- Only crawls pages that belong to the given website (same domain).
- Uses BFS to traverse the "url tree" efficiently.
- The website URL is received as a GET request to /pages?target=url to start the
  crawling.
- Outputs a JSON that lists the identified pages on the domain.

### Future Improvements

These might not be ready for interview so I will start them on a separate branch.

- Concurrent BFS implementation.
- More robust page fetching.

## Tech Stack

**Language:** Go

**Packages:** Chi (V5), Goquery

**Extras:** Postman (testing)


README made using https://readme.so/editor