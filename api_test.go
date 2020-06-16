package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"toury_bakcend/src/config"
)

func TestCreateTourney(t *testing.T) {

	jsonStr := []byte(`[{"name":"Roma","teams_count":4,"required_teams_ids":[2,8]}, {"name":"Matvey","teams_count":6,"required_teams_ids":[1,7,9]}]`)
	req, err := http.NewRequest("POST", config.TourneyCreateUri, bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(createTourney)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"}`
	if recorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expected)
	}
}
