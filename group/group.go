package group

type Group struct {
	url           string
	refresh_token string
	token         string
}

func New(url string) *Group {
	return &Group{
		url: url,
	}
}
