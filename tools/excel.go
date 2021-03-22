package tools

import (
	"fmt"
	"os"
	"strconv"

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

func ReadExcel(filePath string) []ExcelUser {
	//xlsx, err := excelize.OpenFile("./storage/天津数据.xlsx")
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
	return users
}

// 读取excel
//func readExcel() {
//	f, err := excelize.OpenFile("./storage/天津数据.xlsx")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// 从指定的 cell（单元格） 获取数据
//	//cell := f.GetCellValue("Sheet1", "B2")
//	//fmt.Println(cell)
//
//	// 获取 Sheet1 中所有数据
//	// rows: 所有行数
//	rows := f.GetRows("Sheet1")
//	for _, row := range rows {
//		// 一行一行打印出来
//		str := ""
//		for _, cellVal := range row {
//			str += cellVal + "--"
//		}
//		fmt.Println(str)
//	}
//}
