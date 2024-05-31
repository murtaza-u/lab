package main

import (
	"context"
	"crypto/subtle"
	"fmt"
	"log"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

const (
	ISSUER        = "http://localhost:8080/realms/shiganshina"
	CLIENT_ID     = "demo"
	CLIENT_SECRET = "Q9o9JLGrphnbgarLjYzLipBmHJFv0vSx"
	State         = "foobar"
)

type Srv struct {
	oauth2.Config
	provider *oidc.Provider
}

func main() {
	ctx := context.Background()
	p, err := oidc.NewProvider(ctx, ISSUER)
	if err != nil {
		log.Fatalf("failed to setup oidc provider: %s", err.Error())
	}

	srv := Srv{
		Config: oauth2.Config{
			Endpoint:     p.Endpoint(),
			ClientID:     CLIENT_ID,
			ClientSecret: CLIENT_SECRET,
			Scopes:       []string{"openid", "email", "profile"},
			RedirectURL:  "http://localhost:5000/callback",
		},
		provider: p,
	}

	e := echo.New()
	e.GET("/", srv.index)
	e.GET("/login", srv.login)
	e.GET("/callback", srv.callback)
	e.Logger.Fatal(e.Start(":5000"))
}

type Map map[string]any

func (s Srv) index(c echo.Context) error {
	return c.HTML(http.StatusOK, `<a href="/login">Login</a>`)
}

func (s Srv) login(c echo.Context) error {
	url := s.AuthCodeURL(State)
	return c.Redirect(http.StatusSeeOther, url)
}

type callbackParams struct {
	Err     string `query:"error"`
	ErrDesc string `query:"error_description"`
	Code    string `query:"code"`
	State   string `query:"state"`
}

func (s Srv) callback(c echo.Context) error {
	p := new(callbackParams)
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, Map{
			"error": "failed to parse query params",
		})
	}

	if p.Err != "" {
		return c.JSON(http.StatusBadRequest, Map{
			"error": p.Err,
		})
	}

	// prevent timing attacks on state
	if subtle.ConstantTimeCompare([]byte(State), []byte(p.State)) == 0 {
		return c.JSON(http.StatusBadRequest, Map{
			"error": fmt.Sprintf("invalid state: want: %s got: %s",
				State,
				p.State,
			),
		})
	}

	tkn, err := s.Exchange(c.Request().Context(), p.Code)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Map{
			"error": fmt.Sprintf("token exchange failed: %s", err.Error()),
		})
	}

	return c.JSON(http.StatusOK, tkn)
}
