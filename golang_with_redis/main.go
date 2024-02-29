package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go and redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

https://tutorialedge.net/golang/go-redis-tutorial/
https://www.google.com/search?q=message+broker+with+golang&sca_esv=b378cbf59b0b044f&ei=CjnXZdb5CbaTwPAP29W9gAg&udm=&ved=0ahUKEwjWpu3Q9L6EAxW2CRAIHdtqD4AQ4dUDCBA&uact=5&oq=message+broker+with+golang&gs_lp=Egxnd3Mtd2l6LXNlcnAiGm1lc3NhZ2UgYnJva2VyIHdpdGggZ29sYW5nMgYQABgWGB5Ivw1QYFjEDHABeAGQAQCYAbIBoAG-DKoBBDAuMTK4AQPIAQD4AQHCAgoQABhHGNYEGLADwgINEAAYgAQYigUYQxiwA8ICBRAAGIAEwgIKEAAYgAQYigUYQ8ICCBAAGBYYHhgPwgIJEAAYgAQYDRgTwgIIEAAYFhgeGBPCAgcQABiABBgTiAYBkAYK&sclient=gws-wiz-serp
https://habr.com/ru/articles/774636/
https://habr.com/ru/companies/wunderfund/articles/685894/

