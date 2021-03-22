package controller

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	sgxaes "github.com/yottachain/YTSGX/aes"
	"github.com/yottachain/YTSGX/s3server"
	"github.com/yottachain/YTSGX/tools"
)

func GetUserInfo(g *gin.Context) {
	data := tools.ReadUserInfo()
	var uu tools.User
	uu = tools.UserUnmarshal(data)
	g.JSON(http.StatusOK, uu)
}

func UpdateUserInfo(g *gin.Context) {
	var uu tools.User
	userName := g.Query("userName")
	priKey := g.Query("privateKey")
	pubKey := g.Query("publicKey")
	uu.UserName = userName
	uu.PublicKey = pubKey
	uu.PrivateKey = priKey
	data, err := json.Marshal(uu)
	if err != nil {
		logrus.Errorf("Marshal err:%s\n", err)
	}

	tools.UserWrite(data)

	data1 := tools.ReadUserInfo()
	var u1 tools.User
	u1 = tools.UserUnmarshal(data1)

	g.JSON(http.StatusOK, u1)

}

func GetPubKey(g *gin.Context) {
	userName := g.Query("userName")
	priKey, pubKey := tools.CreateKey()
	//priKey := g.Query("privateKey")
	//pubKey := g.Query("publicKey")
	user := tools.User{
		UserName:   userName,
		PrivateKey: priKey,
		PublicKey:  "YTA" + pubKey,
	}

	data, err := json.Marshal(user)
	if err != nil {
		logrus.Errorf("Marshal err:%s\n", err)
	}

	tools.UserWrite(data)

	g.JSON(http.StatusOK, gin.H{"publicKey": user.PublicKey, "privateKey": user.PrivateKey})
}

func AddUser(g *gin.Context) {
	var num uint32
	data := tools.ReadUserInfo()
	var uu tools.User
	uu = tools.UserUnmarshal(data)
	num, err := s3server.AddKey(uu.UserName, uu.PublicKey)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"errMsg": err})
	} else {
		user := tools.User{
			UserName:   uu.UserName,
			Num:        num,
			PrivateKey: uu.PrivateKey,
			PublicKey:  uu.PublicKey,
		}

		data, err := json.Marshal(user)
		if err != nil {
			logrus.Errorf("Marshal err:%s\n", err)
		}

		tools.UserWrite(data)

		g.JSON(http.StatusOK, gin.H{"publicKey:": uu.PublicKey})
	}

	//user := tools.User{
	//	UserName:   userName,
	//	Num:        num,
	//	PrivateKey: priKey,
	//	PublicKey:  pubKey,
	//}
	//data, err := json.Marshal(user)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//tools.UserWrite(data)

	////读 user JSON文件
	//data1 := tools.ReadUserInfo()
	//var uu tools.User
	//uu = tools.UserUnmarshal(data1)
	//
	//g.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "priKey:": uu.PrivateKey + ",pubKey:" + uu.PublicKey})

}

func SaveFileTo(g *gin.Context) {
	fileName := g.Query("fileName")
	var data []byte

	resp, err := http.Get("http://localhost:18080/api/v1/getContents?fileName=" + fileName)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	// 一次性读取
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("错误")
		return
	}

	err = json.Unmarshal(bs, &data)
	tools.WriteBytes(data, fileName)
}

func GetFileContents(g *gin.Context) {

	fileName := g.Query("fileName")

	directory := "./storage/" + fileName

	// dd := "./storage/test/" + fileName

	bytes, err := ioutil.ReadFile(directory)
	if err != nil {
		g.JSON(http.StatusBadGateway, err)
	} else {

		encodeString := base64.StdEncoding.EncodeToString(bytes)

		decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
		if err != nil {
			logrus.Errorf("err:%s\n", err)
		}
		tools.WriteBytes(decodeBytes, fileName)
		// g.JSON(http.StatusOK, encodeString)
		g.String(http.StatusOK, encodeString)
	}

	// fp, err := os.OpenFile(directory, os.O_RDONLY, 0755)
	// // defer fp.Close()
	// if err != nil {
	// 	logrus.Errorf("err:%s\n", err)
	// }
	// data := make([]byte, 1024*1024*8)
	// n, err := fp.Read(data)
	// if err != nil {
	// 	logrus.Errorf("err:%s\n", err)
	// } else {

	// 	ss := len(data[:n])
	// 	fmt.Printf("nnnnnn:%d\n", ss)
	// 	tools.WriteBytes(data[:n], fileName)
	// 	g.JSON(http.StatusOK, data[:n])
	// }

}

func write(directory, fileName string) {
	s, err := os.Stat(directory)
	if err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(directory, os.ModePerm)
			if err != nil {
				//log.Fatal(err)
			}
		} else {
			//log.Fatal(err)
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
		logrus.Errorf("Erre:%s\n", err)
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

func DownloadFileForSGX(g *gin.Context) {
	bucketName := g.Query("bucketName")
	fileName := g.Query("objectKey")
	userdata := tools.ReadUserInfo()
	var user tools.User
	user = tools.UserUnmarshal(userdata)
	//publicKey := user.PublicKey
	userName := user.UserName
	var blockNum int
	blockNum = 0

	key, err := sgxaes.NewKey(user.PrivateKey, user.Num)
	if err != nil {

	}
	index := strings.Index(fileName, "/")
	var filePath string
	if index != -1 {
		directory := "./storage/" + bucketName + "/" + fileName[:index]
		fname := fileName[index+1:]
		filePath = createDirectory(directory, fname)
	} else {
		directory := "./storage/" + bucketName

		filePath = createDirectory(directory, fileName)
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("erra:%s\n", err)
	}
	defer f.Close()
	//write := bufio.NewWriter(f)

	for blockNum != -1 {
		data, err := s3server.DownBlock(userName, bucketName, fileName, blockNum)
		if err != nil {
			blockNum = -1
			g.JSON(http.StatusAccepted, gin.H{"Msg": "[" + fileName + "] download is failure ."})
			return
		} else {
			if len(data) > 0 {

				block := sgxaes.NewEncryptedBlock(data)
				err1 := block.Decode(key, f)

				if err1 != nil {
					fmt.Println(err1)
					g.JSON(http.StatusAccepted, gin.H{"Msg": "[" + fileName + "] download is failure ."})
					return
				} else {
					blockNum++
				}

			} else {
				blockNum = -1
			}
		}
	}
	md5Value := Md5SumFile(filePath)

	g.JSON(http.StatusOK, gin.H{"File md5:": md5Value, "Msg": "[" + fileName + "] download is successful."})

}

func createDirectory(directory, fileName string) string {
	s, err := os.Stat(directory)
	if err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(directory, os.ModePerm)
			if err != nil {
				logrus.Errorf("err1:%s\n", err)
			}
		} else {
			logrus.Errorf("err2:%s\n", err)
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

	return filePath
}

func Md5SumFile(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err.Error()
	}
	md5 := md5.New()
	md5.Write(data)
	md5Data := md5.Sum([]byte(nil))
	return hex.EncodeToString(md5Data)
}
