package appcontext

import "gorm.io/gorm"

type AppCtx interface {
	GetMainDBConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB { return ctx.db }

func NewAppCtx(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}
