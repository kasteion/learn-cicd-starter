package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	expected := ""
	expectedError := ErrNoAuthHeaderIncluded

	headers := http.Header{}

	actual, err := GetAPIKey(headers)

	if actual != expected {
		t.Errorf("got %q; expected %q", actual, expected)
	}

	if !errors.Is(err, expectedError) {
		t.Errorf("got err %q; expected err %q", err, expectedError)
	}
}

func TestGetAPIKey_MalformedAuthHeader(t *testing.T) {
	expected := ""
	expectedError := errors.New("malformed authorization header")

	headers := http.Header{
		"Authorization": []string{"malformed"},
	}

	actual, err := GetAPIKey(headers)

	if actual != expected {
		t.Errorf("got %q; expected %q", actual, expected)
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("got err %q; expected err %q", err, expectedError)
	}
}

func TestGetAPIKey_Success(t *testing.T) {
	expected := "XXX"
	headers := http.Header{
		"Authorization": []string{"ApiKey XXX"},
	}

	actual, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("unexpected error %q", err)
	}

	if actual != expected {
		t.Errorf("got %q; expected %q", actual, expected)
	}

}
