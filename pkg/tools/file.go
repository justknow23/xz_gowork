package tools

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"insurance/pkg/global"
)

// CreateCSVFile 创建CSV文件
func CreateCSVFile(failedUserIds []string, filename string, params ...interface{}) (path string, err error) {
	file := fmt.Sprintf("%s/%s.csv", global.TempPath, fmt.Sprintf(filename, params...))
	f, err := os.Create(file) //创建文件
	if err != nil {
		return "", err
	}
	if f == nil {
		return "", fmt.Errorf("createCSVFile error")
	}
	defer func() {
		err := f.Close()
		fmt.Printf("close file error:%+v\n", err)
	}()

	_, err = f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	if err != nil {
		return "", err
	}
	w := csv.NewWriter(f) //创建一个新的写入文件流
	var data [][]string
	for key, value := range failedUserIds {
		keyStr := strconv.Itoa(key + 1)
		data = append(data, []string{keyStr, value})
	}

	err = w.WriteAll(data) //写入数据
	w.Flush()

	return file, nil
}
