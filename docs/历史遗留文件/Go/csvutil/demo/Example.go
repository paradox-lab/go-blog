/*******************************************************************************
 * Copyright (c) 2021
 * @Description：Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.
 * @Date: 2021/3/12 上午10:32
 *
 * @github地址: https://github.com/golang/go
 * @标准库文档 : https://studygolang.com/pkgdoc
 * @Go指南: http://tour.studygolang.com/welcome/1
 * @awesome库: https://github.com/avelino/awesome-go
 * @awesome库(中文版): https://github.com/jobbole/awesome-go-cn
 *
 * @use version: v1.15.8
 ******************************************************************************

 - Example1: 官方提供的Demo
 - Example2: 不同分隔符的处理
 ******************************************************************************/

package main

import (
	"encoding/csv"
	"fmt"
	"github.com/jszwec/csvutil"
	"io"
	"log"
	"os"
	"time"
)

// 官方提供的Demo
func Example1() {
	//【参考】https://github.com/jszwec/csvutil#unmarshal-

    // csv的数据
	var csvInput = []byte(`
name,age,CreatedAt
jacek,26,2012-04-01T15:00:00Z
john,,0001-01-01T00:00:00Z`,
	)

    // 【Notification】csv字段映射
    // csv的每个字段都要列出来
	type User struct {
		// 结构体字段要求大写，但在csv里的字段是小写，因此需要填上`csv:"name"`做一个映射
		Name      string `csv:"name"`
		// 加上omitempty可处理空值情况
		Age       int    `csv:"age,omitempty"`
		CreatedAt time.Time
	}

	var users []User
	// 【Notification】反序列化
	if err := csvutil.Unmarshal(csvInput, &users); err != nil {
		fmt.Println("error:", err)
	}
    // 遍历和输出反序列化以后的数据
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}

	// Output:
	// {Name:jacek Age:26 CreatedAt:2012-04-01 15:00:00 +0000 UTC}
	// {Name:john Age:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC}
}

// Example2 不同分隔符的处理
func Example2()  {
	//【参考】https://github.com/jszwec/csvutil#different-separatordelimiter-
	//【Notification】 默认处理的分隔符是逗号，以下代码为遇到其他分隔符例如'\t'时的处理

	file, err := os.Open("filename")
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)

	csvReader.Comma = '\t'

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		Name      string `csv:"name"`
		Age       int    `csv:"age,omitempty"`
		CreatedAt time.Time
	}

	var users []User
	for {
		var u User
		if err := dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u.Name)
		fmt.Println(u.Age)
		fmt.Println(u.CreatedAt)

		users = append(users, u)
	}

	// 解码后的数据已全部写入到users切片，后续继续处理users即可
}
