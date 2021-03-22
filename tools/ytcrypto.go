package tools

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
)

var (
	privKey string
	pubKey  string
)

// func CreateKey() (string, string) {

// 	privKey, pubKey = CreateKey()
// 	if privKey == "" || pubKey == "" {
// 		panic("Create key pair failed")
// 	}
// 	fmt.Printf("Createkey pair success,private key is %s and public key is %s\n", privKey, pubKey)
// 	return privKey, pubKey
// }

func CreateKey() (string, string) {
	privKey, _ := ecrypto.GenerateKey()
	privKeyBytes := ecrypto.FromECDSA(privKey)
	rawPrivKeyBytes := append([]byte{}, 0x80)
	rawPrivKeyBytes = append(rawPrivKeyBytes, privKeyBytes...)
	checksum := sha256Sum(rawPrivKeyBytes)
	checksum = sha256Sum(checksum)
	rawPrivKeyBytes = append(rawPrivKeyBytes, checksum[0:4]...)
	privateKey := base58.Encode(rawPrivKeyBytes)

	pubKey := privKey.PublicKey
	pubKeyBytes := ecrypto.CompressPubkey(&pubKey)
	checksum = ripemd160Sum(pubKeyBytes)
	rawPublicKeyBytes := append(pubKeyBytes, checksum[0:4]...)
	publicKey := base58.Encode(rawPublicKeyBytes)
	return privateKey, publicKey
}

func sha256Sum(bytes []byte) []byte {
	h := sha256.New()
	h.Write(bytes)
	return h.Sum(nil)
}
func ripemd160Sum(bytes []byte) []byte {
	h := ripemd160.New()
	h.Write(bytes)
	return h.Sum(nil)
}

func UserWrite(data []byte) {
	fp, err := os.OpenFile("user.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		logrus.Errorf("err:%s\n", err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		logrus.Errorf("err:%s\n", err)
	}
}

//func WriteBytes(data []byte, file string) {
//	directory := "./storage/test"
//	s, err := os.Stat(directory)
//	if err != nil {
//		if !os.IsExist(err) {
//			err = os.MkdirAll(directory, os.ModePerm)
//			if err != nil {
//				//log.Fatal(err)
//			}
//		} else {
//			//log.Fatal(err)
//		}
//	} else {
//		if !s.IsDir() {
//			// logrus.Errorf("err:%s\n", "The specified path is not a directory.")
//		}
//	}
//	if !strings.HasSuffix(directory, "/") {
//		directory = directory + "/"
//	}
//	filePath := directory + file
//	// logrus.Infof("file directory:%s\n", directory)
//	fp, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
//	if err != nil {
//		logrus.Errorf("Erre:%s\n", err)
//	}
//	defer fp.Close()
//
//	_, err = fp.Write(data)
//	if err != nil {
//		logrus.Errorf("err:%s\n", err)
//	}
//}

func ReadUserInfo() []byte {
	fp, err := os.OpenFile("./user.json", os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		logrus.Errorf("err:%s\n", err)
	}
	data := make([]byte, 1024)
	n, err := fp.Read(data)
	if err != nil {
		logrus.Errorf("err:%s\n", err)
	}
	//fmt.Println(string(data[:n]))
	return data[:n]
}

func UserUnmarshal(data []byte) User {
	var user User
	if len(data) == 0 {
		fmt.Println("User JSON is null....")
	}
	err := json.Unmarshal(data, &user)
	if err != nil {
		logrus.Errorf("err:%s\n", err)
	}
	//fmt.Println("UserName:" + user.UserName)
	//fmt.Println("PrivateKey:" + user.PrivateKey)
	//fmt.Println("PublicKey:" + user.PublicKey)
	return user
}

type User struct {
	UserName   string
	Num        uint32
	PrivateKey string
	PublicKey  string
}
