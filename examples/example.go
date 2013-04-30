package main

import (
    "fmt"
    "time"
    "./reddit"
)

func main() {
    entries, err := reddit.GetSubmissions("golang")
    if err != nil {
        panic(err)
    }
    for _, x := range entries {
        t := time.Unix(int64(x.Created_utc), 0)
        fmt.Printf("[%d] %s\n\ton %s\n\tby %s\n\t%s\n\n",
                   x.Score, x.Title, t, x.Author, x.Url)
    }
}
