package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	// Gửi get request đến một url nào đó
	resp, err := http.Get("http://localhost:3000/health")
	if err != nil {
		log.Fatalln(err)
	}

	// Đọc body từ response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert body sang kiểu string
	sb := string(body)
	expected := "OK"

	if sb != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			sb, expected)
	}
}
