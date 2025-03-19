# 📝 一个现代化的 Markdown 笔记仓库后端，基于 Golang + ECHO 构建，提供优雅的在线浏览体验。

## 🌟 项目概述
这是一个基于Go语言开发的轻量级Web服务，专门用于提供Markdown文档的存储和获取功能。该服务提供RESTful API，允许客户端应用程序访问Markdown格式的笔记和文档。

## ✨ 功能特点
- 📄 提供Markdown格式文档的存储和访问
- 🔑 基于UUID的文档访问机制
- 🌐 RESTful API接口设计
- 🔄 跨域资源共享(CORS)支持
- 🚀 轻量级设计，易于部署和使用

## 🛠️ 技术栈
- 💻 编程语言：Go (1.21+)
- 🔧 Web框架：Echo v4.11.4
- 📋 文档格式：Markdown

## 📚 API接口
服务提供以下API接口：

### 1. 获取内容目录
```
GET /contents
```
返回一个JSON对象，包含所有可用文档的信息和对应的UUID。

**响应格式**：`application/json`

### 2. 获取特定文档
```
GET /documents/:uuid
```
根据提供的UUID返回对应的Markdown文档内容。

**响应格式**：`text/markdown`

## 📂 项目结构
```
.
├── main.go          # 主程序入口
├── go.mod           # Go模块定义
├── go.sum           # 依赖版本锁定文件
├── contents.json    # 文档目录配置文件
└── documents/       # 包含所有Markdown格式的文档
    └── [uuid].md    # 以UUID命名的Markdown文档
```

## 🚀 安装与运行
1. ✅ 确保已安装Go 1.21或更高版本
2. 📥 克隆项目仓库
   ```
   git clone https://github.com/yourusername/markdown-notebook-backend.git
   cd markdown-notebook-backend
   ```
3. 📦 安装依赖
   ```
   go mod download
   ```
4. ▶️ 运行服务
   ```
   go run main.go
   ```
5. 🌐 服务将在 http://localhost:8080 上启动

## 🔍 如何使用
服务启动后，可以通过以下方式访问文档：

1. 📋 获取文档目录：
   ```
   curl http://localhost:8080/contents
   ```

2. 📄 获取特定文档（使用从目录中获取的UUID）：
   ```
   curl http://localhost:8080/documents/[uuid]
   ```

## 📝 contents.json格式
`contents.json`文件包含所有可用文档的信息，格式如下：
```json

{
  "文章名1": "uuid-1",
  "文章名2": "uuid-2"
}

```

## ➕ 添加新文档
1. 📄 在`documents/`目录中添加新的Markdown文档，文件名应为UUID格式（例如：`550e8400-e29b-41d4-a716-446655440000.md`）
2. 🔄 在`contents.json`文件中添加新文档的信息和对应的UUID

## 🔒 安全注意事项
- 🌐 当前CORS配置允许所有来源访问（`"*"`）。在生产环境中，建议将其限制为特定的域名
- 🛡️ 考虑添加更多的安全措施，如请求限制、身份验证和授权机制

## 👥 贡献指南
1. 🔱 Fork本仓库
2. 🌿 创建您的特性分支 (`git checkout -b feature/AmazingFeature`)
3. 💾 提交您的更改 (`git commit -m 'Add some AmazingFeature'`)
4. 📤 推送到分支 (`git push origin feature/AmazingFeature`)
5. 🔄 打开一个Pull Request

## 📄 许可证
本项目采用MIT许可证 - 详情请参阅LICENSE文件 