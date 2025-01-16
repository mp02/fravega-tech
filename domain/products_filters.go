package domain

import "errors"

type ProductFilters struct {
	Name       *string   `form:"name"`
	MinPrice   *float64  `form:"min_price"`
	MaxPrice   *float64  `form:"max_price"`
	Categories *[]string `form:"categories"`
	IsDeleted  *bool     `form:"is_deleted"`
}

func (f *ProductFilters) Validate() error {
	if *f.MinPrice > *f.MaxPrice && *f.MaxPrice > 0 {
		return errors.New("min_price cannot be greater than max_price")
	}
	return nil
}

func (f *ProductFilters) AreFiltersEmpty() bool {
	if f.MinPrice == nil && f.MaxPrice == nil && f.Categories == nil && f.IsDeleted == nil {
		return true
	}
	return false
}
