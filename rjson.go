package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Record struct {
	AuthorRaw json.RawMessage `json:"author"`
	Title     string          `json:"title"`
	URL       string          `json:"url"`

	Author Author
}

type Author struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

func Decode(r io.Reader) (x *Record, err error) {
	x = new(Record)
	if err = json.NewDecoder(r).Decode(x); err != nil {
		return
	}
	if err = json.Unmarshal(x.AuthorRaw, &x.Author); err == nil {
		return
	}
	var s string
	if err = json.Unmarshal(x.AuthorRaw, &s); err == nil {
		x.Author.Email = s
		return
	}
	var n uint64
	if err = json.Unmarshal(x.AuthorRaw, &n); err == nil {
		x.Author.ID = n
	}
	return
}

func main() {

	byt_1 := []byte(`{"author": 2,"title": "some things","url": "https://stackoverflow.com"}`)

	byt_2 := []byte(`{"author": "Mad Scientist","title": "some things","url": "https://stackoverflow.com"}`)

	var dat Record

	if err := json.Unmarshal(byt_1, &dat); err != nil {
		panic(err)
	}
	fmt.Printf("%#s\r\n", dat)

	if err := json.Unmarshal(byt_2, &dat); err != nil {
		panic(err)
	}
	log.Println(dat.Author.Email)
	fmt.Printf("%#s\r\n", dat)
}
