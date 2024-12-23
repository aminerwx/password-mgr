package api

import "net/http"

func GetVaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
func CreateVaultHandler(w http.ResponseWriter, r *http.Request) {}
func RemoveVaultHandler(w http.ResponseWriter, r *http.Request) {}
func UpdateVaultHandler(w http.ResponseWriter, r *http.Request) {}
