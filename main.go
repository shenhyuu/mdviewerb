package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// 配置 CORS 中间件
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // 允许所有来源，生产环境建议设置具体的域名
		AllowMethods: []string{http.MethodGet, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// 注册路由处理函数
	e.GET("/documents/:uuid", handleDocument)
	e.GET("/contents", handleContents)

	// 启动服务器
	e.Logger.Fatal(e.Start(":8080"))
}

func handleContents(c echo.Context) error {
	// 读取 contents.json 文件
	content, err := os.ReadFile("contents.json")
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("读取 contents.json 失败: %v", err))
	}

	// 设置响应头
	c.Response().Header().Set(echo.HeaderContentType, "application/json")
	return c.Blob(http.StatusOK, "application/json", content)
}

func handleDocument(c echo.Context) error {
	uuid := c.Param("uuid")
	found := false

	// 遍历 documents 目录
	err := filepath.Walk("documents", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查是否是 markdown 文件
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(path), ".md") {
			// 获取文件名（不含扩展名）
			fileName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

			// 如果文件名匹配 UUID
			if fileName == uuid {
				found = true
				// 读取文件内容
				content, err := os.ReadFile(path)
				if err != nil {
					return c.String(http.StatusInternalServerError, fmt.Sprintf("读取文件失败: %v", err))
				}

				// 设置响应头
				c.Response().Header().Set(echo.HeaderContentType, "text/markdown")
				return c.Blob(http.StatusOK, "text/markdown", content)
			}
		}
		return nil
	})

	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("遍历目录失败: %v", err))
	}

	// 如果没有找到匹配的文件
	if !found {
		return c.String(http.StatusNotFound, "未找到指定的文档")
	}

	return nil
}
