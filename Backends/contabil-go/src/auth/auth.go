package auth

import (
	"fmt"

	"github.com/TalisonK/TalisonContabil/src/config"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

var params = config.GetAuthConfig()

const (
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth() {

	store := sessions.NewCookieStore([]byte(params.Key))

	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		google.New(params.Google_Client_id, params.Google_Client_Secret, fmt.Sprintf("http://localhost:%s/auth/google/callback", config.GetServerPort())),
	)

}
