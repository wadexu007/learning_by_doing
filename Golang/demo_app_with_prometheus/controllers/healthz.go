package controllers

import (
	"net/http"
)

// func randInt(min int, max int) int {
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	return min + rand.Intn(max-min)
// }

func HealthzHandler(w http.ResponseWriter, r *http.Request) {

	// add 10 seconds delay
	// delay := randInt(10, 2000)
	// time.Sleep(time.Millisecond * time.Duration(delay))

	w.Write([]byte("[INFO] Ping ok!\n"))
}
