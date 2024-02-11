package app

import (
	ItemController "aiconec/commerce-adapter/internal/domain/item/controller"
	"aiconec/commerce-adapter/internal/port"

	"github.com/emirpasic/gods/maps/hashmap"
)

func (r *AppRegistry) loadControllers() *hashmap.Map {
	var ok bool
	var exist bool
	hm := hashmap.New()

	var _itemUc any
	var itemUc port.ServiceUsecase
	_itemUc, exist = r.usecases.Get("item")
	itemUc, ok = _itemUc.(port.ServiceUsecase)
	if exist && ok {
		hm.Put("item", ItemController.New(r.cfg, itemUc))
	}

	return hm
}
