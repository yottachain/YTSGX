package s3server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AddKey(username, pubKey string) (uint32, error) {
	var num uint32

	url := "https://47.115.114.243:8080/api/v1/addPubKey?pubkey=" + pubKey + "&username=" + username

	resp, err := http.Get(url)
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
