package tools

import (
	"encoding/csv"
	"fmt"
	xzcsv "gitlab.idc.xiaozhu.com/xz-go/common/csv"
	"gitlab.idc.xiaozhu.com/xz-go/common/log"
	"os"
	"strconv"
)

// CreateCSVFile 创建CSV文件无Title
func CreateCSVFile(data []string, filename, filepath string, params ...interface{}) (path string, err error) {
	file := fmt.Sprintf("%s/%s", filepath, fmt.Sprintf(filename, params...))
	f, err := os.Create(file) //创建文件
	if err != nil {
		return "", err
	}
	if f == nil {
		return "", fmt.Errorf("createCSVFile error")
	}
	defer func() {
		err = f.Close()
		fmt.Printf("close file error:%+v\n", err)
	}()

	_, err = f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	if err != nil {
		return "", err
	}
	w := csv.NewWriter(f) //创建一个新的写入文件流
	var datas [][]string
	for key, value := range data {
		keyStr := strconv.Itoa(key + 1)
		datas = append(datas, []string{keyStr, value})
	}

	err = w.WriteAll(datas) //写入数据
	w.Flush()

	return file, nil
}

// CreateCsv
//自动设置UTF-8输出格式
//支持appending模式追加输入
func CreateCsv(data [][]string, filename, filepath string, params ...interface{}) (string, error) {
	file := fmt.Sprintf("%s/%s", filepath, fmt.Sprintf(filename, params...))
	f, err := xzcsv.CreateCSV(file)
	if err != nil {
		return file, err
	}
	defer func(f *xzcsv.File) {
		err = f.Close()
		if err != nil {
			log.Infof("CreateCsv Err:%+v", err)
		}
	}(f)
	err = f.WriteAll(data)
	if err != nil {
		return file, err
	}

	return file, nil

}

// ReadCsv 读取Csv内容
func ReadCsv(filewithpath string) ([][]string, error) {
	var rows [][]string
	f, err := xzcsv.OpenCSV(filewithpath)
	if err != nil {
		log.Errorf("OpenCSV fails with err=%v", err)
		return rows, err
	}
	defer func(f *xzcsv.File) {
		err = f.Close()
		if err != nil {
			log.Infof("ReadCsv Err:%+v", err)
		}
	}(f)
	all, err := f.ReadAll()
	if err != nil {
		log.Errorf("ReadAll fails with err=%v", err)
		return rows, err
	}
	rows = all

	return rows, nil
}
