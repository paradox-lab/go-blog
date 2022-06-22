/*******************************************************************************
 * Copyright (c) 2021
 * @Description：Python由荷兰数学和计算机科学研究学会的Guido van Rossum 于1990 年代初设计，作为一门叫做ABC语言的替代品。
 * @Date: 2021/3/16 上午11:38
 *
 * @Python官网 - https://www.python.org/
 * @Python中文文档 - https://docs.python.org/zh-cn/3/
 * @github地址 - https://github.com/python/cpython
 * @awesome库 - https://github.com/vinta/awesome-python
 * @awesome库中文版 - https://github.com/jobbole/awesome-python-cn
 *
 * Python开发指导
 * @Docs - https://docs.python-guide.org/
 * @github地址 - https://github.com/realpython/python-guide
 * @Docs中文版 - https://pythonguidecn.readthedocs.io/zh/latest/
 * @github中文版 - https://github.com/Prodesire/Python-Guide-CN
 *
 * @latest version: v3.9.2
 *
 *
 *
 *
 ******************************************************************************/

package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc/pb"
	"log"
	"time"
)

func main() {

	conn, err := NewGrpcConn()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := NewGrpcClient(conn)
	UploadJdCfst(c, "C:\\Users\\39713\\Desktop\\新建文件夹\\jd_opf_22160995_20210311.csv")
}

func NewGrpcConn() (*grpc.ClientConn, error) {
	// Set up a connection to the server.
	return grpc.Dial("localhost:50062", grpc.WithInsecure(), grpc.WithBlock())
}

func NewGrpcClient(conn *grpc.ClientConn) pb.SeleniumClient {
	return pb.NewSeleniumClient(conn)
}

func UploadJdCfst(c pb.SeleniumClient, file string) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	r, err := c.UploadJdCfst(ctx, &pb.RunRequst{File: file})
	if err != nil {
		log.Fatalf("could not upload storage blob: %v", err)
	}

	if r.String() != "" {
		log.Printf("Aje: %s", r.String())
	}
}
