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

func TestForm_Has(t *testing.T) {
	r := httptest.NewRequest("POST", "/pohuy", nil)
	form := New(r.PostForm)

	if form.Has("SFsafs") {
		t.Error("Form shows has field when it does not")
	}

	postedData := url.Values{}
	postedData.Add("Anton Gandon", "Anton Gandon")
	form = New(postedData)

	if !form.Has("Anton Gandon") {
		t.Error("Got false when we should have been true")
	}
}

func TestForm_MinLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/pohuy", nil)
	form := New(r.PostForm)

	if form.MinLength("Anton Gandon", 3) {
		t.Error("Form shows that field have required length when it's not")
	}

	isError := form.Errors.Get("Anton Gandon")
	if isError == "" {
		t.Error("Should have an error, but did not get one")
	}

	postedData := url.Values{}
	postedData.Add("Anton Gandon", "Anton Gandon")
	form = New(postedData)

	if !form.MinLength("Anton Gandon", 3) {
		t.Error("Form shows that field doesnt have required length when it is")
	}

	isError = form.Errors.Get("Anton Gandon")
	if isError != "" {
		t.Error("Have an error, when not suposed to")
	}

	postedData = url.Values{}
	postedData.Add("Anton Super Gandon", "joke")
	form = New(postedData)

	if form.MinLength("Anton Super Gandon", 300) {
		t.Error("Form shows that field have required length when it doesnt")
	}
}

func TestForm_IsEmail(t *testing.T) {
	postedData := url.Values{}
	form := New(postedData)

	form.IsEmail("")
	if form.Valid() {
		t.Error("Show that email is valid when email is not exist")
	}

	postedData = url.Values{}
	postedData.Add("email", "Kekw@anytext.com")
	form = New(postedData)

	form.IsEmail("email")
	if !form.Valid() {
		t.Error("Show that email invalid when it valid")
	}

	postedData = url.Values{}
	postedData.Add("email", "kekw")
	form = New(postedData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("Show that email valid when it invalid")
	}
}
