package fireball

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

)

func RecordJSONResponse(t *testing.T, resp Response, v interface{}) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	resp.Write(recorder, nil)
	if v != nil {
		if err := json.Unmarshal(recorder.Body.Bytes(), v); err != nil {
			t.Fatal(err)
		}
	}

	return recorder
}
