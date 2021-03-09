package controller

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/yottachain/YTSGX/s3server"
	"github.com/yottachain/YTSGX/tools"
)

//GetInfo test demo
func AddUser(g *gin.Context) {
	var num uint32

	// priKey, pubKey := tools.CreateKey()
	priKey := "5KcTgaryRNpRhQ33NYMkHDVzGSyXF35hYXKR34aZkNhzU2S1iBK"
	pubKey := "8fYbPL1vD3RZjFDhax7bruKymvTvrXUzY7FSCa4Qjfs9NQy2aR"
	userName := "yottanewsabc"
	num, err := s3server.AddKey(userName, pubKey)

	if err != nil {

	} else {
		user := tools.User{
			UserName:   userName,
			Num:        num,
			PrivateKey: priKey,
			PublicKey:  pubKey,
		}

		data, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}

		tools.UserWrite(data)
	}

	data := tools.ReadUserInfo()
	var uu tools.User
	uu = tools.UserUnmarshal(data)

	g.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "priKey:": uu.PrivateKey + ",pubKey:" + uu.PublicKey})

}

func write(directory, fileName string) {
	s, err := os.Stat(directory)
	if err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(directory, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	} else {
		if !s.IsDir() {
			// logrus.Errorf("err:%s\n", "The specified path is not a directory.")
		}
	}
	if !strings.HasSuffix(directory, "/") {
		directory = directory + "/"
	}
	filePath := directory + fileName
	// logrus.Infof("file directory:%s\n", directory)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	x := []byte("黄河之水天上来，哈哈哈哈哈哈哈哈哈哈哈！！！！！！！！！")
	key := []byte("hgfedcba87654321")
	x1 := encryptAES(x, key)
	f.Write(x1)
}

func encryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func decryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func read(directory, fileName string) {
	// directory = "/root/mnt/tmp/"
	// fileName = "test.txt"
	if !strings.HasSuffix(directory, "/") {
		directory = directory + "/"
	}
	path := directory + fileName
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("bytes read: ", bytesread)

	key := []byte("hgfedcba87654321")
	x2 := decryptAES(buffer, key)
	fmt.Println("bytestream to string: ", string(x2))
}
