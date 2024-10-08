package restclient

import (
	"bytes"
	"net/http"
	"testing"
)

func TestStartMockups(t *testing.T) {
	StartMockups()
	if !enableMocks {
		t.Error("expected enableMocks to be true, got false")
	}
}

func TestStopMockups(t *testing.T) {
	StopMockups()
	if enableMocks {
		t.Error("expected enableMocks to be false, got true")
	}
}

func TestFlushMockups(t *testing.T) {
	mocks["test"] = &Mock{}
	FlushMockups()
	if len(mocks) != 0 {
		t.Error("expected mocks to be empty, got non-empty")
	}
}

func TestAddMockup(t *testing.T) {
	mock := Mock{
		Url:        "http://example.com",
		HttpMethod: http.MethodPost,
		Response:   &http.Response{StatusCode: 200},
		Err:        nil,
	}
	AddMockup(mock)
	if len(mocks) != 1 {
		t.Error("expected mocks to have one entry, got", len(mocks))
	}
}

func TestPostWithMock(t *testing.T) {
	StartMockups()
	defer StopMockups()

	mock := Mock{
		Url:        "http://example.com",
		HttpMethod: http.MethodPost,
		Response:   &http.Response{StatusCode: 200},
		Err:        nil,
	}
	AddMockup(mock)

	response, err := Post("http://example.com", nil, nil)
	if err != nil {
		t.Error("expected no error, got", err)
	}
	if response.StatusCode != 200 {
		t.Error("expected status code 200, got", response.StatusCode)
	}
}

func TestPostWithoutMock(t *testing.T) {
	StopMockups()

	_, err := Post("http://example.com", nil, nil)
	if err == nil {
		t.Error("expected an error, got none")
	}
}

func TestPostWithRealRequest(t *testing.T) {
	StopMockups()

	body := map[string]string{"key": "value"}
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			buf := new(bytes.Buffer)
			buf.ReadFrom(r.Body)
			if buf.String() != `{"key":"value"}` {
				t.Error("expected body to be {\"key\":\"value\"}, got", buf.String())
			}
			w.WriteHeader(http.StatusOK)
		}),
	}
	go server.ListenAndServe()
	defer server.Close()

	response, err := Post("http://localhost:8080", body, headers)
	if err != nil {
		t.Error("expected no error, got", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Error("expected status code 200, got", response.StatusCode)
	}
}
