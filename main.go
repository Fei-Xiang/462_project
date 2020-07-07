package main

import (
    "bytes"
    "fmt"
    "io"
    "time"
    "os"

    "github.com/juju/ratelimit"
)

func main() {
    // Source holding 1MB
    i:=0
    f, err := os.Open("testing.mov")

    check(err)
    stat, err := f.Stat()
    if err != nil {
        // Could not obtain stat, handle error
    }
    len := stat.Size()

    for i<5{

        i++
    }
    size := make([]byte, len)
    src, err := f.Read(size)
    // Destination
    dst := &bytes.Buffer{}

    // Bucket adding 1000KB every second, holding max 1000KB
    bucket := ratelimit.NewBucketWithRate(1000*1024, 1000*1024)

    start := time.Now()

    // Copy source to destination, but wrap our reader with rate limited one
    io.Copy(dst, ratelimit.Reader(src, bucket))

    fmt.Printf("Copied %d bytes in %s\n", dst.Len(), time.Since(start))
}