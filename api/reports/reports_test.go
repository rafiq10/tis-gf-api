package reports_mgt

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"tis-gf-api/api/config"
)

func TestGetReportsLoadedResume(t *testing.T) {
	t.Run("returns reports", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, config.API_VERSION+"/reports/mgt/reports-loaded-resume?year=2021&month=5", nil)
		response := httptest.NewRecorder()

		GetReportsLoadedResume(response, request)
		got := response.Body.String()
		want := `[{"country":"CHI","funnel_numOfertas":114,"CPY_numPY":1416,"RfCartera_numPY":334,"CLI_numPY":0,"PRO_numPY":0,"TESO_numConceptos":0},{"country":"COL","funnel_numOfertas":39,"CPY_numPY":22,"RfCartera_numPY":9,"CLI_numPY":0,"PRO_numPY":0,"TESO_numConceptos":0},{"country":"ECU","funnel_numOfertas":14,"CPY_numPY":123,"RfCartera_numPY":29,"CLI_numPY":9,"PRO_numPY":26,"TESO_numConceptos":0},{"country":"ESP","funnel_numOfertas":427,"CPY_numPY":3834,"RfCartera_numPY":612,"CLI_numPY":710,"PRO_numPY":452,"TESO_numConceptos":0},{"country":"PER","funnel_numOfertas":341,"CPY_numPY":1882,"RfCartera_numPY":437,"CLI_numPY":228,"PRO_numPY":0,"TESO_numConceptos":0}]` + "\n"
		if got != want {
			t.Errorf("wanted: %q but got: %q", want, got)
		}
	})
}
