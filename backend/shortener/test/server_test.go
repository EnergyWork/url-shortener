package test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"
	u "url_shortener/backend/lib"
	"url_shortener/backend/shortener/api"
	"url_shortener/backend/shortener/internal/lib"
)

func TestServer(t *testing.T) {
	config := u.NewConfig("../cmd/config.yml")
	server := lib.NewServer(config)
	server.ConfigureRouter()

	if err := server.ConnectToDB(); err != nil {
		t.Fatal(err)
	}
	if err := server.Run(); err != nil {
		t.Fatal(err)
	}

	// test request

	url := "http://localhost:9000/create"

	r := api.ReqCreateShortUrl{
		UrlLong: "https://yandex.ru",
	}
	js, _ := json.Marshal(r)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(js))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Test", "testheader")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	header := resp.Header
	t.Log("response Header: ", header)

	body, _ := ioutil.ReadAll(resp.Body)
	t.Logf("response Body: %s", string(body))

	server.Shutdown(context.Background())
}

func TestValidate(t *testing.T) {
	url := "https://yandex.ru"
	urlPattern := `^((https?|ftp|file):\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	re := regexp.MustCompile(urlPattern)
	if !re.MatchString(url) {
		t.Fatal("wtf url")
	}
	t.Log("good")
}
