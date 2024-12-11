package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _PolygonsMianCurrencyTxMgr struct {
	*_BaseMgr
}

// PolygonsMianCurrencyTxMgr open func
func PolygonsMianCurrencyTxMgr(db *gorm.DB) *_PolygonsMianCurrencyTxMgr {
	if db == nil {
		panic(fmt.Errorf("PolygonsMianCurrencyTxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PolygonsMianCurrencyTxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("polygons_mian_currency_tx"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PolygonsMianCurrencyTxMgr) GetTableName() string {
	return "polygons_mian_currency_tx"
}

// Get 获取
func (obj *_PolygonsMianCurrencyTxMgr) Get() (result PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PolygonsMianCurrencyTxMgr) Gets() (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PolygonsMianCurrencyTxMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHeigth heigth获取 区块高度
func (obj *_PolygonsMianCurrencyTxMgr) WithHeigth(heigth int64) Option {
	return optionFunc(func(o *options) { o.query["heigth"] = heigth })
}

// WithType type获取 交易类型
func (obj *_PolygonsMianCurrencyTxMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithStatus status获取 交易状态:0-等待链确认 1-达到安全高度 2-交易失败
func (obj *_PolygonsMianCurrencyTxMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTxid txid获取 交易HASH
func (obj *_PolygonsMianCurrencyTxMgr) WithTxid(txid string) Option {
	return optionFunc(func(o *options) { o.query["txid"] = txid })
}

// WithFromAddress from_address获取 from地址
func (obj *_PolygonsMianCurrencyTxMgr) WithFromAddress(fromAddress string) Option {
	return optionFunc(func(o *options) { o.query["from_address"] = fromAddress })
}

// WithToAddress to_address获取 to地址
func (obj *_PolygonsMianCurrencyTxMgr) WithToAddress(toAddress string) Option {
	return optionFunc(func(o *options) { o.query["to_address"] = toAddress })
}

// WithAmount amount获取 交易金额
func (obj *_PolygonsMianCurrencyTxMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithFee fee获取 手续费
func (obj *_PolygonsMianCurrencyTxMgr) WithFee(fee float64) Option {
	return optionFunc(func(o *options) { o.query["fee"] = fee })
}

// WithRemark remark获取 备注
func (obj *_PolygonsMianCurrencyTxMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithUpdatedTime updated_time获取 更新时间
func (obj *_PolygonsMianCurrencyTxMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// WithCreatedTime created_time获取 创建时间
func (obj *_PolygonsMianCurrencyTxMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// GetByOption 功能选项模式获取
func (obj *_PolygonsMianCurrencyTxMgr) GetByOption(opts ...Option) (result PolygonsMianCurrencyTx, err error) {
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
func (obj *_PolygonsMianCurrencyTxMgr) GetByOptions(opts ...Option) (results []*PolygonsMianCurrencyTx, err error) {
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
func (obj *_PolygonsMianCurrencyTxMgr) GetFromID(id uint64) (result PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromID(ids []uint64) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromHeigth 通过heigth获取内容 区块高度
func (obj *_PolygonsMianCurrencyTxMgr) GetFromHeigth(heigth int64) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("heigth = ?", heigth).Find(&results).Error

	return
}

// GetBatchFromHeigth 批量查找 区块高度
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromHeigth(heigths []int64) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("heigth IN (?)", heigths).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 交易类型
func (obj *_PolygonsMianCurrencyTxMgr) GetFromType(_type int) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 交易类型
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromType(_types []int) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("type IN (?)", _types).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 交易状态:0-等待链确认 1-达到安全高度 2-交易失败
func (obj *_PolygonsMianCurrencyTxMgr) GetFromStatus(status int) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 交易状态:0-等待链确认 1-达到安全高度 2-交易失败
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromStatus(statuss []int) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTxid 通过txid获取内容 交易HASH
func (obj *_PolygonsMianCurrencyTxMgr) GetFromTxid(txid string) (result PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("txid = ?", txid).Find(&result).Error

	return
}

// GetBatchFromTxid 批量查找 交易HASH
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromTxid(txids []string) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("txid IN (?)", txids).Find(&results).Error

	return
}

// GetFromFromAddress 通过from_address获取内容 from地址
func (obj *_PolygonsMianCurrencyTxMgr) GetFromFromAddress(fromAddress string) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_address = ?", fromAddress).Find(&results).Error

	return
}

// GetBatchFromFromAddress 批量查找 from地址
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromFromAddress(fromAddresss []string) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_address IN (?)", fromAddresss).Find(&results).Error

	return
}

// GetFromToAddress 通过to_address获取内容 to地址
func (obj *_PolygonsMianCurrencyTxMgr) GetFromToAddress(toAddress string) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to_address = ?", toAddress).Find(&results).Error

	return
}

// GetBatchFromToAddress 批量查找 to地址
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromToAddress(toAddresss []string) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("to_address IN (?)", toAddresss).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 交易金额
func (obj *_PolygonsMianCurrencyTxMgr) GetFromAmount(amount float64) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找 交易金额
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromAmount(amounts []float64) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("amount IN (?)", amounts).Find(&results).Error

	return
}

// GetFromFee 通过fee获取内容 手续费
func (obj *_PolygonsMianCurrencyTxMgr) GetFromFee(fee float64) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee = ?", fee).Find(&results).Error

	return
}

// GetBatchFromFee 批量查找 手续费
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromFee(fees []float64) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee IN (?)", fees).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_PolygonsMianCurrencyTxMgr) GetFromRemark(remark string) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromRemark(remarks []string) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 更新时间
func (obj *_PolygonsMianCurrencyTxMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 更新时间
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 创建时间
func (obj *_PolygonsMianCurrencyTxMgr) GetFromCreatedTime(createdTime time.Time) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 创建时间
func (obj *_PolygonsMianCurrencyTxMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PolygonsMianCurrencyTxMgr) FetchByPrimaryKey(id uint64) (result PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByTxidIndex primay or index 获取唯一内容
func (obj *_PolygonsMianCurrencyTxMgr) FetchUniqueByTxidIndex(txid string) (result PolygonsMianCurrencyTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("txid = ?", txid).Find(&result).Error

	return
}
