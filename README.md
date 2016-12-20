# GoLang Yahoo Finance API (goyhfin)
Simple API for fetching historical (15 minute delayed) financial (stocks, currencies, etc.) quote data. This can also be extended to fetch more technical indicators or fundamentals from the same endpoint.
###Setup
`go get github.com/svarlamov/goyhfin`
100% GoLang, so functions on any GoLang-compatible operating system
###Usage Example
main.go
```go
package main

import (
        "fmt"
        "github.com/svarlamov/goyhfin"
)

func main() {
        resp, err := goyhfin.GetTickerData("AAPL", goyhfin.OneMonth, goyhfin.OneDay, false)
        if err != nil {
                // NOTE: For library-specific errors, you can check the err against the errors exposed in goyhfin/errors.go
                fmt.Println("Error fetching Yahoo Finance data:", err)
                panic(err)
        }
        for ind := range resp.Quotes {
                fmt.Println("The day's high was", resp.Quotes[ind].High, "on the", resp.Quotes[ind].OpensAt.Day(), "day of", resp.Quotes[ind].OpensAt.Month(), "of", resp.Quotes[ind].OpensAt.Year())
        }
}
```
Example Running the Above main.go:
```bash
svarlamov$ go run main.go
The day's high was 111.98999786376953 on the 21 day of November of 2016
The day's high was 112.41999816894531 on the 22 day of November of 2016
The day's high was 111.51000213623047 on the 23 day of November of 2016
The day's high was 111.87000274658203 on the 25 day of November of 2016
The day's high was 112.47000122070312 on the 28 day of November of 2016
The day's high was 112.02999877929688 on the 29 day of November of 2016
The day's high was 112.19999694824219 on the 30 day of November of 2016
The day's high was 110.94000244140625 on the 1 day of December of 2016
The day's high was 110.08999633789062 on the 2 day of December of 2016
The day's high was 110.02999877929688 on the 5 day of December of 2016
The day's high was 110.36000061035156 on the 6 day of December of 2016
The day's high was 111.19000244140625 on the 7 day of December of 2016
The day's high was 112.43000030517578 on the 8 day of December of 2016
The day's high was 114.69999694824219 on the 9 day of December of 2016
The day's high was 115 on the 12 day of December of 2016
The day's high was 115.91999816894531 on the 13 day of December of 2016
The day's high was 116.19999694824219 on the 14 day of December of 2016
The day's high was 116.7300033569336 on the 15 day of December of 2016
The day's high was 116.5 on the 16 day of December of 2016
The day's high was 117.37999725341797 on the 19 day of December of 2016
```
###Contribute/Enquire
Create an issue or pull request to get started
###Disclaimer
The Yahoo Finance API that is implemented is not explicitly documented, maintained, or guaranteed by Yahoo, so it should not be used in mission-critical systems