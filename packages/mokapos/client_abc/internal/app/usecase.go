package app

import (
	ItemUsecase "aiconec/commerce-adapter/internal/domain/item/usecase"

	"github.com/emirpasic/gods/maps/hashmap"
)

func (r *AppRegistry) loadUsecases() *hashmap.Map {
	hm := hashmap.New()

	itemCtrl := ItemUsecase.New(r.cfg)

	hm.Put("item", itemCtrl)

	return hm
}
