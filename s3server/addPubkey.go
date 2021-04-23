package s3server

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

func AddKey(username, pubKey string) (uint32, error) {
	var num uint32

	url := "https://39.108.113.193:8080/api/v1/addPubkey?publicKey=" + pubKey + "&userName=" + username

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}

	resp, err := c.Get(url)
	if err != nil {

	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &num)
	} else {

	}

	return num, err
}
