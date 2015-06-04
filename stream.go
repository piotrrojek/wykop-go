package wykop

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Entry struct {
	App               string         `json:"app,omitempty"`
	ID                int            `json:"id,omitempty"`
	Author            string         `json:"author,omitempty"`
	AuthorAvatar      string         `json:"author_avatar,omitempty"`
	AuthorAvatarBig   string         `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed   string         `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo    string         `json:"author_avatar_lo,omitempty"`
	AuthorGroup       int            `json:"author_group,omitempty"`
	Date              string         `json:"date,omitempty"`
	Body              string         `json:"body,omitempty"`
	URL               string         `json:"url,omitempty"`
	Receiver          string         `json:"receiver,omitempty"`
	ReceiverSex       string         `json:"receiver_sex,omitempty"`
	ReceiverAvatar    string         `json:"receiver_avatar,omitempty"`
	ReceiverAvatarMed string         `json:"receiver_avatar_med,omitempty"`
	ReceiverAvatarLo  string         `json:"receiver_avatar_lo,omitempty"`
	ReceiverAvatarBig string         `json:"receiver_avatar_big,omitempty"`
	ReceiverGroup     int            `json:"receiver_group,omitempty"`
	Comments          []EntryComment `json:"comments,omitempty"`
	CommentCount      int            `json:"comment_count"`
	VoteCount         int            `json:"vote_count,omitempty"`
	UserVote          int            `json:"user_vote,omitempty"`
	Voters            []Dig          `json:"voters,omitempty"`
	UserFavorite      bool           `json:"user_favorite,omitempty"`
	Embed             Embed          `json:"embed,omitempty"`
	Blocked           bool           `json:"blocked,omitempty"`
	Type              string         `json:"type,omitempty"`
	Deleted           bool           `json:"deleted,omitempty"`
	AuthorSex         string         `json:"author_sex"`
	ViolationURL      string         `json:"violation_url,omitempty"`
	Source            string         `json:"source"`
	CanComment        bool           `json:"can_comment"`
}

type EntryComment struct {
	ID              int    `json:"id,omitempty"`
	Author          string `json:"author,omitempty"`
	AuthorSex       string `json:"author_sex"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	Date            string `json:"date,omitempty"`
	Body            string `json:"body,omitempty"`
	VoteCount       int    `json:"vote_count,omitempty"`
	UserVote        int    `json:"user_vote,omitempty"`
	Voters          []Dig  `json:"voters,omitempty"`
	Embed           Embed  `json:"embed,omitempty"`
	Entry           Entry  `json:"entry,omitempty"`
	ViolationUrl    string `json:"violation_url,omitempty"`
	App             string `json:"app,omitempty"`
	Deleted         bool   `json:"deleted,omitempty"`
	Source          string `json:"source,omitempty"`
	Type            string `json:"entry_comment,omitempty"`
	Blocked         bool   `json:"blocked,omitempty"`
}

type Embed struct {
	Type    string `json:"type,omitempty"`
	Preview string `json:"preview,omitempty"`
	URL     string `json:"url,omitempty"`
	Source  string `json:"source,omitempty"`
	Plus18  bool   `json:"plus18,omitempty"`
}

type Dig struct {
	Author          string `json:"author,omitempty"`
	AuthorAvatar    string `json:"author_avatar,omitempty"`
	AuthorAvatarBig string `json:"author_avatar_big,omitempty"`
	AuthorAvatarMed string `json:"author_avatar_med,omitempty"`
	AuthorAvatarLo  string `json:"author_avatar_lo,omitempty"`
	AuthorGroup     int    `json:"author_group,omitempty"`
	AuthorSex       string `json:"author_sex,omitempty"`
	Date            string `json:"date,omitempty"`
}

func (s *Session) GetStreamIndex() ([]Entry, error) {
	url := baseUri + "/Stream/Index/appkey," + s.appkey
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

func (s *Session) GetStreamHot(page, period int) ([]Entry, error) {
	url := baseUri + "/Stream/Hot/appkey," + s.appkey + ",page," + strconv.Itoa(page) + ",period," + strconv.Itoa(period)
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

func (s *Session) GetStreamEntry(index int) (Entry, error) {
	url := baseUri + "/Entries/Index/" + strconv.Itoa(index) + "/appkey," + s.appkey
	resp, err := getRequest(url, s.appsecret, s.ClearOutput)
	if err != nil {
		return Entry{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Entry{}, err
	}
	var e Entry
	json.Unmarshal(body, &e)
	return e, nil
}
