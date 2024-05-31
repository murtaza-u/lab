package main

import (
	"bytes"
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	clientID        = "XXXXXXXXXXXXXXXXXXXX"
	clientSecret    = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	authBaseCodeURL = "https://github.com/login/oauth/authorize"
	accessTknURL    = "https://github.com/login/oauth/access_token"
	userURL         = "https://api.github.com/user"
)

var state = ""

func randomBytesInHex(count int) (string, error) {
	buf := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

func getAuthCodeURL(state string) string {
	return fmt.Sprintf("%s?scope=%s&client_id=%s&state=%s",
		authBaseCodeURL, "read:user", clientID, state)
}

type AcccessTknParams struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type AccessTknResponse struct {
	Token string `json:"access_token"`
	Scope string `json:"scope"`
	Typ   string `json:"token_type"`
}

type User struct {
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
}

func callback(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	code := q.Get("code")
	state := q.Get("state")
	if subtle.ConstantTimeCompare([]byte(state), []byte(state)) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid state"))
		return
	}
	fmt.Printf("Code: %s\n", code)
	fmt.Println("Exchanging authorization code with access token")

	params := AcccessTknParams{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         code,
	}
	body, err := json.Marshal(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive access token: %s",
			err.Error())))
		return
	}
	fmt.Println("Marshalled access token params")

	req, err := http.NewRequest(http.MethodPost, accessTknURL, bytes.NewReader(body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive access token: %s",
			err.Error())))
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive access token: %s",
			err.Error())))
		return
	}
	fmt.Println("Received response from access token API")

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(
			"failed to retreive access token: status %s", resp.Status)))
		return
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive access token: %s",
			err.Error())))
		return
	}
	tkn := new(AccessTknResponse)
	err = json.Unmarshal(body, tkn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive access token: %s",
			err.Error())))
		return
	}
	fmt.Println("Marshalled access token response")
	fmt.Printf("Access token: %s | type: %s | scope: %s\n",
		tkn.Token, tkn.Typ, tkn.Scope)

	req, err = http.NewRequest(http.MethodGet, userURL, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive user: %s", err.Error())))
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tkn.Token))

	client = new(http.Client)
	resp, err = client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive user: %s", err.Error())))
		return
	}
	fmt.Println("Received response from user API")

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive user: %s", resp.Status)))
		return
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive user: %s", err.Error())))
		return
	}

	user := new(User)
	err = json.Unmarshal(body, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retreive user: %s", err.Error())))
		return
	}
	fmt.Println("Marshalled user response")
	fmt.Printf("Name: %s | E-Mail: %s | Avatar URL: %s | Bio: %s\n",
		user.Name, user.Email, user.AvatarURL, user.Bio)
}

func main() {
	var err error
	state, err = randomBytesInHex(24)
	if err != nil {
		log.Fatal("failed to generate state")
	}
	fmt.Printf("Authorization Code URL: %s\n", getAuthCodeURL(state))

	http.HandleFunc("/callback", callback)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
