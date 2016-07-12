# gofstream
Turn a file into a stream (channel) of lines.

# Usage

Simple utility library.  Just supply a file path and desired size for the channel buffer and then iterate over the resulting lines on the return channel.

# Example

Assume a file `somedata.txt` contains the desired lines to process.

```sh
printf "111\n222\n333\n444\n555\n" > somedata.txt
```

```go
package main

import (
    "fmt"
    "log"
    "github.com/adamdrake/gofstream"
)

func main() {
    lines, err := fstream.New("somedata.txt", 0)
    if err != nil {
        log.Fatal(err)
    }
    for r := range lines {
        // Do something with each line
        fmt.Println(r)
    }
}
```