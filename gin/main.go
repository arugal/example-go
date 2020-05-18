// Copyright 2020 arugal, zhangwei24@apache.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// https://www.yoytang.com/go-gin-doc.html

func main() {
	router := gin.Default()

	router.GET("/", WebRoot)
	router.POST("/", WebRoot)

	// 动态参数
	router.GET("/user/:name/:action", UserAction)

	// 路由组
	v1 := router.Group("/v1")
	{
		v1.GET("/add", func(context *gin.Context) {
			context.String(http.StatusOK, "add")
		})
	}

	v2 := router.Group("v2")
	{
		v2.GET("/add", func(context *gin.Context) {
			context.String(http.StatusOK, "add")
		})
	}

	// 单个路由中间件
	router.GET("/middleware", middleware1, middleware2, func(context *gin.Context) {
		context.String(http.StatusOK, "middleware")
	})

	// 路由组使用中间件
	v3 := router.Group("/v3", middleware1)
	{
		v3.GET("/add", func(context *gin.Context) {
			context.String(http.StatusOK, "add")
		})
	}

	// url 查询参数
	router.GET("/welcome", func(context *gin.Context) {
		// 获取参数内容
		// 获取的所有参数内容的类型都是string
		// 如果不存在,使用第二个当作默认内容
		firstName := context.DefaultQuery("firstName", "Guest")
		lastName := context.Query("lastName")

		context.String(http.StatusOK, "Hello %s%s", firstName, lastName)
	})

	// 表单和Body参数
	router.POST("/welcome", func(context *gin.Context) {
		firstName := context.PostForm("firstName")
		lastName := context.DefaultPostForm("lastName", "Guest")

		context.JSON(http.StatusOK, gin.H{
			"firstName": firstName,
			"lastName":  lastName,
		})
	})

	_ = router.Run()
}

func WebRoot(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello, world")
}

func UserAction(ctx *gin.Context) {
	name := ctx.Param("name")
	action := ctx.Param("action")

	msg := name + " is " + action
	ctx.String(http.StatusOK, msg)
}

func middleware1(ctx *gin.Context) {
	log.Println("in middleware1")

	ctx.Next()
}

func middleware2(ctx *gin.Context) {
	log.Println("before in middleware2")
	ctx.Next()
	log.Println("after in middleware2")
}
