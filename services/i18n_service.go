package services

import (
	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type I18n struct {
	i18nBundle goi18n.Bundle
}

func NewI18n() I18n {
	bundle := goi18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("./resources/langs/active.en.toml")
	bundle.MustLoadMessageFile("./resources/langs/active.fr.toml")
	return I18n{i18nBundle: *bundle}
}

var I18nSet = wire.NewSet(NewI18n)

func GetLocale(ctx *fiber.Ctx) string {
	return ctx.Get("Accept-Language")
}

func (i18n I18n) Localize(ctx *fiber.Ctx, messageId string, template map[string]string) string {
	localizer := goi18n.NewLocalizer(&i18n.i18nBundle, GetLocale(ctx))
	return localizer.MustLocalize(&goi18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: template,
	})
}
