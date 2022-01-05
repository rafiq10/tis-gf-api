package reports

import (
	"log"
	"net/http"

	"tis-gf-api/mydb"
	"tis-gf-api/utils"
)

type IReportsDCC interface {
	GetDCC(w http.ResponseWriter, r *http.Request)
}
type Reports struct {
	l *log.Logger
}

type ReportsDCC struct {
	l *log.Logger
}

func NewReports(l *log.Logger) *Reports {
	return &Reports{l}
}

func NewReportsDCC(l *log.Logger) *ReportsDCC {
	return &ReportsDCC{l}
}

func (acc *Reports) GetReportsLoadedResume(w http.ResponseWriter, r *http.Request) {
	myYear := r.FormValue("year")
	myMonth := r.FormValue("month")

	err := mydb.ExecuteSQL("select country,funnel_numOfertas,CPY_numPY,RfCartera_numPY,CLI_numPY,PRO_numPY,TESO_numConceptos from REPORTS_Get_Reports_Loaded_Resume("+myYear+","+myMonth+") order by country", w)
	if err != nil {
		log.Fatal(err)
	}
}

func (rep *ReportsDCC) GetDCC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var errMsg []utils.ErrMsg

	if r.FormValue("country") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: country"})
	}
	if r.FormValue("pdcyear") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: pdcyear"})
	}
	if r.FormValue("pdcmonth") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: pdcmonth"})
	}
	if r.FormValue("reportyear") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: reportyear"})
	}
	if r.FormValue("reportmonth") == "" {
		errMsg = append(errMsg, utils.ErrMsg{ErrTxt: "Empty field: reportmonth"})
	}

	if len(errMsg) != 0 {
		utils.RespondWithError(w, http.StatusBadRequest, errMsg)
		return
	}
	country := r.FormValue("country")
	pdcYear := r.FormValue("pdcyear")
	pdcMonth := r.FormValue("pdcmonth")
	reportYear := r.FormValue("reportyear")
	reportMonth := r.FormValue("reportmonth")

	mysql := "exec [sp_Reports_TESO_Get_DCC_JP]'" + country + "','" + reportYear + "','" + reportMonth + "','" + pdcYear + "','" + pdcMonth + "'"

	err := mydb.ExecuteSQL(mysql, w)
	if err != nil {
		log.Fatal(err, mysql)
	}

}
