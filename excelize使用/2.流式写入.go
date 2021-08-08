package main

import (
	"fmt"

	excelize "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {

	//stringDic := map[string][5]string{}
	//intDic := map[string][5]int{}
	//stringDic["name"] = [5]string{"luenci", "lynn", "xiaoming", "xiaohong", "xiaobai"}
	//intDic["age"] = [5]int{11, 12, 13, 14, 15}

	stringDic := [6]string{"name", "luenci", "lynn", "xiaoming", "xiaohong", "xiaobai"}

	//jsonData, err := json.Marshal(stringDic)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//fmt.Println(stringDic)
	//fmt.Println(string(jsonData))
	//fmt.Println(intDic)
	//fmt.Println(stringDic["name"][0])

	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}

	for rowID := 1; rowID <= 3; rowID++ {
		row := make([]interface{}, 1)
		for colID := 1; colID <= 6; colID++ {
			row[0] = stringDic[colID-1]

			// 判断 (1, rowID) 返回单元格的位置 为下一步向此位置写数据
			cell, _ := excelize.CoordinatesToCellName(rowID, colID)
			//
			fmt.Println(cell)
			fmt.Println(row)
			if err := streamWriter.SetRow(cell, row); err != nil {
				fmt.Println(err)
			}
		}

	}
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}
	if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)

	}

	//names := []string{"luenci", "lynn", "xiaoming", "xiaohong", "xiaobai"}
	//names = append(names, "kk")

	//
	//for i := 0; i < len(names); i++ {
	//	row := make([]interface{}, 3)
	//
	//}

	//file := excelize.NewFile()
	//streamWriter, err := file.NewStreamWriter("Sheet1")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for rowID := 2; rowID <= 1024000; rowID++ {
	//	row := make([]interface{}, 50)
	//	for colID := 0; colID < 50; colID++ {
	//		row[colID] = rand.Intn(640000)
	//	}

	// 判断 (1, rowID) 返回单元格的位置 为下一步向此位置写数据
	//	cell, _ := excelize.CoordinatesToCellName(1, rowID)

	//	if err := streamWriter.SetRow(cell, row); err != nil {
	//		fmt.Println(err)
	//	}
	//}
	//if err := streamWriter.Flush(); err != nil {
	//	fmt.Println(err)
	//}
	//if err := file.SaveAs("Book1.xlsx"); err != nil {
	//	fmt.Println(err)
	//}
}
