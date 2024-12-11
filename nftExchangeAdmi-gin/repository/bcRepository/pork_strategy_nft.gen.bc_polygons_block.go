package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _BcPolygonsBlockMgr struct {
	*_BaseMgr
}

// BcPolygonsBlockMgr open func
func BcPolygonsBlockMgr(db *gorm.DB) *_BcPolygonsBlockMgr {
	if db == nil {
		panic(fmt.Errorf("BcPolygonsBlockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BcPolygonsBlockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bc_polygons_block"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BcPolygonsBlockMgr) GetTableName() string {
	return "bc_polygons_block"
}

// Get 获取
func (obj *_BcPolygonsBlockMgr) Get() (result BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_BcPolygonsBlockMgr) Gets() (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithHeight height获取 高度
func (obj *_BcPolygonsBlockMgr) WithHeight(height uint64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithParentHash parent_hash获取 父hash
func (obj *_BcPolygonsBlockMgr) WithParentHash(parentHash string) Option {
	return optionFunc(func(o *options) { o.query["parent_hash"] = parentHash })
}

// WithBlockHash block_hash获取 区块hash
func (obj *_BcPolygonsBlockMgr) WithBlockHash(blockHash string) Option {
	return optionFunc(func(o *options) { o.query["block_hash"] = blockHash })
}

// WithTransaction transaction获取 交易
func (obj *_BcPolygonsBlockMgr) WithTransaction(transaction int) Option {
	return optionFunc(func(o *options) { o.query["transaction"] = transaction })
}

// WithContractTransaction contract_transaction获取 合约内部交易
func (obj *_BcPolygonsBlockMgr) WithContractTransaction(contractTransaction int) Option {
	return optionFunc(func(o *options) { o.query["contract_transaction"] = contractTransaction })
}

// WithBlockTimeStamp block_time_stamp获取 出块时间
func (obj *_BcPolygonsBlockMgr) WithBlockTimeStamp(blockTimeStamp int64) Option {
	return optionFunc(func(o *options) { o.query["block_time_stamp"] = blockTimeStamp })
}

// WithCreatedTime created_time获取 创建时间
func (obj *_BcPolygonsBlockMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// GetByOption 功能选项模式获取
func (obj *_BcPolygonsBlockMgr) GetByOption(opts ...Option) (result BcPolygonsBlock, err error) {
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
func (obj *_BcPolygonsBlockMgr) GetByOptions(opts ...Option) (results []*BcPolygonsBlock, err error) {
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

// GetFromHeight 通过height获取内容 高度
func (obj *_BcPolygonsBlockMgr) GetFromHeight(height uint64) (result BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&result).Error

	return
}

// GetBatchFromHeight 批量查找 高度
func (obj *_BcPolygonsBlockMgr) GetBatchFromHeight(heights []uint64) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height IN (?)", heights).Find(&results).Error

	return
}

// GetFromParentHash 通过parent_hash获取内容 父hash
func (obj *_BcPolygonsBlockMgr) GetFromParentHash(parentHash string) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_hash = ?", parentHash).Find(&results).Error

	return
}

// GetBatchFromParentHash 批量查找 父hash
func (obj *_BcPolygonsBlockMgr) GetBatchFromParentHash(parentHashs []string) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("parent_hash IN (?)", parentHashs).Find(&results).Error

	return
}

// GetFromBlockHash 通过block_hash获取内容 区块hash
func (obj *_BcPolygonsBlockMgr) GetFromBlockHash(blockHash string) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("block_hash = ?", blockHash).Find(&results).Error

	return
}

// GetBatchFromBlockHash 批量查找 区块hash
func (obj *_BcPolygonsBlockMgr) GetBatchFromBlockHash(blockHashs []string) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("block_hash IN (?)", blockHashs).Find(&results).Error

	return
}

// GetFromTransaction 通过transaction获取内容 交易
func (obj *_BcPolygonsBlockMgr) GetFromTransaction(transaction int) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("transaction = ?", transaction).Find(&results).Error

	return
}

// GetBatchFromTransaction 批量查找 交易
func (obj *_BcPolygonsBlockMgr) GetBatchFromTransaction(transactions []int) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("transaction IN (?)", transactions).Find(&results).Error

	return
}

// GetFromContractTransaction 通过contract_transaction获取内容 合约内部交易
func (obj *_BcPolygonsBlockMgr) GetFromContractTransaction(contractTransaction int) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_transaction = ?", contractTransaction).Find(&results).Error

	return
}

// GetBatchFromContractTransaction 批量查找 合约内部交易
func (obj *_BcPolygonsBlockMgr) GetBatchFromContractTransaction(contractTransactions []int) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_transaction IN (?)", contractTransactions).Find(&results).Error

	return
}

// GetFromBlockTimeStamp 通过block_time_stamp获取内容 出块时间
func (obj *_BcPolygonsBlockMgr) GetFromBlockTimeStamp(blockTimeStamp int64) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("block_time_stamp = ?", blockTimeStamp).Find(&results).Error

	return
}

// GetBatchFromBlockTimeStamp 批量查找 出块时间
func (obj *_BcPolygonsBlockMgr) GetBatchFromBlockTimeStamp(blockTimeStamps []int64) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("block_time_stamp IN (?)", blockTimeStamps).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 创建时间
func (obj *_BcPolygonsBlockMgr) GetFromCreatedTime(createdTime time.Time) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 创建时间
func (obj *_BcPolygonsBlockMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_BcPolygonsBlockMgr) FetchByPrimaryKey(height uint64) (result BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&result).Error

	return
}

// FetchUniqueByHeightIndex primay or index 获取唯一内容
func (obj *_BcPolygonsBlockMgr) FetchUniqueByHeightIndex(height uint64) (result BcPolygonsBlock, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&result).Error

	return
}
