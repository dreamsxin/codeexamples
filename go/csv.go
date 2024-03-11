package main

import (
	"encoding/csv"
	"log"
	"os"

	"fmt"
	"strings"
)

func join(fileName1 string, fileName2 string) (row []string) {
	fs1, _ := os.Open(fileName1)
	r1 := csv.NewReader(fs1)
	content1, err := r1.ReadAll()
	if err != nil {
		log.Fatalf("can not read 1.csv, err is %+v", err)
	}

	fs2, _ := os.Open(fileName2)
	r2 := csv.NewReader(fs2)
	content2, err := r2.ReadAll()
	if err != nil {
		log.Fatalf("can not read 2.csv, err is %+v", err)
	}
	row = make([]string, 10)
	j := 0
	for i, row1 := range content1[1:] {
		for _, row2 := range content2[1:] {
			if row1[0] == row2[0] {
				row[j] = strings.Join(content1[i], ",")
				j++
			}
		}
	}
	return
}

func save(filename string, row []string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//防止乱码
	f.WriteString("\xEF\xBB\xBF")
	//w := csv.NewWriter(f)
	for _, str := range row {
		f.WriteString(str + "\n")
	}
	//f.Flush()
}

func main() {
	row := join("1.csv", "2.csv")
	fmt.Println(row)
	save("join.csv", row)
	fmt.Println("合并已完成！")
}
