package main

import (
	"dream/serve/routes"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
)

var g errgroup.Group

func main() {
	// 初始化全局变量
	routes.InitGlobalVariable()
	fmt.Print("初始化全局变量成功")

	//前台接口服务
	g.Go(func() error {
		return routes.FrontendServer().ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
