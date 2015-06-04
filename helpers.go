package wykop

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
)

const (
	baseUri = "http://a.wykop.pl"
)

func apiSign(appsecret, url string) string {
	source := append([]byte(appsecret), url...)
	result := md5.New()

	result.Write(source)
	return hex.EncodeToString(result.Sum(nil))
}

func getRequest(url, appsecret string, clearoutput bool) (resp *http.Response, err error) {

	if clearoutput {
		url += ",output,clear"
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	header := apiSign(appsecret, url)
	req.Header.Add("apisign", header)
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
