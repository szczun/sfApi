package main

import (
    "fmt"

    "github.com/gocolly/colly"
)

func main() {
    const url string = "https://www.sofascore.com/pl/turniej/pilka-nozna/england/premier-league/17#id:61627"
    c := colly.NewCollector()
    c.OnHTML("a[data-testid='standings_row']", func(e *colly.HTMLElement) {
        name := e.ChildText("div.Box.Flex.kQcHaX.jLRkRA.sc-4430bda6-0.dMmcA-D div.Box.ljKzDM div.Box.Flex.gHEYkP.jLRkRA div.Text.fsoviT")
        fmt.Println(name)
    })
    fmt.Println("xd")
    c.Visit(url)
}
