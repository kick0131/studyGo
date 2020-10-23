# study Go
Multi module sample

### Get Start

1. clone project
    ```
    https://github.com/kick0131/studyGo.git
    ```
1. go run
    ```
    go run ./main.go
    ```
1. stop program
    ctrl + c

---
# Go Test

## Simple test
### Rule
- filename must be `xxx_test.go`
- funcformat must be `Testxxxx(t *testing.T)`

## Run test

### all package
```
go test -v ./...
```

# Another Go Commands

### list go package
`go list -m all`

### Go Doc
`godoc -http=:(port)`


