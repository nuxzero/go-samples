package foo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	Greeting(response, request)

	got := response.Body.String()
	want := "Hello, World!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestSayGoodbye(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/goodbye", nil)
	response := httptest.NewRecorder()

	SayGoodbye(response, request)

	got := response.Body.String()
	want := "Goodbye, World!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestTalk(t *testing.T) {
	body := strings.NewReader("message=How are you?")
	request, _ := http.NewRequest(http.MethodPost, "/talk", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response := httptest.NewRecorder()

	Talk(response, request)

	got := response.Body.String()
	want := "Talk, How are you?"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
