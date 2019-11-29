# gin-web-api
golang gin web api framework 基于Gin开发的api框架，开箱即用

## 配置文件
```text
\config\config.yaml
```
## 目录说明
```text
\config
    └  config.go 配置结构
    └ config.yaml 配置信息
\controller
    └ home.go 首页控制器
    └ manage.go 管理员控制器
\db
    └ db.go 数据库实例
\lib
    └ auth.go 授权验证
    └ cron.go 定时任务
    └ router.go 路由注册及应用初始化
\model
    └ admin.go 管理员
    └ datetime.go 时间格式化
    └ user.go 用户
\resp
    └ response.go 响应
\test
    └ main_test.go 单元测试
\utils
    └ curl.go curl模拟请求，包含get、post、json
    └ upload.go 文件上传，base64文件上传
    └ wechat.go 微信sdk
main.go 应用入口
```
## 使用方法
```text
Linux 环境 go run main.go
Windows 环境 go run win.go
```
## 平滑重启
kill -1 pid
