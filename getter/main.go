package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

const AUTH_URL = ""
const SEASON_URL = "https://members-ng.iracing.com/data/season/list?season_year=2024&season_quarter=3"

func getCredentials() (string, string) {
	var email, password string
	fmt.Print("Email:")
	fmt.Scan(&email)
	fmt.Print("Password:")
	fmt.Scan(&password)
	return email, password
}

func makeAuthPayload(email string, password string) map[string]string {
	var joined = password + strings.ToLower(email)
	var s256 = sha256.Sum256([]byte(joined))
	var encoded = base64.StdEncoding.EncodeToString(s256[:])
	return map[string]string{
		"email":    email,
		"password": encoded}
}

func main() {
	fmt.Print(makeAuthPayload(getCredentials()))
}
