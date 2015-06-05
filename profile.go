package wykop

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

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
