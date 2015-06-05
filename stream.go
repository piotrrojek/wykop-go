package wykop

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

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
