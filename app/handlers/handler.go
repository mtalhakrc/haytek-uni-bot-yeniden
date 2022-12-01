package handlers

import (
	"context"
	"errors"
	"github.com/haytek-uni-bot-yeniden/app/service"
	"github.com/haytek-uni-bot-yeniden/common/model"
	"github.com/haytek-uni-bot-yeniden/pkg/app"
	baseservice "github.com/haytek-uni-bot-yeniden/pkg/service"
	"github.com/uptrace/bun"
	"strings"
)

type UserHandler struct {
	baseservice.IBaseService[model.User]
	userService   service.IUserService
	sheetsService service.ISheetsService
}

func NewUserHandler(db *bun.DB, s service.ISheetsService) *UserHandler {
	return &UserHandler{
		IBaseService:  baseservice.NewBaseService[model.User](db),
		userService:   service.NewUserService(db),
		sheetsService: s,
	}
}
func (u UserHandler) Kaydol(ctx *app.Ctx, params []string) (string, error) {

	username := ctx.SentFrom().String()
	userid := ctx.SentFrom().ID

	//zaten kaydoldu mu kontrol et
	_, err := u.userService.GetByUserID(context.Background(), userid)
	if err == nil {
		err = errors.New("kaydınız zaten var")
		return "", err
	}

	name := strings.Join(params, " ")

	if name == "" {
		return "", errors.New("komuttan sonra isminizi belirtiniz(ör: /kaydol Talha Karaca)")
	}
	if !u.sheetsService.TestSheetExist(name) {
		return "", errors.New("isminiz çetelede bulunamadı, çetelenizde bu isme ait bir sayfa olduğundan emin olup tekrar deneyiniz")
	}

	user := model.User{
		Name:     name,
		Username: username,
		UserID:   userid,
		Type:     model.UserTypeNormal,
	}

	err = u.Create(context.Background(), &user)
	if err != nil {
		return "", err
	}
	return "kaydınız başarı ile gerçekleştirildi", nil
}

func (u UserHandler) KayitSil(ctx *app.Ctx, params []string) (string, error) {
	err := u.DeleteByUserID(context.Background(), ctx.SentFrom().ID)
	return "kaydınız başarı ile silindi", err
}
