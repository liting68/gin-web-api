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
\controller
    └ crontab.go 定时任务
    └ manage.go 管理员功能模块
    └ router.go 路由注册模块
    └ user.go 用户功能模块
\db
    └ db.go 数据库实例
\model
    └ admin.go 管理员模块
    └ bearer.go 系统角色模块
    └ crontab.go 定时任务模块
    └ datetime.go 模型数据日期时间格式化模块
    └ user.go 用户模块
\route
    └ auth.go 鉴权验证
    └ resp.go 响应数据格式及状态码
    └ router.go 路由初始化
\test
    └ main_test.go 单元测试模块
\utils
    └ aesencrypt.go 加密工具
    └ captcha.go 验证码工具
    └ curl.go 模拟请求工具
    └ file.go 文件处理工具
    └ isgbk.go gbk编码判断工具
    └ isutf8.go utf8编码判断工具
    └ time.go 日期时间转换工具
    └ upload.go 文件上传工具
    └ wechat.go 微信sdk
main.go 应用入口
```
## 使用方法
```text
go run main.go
```


