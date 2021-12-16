package accounting

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"tis-gf-api/models"
	"tis-gf-api/mydb"
	"tis-gf-api/utils"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Accounting struct {
	l *log.Logger
}

func NewAccounting(l *log.Logger) *Accounting {
	return &Accounting{l}
}

func (a *Accounting) GetAllFS(w http.ResponseWriter, r *http.Request) {
	db, err := mydb.GetDb()
	if err != nil {
		log.Fatal(err)
	}

	if r.FormValue("country") == "" {
		getAllFinancialStatements(w, r, db)
	} else {
		getFinancialStatementsByCountry(w, r, db)
	}
}

func getAllFinancialStatements(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.Header().Set("content-type", "application/json")
	fs, err := models.AccountingFinancialStatements().All(context.Background(), db)
	if err != nil {
		http.Error(w, "Unable to get Financial Statement", http.StatusBadRequest)
	}
	utils.ToJSON(w, fs)
}

func getFinancialStatementsByCountry(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	w.Header().Set("content-type", "application/json")
	country := r.FormValue(models.AccountingFinancialStatementColumns.Country)

	exists, err := models.TblCountriesEspExists(context.Background(), db, country)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		utils.ToJSON(w, nil)
		return
	}
	fs, err := models.AccountingFinancialStatements(models.AccountingFinancialStatementWhere.Country.EQ(country)).All(context.Background(), db)

	if err != nil {
		http.Error(w, "Unable to get Financial Statement", http.StatusBadRequest)
	}
	err = utils.ToJSON(w, fs)
	if err != nil {
		http.Error(w, "Unable to read to json", http.StatusInternalServerError)
	}

}

func (a *Accounting) AddFinancialStatement(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB
	var fs models.AccountingFinancialStatement
	var err error

	buf, _ := ioutil.ReadAll(r.Body)
	data := ioutil.NopCloser(bytes.NewBuffer(buf))
	data2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	bodyBytes, _ := ioutil.ReadAll(data)
	bodyStr := string(bodyBytes)

	isJ := utils.IsJSON(bodyStr)
	if !isJ {
		fmt.Println("is NOT JSON")
		var errMsg []utils.ErrMsg
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "is not JSON: " + bodyStr})
		utils.RespondWithError(w, http.StatusBadRequest, errMsg)
		return
	}

	hasJsonField := utils.HasJsonStructFields(fs, bodyStr)

	if !hasJsonField {
		fmt.Println("NOT hasJsonField: " + bodyStr)
		var errMsg []utils.ErrMsg
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "NOT hasJsonField: " + bodyStr})
		utils.RespondWithError(w, http.StatusBadRequest, errMsg)
		return
	}
	r.Body = data2
	// bodyBytes, err := io.ReadAll(r.Body)
	err = utils.FromJSON(r.Body, &fs)
	if err != nil {
		fmt.Println("unable to unmarshall POST request utils.FromJSON:")
		var errMsg []utils.ErrMsg
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "unable to unmarshall POST: " + err.Error()})
		utils.RespondWithError(w, http.StatusBadRequest, errMsg)
		return
	}
	defer r.Body.Close()

	db, err = mydb.GetDb()

	err = fs.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		var errMsg []utils.ErrMsg
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "unable to insert into the database" + err.Error()})
		utils.RespondWithError(w, http.StatusBadRequest, errMsg)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, fs)

}

func (a *Accounting) DeleteFs(w http.ResponseWriter, r *http.Request) {
	var db *sql.DB
	var fs models.AccountingFinancialStatement
	var err error
	defer r.Body.Close()

	db, err = mydb.GetDb()
	if err != nil {
		var errMsg []utils.ErrMsg
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: err.Error()})
		utils.RespondWithError(w, http.StatusBadRequest, errMsg)
		return
	}
	err = utils.FromJSON(r.Body, &fs)
	fs.Delete(context.Background(), db)
	utils.RespondWithJSON(w, http.StatusNotFound, map[string]string{"result": "success"})
}
