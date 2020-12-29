package middlewares

import "diy-server/utils"

// 初始化log 与 jwt 工具
var (
	log = utils.Log
	jwt = new(utils.Jwt)
)