package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _PolygonsContractMgr struct {
	*_BaseMgr
}

// PolygonsContractMgr open func
func PolygonsContractMgr(db *gorm.DB) *_PolygonsContractMgr {
	if db == nil {
		panic(fmt.Errorf("PolygonsContractMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PolygonsContractMgr{_BaseMgr: &_BaseMgr{DB: db.Table("polygons_contract"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PolygonsContractMgr) GetTableName() string {
	return "polygons_contract"
}

// Get 获取
func (obj *_PolygonsContractMgr) Get() (result PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PolygonsContractMgr) Gets() (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PolygonsContractMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取 合约类型：0-预言机订阅合约 1-721合约 2-挂单交易所合约 3-竞拍交易所合约 4-授权合约
func (obj *_PolygonsContractMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithContractAddress contract_address获取 合约地址
func (obj *_PolygonsContractMgr) WithContractAddress(contractAddress string) Option {
	return optionFunc(func(o *options) { o.query["contract_address"] = contractAddress })
}

// WithName name获取 合约名称
func (obj *_PolygonsContractMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithOwnerAddress owner_address获取 合约所有者地址
func (obj *_PolygonsContractMgr) WithOwnerAddress(ownerAddress string) Option {
	return optionFunc(func(o *options) { o.query["owner_address"] = ownerAddress })
}

// WithStatus status获取 状态：0-停用 1-启用
func (obj *_PolygonsContractMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithPropID prop_id获取 合约绑定的道具ID：-1 没有道具
func (obj *_PolygonsContractMgr) WithPropID(propID int64) Option {
	return optionFunc(func(o *options) { o.query["prop_id"] = propID })
}

// WithCreatedTime created_time获取 创建时间
func (obj *_PolygonsContractMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// GetByOption 功能选项模式获取
func (obj *_PolygonsContractMgr) GetByOption(opts ...Option) (result PolygonsContract, err error) {
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
func (obj *_PolygonsContractMgr) GetByOptions(opts ...Option) (results []*PolygonsContract, err error) {
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
func (obj *_PolygonsContractMgr) GetFromID(id uint64) (result PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_PolygonsContractMgr) GetBatchFromID(ids []uint64) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 合约类型：0-预言机订阅合约 1-721合约 2-挂单交易所合约 3-竞拍交易所合约 4-授权合约
func (obj *_PolygonsContractMgr) GetFromType(_type int) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 合约类型：0-预言机订阅合约 1-721合约 2-挂单交易所合约 3-竞拍交易所合约 4-授权合约
func (obj *_PolygonsContractMgr) GetBatchFromType(_types []int) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromContractAddress 通过contract_address获取内容 合约地址
func (obj *_PolygonsContractMgr) GetFromContractAddress(contractAddress string) (result PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_address = ?", contractAddress).Find(&result).Error

	return
}

// GetBatchFromContractAddress 批量查找 合约地址
func (obj *_PolygonsContractMgr) GetBatchFromContractAddress(contractAddresss []string) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_address IN (?)", contractAddresss).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 合约名称
func (obj *_PolygonsContractMgr) GetFromName(name string) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 合约名称
func (obj *_PolygonsContractMgr) GetBatchFromName(names []string) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("name IN (?)", names).Find(&results).Error

	return
}

// GetFromOwnerAddress 通过owner_address获取内容 合约所有者地址
func (obj *_PolygonsContractMgr) GetFromOwnerAddress(ownerAddress string) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("owner_address = ?", ownerAddress).Find(&results).Error

	return
}

// GetBatchFromOwnerAddress 批量查找 合约所有者地址
func (obj *_PolygonsContractMgr) GetBatchFromOwnerAddress(ownerAddresss []string) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("owner_address IN (?)", ownerAddresss).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态：0-停用 1-启用
func (obj *_PolygonsContractMgr) GetFromStatus(status int) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态：0-停用 1-启用
func (obj *_PolygonsContractMgr) GetBatchFromStatus(statuss []int) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromPropID 通过prop_id获取内容 合约绑定的道具ID：-1 没有道具
func (obj *_PolygonsContractMgr) GetFromPropID(propID int64) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("prop_id = ?", propID).Find(&results).Error

	return
}

// GetBatchFromPropID 批量查找 合约绑定的道具ID：-1 没有道具
func (obj *_PolygonsContractMgr) GetBatchFromPropID(propIDs []int64) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("prop_id IN (?)", propIDs).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 创建时间
func (obj *_PolygonsContractMgr) GetFromCreatedTime(createdTime time.Time) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 创建时间
func (obj *_PolygonsContractMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PolygonsContractMgr) FetchByPrimaryKey(id uint64) (result PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByContractIndex primay or index 获取唯一内容
func (obj *_PolygonsContractMgr) FetchUniqueByContractIndex(contractAddress string) (result PolygonsContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_address = ?", contractAddress).Find(&result).Error

	return
}
