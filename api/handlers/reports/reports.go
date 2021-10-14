package reports

import (
	"log"
	"net/http"

	reports "tis-gf-api/api/reports"
)

type Reports struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

type ReportsAging struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

type ReportsDCC struct {
	l *log.Logger
	// afs *models.AccountingFinancialStatement
}

func NewReports(l *log.Logger) *Reports {
	return &Reports{l}
}
func NewReportsAgingCLI(l *log.Logger) *ReportsAging {
	return &ReportsAging{l}
}

func NewReportsDCC(l *log.Logger) *ReportsDCC {
	return &ReportsDCC{l}
}

func (acc *Reports) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == http.MethodGet {
		err = reports.GetReportsLoadedResume(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func (rep *ReportsAging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == http.MethodGet {
		err = reports.GetAgingCLI(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (rep *ReportsDCC) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method == http.MethodGet {
		err = reports.GetDCC(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
