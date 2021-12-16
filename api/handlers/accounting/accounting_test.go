package accounting

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tis-gf-api/api/config"
	"tis-gf-api/models"
	"tis-gf-api/utils"
)

func TestGetAllFS(t *testing.T) {
	countries := []string{"ESP", "CHI", "COL", "PER", "ECU"}
	t.Run("it returns 200 on /accounting/financial-statements", func(t *testing.T) {
		checkGETStatusCodeForCountry(t, "", http.StatusOK)
	})

	for _, c := range countries {
		t.Run("it returns 200 on /accounting/financial-statements?country="+c, func(t *testing.T) {
			checkGETStatusCodeForCountry(t, c, http.StatusOK)
		})
		t.Run("check if fits fs on /accounting/financial-statements?country="+c, func(t *testing.T) {
			checkGETDecodesToFS(t, c)
		})
	}

	t.Run("returns 404 on /accounting/financial-statements?country=POL", func(t *testing.T) {
		checkGETStatusCodeForCountry(t, "POL", http.StatusNotFound)
	})

	t.Run("check if fits fs on /accounting/financial-statements", func(t *testing.T) {
		checkGETDecodesToFS(t, "")
	})
}

func TestAddFinancialStatement(t *testing.T) {
	data := []byte(`{"country":"ECU","reportType":"BLC","ReportYear":2021,"ReportMonth": 6,"AccNum":"610025","Amount":"199.01"}`)

	t.Run("test accepted when posting valid data", func(t *testing.T) {
		checkPOSTStatusCode(t, http.StatusCreated, data)
	})
	t.Run("deleting after posting valid data", func(t *testing.T) {
		checkDELETEStatusCode(t, http.StatusNotFound, data)
	})

	t.Run("test that error occurs when posting empty data", func(t *testing.T) {
		data := []byte(`{}`)
		checkPOSTStatusCode(t, http.StatusBadRequest, data)
	})

	t.Run("test post data without Amount field gives error", func(t *testing.T) {
		d := []byte(`{"country":"ECU","reportType":"BLC","ReportYear":2021,"ReportMonth": 6,"AccNum":"610025"`)
		checkPOSTStatusCode(t, http.StatusBadRequest, d)
	})
}

func checkGETStatusCodeForCountry(t *testing.T, country string, statusCode int) {
	t.Helper()

	var path string

	acc := &Accounting{}
	if country == "" {
		path = config.API_VERSION + "/accounting/financial-statements"
	} else {
		path = config.API_VERSION + "/accounting/financial-statements?country=" + country
	}
	request := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()
	acc.GetAllFS(response, request)
	utils.AssertResponseStatusCode(t, response.Code, statusCode)
	utils.AssertResponseHeader(t, response, "content-type", "application/json")
}

func checkGETDecodesToFS(t *testing.T, country string) {
	t.Helper()

	var path string
	var got []models.AccountingFinancialStatement

	acc := &Accounting{}
	if country == "" {
		path = config.API_VERSION + "/accounting/financial-statements"
	} else {
		path = config.API_VERSION + "/accounting/financial-statements?country=" + country
	}
	request := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()
	acc.GetAllFS(response, request)

	err := json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Errorf("Unable to decode response from SQL server %q to accounting_FinancialStatement type, %v", response.Body, err)
	}
}

func checkPOSTStatusCode(t *testing.T, statusCode int, data []byte) {
	t.Helper()

	path := config.API_VERSION + "/accounting/financial-statements"
	response := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	acc := &Accounting{}
	acc.AddFinancialStatement(response, request)
	utils.AssertResponseStatusCode(t, response.Code, statusCode)
}

func checkDELETEStatusCode(t *testing.T, statusCode int, data []byte) {
	t.Helper()

	path := config.API_VERSION + "/accounting/financial-statements"
	response := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodDelete, path, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	acc := &Accounting{}
	acc.DeleteFs(response, request)
	utils.AssertResponseStatusCode(t, response.Code, statusCode)
}
