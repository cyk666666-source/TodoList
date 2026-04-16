# TodoList - 基于 Gin + GORM 的全栈待办事项管理系统

> 一个学习 Go Web 开发的实战项目，基于 Gin 框架 + GORM 实现，前后端一体化的待办事项管理工具。

## 🚀 项目简介
本项目是 Go Web 开发实战练手项目，完整实现待办事项的增删改查功能，包含后端 API 服务与前端页面展示：
- 后端：基于 Gin 框架搭建 RESTful API
- 数据库：使用 GORM 操作 MySQL
- 前端：内置静态页面，开箱即用
- 功能：新增、查询、修改、删除、状态切换

## 🛠️ 技术栈
### 后端
- 语言：Go 1.20+
- Web 框架：Gin
- ORM：GORM
- 数据库：MySQL
- 依赖管理：Go Modules

### 前端
- 技术：HTML + CSS + JS
- 静态资源：内置在 static/ 与 templates/

## 📁 项目结构
```text
TodoList/
├── controller/    # 控制器层：处理请求与响应
├── dao/           # 数据访问层：数据库操作
├── models/        # 模型层：数据结构与表结构映射
├── routers/       # 路由层：API 路由配置
├── static/        # 前端静态资源（JS/CSS/图标）
├── templates/     # 前端 HTML 模板
├── main.go        # 项目入口
├── go.mod         # Go 依赖文件
├── go.sum         # 依赖校验文件
└── .gitignore     # Git 忽略配置
```

## ⚙️ 环境要求
- Go 1.20及以上版本
- MySQL 5.7及以上版本
## ✨ 功能清单
```text
- 添加待办事项
- 查看所有事项
- 编辑事项内容
- 删除事项
- 标记完成/未完成
- 前端页面可视化操作
```
## 🙏 说明
- 本项目用于Go Web开发学习与实践，代码简洁易懂，适合新手入门Gin+GORM
