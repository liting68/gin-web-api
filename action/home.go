package action

//HomeAction 首页控制器
type HomeAction struct{}

//Login 登录
func (action *HomeAction) Login(c *gin.Context) {
	result.PrintJSON(c, result.OK, "")
}

func (action *HomeAction) InitInfo(c *gin.Context) {
	result.PrintJSON(c, result.OK, "")
}