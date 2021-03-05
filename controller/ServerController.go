package controller

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//GetInfo test demo
func GetInfo(g *gin.Context) {

	directory := "/root/mnt/tmp/"
	fileName := "test.txt"
	s, err := os.Stat(directory)
	if err != nil {
		if !os.IsExist(err) {
			err = os.MkdirAll(directory, os.ModePerm)
			if err != nil {
				logrus.Errorf("err:%s\n", err)
			}
		} else {
			logrus.Errorf("err:%s\n", err)
		}
	} else {
		if !s.IsDir() {
			logrus.Errorf("err:%s\n", "The specified path is not a directory.")
		}
	}
	if !strings.HasSuffix(directory, "/") {
		directory = directory + "/"
	}
	filePath := directory + fileName
	logrus.Infof("file directory:%s\n", directory)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		logrus.Errorf("err:%s\n", err)
	}
	defer f.Close()

	f.Write([]byte("黄河之水天上来，奔流到海不复回。\r\n"))
}
