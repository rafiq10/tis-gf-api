package accounting

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"tis-gf-api/api/config"
	"tis-gf-api/models"
	"tis-gf-api/mydb"
	"tis-gf-api/utils"
)

func TestMain(m *testing.M) {
	// clearTests()
	initTests()
	exitCode := m.Run()
	clearTests()
	os.Exit(exitCode)
}

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
	type fnCheck func(t *testing.T, statusCode int, data []byte)
	chkPostStCd := func() func(*testing.T, int, []byte) {
		return checkPOSTStatusCode
	}
	chkDelStCd := func() func(*testing.T, int, []byte) {
		return checkDELETEStatusCode
	}

	data := []byte(`{"country":"ECU","reportType":"BLC","ReportYear":2021,"ReportMonth": 6,"AccNum":"610025","Amount":"199.01"}`)
	dataWithoutAmount := []byte(`{"country":"ECU","reportType":"BLC","ReportYear":2021,"ReportMonth": 6,"AccNum":"610025"`)
	dataEmpty := []byte(`{}`)
	tests := []struct {
		testName string
		stCode   int
		data     []byte
		checks   fnCheck
	}{
		{"test accepted when posting valid data", http.StatusCreated, data, chkPostStCd()},
		{"deleting after posting valid data", http.StatusNotFound, data, chkDelStCd()},
		{"test that error occurs when posting empty data", http.StatusBadRequest, dataEmpty, chkPostStCd()},
		{"test post data without Amount field gives error", http.StatusBadRequest, dataWithoutAmount, chkPostStCd()},
	}
	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) { tc.checks(t, tc.stCode, tc.data) })
	}
}

func checkGETStatusCodeForCountry(t *testing.T, country string, statusCode int) {
	t.Helper()

	var path string

	acc, dbTeardown := getTestingAcc(t)
	defer dbTeardown()

	if country == "" {
		path = config.API_VERSION + "/accounting/financial-statements"
	} else {
		path = config.API_VERSION + "/accounting/financial-statements?country=" + country
	}
	request := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()
	// defer response.Result().Body.Close()

	acc.GetAllFS(response, request)
	utils.AssertResponseStatusCode(t, response.Code, statusCode)
	utils.AssertResponseHeader(t, response, "content-type", "application/json")
}

func checkGETDecodesToFS(t *testing.T, country string) {
	t.Helper()

	var path string
	var got []models.AccountingFinancialStatement

	acc, dbTeardown := getTestingAcc(t)
	defer dbTeardown()

	if country == "" {
		path = config.API_VERSION + "/accounting/financial-statements"
	} else {
		path = config.API_VERSION + "/accounting/financial-statements?country=" + country
	}
	request := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()
	// defer response.Result().Body.Close()

	acc.GetAllFS(response, request)

	err := json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Errorf("Unable to decode response from SQL server %q to accounting_FinancialStatement type, %v", response.Body, err)
	}
}

func checkPOSTStatusCode(t *testing.T, statusCode int, data []byte) {

	t.Helper()

	path := config.API_VERSION + "/accounting/financial-statements"
	acc, dbTeardown := getTestingAcc(t)
	defer dbTeardown()

	response := httptest.NewRecorder()

	request := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")
	// defer response.Result().Body.Close()

	acc.AddFinancialStatement(response, request)
	utils.AssertResponseStatusCode(t, response.Code, statusCode)
}

func checkDELETEStatusCode(t *testing.T, statusCode int, data []byte) {
	t.Helper()

	path := config.API_VERSION + "/accounting/financial-statements"
	acc, dbTeardown := getTestingAcc(t)
	defer dbTeardown()

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete, path, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")

	acc.DeleteFs(response, request)
	utils.AssertResponseStatusCode(t, response.Code, statusCode)
}

func initTests() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while opening SQL file: '%s'; error message: %s", path, err.Error())
	}
	path = path + "/init_db.sql"
	err = mydb.ExecSqlFromFile(path)
	if err != nil {
		log.Fatalf("Error while executing SQL from file: '%s', error message: %s", path, err.Error())
	}
}

func clearTests() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while opening SQL file: '%s'; error message: %s", path, err.Error())
	}
	path = path + "/clear_db.sql"
	err = mydb.ExecSqlFromFile(path)
	if err != nil {
		log.Fatalf("Error while executing SQL from file: '%s', error message: %s", path, err.Error())
	}
}

func getTestingAcc(t *testing.T) (acc *Accounting, storeTeardown func()) {
	t.Helper()
	l := log.New(os.Stdout, "tis-gf-api", log.LstdFlags)
	db, err := mydb.GetDb()
	if err != nil {
		l.Fatalf("mydb.GetDb() err=%s", err.Error())
		return nil, nil
	}

	acc = NewAccounting(l, db)
	return acc, func() { db.Close() }
}
