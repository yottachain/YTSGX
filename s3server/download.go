package s3server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func DownBlock(publicKey, bucketName, fileName string, blockNum int) ([]byte, error) {
	var data []byte

	resp, err := http.Get("https://localhost:8080/api/v1/getBlockForSGX?publicKey=" + publicKey + "&bucketName=" + bucketName + "&fileName=" + fileName + "&blockNum=" + strconv.Itoa(blockNum))
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			data = body
		}

	}

	return data, nil
}
