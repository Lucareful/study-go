package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _StudentsMgr struct {
	*_BaseMgr
}

// StudentsMgr open func
func StudentsMgr(db *gorm.DB) *_StudentsMgr {
	if db == nil {
		panic(fmt.Errorf("StudentsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StudentsMgr{_BaseMgr: &_BaseMgr{DB: db.Model(Students{}), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StudentsMgr) GetTableName() string {
	return "students"
}

// Get 获取
func (obj *_StudentsMgr) Get() (result Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StudentsMgr) Gets() (results []*Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 Primary key
func (obj *_StudentsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCreatedAt created_at获取 created time
func (obj *_StudentsMgr) WithCreatedAt(createdAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_at"] = createdAt })
}

// WithUpdatedAt updated_at获取 updated at
func (obj *_StudentsMgr) WithUpdatedAt(updatedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_at"] = updatedAt })
}

// WithDeletedAt deleted_at获取 deleted time
func (obj *_StudentsMgr) WithDeletedAt(deletedAt time.Time) Option {
	return optionFunc(func(o *options) { o.query["deleted_at"] = deletedAt })
}

// WithName name获取
func (obj *_StudentsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAge age获取
func (obj *_StudentsMgr) WithAge(age int64) Option {
	return optionFunc(func(o *options) { o.query["age"] = age })
}

// GetByOption 功能选项模式获取
func (obj *_StudentsMgr) GetByOption(opts ...Option) (result Students, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_StudentsMgr) GetByOptions(opts ...Option) (results []*Students, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 Primary key
func (obj *_StudentsMgr) GetFromID(id int64) (result Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 Primary key
func (obj *_StudentsMgr) GetBatchFromID(ids []int64) (results []*Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" IN (?)", ids).Find(&results).Error

	return
}

// GetFromCreatedAt 通过created_at获取内容 created time
func (obj *_StudentsMgr) GetFromCreatedAt(createdAt time.Time) (result Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" = ?", createdAt).Find(&result).Error

	return
}

// GetBatchFromCreatedAt 批量查找 created time
func (obj *_StudentsMgr) GetBatchFromCreatedAt(createdAts []time.Time) (results []*Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" IN (?)", createdAts).Find(&results).Error

	return
}

// GetFromUpdatedAt 通过updated_at获取内容 updated at
func (obj *_StudentsMgr) GetFromUpdatedAt(updatedAt time.Time) (result Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" = ?", updatedAt).Find(&result).Error

	return
}

// GetBatchFromUpdatedAt 批量查找 updated at
func (obj *_StudentsMgr) GetBatchFromUpdatedAt(updatedAts []time.Time) (results []*Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" IN (?)", updatedAts).Find(&results).Error

	return
}

// GetFromDeletedAt 通过deleted_at获取内容 deleted time
func (obj *_StudentsMgr) GetFromDeletedAt(deletedAt time.Time) (result Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" = ?", deletedAt).Find(&result).Error

	return
}

// GetBatchFromDeletedAt 批量查找 deleted time
func (obj *_StudentsMgr) GetBatchFromDeletedAt(deletedAts []time.Time) (results []*Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where(" IN (?)", deletedAts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_StudentsMgr) GetFromName(name string) (results []*Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_StudentsMgr) GetBatchFromName(names []string) (results []*Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAge 通过age获取内容
func (obj *_StudentsMgr) GetFromAge(age int64) (results []*Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`age` = ?", age).Find(&results).Error

	return
}

// GetBatchFromAge 批量查找
func (obj *_StudentsMgr) GetBatchFromAge(ages []int64) (results []*Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`age` IN (?)", ages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_StudentsMgr) FetchByPrimaryKey(id int64) (result Students, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("`id` = ?", id).Find(&result).Error

	return
}
