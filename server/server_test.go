package server

import (
	"bytes"
	"log"
	"merchant-service/api/v1/handlers"
	"merchant-service/api/v1/services"
	"merchant-service/config"
	"merchant-service/database"
	"merchant-service/repository"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetMerchantList(t *testing.T) {
	merchantHandler := getMerchantHandler()

	req, err := http.NewRequest("GET", "/api/v1/merchant/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("page", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(merchantHandler.GetMerchantList)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"success":true,"responseCode":200,"data":{"merchants":[{"id":7,"name":"merchant 6","code":"merchant6","status":1,"created_at":"2022-06-04T10:07:55Z","updated_at":"2022-06-04T10:07:55Z"},{"id":6,"name":"merchant 5","code":"merchant5","status":1,"created_at":"2022-06-04T09:07:44Z","updated_at":"2022-06-04T09:07:44Z"}],"pagination":{"next_page":true,"previous_page":false,"page":1}},"error":null}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreatetMerchant(t *testing.T) {
	merchantHandler := getMerchantHandler()

	var jsonStr = []byte(`{"name": "merchant 7","code": "merchant7","status": 1}`)

	req, err := http.NewRequest("POST", "/api/v1/merchant/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(merchantHandler.CreatetMerchant)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success":true,"responseCode":200,"data":"Merchant created successfully.","error":null}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetMerchantTeamList(t *testing.T) {
	merchantTeamHandler := getMerchantTeamHandler()

	req, err := http.NewRequest("GET", "/api/v1/merchant-team/list", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("page", "1")
	q.Add("merchant_id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(merchantTeamHandler.GetMerchantTeamList)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"success":true,"responseCode":200,"data":{"merchant_teams":[{"id":3,"merchant_id":1,"name":"team3","email":"merchant1.team3@gmail.com","status":1,"created_at":"2022-06-04T02:00:33Z","updated_at":"2022-06-04T02:00:33Z"},{"id":2,"merchant_id":1,"name":"team2","email":"merchant1.team2@gmail.com","status":1,"created_at":"2022-06-04T01:58:30Z","updated_at":"2022-06-04T01:59:58Z"}],"pagination":{"next_page":true,"previous_page":false,"page":1}},"error":null}
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreatetMerchantTeam(t *testing.T) {
	merchantTeamHandler := getMerchantTeamHandler()

	var jsonStr = []byte(`{"merchant_id":2,"name":"team6","email":"merchant2.team6@gmail.com","status": 1}`)

	req, err := http.NewRequest("POST", "/api/v1/merchant-team/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(merchantTeamHandler.CreateMerchantTeam)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"success":true,"responseCode":200,"data":"Merchant team created successfully.","error":null}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func getMerchantHandler() *handlers.MerchantHandler {
	os.Setenv("ENVIRONMENT", "test")
	err := config.SetupConfig()
	if err != nil {
		log.Fatal("Can not load config.")
	}
	config.LoadConfig()

	dbConn := database.DBConnection().GetConnection()

	if dbConn == nil {
		log.Fatal("Expecting db connection object but received nil")
	}
	merchantRepo := repository.NewMerchantRepository(dbConn)
	merchantService := services.NewMerchantService(merchantRepo)
	return handlers.MerchantHttpHandler(merchantService)
}

func getMerchantTeamHandler() *handlers.MerchantTeamHandler {
	os.Setenv("ENVIRONMENT", "test")
	err := config.SetupConfig()
	if err != nil {
		log.Fatal("Can not load config.")
	}
	config.LoadConfig()

	dbConn := database.DBConnection().GetConnection()

	if dbConn == nil {
		log.Fatal("Expecting db connection object but received nil")
	}
	merchantTeamRepo := repository.NewMerchantTeamRepository(dbConn)
	merchantTeamService := services.NewMerchantTeamService(merchantTeamRepo)
	return handlers.MerchantTeamHttpHandler(merchantTeamService)
}
