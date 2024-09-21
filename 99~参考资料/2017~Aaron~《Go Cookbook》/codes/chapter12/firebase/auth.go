package firebase

import (
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/zabawaba99/firego.v1"
)

// Authenticate grabs oauth scopes using a generated
// service_account.json file from
// https://console.firebase.google.com/project/go-cookbook/settings/serviceaccounts/adminsdk
func Authenticate() (Client, error) {
	d, err := ioutil.ReadFile("/tmp/service_account.json")
	if err != nil {
		return nil, err
	}

	conf, err := google.JWTConfigFromJSON(d, "https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/firebase.database")
	if err != nil {
		return nil, err
	}
	f := firego.New("https://go-cookbook.firebaseio.com", conf.Client(oauth2.NoContext))
	return &firebaseClient{f}, err
}
