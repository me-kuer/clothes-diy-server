package home

import (
	"diy-server/models"
	"diy-server/utils"
)

var(
	db = models.Engine
	log = utils.Log
	jwt = new(utils.Jwt)
)
