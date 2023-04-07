package controllers

import (
	"github.com/google/wire"
	"go-fiber-demo/services"
)

type BaseController struct {
	i18n services.I18n
}

func NewBaseController(i18n services.I18n) BaseController {
	return BaseController{i18n: i18n}
}

var BaseControllerSet = wire.NewSet(NewBaseController, services.I18nSet)
