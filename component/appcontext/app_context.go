package appcontext

import (
	"gorm.io/gorm"
	"ocean-app-be/component/uploadprovider"
	"ocean-app-be/pubsub"
)

type AppCtx interface {
	GetMainDBConnection() *gorm.DB

	SecretKey() string

	UploadProvider() uploadprovider.UploadProvider

	GetPubsub() pubsub.Pubsub
}

type appCtx struct {
	db             *gorm.DB
	secretKey      string
	uploadProvider uploadprovider.UploadProvider
	ps             pubsub.Pubsub
}

func NewAppCtx(db *gorm.DB, secretKey string, uploadProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, uploadProvider: uploadProvider}
}
func (ctx *appCtx) GetMainDBConnection() *gorm.DB { return ctx.db }

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}
func (ctx *appCtx) UploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) GetPubsub() pubsub.Pubsub {
	return ctx.ps
}
