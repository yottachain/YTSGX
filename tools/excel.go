package tools

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type ExcelUser struct {
	Sex       string
	Age       string
	Allergen  string
	Heart     string
	BadHabits string
	BloddFat  string
}

func ReadExcel(filePath string) ([]ExcelUser, error) {
	//xlsx, err := excelize.OpenFile("./storage/天津数据.xlsx")
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		logrus.Errorf("err:%s\n", err)
		return nil, err
		//os.Exit(1)
	} else {
		// Get sheet index.
		index := xlsx.GetSheetIndex("Sheet1")
		// Get all the rows in a sheet.
		rows := xlsx.GetRows("Sheet" + strconv.Itoa(index))

		var users []ExcelUser
		for n, row := range rows {
			if n >= 2 {
				excelUser := ExcelUser{}
				excelUser.Sex = row[0]
				excelUser.Age = row[1]
				excelUser.Allergen = row[2]
				excelUser.Heart = row[3]
				excelUser.BadHabits = row[4]
				excelUser.BloddFat = row[5]
				users = append(users, excelUser)

			}
		}
		return users, nil
	}
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}

		if fi.IsDir() { // 忽略目录
			return nil
		}

		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}

		return nil
	})

	return files, err
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
//func ListDir(dirPth string, suffix string) (files []string, err error) {
//	files = make([]string, 0, 10)
//
//	dir, err := ioutil.ReadDir(dirPth)
//	if err != nil {
//		return nil, err
//	}
//
//	PthSep := string(os.PathSeparator)
//	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
//
//	for _, fi := range dir {
//		if fi.IsDir() { // 忽略目录
//			continue
//		}
//		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
//			files = append(files, dirPth+PthSep+fi.Name())
//		}
//	}
//
//	return files, nil
//}
