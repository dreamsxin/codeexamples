package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"os"
	"strconv"
	"fmt"
)

func join(fileName1 string, fileName2 string, col int) {

	if col < 0 {
		log.Fatalf("col can not less 0")
	}
	fs1, err := excelize.OpenFile(fileName1)
	if err != nil {
		log.Fatalf("can not read %v, err is %+v", fileName1, err)
	}
	rows1, err := fs1.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("GetRows error %v, err is %+v", fileName1, err)
	}
	fs2, err := excelize.OpenFile(fileName2)
	if err != nil {
		log.Fatalf("can not read %v, err is %+v", fileName2, err)
	}
	rows2, err := fs2.GetRows("Sheet1")
	if err != nil {
		log.Fatalf("GetRows error %v, err is %+v", fileName2, err)
	}
	//f := excelize.NewFile()
	//excel错误样式，黄色背景
	style, _ := fs1.NewStyle(`{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],"fill":{"type":"pattern","color":["#ffeb00"],"pattern":1},"alignment":{"horizontal":"left","ident":1,"vertical":"center","wrap_text":true}}`)

	for rindex, row1 := range rows1 {
		if (col >= len(row1)) {
			continue
		}
		for _, row2 := range rows2 {
			if (col >= len(row2)) {
				continue
			}
			if row1[col] == row2[col] {
				fs1.SetCellStyle("Sheet1", fmt.Sprintf("A%v", rindex+1), fmt.Sprintf("AA%v", rindex+1), style)
				//f.SetSheetRow("Sheet1", fmt.Sprint("A", j), row1)
			}
		}
	}
	if err := fs1.Save(); err != nil {
		println(err.Error())
	}
}

func main() {
	args := os.Args

	if args == nil || len(args) < 3 {
		join("1.xlsx", "2.xlsx", 1)
	} else {
		fmt.Println("转换文件 ", args[1], args[2])
		if len(args) > 3 {
			col, _ := strconv.Atoi(args[3])
			join(args[1], args[2], col)
		} else {
			join(args[1], args[2], 0)
		}
	}
	fmt.Println("标记完成！")
}
