package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/iobear/uxt/uxt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch {
	case strings.HasPrefix(path, "/current"):
		fmt.Fprint(w, uxt.GetCurrentUnixTime())
	case strings.HasPrefix(path, "/since/"):
		parts := strings.Split(path, "/")
		if len(parts) != 3 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		unixTime, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			http.Error(w, "Invalid timestamp", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, uxt.GetTimeSince(unixTime))
	case strings.HasPrefix(path, "/convert/"):
		parts := strings.Split(path, "/")
		if len(parts) < 3 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		unixTime, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			http.Error(w, "Invalid timestamp", http.StatusBadRequest)
			return
		}
		format := ""
		if len(parts) == 4 {
			format = parts[3]
		}
		result, err := uxt.ConvertUnixTimeToFormattedString(unixTime, format)
		if err != nil {
			http.Error(w, "Error converting timestamp", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, result)
	case strings.HasPrefix(path, "/plus/"):
		parts := strings.Split(path, "/")
		if len(parts) != 3 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}
		result, err := uxt.AdjustCurrentUnixTime(value)
		if err != nil {
			http.Error(w, "Error calculating time", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, result)
	case strings.HasPrefix(path, "/minus/"):
		parts := strings.Split(path, "/")
		if len(parts) != 3 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}
		result, err := uxt.AdjustCurrentUnixTime(-value) // Note the minus sign here to subtract
		if err != nil {
			http.Error(w, "Error calculating time", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, result)
	case strings.HasPrefix(path, "/rfc3339/"):
		parts := strings.Split(path, "/")
		if len(parts) != 3 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		value, err := strconv.ParseInt(parts[2], 10, 64)
		if err != nil {
			http.Error(w, "Invalid number", http.StatusBadRequest)
			return
		}
		result, err := uxt.ConvertUnixTimeToFormattedString(value, time.RFC3339)
		if err != nil {
			http.Error(w, "Error converting time", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, result)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func RunServer(port string) {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
