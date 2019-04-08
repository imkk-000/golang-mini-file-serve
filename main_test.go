package main_test

import (
	. "miniserve"
	"net/http/httptest"
	"testing"
)

func TestCallAPI(t *testing.T) {
	expectedContentType := "application/octet-stream"
	expectedContentDisposition := "attachment; filename=file.zip"
	server := httptest.NewServer(APIHandler{})
	defer server.Close()
	client := server.Client()

	response, _ := client.Get(server.URL)
	actualContentType := response.Header.Get("Content-Type")
	actualContentDisposition := response.Header.Get("Content-Disposition")

	if expectedContentType != actualContentType {
		t.Errorf("expect '%s' but it got '%s'", expectedContentType, actualContentType)
	}
	if expectedContentDisposition != actualContentDisposition {
		t.Errorf("expect '%s' but it got '%s'", expectedContentDisposition, actualContentDisposition)
	}
}
