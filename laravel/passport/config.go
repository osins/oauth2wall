package passport

type Config struct {
	AuthURL      string
	TokenURL     string
	UserURL      string
	CallbackURL  string
	ClientID     string
	ClientSecret string
	SessionKey   string
}
