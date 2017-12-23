package subtype

import (
	"github.com/rancher/norman/types"
)

type Store struct {
	types.Store
	subType string
}

func NewSubTypeStore(subType string, store types.Store) *Store {
	return &Store{
		Store:   store,
		subType: subType,
	}
}

func (p *Store) List(apiContext *types.APIContext, schema *types.Schema, opt types.QueryOptions) ([]map[string]interface{}, error) {
	opt.Conditions = append(opt.Conditions, types.NewConditionFromString("type", types.ModifierEQ, p.subType))
	return p.Store.List(apiContext, schema, opt)
}

func (p *Store) Watch(apiContext *types.APIContext, schema *types.Schema, opt types.QueryOptions) (chan map[string]interface{}, error) {
	opt.Conditions = append(opt.Conditions, types.NewConditionFromString("type", types.ModifierEQ, p.subType))
	return p.Store.Watch(apiContext, schema, opt)
}