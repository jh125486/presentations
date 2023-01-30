package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var print = fmt.Println

// START OMIT

func main() {
	print := fmt.Println
	print(OneDoesNot("learn go concurrency in one day"))
	print(Aliens("added more goroutines, program runs slower"))
	print(ToyStory("goroutines"))
	print(WinterIsComing("go generics"))
	print(BadTime("write go like java"))
	print(GrusPlan("everyone loves memes", "make a meme generator", "no one likes the memes"))
}

// END OMIT

func OneDoesNot(xyz string) (string, error) {
	return meme(61579, "one does not simply "+xyz)
}

func Aliens(xyz string) (string, error) {
	return meme(101470, xyz, "aliens")
}

func ToyStory(xyz string) (string, error) {
	return meme(347390, fmt.Sprintf("%v, %v everywhere", xyz, xyz))
}

func WinterIsComing(xyz string) (string, error) {
	return meme(61546, "brace yourself", xyz+" is coming")
}

func BadTime(xyz string) (string, error) {
	return meme(100951, "if you "+xyz, "you're gonna have a bad time")
}

func GrusPlan(x, y, z string) (string, error) {
	return meme(131940431, x, y, z, z)
}

var ErrMemeFailed = errors.New("meme failed")

const (
	u, p            = "DigitalDivas2023", "DigitalDivas2023"
	captionEndpoint = "https://api.imgflip.com/caption_image"
)

type CaptionResp struct {
	Success bool `json:"success"`
	Data    struct {
		URL     string `json:"url"`
		PageURL string `json:"page_url"`
	} `json:"data"`
	Error string `json:"error_message,omitempty"`
}

func meme(id int, text ...string) (string, error) {
	v := url.Values{}
	v.Set("template_id", strconv.Itoa(id))
	v.Set("username", u)
	v.Set("password", p)
	for i := range text {
		v.Set(fmt.Sprintf("boxes[%d][text]", i), strings.ToUpper(text[i]))
	}

	resp, err := http.PostForm(captionEndpoint, v)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var data CaptionResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", nil
	}
	if !data.Success {
		return "", fmt.Errorf("%w: %v", ErrMemeFailed, data.Error)
	}

	return data.Data.URL, nil
}
