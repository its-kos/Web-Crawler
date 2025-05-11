
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

- TODO
- TODO


## Tech Stack

**Language:** Go

**Packages:** Chi (V5)

**Extras:** -


README made using https://readme.so/editor