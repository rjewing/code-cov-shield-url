# code-cov-shield-url
builds a shields.io url for code coverage

# Usage
```
$ go build main.go
$ ./main --coverage 72
https://img.shields.io/badge/coverage-72%25-green

```

To download image, run:
```
$ curl --output image.svg "`./main --coverage 72`"
```

Paired with go test:
```
$ go test -coverprofile cover.out
$ ./main --coverage `go tool cover -func cover.out | grep "^total:" | tr -s "\t" " " | awk '{ printf "%.0f", $3 }'`
https://img.shields.io/badge/coverage-72%25-green
```
