package main

import (
	"fmt"

	excelize "github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("demo.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "F1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cell)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	// 写内容
	f.SetCellValue("Sheet1", "A3", "Hello world.")
	f.SetCellValue("Sheet1", "B3", 100)



}
