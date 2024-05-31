package main

import (
	"crypto/subtle"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	AuthCodeBaseURL = "https://accounts.google.com/o/oauth2/v2/auth"
	TokenURL        = "https://oauth2.googleapis.com/token"
	ClientID        = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.apps.googleusercontent.com"
	ClientSecret    = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	State           = "cdcdjcdk3302130vd"
	RedirectURI     = "http://localhost:5000/callback"
	Scope           = "openid email"
)

type tknResponse struct {
	AccessTkn string        `json:"access_token"`
	ExpiresIn time.Duration `json:"expires_in"`
	IDTkn     string        `json:"id_token"`
	Scope     string        `json:"scope"`
	TknType   string        `json:"token_type"`
}

func main() {
	authCodeURL := fmt.Sprintf(
		"%s?client_id=%s&response_type=code&scope=%s&redirect_uri=%s&state=%s",
		AuthCodeBaseURL,
		url.QueryEscape(ClientID),
		url.QueryEscape(Scope),
		url.QueryEscape(RedirectURI),
		url.QueryEscape(State),
	)
	fmt.Printf("auth code URL: %s\n", authCodeURL)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if err := q.Get("error"); err != "" {
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte(fmt.Sprintf("%s", err)))
			return
		}
		qState := []byte(q.Get("state"))
		// prevent timing attacks on state
		if subtle.ConstantTimeCompare([]byte(State), qState) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("state does not match"))
			return
		}
		code := q.Get("code")

		q = make(url.Values)
		q.Set("code", code)
		q.Set("client_id", url.QueryEscape(ClientID))
		q.Set("client_secret", url.QueryEscape(ClientSecret))
		q.Set("grant_type", "authorization_code")
		q.Set("redirect_uri", url.QueryEscape("http://localhost:5000"))

		resp, err := http.PostForm(TokenURL, q)
		if err != nil {
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte(err.Error()))
			return
		}
		// if resp.StatusCode != http.StatusOK {
		// 	w.WriteHeader(http.StatusExpectationFailed)
		// 	w.Write([]byte(fmt.Sprintf("not OK status code: %s", resp.Status)))
		// 	return
		// }
		// tknResp := new(tknResponse)
		// if err := json.NewDecoder(resp.Body).Decode(tknResp); err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte("failed to unmarshal token response body"))
		// 	return
		// }
		tknBody, err := io.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to read token response body"))
			return
		}
		w.Write(tknBody)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Google OIDC"))
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}
