# go-urljoin

This is a port attempt for https://github.com/jfromaniello/url-join

## Install

```
go get -u github.com/TheBoringDude/go-urljoin
```

## Usage

```go

url, err := urljoin.UrlJoin("https://www.google.com", "search", "?q=hello-world")
if err != nil {
    log.Fatal(err)
}

fmt.Println(url) // https://www.google.com/search?q=hello-world

```

##

### &copy; 2021 | TheBoringDude
