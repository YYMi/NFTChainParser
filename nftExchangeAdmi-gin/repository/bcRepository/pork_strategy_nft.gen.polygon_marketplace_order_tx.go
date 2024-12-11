package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _PolygonMarketplaceOrderTxMgr struct {
	*_BaseMgr
}

// PolygonMarketplaceOrderTxMgr open func
func PolygonMarketplaceOrderTxMgr(db *gorm.DB) *_PolygonMarketplaceOrderTxMgr {
	if db == nil {
		panic(fmt.Errorf("PolygonMarketplaceOrderTxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PolygonMarketplaceOrderTxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("polygon_marketplace_order_tx"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PolygonMarketplaceOrderTxMgr) GetTableName() string {
	return "polygon_marketplace_order_tx"
}

// Get 获取
func (obj *_PolygonMarketplaceOrderTxMgr) Get() (result PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PolygonMarketplaceOrderTxMgr) Gets() (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PolygonMarketplaceOrderTxMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHeight height获取 区块高度
func (obj *_PolygonMarketplaceOrderTxMgr) WithHeight(height int64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithMethodID method_id获取 交易方法id
func (obj *_PolygonMarketplaceOrderTxMgr) WithMethodID(methodID string) Option {
	return optionFunc(func(o *options) { o.query["method_id"] = methodID })
}

// WithOrderID order_id获取 交易订单号
func (obj *_PolygonMarketplaceOrderTxMgr) WithOrderID(orderID string) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithContractAdd contract_add获取 合约地址
func (obj *_PolygonMarketplaceOrderTxMgr) WithContractAdd(contractAdd string) Option {
	return optionFunc(func(o *options) { o.query["contract_add"] = contractAdd })
}

// WithFromAddr from_addr获取 from地址
func (obj *_PolygonMarketplaceOrderTxMgr) WithFromAddr(fromAddr string) Option {
	return optionFunc(func(o *options) { o.query["from_addr"] = fromAddr })
}

// WithTxID tx_id获取 交易hash
func (obj *_PolygonMarketplaceOrderTxMgr) WithTxID(txID string) Option {
	return optionFunc(func(o *options) { o.query["tx_id"] = txID })
}

// WithStatus status获取 交易状态（0-等待链确认 1-到达安全高度 2-交易失败）
func (obj *_PolygonMarketplaceOrderTxMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithFee fee获取 交易收费
func (obj *_PolygonMarketplaceOrderTxMgr) WithFee(fee float64) Option {
	return optionFunc(func(o *options) { o.query["fee"] = fee })
}

// WithMethodParam method_param获取 调用参数
func (obj *_PolygonMarketplaceOrderTxMgr) WithMethodParam(methodParam string) Option {
	return optionFunc(func(o *options) { o.query["method_param"] = methodParam })
}

// WithRemark remark获取 备注
func (obj *_PolygonMarketplaceOrderTxMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithTimeStamp time_stamp获取 链确认时间
func (obj *_PolygonMarketplaceOrderTxMgr) WithTimeStamp(timeStamp int64) Option {
	return optionFunc(func(o *options) { o.query["time_stamp"] = timeStamp })
}

// WithUpdatedTime updated_time获取 更新时间
func (obj *_PolygonMarketplaceOrderTxMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// WithCreatadTime creatad_time获取 创建时间
func (obj *_PolygonMarketplaceOrderTxMgr) WithCreatadTime(creatadTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["creatad_time"] = creatadTime })
}

// GetByOption 功能选项模式获取
func (obj *_PolygonMarketplaceOrderTxMgr) GetByOption(opts ...Option) (result PolygonMarketplaceOrderTx, err error) {
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
func (obj *_PolygonMarketplaceOrderTxMgr) GetByOptions(opts ...Option) (results []*PolygonMarketplaceOrderTx, err error) {
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
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromID(id uint64) (result PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromID(ids []uint64) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容 区块高度
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromHeight(height int64) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量查找 区块高度
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromHeight(heights []int64) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height IN (?)", heights).Find(&results).Error

	return
}

// GetFromMethodID 通过method_id获取内容 交易方法id
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromMethodID(methodID string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_id = ?", methodID).Find(&results).Error

	return
}

// GetBatchFromMethodID 批量查找 交易方法id
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromMethodID(methodIDs []string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_id IN (?)", methodIDs).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 交易订单号
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromOrderID(orderID string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_id = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 交易订单号
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromOrderID(orderIDs []string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_id IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromContractAdd 通过contract_add获取内容 合约地址
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromContractAdd(contractAdd string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_add = ?", contractAdd).Find(&results).Error

	return
}

// GetBatchFromContractAdd 批量查找 合约地址
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromContractAdd(contractAdds []string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_add IN (?)", contractAdds).Find(&results).Error

	return
}

// GetFromFromAddr 通过from_addr获取内容 from地址
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromFromAddr(fromAddr string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_addr = ?", fromAddr).Find(&results).Error

	return
}

// GetBatchFromFromAddr 批量查找 from地址
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromFromAddr(fromAddrs []string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_addr IN (?)", fromAddrs).Find(&results).Error

	return
}

// GetFromTxID 通过tx_id获取内容 交易hash
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromTxID(txID string) (result PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tx_id = ?", txID).Find(&result).Error

	return
}

// GetBatchFromTxID 批量查找 交易hash
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromTxID(txIDs []string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tx_id IN (?)", txIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 交易状态（0-等待链确认 1-到达安全高度 2-交易失败）
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromStatus(status int) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 交易状态（0-等待链确认 1-到达安全高度 2-交易失败）
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromStatus(statuss []int) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromFee 通过fee获取内容 交易收费
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromFee(fee float64) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee = ?", fee).Find(&results).Error

	return
}

// GetBatchFromFee 批量查找 交易收费
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromFee(fees []float64) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee IN (?)", fees).Find(&results).Error

	return
}

// GetFromMethodParam 通过method_param获取内容 调用参数
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromMethodParam(methodParam string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_param = ?", methodParam).Find(&results).Error

	return
}

// GetBatchFromMethodParam 批量查找 调用参数
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromMethodParam(methodParams []string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_param IN (?)", methodParams).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromRemark(remark string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromRemark(remarks []string) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromTimeStamp 通过time_stamp获取内容 链确认时间
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromTimeStamp(timeStamp int64) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_stamp = ?", timeStamp).Find(&results).Error

	return
}

// GetBatchFromTimeStamp 批量查找 链确认时间
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromTimeStamp(timeStamps []int64) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("time_stamp IN (?)", timeStamps).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 更新时间
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 更新时间
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

// GetFromCreatadTime 通过creatad_time获取内容 创建时间
func (obj *_PolygonMarketplaceOrderTxMgr) GetFromCreatadTime(creatadTime time.Time) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creatad_time = ?", creatadTime).Find(&results).Error

	return
}

// GetBatchFromCreatadTime 批量查找 创建时间
func (obj *_PolygonMarketplaceOrderTxMgr) GetBatchFromCreatadTime(creatadTimes []time.Time) (results []*PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creatad_time IN (?)", creatadTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PolygonMarketplaceOrderTxMgr) FetchByPrimaryKey(id uint64) (result PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByTxidIndex primay or index 获取唯一内容
func (obj *_PolygonMarketplaceOrderTxMgr) FetchUniqueByTxidIndex(txID string) (result PolygonMarketplaceOrderTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tx_id = ?", txID).Find(&result).Error

	return
}
