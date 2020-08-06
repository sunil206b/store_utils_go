package http_utils

import (
	"encoding/json"
	"github.com/sunil206b/store_utils_go/errors"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func ResponseError(w http.ResponseWriter, errMsg errors.RestErr) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errMsg.StatusCode)
	json.NewEncoder(w).Encode(errMsg)
}