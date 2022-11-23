package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when should have been valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/pohui", nil)
	form := New(r.PostForm)

	form.Required("Anton Gandon", "MEOW", "YoUr LiFe Is A jOkE")
	if len(form.Errors) == 0 {
		t.Error("Shows that form valid when required fields is missing")
	}

	postedData := url.Values{}
	postedData.Add("Anton Gandon", "Anton Gandon")
	postedData.Add("MEOW", "Anton Gandon")
	postedData.Add("YoUr LiFe Is A jOkE", "Anton Gandon")

	r, _ = http.NewRequest("POST", "/pohui", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("Anton Gandon", "MEOW", "YoUr LiFe Is A jOkE")
	if len(form.Errors) != 0 {
		t.Error("shows does not have fields when it does")
	}
}
