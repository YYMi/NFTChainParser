package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _BcConfigMgr struct {
	*_BaseMgr
}

// BcConfigMgr open func
func BcConfigMgr(db *gorm.DB) *_BcConfigMgr {
	if db == nil {
		panic(fmt.Errorf("BcConfigMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BcConfigMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bc_config"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BcConfigMgr) GetTableName() string {
	return "bc_config"
}

// Get 获取
func (obj *_BcConfigMgr) Get() (result BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BcConfigMgr) Gets() (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_BcConfigMgr) WithID(id uint32) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithModule module获取 模块名称
func (obj *_BcConfigMgr) WithModule(module string) Option {
	return optionFunc(func(o *options) { o.query["module"] = module })
}

// WithName name获取 配置变量名
func (obj *_BcConfigMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithValue value获取 值
func (obj *_BcConfigMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithDescription description获取 描述
func (obj *_BcConfigMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithUpdatedTime updated_time获取 更新时间
func (obj *_BcConfigMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// WithCreatedTime created_time获取 创建时间
func (obj *_BcConfigMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// GetByOption 功能选项模式获取
func (obj *_BcConfigMgr) GetByOption(opts ...Option) (result BcConfig, err error) {
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
func (obj *_BcConfigMgr) GetByOptions(opts ...Option) (results []*BcConfig, err error) {
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

// GetFromID 通过id获取内容
func (obj *_BcConfigMgr) GetFromID(id uint32) (result BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_BcConfigMgr) GetBatchFromID(ids []uint32) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromModule 通过module获取内容 模块名称
func (obj *_BcConfigMgr) GetFromModule(module string) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module = ?", module).Find(&results).Error

	return
}

// GetBatchFromModule 批量查找 模块名称
func (obj *_BcConfigMgr) GetBatchFromModule(modules []string) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("module IN (?)", modules).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 配置变量名
func (obj *_BcConfigMgr) GetFromName(name string) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 配置变量名
func (obj *_BcConfigMgr) GetBatchFromName(names []string) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容 值
func (obj *_BcConfigMgr) GetFromValue(value string) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找 值
func (obj *_BcConfigMgr) GetBatchFromValue(values []string) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("value IN (?)", values).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 描述
func (obj *_BcConfigMgr) GetFromDescription(description string) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 描述
func (obj *_BcConfigMgr) GetBatchFromDescription(descriptions []string) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("description IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 更新时间
func (obj *_BcConfigMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 更新时间
func (obj *_BcConfigMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 创建时间
func (obj *_BcConfigMgr) GetFromCreatedTime(createdTime time.Time) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 创建时间
func (obj *_BcConfigMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BcConfigMgr) FetchByPrimaryKey(id uint32) (result BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByNameValueIndex primay or index 获取唯一内容
func (obj *_BcConfigMgr) FetchUniqueIndexByNameValueIndex(name string, value string) (result BcConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ? AND value = ?", name, value).Find(&result).Error

	return
}
