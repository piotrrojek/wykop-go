package wykop

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

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
