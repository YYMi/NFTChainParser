package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _PolygonChainlinkContractTxMgr struct {
	*_BaseMgr
}

// PolygonChainlinkContractTxMgr open func
func PolygonChainlinkContractTxMgr(db *gorm.DB) *_PolygonChainlinkContractTxMgr {
	if db == nil {
		panic(fmt.Errorf("PolygonChainlinkContractTxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PolygonChainlinkContractTxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("polygon_chainlink_contract_tx"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PolygonChainlinkContractTxMgr) GetTableName() string {
	return "polygon_chainlink_contract_tx"
}

// Get 获取
func (obj *_PolygonChainlinkContractTxMgr) Get() (result PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PolygonChainlinkContractTxMgr) Gets() (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PolygonChainlinkContractTxMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHeight height获取 区块高度
func (obj *_PolygonChainlinkContractTxMgr) WithHeight(height int64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithMethodID method_id获取 交易方法id
func (obj *_PolygonChainlinkContractTxMgr) WithMethodID(methodID string) Option {
	return optionFunc(func(o *options) { o.query["method_id"] = methodID })
}

// WithContractAdd contract_add获取 合约地址
func (obj *_PolygonChainlinkContractTxMgr) WithContractAdd(contractAdd string) Option {
	return optionFunc(func(o *options) { o.query["contract_add"] = contractAdd })
}

// WithFromAddr from_addr获取 from地址
func (obj *_PolygonChainlinkContractTxMgr) WithFromAddr(fromAddr string) Option {
	return optionFunc(func(o *options) { o.query["from_addr"] = fromAddr })
}

// WithTxID tx_id获取 交易hash
func (obj *_PolygonChainlinkContractTxMgr) WithTxID(txID string) Option {
	return optionFunc(func(o *options) { o.query["tx_id"] = txID })
}

// WithStatus status获取 交易状态（0-等待链确认 1-到达安全高度 2-交易失败）
func (obj *_PolygonChainlinkContractTxMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithFee fee获取 交易收费
func (obj *_PolygonChainlinkContractTxMgr) WithFee(fee float64) Option {
	return optionFunc(func(o *options) { o.query["fee"] = fee })
}

// WithMethodParam method_param获取 调用参数
func (obj *_PolygonChainlinkContractTxMgr) WithMethodParam(methodParam string) Option {
	return optionFunc(func(o *options) { o.query["method_param"] = methodParam })
}

// WithRemark remark获取 备注
func (obj *_PolygonChainlinkContractTxMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithUpdatedTime updated_time获取 更新时间
func (obj *_PolygonChainlinkContractTxMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// WithCreatadTime creatad_time获取 创建时间
func (obj *_PolygonChainlinkContractTxMgr) WithCreatadTime(creatadTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["creatad_time"] = creatadTime })
}

// GetByOption 功能选项模式获取
func (obj *_PolygonChainlinkContractTxMgr) GetByOption(opts ...Option) (result PolygonChainlinkContractTx, err error) {
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
func (obj *_PolygonChainlinkContractTxMgr) GetByOptions(opts ...Option) (results []*PolygonChainlinkContractTx, err error) {
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
func (obj *_PolygonChainlinkContractTxMgr) GetFromID(id uint64) (result PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromID(ids []uint64) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容 区块高度
func (obj *_PolygonChainlinkContractTxMgr) GetFromHeight(height int64) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量查找 区块高度
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromHeight(heights []int64) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height IN (?)", heights).Find(&results).Error

	return
}

// GetFromMethodID 通过method_id获取内容 交易方法id
func (obj *_PolygonChainlinkContractTxMgr) GetFromMethodID(methodID string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_id = ?", methodID).Find(&results).Error

	return
}

// GetBatchFromMethodID 批量查找 交易方法id
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromMethodID(methodIDs []string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_id IN (?)", methodIDs).Find(&results).Error

	return
}

// GetFromContractAdd 通过contract_add获取内容 合约地址
func (obj *_PolygonChainlinkContractTxMgr) GetFromContractAdd(contractAdd string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_add = ?", contractAdd).Find(&results).Error

	return
}

// GetBatchFromContractAdd 批量查找 合约地址
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromContractAdd(contractAdds []string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_add IN (?)", contractAdds).Find(&results).Error

	return
}

// GetFromFromAddr 通过from_addr获取内容 from地址
func (obj *_PolygonChainlinkContractTxMgr) GetFromFromAddr(fromAddr string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_addr = ?", fromAddr).Find(&results).Error

	return
}

// GetBatchFromFromAddr 批量查找 from地址
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromFromAddr(fromAddrs []string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_addr IN (?)", fromAddrs).Find(&results).Error

	return
}

// GetFromTxID 通过tx_id获取内容 交易hash
func (obj *_PolygonChainlinkContractTxMgr) GetFromTxID(txID string) (result PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tx_id = ?", txID).Find(&result).Error

	return
}

// GetBatchFromTxID 批量查找 交易hash
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromTxID(txIDs []string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tx_id IN (?)", txIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 交易状态（0-等待链确认 1-到达安全高度 2-交易失败）
func (obj *_PolygonChainlinkContractTxMgr) GetFromStatus(status int) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 交易状态（0-等待链确认 1-到达安全高度 2-交易失败）
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromStatus(statuss []int) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromFee 通过fee获取内容 交易收费
func (obj *_PolygonChainlinkContractTxMgr) GetFromFee(fee float64) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee = ?", fee).Find(&results).Error

	return
}

// GetBatchFromFee 批量查找 交易收费
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromFee(fees []float64) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee IN (?)", fees).Find(&results).Error

	return
}

// GetFromMethodParam 通过method_param获取内容 调用参数
func (obj *_PolygonChainlinkContractTxMgr) GetFromMethodParam(methodParam string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_param = ?", methodParam).Find(&results).Error

	return
}

// GetBatchFromMethodParam 批量查找 调用参数
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromMethodParam(methodParams []string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_param IN (?)", methodParams).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_PolygonChainlinkContractTxMgr) GetFromRemark(remark string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromRemark(remarks []string) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 更新时间
func (obj *_PolygonChainlinkContractTxMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 更新时间
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

// GetFromCreatadTime 通过creatad_time获取内容 创建时间
func (obj *_PolygonChainlinkContractTxMgr) GetFromCreatadTime(creatadTime time.Time) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creatad_time = ?", creatadTime).Find(&results).Error

	return
}

// GetBatchFromCreatadTime 批量查找 创建时间
func (obj *_PolygonChainlinkContractTxMgr) GetBatchFromCreatadTime(creatadTimes []time.Time) (results []*PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creatad_time IN (?)", creatadTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PolygonChainlinkContractTxMgr) FetchByPrimaryKey(id uint64) (result PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByTxidIndex primay or index 获取唯一内容
func (obj *_PolygonChainlinkContractTxMgr) FetchUniqueByTxidIndex(txID string) (result PolygonChainlinkContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tx_id = ?", txID).Find(&result).Error

	return
}
