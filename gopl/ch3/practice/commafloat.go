package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(commaFloat("11323123.323213210"))
}

func commaFloat(num string) string {
	//因为float转换为字符串要指定小数点后位数,函数改为接收字符串参数
	// stringFloat := strconv.FormatFloat(num, 'f', 10, 64)

	if num == "" {
		return ""
	}

	numBytes := []byte(num)
	sign := numBytes[0]
	var buf bytes.Buffer
	//记录符号
	if sign == '+' || sign == '-' {
		buf.WriteByte(sign)
		numBytes = numBytes[1:]
	}

	// 根据小数点,进行数据拆分
	after := make([]byte, 0) //小数点后数据
	for i := 0; i < len(numBytes); i++ {
		if numBytes[i] == '.' {
			after = numBytes[i:] //小数点后数字不做处理
			numBytes = numBytes[:i]
		}
	}
	// 处理小数点前数字
	for start := 0; start < len(numBytes); start++ {
		buf.WriteByte(numBytes[start])
		if (len(numBytes)-start-1)%3 == 0 && (start < len(numBytes)-1) {
			buf.WriteByte(',')
		}
	}
	buf.WriteString(string(after))
	return buf.String()
}
