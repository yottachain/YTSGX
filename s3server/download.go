package s3server

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
)

func DownBlock(publicKey, bucketName, fileName string, blockNum int) ([]byte, error) {
	var data []byte

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie接口
	cookieJar, _ := cookiejar.New(nil)
	c := &http.Client{
		Jar:       cookieJar,
		Transport: tr,
	}
	str2 := fmt.Sprintf("%d", blockNum)
	resp, err := c.Get("https://47.115.114.243:8080/api/v1/getBlockForSGX?publicKey=" + publicKey + "&bucketName=" + bucketName + "&fileName=" + fileName + "&blockNum=" + str2)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			err = json.Unmarshal(body, &data)
			//data = body
		}

	}

	return data, nil
}
