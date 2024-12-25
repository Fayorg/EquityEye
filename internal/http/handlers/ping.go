package handlers

import (
	"EquityEye/pkg/helpers"
	"EquityEye/types"
	"net/http"
)

func PingHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helpers.Encode(w, r, http.StatusOK, &types.ApiResponse{
			Status:  http.StatusOK,
			Message: "pong",
		})
	})
}
