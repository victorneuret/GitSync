package githubLogin

import (
	"github.com/dghubble/gologin"
	"github.com/dghubble/gologin/github"
	"github.com/victorneuret/GitSync/config"
	"github.com/victorneuret/GitSync/database"
	"golang.org/x/oauth2"
	githubOAuth2 "golang.org/x/oauth2/github"
	oauth2Login "github.com/dghubble/gologin/oauth2"
	"net/http"
)

func Setup() {
	oauth2Config := &oauth2.Config{
		ClientID:     config.Config.GithubOAuth.ClientID,
		ClientSecret: config.Config.GithubOAuth.ClientSecret,
		RedirectURL:  "http://localhost:8080/github/callback",
		Endpoint:     githubOAuth2.Endpoint,
		Scopes:       []string{"read:user", "user:email", "repo"},
	}
	stateConfig:= gologin.DebugOnlyCookieConfig
	http.Handle("/github/login", github.StateHandler(stateConfig, github.LoginHandler(oauth2Config, nil)))
	http.Handle("/github/callback", github.StateHandler(stateConfig, github.CallbackHandler(oauth2Config, issueSession(), nil)))
}

func issueSession() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		token, _ := oauth2Login.TokenFromContext(ctx)
		fromContext, e := github.UserFromContext(ctx)
		githubUser, err := fromContext, e
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !database.DB.Where(&database.User{Login: *githubUser.Login}).First(&database.User{}).RecordNotFound() {
			http.Redirect(w, req, "/", http.StatusFound)
			return
		}
		user := database.User{
			Name: *githubUser.Name,
			Login: *githubUser.Login,
			Email: *githubUser.Email,
			AvatarURL: *githubUser.AvatarURL,
			Token: token.AccessToken,
		}
		database.DB.Create(&user)

		http.Redirect(w, req, "/", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}