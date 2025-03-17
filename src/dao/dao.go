package dao

import (
	"context"

	"github.com/xavier-web3/XavSwapBase/stores/xkv"
	"gorm.io/gorm"
)

type Dao struct {
	ctx    context.Context
	DB     *gorm.DB
	vStore *xkv.Store
}
func New(ctx context.Context, db *gorm.DB, vStore *xkv.Store) *dao {
	return &Dao{
		ctx:    ctx,
		DB:     db,
		vStore: vStore,
	}
}
