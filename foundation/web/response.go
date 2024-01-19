package web

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	StatusCode   string `json:"status_code"`
	TraceID      string `json:"trace_id"`
	CorelationID string `json:"corelation_id"`
	Data         any    `json:"data"`
}

// Respond converts a Go value to JSON and sends it to the client.
func Respond(ctx context.Context, w http.ResponseWriter, data any, statusCode int) error {
	SetStatusCode(ctx, statusCode)
	if statusCode == http.StatusNoContent {
		w.WriteHeader(statusCode)
		return nil
	}

	resp := Response{
		StatusCode:   strconv.Itoa(statusCode),
		TraceID:      GetTraceID(ctx),
		CorelationID: GetCoRelationID(ctx),
		Data:         data,
	}

	jsonData, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
