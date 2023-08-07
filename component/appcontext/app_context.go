package appcontext

import "gorm.io/gorm"

type AppCtx interface {
	GetMainDBConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppCtx(db *gorm.DB) *appCtx {
	return NewAppCtx(db)
}
func (ctx *appCtx) GetMainDBConnection() *gorm.DB { return ctx.db }
