package wykop

import (
	"encoding/json"
	"io/ioutil"
)

type Profile struct {
	Login             string `json:"login"`
	RegistrationEmail string `json:"email"`
	Email             string `json:"public_email"`
	Name              string `json:"name"`
	WWW               string `json:"www"`
	Jabber            string `json:"jabber"`
	GG                string `json:"gg"`
	City              string `json:"city"`
	About             string `json:"about"`
	AuthorGroup       int    `json:"author_group"`
	LinksAdded        int    `json:"links_added"`
	LinksPublished    int    `json:"links_published"`
	Comments          int    `json:"comments"`
	Rank              int    `json:"rank"`
	Followers         int    `json:"followers"`
	Following         int    `json:"following"`
	Entries           int    `json:"entries"`
	EntriesComments   int    `json:"entries_comments"`
	Diggs             int    `json:"diggs"`
	Buries            int    `json:"buries"`
	Groups            int    `json:"groups"`
	RelatedLinks      int    `json:"related_links"`
	SignupDate        string `json:"signup_date"`
	Avatar            string `json:"avatar"`
	AvatarBig         string `json:"avatar_big"`
	AvatarMed         string `json:"avatar_med"`
	AvatarLo          string `json:"avatar_lo"`
	IsObserved        bool   `json:"is_observed"`
	Sex               string `json:"sex"`
	URL               string `json:"url"`
	ViolationURL      string `json:"violation_url"`
}

func (s *Session) GetProfile(username string) (Profile, error) {
	url := baseUri + "/Profile/Index/" + username + "/appkey," + s.appkey
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return Profile{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Profile{}, err
	}
	var p Profile
	json.Unmarshal(body, &p)
	return p, nil
}
