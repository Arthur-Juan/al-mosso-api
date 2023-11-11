package interfaces

import "gorm.io/gorm"

type IDatabase interface {
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Save(velue interface{}) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Select(query interface{}, args ...interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
}
