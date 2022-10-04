# gin-web-api
Golang Web API framework 基于Golang 开发的Web API框架

## 配置文件
```text
\config\config.yaml
```
## 目录说明
```text
\config
    └ config.go 配置结构
    └ config.yaml 配置文件
\db
    └ db.go 数据库操作实例
\controller
    └ crontab.go 定时任务调用控制器
    └ router.go 路由调用控制器
    └ manage.go 管理员功能调用控制器
    └ user.go 用户功能调用控制器
\model
    └ admin.go 管理员模型
    └ bearer.go 系统角色模型
    └ crontab.go 定时任务模型
    └ datetime.go 日期时间序列化模块
    └ user.go 用户模型
\route
    └ auth.go 鉴权验证模块
    └ resp.go 状态码及响应数据模块
    └ router.go 路由初始化模块
\utils
    └ aesencrypt.go 加密工具
    └ captcha.go 验证码工具
    └ curl.go 模拟请求工具
    └ file.go 文件处理工具
    └ isgbk.go GBK编码判断工具
    └ isutf8.go UTF8编码判断工具
    └ time.go 日期时间转换工具
    └ upload.go 文件上传工具
    └ wechat.go 微信服务端SDK
\test
    └ main_test.go 单元测试模块
main.go 系统入口
```
## 使用方法
```text
go run main.go
```


