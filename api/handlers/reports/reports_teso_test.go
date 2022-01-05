package reports

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"tis-gf-api/api/config"
	"tis-gf-api/utils"
)

type ReportsDccFake struct {
	l           *log.Logger
	country     string
	reportYear  string
	reportMonth string
	pdcYear     string
	pdcMonth    string
	t           *testing.T
}

func TestReportsDCC_GetDCC(t *testing.T) {
	myLog := log.New(os.Stdout, "tis-gf-api", log.LstdFlags)

	type fields struct {
		l           *log.Logger
		country     string
		reportYear  string
		reportMonth string
		pdcYear     string
		pdcMonth    string
	}

	tests := []struct {
		name         string
		fields       fields
		wantStatus   int
		wantRespText string
	}{
		{"test ESP 2021 Nov gets statusOK",
			fields{
				l:           myLog,
				country:     "ESP",
				reportYear:  "2021",
				reportMonth: "11",
				pdcYear:     "2021",
				pdcMonth:    "12",
			},
			http.StatusOK,
			"",
		},
		{"test empty country gets StatusBadRequest",
			fields{
				l:           myLog,
				country:     "",
				reportYear:  "2021",
				reportMonth: "11",
				pdcYear:     "2021",
				pdcMonth:    "12",
			},
			http.StatusBadRequest,
			`[{"error":"Empty field: country"}]`,
		},
		{"test empty reportYear gets StatusBadRequest",
			fields{
				l:           myLog,
				country:     "ECU",
				reportYear:  "",
				reportMonth: "11",
				pdcYear:     "2021",
				pdcMonth:    "12",
			},
			http.StatusBadRequest,
			`[{"error":"Empty field: reportyear"}]`,
		},
		{"test empty country and reportYear gets StatusBadRequest",
			fields{
				l:           myLog,
				country:     "",
				reportYear:  "",
				reportMonth: "11",
				pdcYear:     "2021",
				pdcMonth:    "12",
			},
			http.StatusBadRequest,
			`[{"error":"Empty field: country"},{"error":"Empty field: reportyear"}]`,
		},
		{"test empty all fields gets StatusBadRequest",
			fields{
				l:           myLog,
				country:     "",
				reportYear:  "",
				reportMonth: "",
				pdcYear:     "",
				pdcMonth:    "",
			},
			http.StatusBadRequest,
			`[{"error":"Empty field: country"},{"error":"Empty field: pdcyear"},{"error":"Empty field: pdcmonth"},{"error":"Empty field: reportyear"},{"error":"Empty field: reportmonth"}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := &ReportsDccFake{
				l:           tt.fields.l,
				country:     tt.fields.country,
				reportYear:  tt.fields.reportYear,
				reportMonth: tt.fields.reportMonth,
				pdcYear:     tt.fields.pdcYear,
				pdcMonth:    tt.fields.pdcMonth,
				t:           t,
			}

			path := config.API_VERSION + "/reports/mgt/reports-dcc?country=" + rep.country +
				"&reportyear=" + rep.reportYear + "&reportmonth=" + rep.reportMonth + "&pdcyear=" + rep.pdcYear +
				"&pdcmonth=" + rep.pdcMonth

			// fmt.Println(path)

			dcc := &ReportsDCC{}
			request := httptest.NewRequest(http.MethodGet, path, nil)
			response := httptest.NewRecorder()
			dcc.GetDCC(response, request)

			utils.AssertResponseStatusCode(t, response.Code, tt.wantStatus)
			utils.AssertResponseHeader(t, response, "content-type", "application/json")
			AssertResponseText_DCC(t, *response, tt.wantRespText)

		})
	}
}

func AssertResponseText_DCC(t *testing.T, response httptest.ResponseRecorder, wantRespText string) {
	t.Helper()
	if wantRespText != "" {
		respResult := response.Result()
		defer respResult.Body.Close()
		data, err := ioutil.ReadAll(respResult.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
		utils.AssertResponseText(t, string(data), wantRespText)
	}
}
