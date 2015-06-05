package wykop

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
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

type Comment struct {
	Body            string `json:"body,omitempty"`
	Status          string `json:"status,omitempty"`
	Type            string `json:"type,omitempty"`
	App             string `json:"app,omitempty"`
	Date            string `json:"date,omitempty"`
	ID              int    `json:"id,omitempty"`
	ParentID        int    `json:"parent_id,omitempty"`
	Author          string `json:"author,omitempty"`
	AuthorSex       string `json:"author_sex,omitempty"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	Blocked         bool   `json:"blocked,omitempty"`
	UserFavorite    bool   `json:"user_favorite,omitempty"`
	CanVote         bool   `json:"can_vote,omitempty"`
	Deleted         bool   `json:"deleted,omitempty"`
	VoteCount       int    `json:"vote_count,omitempty"`
	VoteCountPlus   int    `json:"vote_count_plus,omitempty"`
	VoteCountMinus  int    `json:"vote_count_minus,omitempty"`
	ViolationURL    string `json:"violation_url,omitempty"`
	Link            Link   `json:"link,omitempty"`
	Source          string `json:"source,omitempty"`
	Embed           Embed  `json:"embed,omitempty"`
}

type RelatedLink struct {
	Title           string `json:"title,omitempty"`
	Author          string `json:"author,omitempty"`
	AuthorSex       string `json:"author_sex,omitempty"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	URL             string `json:"url,omitempty"`
	ViolationURL    string `json:"violation_url,omitempty"`
	VoteCount       int    `json:"vote_count,omitempty"`
	UserVote        int    `json:"user_vote,omitempty"`
	Link            Link   `json:"link,omitempty"`
	Type            string `json:"type,omitempty"`
	ID              int    `json:"id,omitempty"`
	Plus18          bool   `json:"plus18,omitempty"`
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

func (s *Session) GetProfileAdded(username string, page int) ([]Link, error) {
	url := baseUri + "/Profile/Added/" + username + "/appkey," + s.appkey + ",page," + strconv.Itoa(page)
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return []Link{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Link{}, err
	}
	var ls []Link
	json.Unmarshal(body, &ls)
	return ls, nil
}

func (s *Session) GetProfilePublished(username string, page int) ([]Link, error) {
	url := baseUri + "/Profile/Published/" + username + "/appkey," + s.appkey + ",page," + strconv.Itoa(page)
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return []Link{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Link{}, err
	}
	var ls []Link
	json.Unmarshal(body, &ls)
	return ls, nil
}

func (s *Session) GetProfileCommented(username string, page int) ([]Link, error) {
	url := baseUri + "/Profile/Commented/" + username + "/appkey," + s.appkey + ",page," + strconv.Itoa(page)
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return []Link{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Link{}, err
	}
	var ls []Link
	json.Unmarshal(body, &ls)
	return ls, nil
}

func (s *Session) GetProfileComments(username string, page int) ([]Comment, error) {
	url := baseUri + "/Profile/Comments/" + username + "/appkey," + s.appkey + ",page," + strconv.Itoa(page)
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return []Comment{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Comment{}, err
	}
	var cs []Comment
	json.Unmarshal(body, &cs)
	return cs, nil
}

func (s *Session) GetProfileDigged(username string, page int) ([]Link, error) {
	url := baseUri + "/Profile/Digged/" + username + "/appkey," + s.appkey + ",page," + strconv.Itoa(page)
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return []Link{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Link{}, err
	}
	var ls []Link
	json.Unmarshal(body, &ls)
	return ls, nil
}

func (s *Session) GetProfileBuried(username string, page int) ([]Link, error) {
	url := baseUri + "/Profile/Buried/" + username + "/appkey," + s.appkey + ",page," + strconv.Itoa(page)
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return []Link{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Link{}, err
	}
	var ls []Link
	json.Unmarshal(body, &ls)
	return ls, nil
}

func (s *Session) GetProfileEntries(username string, page int) ([]Entry, error) {
	url := baseUri + "/Profile/Entries/" + username + "/appkey," + s.appkey + ",page," + strconv.Itoa(page)
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return []Entry{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Entry{}, err
	}
	var es []Entry
	json.Unmarshal(body, &es)
	return es, nil
}

func (s *Session) GetProfileEntriesComments(username string, page int) ([]EntryComment, error) {
	url := baseUri + "/Profile/EntriesComments/" + username + "/appkey," + s.appkey + ",page," + strconv.Itoa(page)
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return []EntryComment{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []EntryComment{}, err
	}
	var es []EntryComment
	json.Unmarshal(body, &es)
	return es, nil
}

func (s *Session) GetProfileRelated(username string, page int) ([]RelatedLink, error) {
	url := baseUri + "/Profile/Related/" + username + "/appkey," + s.appkey + ",page," + strconv.Itoa(page)
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return []RelatedLink{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []RelatedLink{}, err
	}
	var rls []RelatedLink
	json.Unmarshal(body, &rls)
	return rls, nil
}
