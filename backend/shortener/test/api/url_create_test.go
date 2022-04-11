package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"url_shortener/backend/shortener/api"
)

func TestUrlCreate(t *testing.T) {
	// test request
	url := "http://localhost:9000/create"
	r := api.ReqCreateShortUrl{
		UrlLong: "https://yandex.ru",
	}
	js, _ := json.Marshal(r)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(js))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Request-Name", "url/short/create")
	req.Header.Add("X-Signature", "test_signature_sd109kjd93j1")
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
}
