package wykop

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Link struct {
	App             string `json:"app,omitempty"`
	ID              int    `json:"id,omitempty"`
	Title           string `json:"title,omitempty"`
	Description     string `json:"description,omitempty"`
	Tags            string `json:"tags,omitempty"`
	URL             string `json:"url,omitempty"`
	SourceURL       string `json:"source_url,omitempty"`
	VoteCount       int    `json:"vote_count,omitempty"`
	CommentCount    int    `json:"comment_count,omitempty"`
	ReportCount     int    `json:"report_count,omitempty"`
	Date            string `json:"date,omitempty"`
	Author          string `json:"author,omitempty"`
	AuthorSex       string `json:"author_sex,omitempty"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	Preview         string `json:"preview,omitempty"`
	UserLists       []int  `json:"user_lists,omitempty"`
	Plus18          bool   `json:"plus18,omitempty"`
	Status          string `json:"status,omitempty"`
	CanVote         bool   `json:"can_vote,omitempty"`
	HasOwnContent   bool   `json:"has_own_content,omitempty"`
	CategoryName    string `json:"category_name,omitempty"`
	UserVote        bool   `json:"user_vote,omitempty"`
	UserObserve     bool   `json:"user_observe,omitempty"`
	UserFavorite    bool   `json:"user_favorite,omitempty"`
	ViolationURL    string `json:"violation_url,omitempty"`
	Info            string `json:"info,omitempty"`
	RelatedCount    int    `json:"related_count,omitempty"`
	IsHot           bool   `json:"is_hot,omitempty"`
	Type            string `json:"type,omitempty"`
}

func (s *Session) GetLinksPromoted(page int) ([]Link, error) {
	url := baseUri + "/Links/Promoted/appkey," + s.appkey + ",page," + strconv.Itoa(page)
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

func (s *Session) GetLinksUpcoming(page int, sort string) ([]Link, error) {
	url := baseUri + "/Links/Upcoming/appkey," + s.appkey + ",page," + strconv.Itoa(page) + ",sort," + sort
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

func (s *Session) GetPopularPromoted() ([]Link, error) {
	url := baseUri + "/Popular/Promoted/appkey," + s.appkey
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

func (s *Session) GetPopularUpcoming() ([]Link, error) {
	url := baseUri + "/Popular/Upcoming/appkey," + s.appkey
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
