package tools

import (
	"log"
	"strconv"
	"time"
)

//中国居民身份证 工具类   18位数 兼容15位
//通过身份证号，获取出生年份，月份，日，和性别，生日，年龄

type IDCardInfo struct {
	IDCardNo string
	Year     string
	Month    string
	Day      string
	BirthDay string
	Sex      uint
	Age      uint
}

var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var valid_value = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

// Citizen15To18 Convert citizen 15 to 18.
func Citizen15To18(citizenNo15 []byte) []byte {
	nLen := len(citizenNo15)
	if nLen != 15 {
		return nil
	}

	citizenNo18 := make([]byte, 0)
	citizenNo18 = append(citizenNo18, citizenNo15[:6]...)
	citizenNo18 = append(citizenNo18, '1', '9')
	citizenNo18 = append(citizenNo18, citizenNo15[6:]...)

	sum := 0
	for i, v := range citizenNo18 {
		n, _ := strconv.Atoi(string(v))
		sum += n * weight[i]
	}
	mod := sum % 11
	citizenNo18 = append(citizenNo18, valid_value[mod])

	return citizenNo18
}

// NewIDCard 实例化居民身份证结构体
func NewIDCard(IDCardNo string) *IDCardInfo {
	if IDCardNo == "" {
		return nil
	}
	if len(IDCardNo) != 18 {
		IDCardNoByte := Citizen15To18([]byte(IDCardNo))
		IDCardNo = string(IDCardNoByte)
	}
	ins := IDCardInfo{
		IDCardNo: IDCardNo,
	}

	ins.Year = ins.GetYear()
	ins.Month = ins.GetMonth()
	ins.Day = ins.GetDay()
	ins.Sex = ins.GetSex()
	ins.BirthDay = ins.GetBirthDayStr()
	ins.Age = ins.GetAge()

	return &ins
}

// GetBirthDay 根据身份证号获取生日（时间格式）
func (s *IDCardInfo) GetBirthDay() *time.Time {
	if s == nil {
		return nil
	}

	dayStr := s.IDCardNo[6:14] + "000001"
	birthDay, err := time.Parse("20060102150405", dayStr)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &birthDay
}

// GetBirthDayStr 根据身份证号获取生日（字符串格式 yyyy-MM-dd HH:mm:ss）
func (s *IDCardInfo) GetBirthDayStr() string {
	defaultDate := "1999-01-01 00:00:01"
	if s == nil {
		return defaultDate
	}

	birthDay := s.GetBirthDay()
	if birthDay == nil {
		return defaultDate
	}

	return birthDay.Format("2006-01-02 15:04:05")
}

// GetYear 根据身份证号获取生日的年份
func (s *IDCardInfo) GetYear() string {
	if s == nil {
		return ""
	}

	return s.IDCardNo[6:10]
}

// GetMonth 根据身份证号获取生日的月份
func (s *IDCardInfo) GetMonth() string {
	if s == nil {
		return ""
	}

	return s.IDCardNo[10:12]
}

// GetDay 根据身份证号获取生日的日份
func (s *IDCardInfo) GetDay() string {
	if s == nil {
		return ""
	}

	return s.IDCardNo[12:14]
}

// GetSex 根据身份证号获取性别
func (s *IDCardInfo) GetSex() uint {
	var unknown uint = 3
	if s == nil {
		return unknown
	}

	sexStr := s.IDCardNo[16:17]
	if sexStr == "" {
		return unknown
	}

	i, err := strconv.Atoi(sexStr)
	if err != nil {
		return unknown
	}

	if i%2 != 0 {
		return 1
	}

	return 0
}

// GetAge 根据身份证号获取年龄
func (s *IDCardInfo) GetAge() uint {
	if s == nil {
		return 19
	}

	birthDay := s.GetBirthDay()
	if birthDay == nil {
		return 19
	}

	now := time.Now()

	age := now.Year() - birthDay.Year()
	if now.Month() > birthDay.Month() {
		age = age - 1
	}

	if age <= 0 {
		return 19
	}

	if age <= 0 || age >= 150 {
		return 19
	}

	return uint(age)
}
