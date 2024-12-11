package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _PolygonsErc721ContractTxMgr struct {
	*_BaseMgr
}

// PolygonsErc721ContractTxMgr open func
func PolygonsErc721ContractTxMgr(db *gorm.DB) *_PolygonsErc721ContractTxMgr {
	if db == nil {
		panic(fmt.Errorf("PolygonsErc721ContractTxMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PolygonsErc721ContractTxMgr{_BaseMgr: &_BaseMgr{DB: db.Table("polygons_erc721_contract_tx"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PolygonsErc721ContractTxMgr) GetTableName() string {
	return "polygons_erc721_contract_tx"
}

// Get 获取
func (obj *_PolygonsErc721ContractTxMgr) Get() (result PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PolygonsErc721ContractTxMgr) Gets() (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_PolygonsErc721ContractTxMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHeight height获取 区块高度
func (obj *_PolygonsErc721ContractTxMgr) WithHeight(height int64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithMethodID method_id获取 方法ID：
// 081812fc : getApproved(uint256)
// 42842e0e : safeTransferFrom(address,address,uint256)
// 4b0795a8 : s_lastError()
// e05ce1bb : upDateMintStartTime(uint256)
// 03b4f88b : authorizedUpdateTokenURI(uint256)
// d5abeb01 : maxSupply()
// 6c83bbfa : frontCoverUrl()
// 6352211e : ownerOf(uint256)
// 46559e2f : authorizedAddr(address)
// 0ca76175 : handleOracleFulfillment(bytes32,bytes,bytes)
// 09c1ba2e : subscriptionId()
// c0f69491 : updataSubscriptionId(uint64)
// e985e9c5 : isApprovedForAll(address,address)
// 06fdde03 : name()
// 095ea7b3 : approve(address,uint256)
// 69e15404 : feeAmount()
// 9a4c483d : upDateFrontCoverUrl(string)
// 8d91b620 : avatarFrame()
// 8da5cb5b : owner()
// 14f710fe : mintNFT()
// 846c67ec : requestIdByAdd(bytes32)
// 01ffc9a7 : supportsInterface(bytes4)
// 8e218783 : upDateAavatarFrame(string)
// 12065fe0 : getBalance()
// 23b872dd : transferFrom(address,address,uint256)
// 9ea55bb0 : updateFeeAmount(uint256)
// 70a08231 : balanceOf(address)
// c87b56dd : tokenURI(uint256)
// b88d4fde : safeTransferFr
func (obj *_PolygonsErc721ContractTxMgr) WithMethodID(methodID string) Option {
	return optionFunc(func(o *options) { o.query["method_id"] = methodID })
}

// WithContractAddr contract_addr获取 合约地址
func (obj *_PolygonsErc721ContractTxMgr) WithContractAddr(contractAddr string) Option {
	return optionFunc(func(o *options) { o.query["contract_addr"] = contractAddr })
}

// WithFromAddr from_addr获取 from地址
func (obj *_PolygonsErc721ContractTxMgr) WithFromAddr(fromAddr string) Option {
	return optionFunc(func(o *options) { o.query["from_addr"] = fromAddr })
}

// WithTxid txid获取 交易HASH
func (obj *_PolygonsErc721ContractTxMgr) WithTxid(txid string) Option {
	return optionFunc(func(o *options) { o.query["txid"] = txid })
}

// WithStatus status获取 交易状态：0-等待链确认 1-达到安全高度 2-交易失败
func (obj *_PolygonsErc721ContractTxMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithFee fee获取 手续费（单位ETH）
func (obj *_PolygonsErc721ContractTxMgr) WithFee(fee float64) Option {
	return optionFunc(func(o *options) { o.query["fee"] = fee })
}

// WithMethodParam method_param获取 调用参数
func (obj *_PolygonsErc721ContractTxMgr) WithMethodParam(methodParam string) Option {
	return optionFunc(func(o *options) { o.query["method_param"] = methodParam })
}

// WithRemark remark获取 备注
func (obj *_PolygonsErc721ContractTxMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithUpdatedTime updated_time获取 更新时间
func (obj *_PolygonsErc721ContractTxMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// WithCreatedTime created_time获取 创建时间
func (obj *_PolygonsErc721ContractTxMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// GetByOption 功能选项模式获取
func (obj *_PolygonsErc721ContractTxMgr) GetByOption(opts ...Option) (result PolygonsErc721ContractTx, err error) {
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
func (obj *_PolygonsErc721ContractTxMgr) GetByOptions(opts ...Option) (results []*PolygonsErc721ContractTx, err error) {
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
func (obj *_PolygonsErc721ContractTxMgr) GetFromID(id uint64) (result PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromID(ids []uint64) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容 区块高度
func (obj *_PolygonsErc721ContractTxMgr) GetFromHeight(height int64) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量查找 区块高度
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromHeight(heights []int64) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height IN (?)", heights).Find(&results).Error

	return
}

// GetFromMethodID 通过method_id获取内容 方法ID：
// 081812fc : getApproved(uint256)
// 42842e0e : safeTransferFrom(address,address,uint256)
// 4b0795a8 : s_lastError()
// e05ce1bb : upDateMintStartTime(uint256)
// 03b4f88b : authorizedUpdateTokenURI(uint256)
// d5abeb01 : maxSupply()
// 6c83bbfa : frontCoverUrl()
// 6352211e : ownerOf(uint256)
// 46559e2f : authorizedAddr(address)
// 0ca76175 : handleOracleFulfillment(bytes32,bytes,bytes)
// 09c1ba2e : subscriptionId()
// c0f69491 : updataSubscriptionId(uint64)
// e985e9c5 : isApprovedForAll(address,address)
// 06fdde03 : name()
// 095ea7b3 : approve(address,uint256)
// 69e15404 : feeAmount()
// 9a4c483d : upDateFrontCoverUrl(string)
// 8d91b620 : avatarFrame()
// 8da5cb5b : owner()
// 14f710fe : mintNFT()
// 846c67ec : requestIdByAdd(bytes32)
// 01ffc9a7 : supportsInterface(bytes4)
// 8e218783 : upDateAavatarFrame(string)
// 12065fe0 : getBalance()
// 23b872dd : transferFrom(address,address,uint256)
// 9ea55bb0 : updateFeeAmount(uint256)
// 70a08231 : balanceOf(address)
// c87b56dd : tokenURI(uint256)
// b88d4fde : safeTransferFr
func (obj *_PolygonsErc721ContractTxMgr) GetFromMethodID(methodID string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_id = ?", methodID).Find(&results).Error

	return
}

// GetBatchFromMethodID 批量查找 方法ID：
// 081812fc : getApproved(uint256)
// 42842e0e : safeTransferFrom(address,address,uint256)
// 4b0795a8 : s_lastError()
// e05ce1bb : upDateMintStartTime(uint256)
// 03b4f88b : authorizedUpdateTokenURI(uint256)
// d5abeb01 : maxSupply()
// 6c83bbfa : frontCoverUrl()
// 6352211e : ownerOf(uint256)
// 46559e2f : authorizedAddr(address)
// 0ca76175 : handleOracleFulfillment(bytes32,bytes,bytes)
// 09c1ba2e : subscriptionId()
// c0f69491 : updataSubscriptionId(uint64)
// e985e9c5 : isApprovedForAll(address,address)
// 06fdde03 : name()
// 095ea7b3 : approve(address,uint256)
// 69e15404 : feeAmount()
// 9a4c483d : upDateFrontCoverUrl(string)
// 8d91b620 : avatarFrame()
// 8da5cb5b : owner()
// 14f710fe : mintNFT()
// 846c67ec : requestIdByAdd(bytes32)
// 01ffc9a7 : supportsInterface(bytes4)
// 8e218783 : upDateAavatarFrame(string)
// 12065fe0 : getBalance()
// 23b872dd : transferFrom(address,address,uint256)
// 9ea55bb0 : updateFeeAmount(uint256)
// 70a08231 : balanceOf(address)
// c87b56dd : tokenURI(uint256)
// b88d4fde : safeTransferFr
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromMethodID(methodIDs []string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_id IN (?)", methodIDs).Find(&results).Error

	return
}

// GetFromContractAddr 通过contract_addr获取内容 合约地址
func (obj *_PolygonsErc721ContractTxMgr) GetFromContractAddr(contractAddr string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_addr = ?", contractAddr).Find(&results).Error

	return
}

// GetBatchFromContractAddr 批量查找 合约地址
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromContractAddr(contractAddrs []string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_addr IN (?)", contractAddrs).Find(&results).Error

	return
}

// GetFromFromAddr 通过from_addr获取内容 from地址
func (obj *_PolygonsErc721ContractTxMgr) GetFromFromAddr(fromAddr string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_addr = ?", fromAddr).Find(&results).Error

	return
}

// GetBatchFromFromAddr 批量查找 from地址
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromFromAddr(fromAddrs []string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("from_addr IN (?)", fromAddrs).Find(&results).Error

	return
}

// GetFromTxid 通过txid获取内容 交易HASH
func (obj *_PolygonsErc721ContractTxMgr) GetFromTxid(txid string) (result PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("txid = ?", txid).Find(&result).Error

	return
}

// GetBatchFromTxid 批量查找 交易HASH
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromTxid(txids []string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("txid IN (?)", txids).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 交易状态：0-等待链确认 1-达到安全高度 2-交易失败
func (obj *_PolygonsErc721ContractTxMgr) GetFromStatus(status int) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 交易状态：0-等待链确认 1-达到安全高度 2-交易失败
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromStatus(statuss []int) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("status IN (?)", statuss).Find(&results).Error

	return
}

// GetFromFee 通过fee获取内容 手续费（单位ETH）
func (obj *_PolygonsErc721ContractTxMgr) GetFromFee(fee float64) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee = ?", fee).Find(&results).Error

	return
}

// GetBatchFromFee 批量查找 手续费（单位ETH）
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromFee(fees []float64) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee IN (?)", fees).Find(&results).Error

	return
}

// GetFromMethodParam 通过method_param获取内容 调用参数
func (obj *_PolygonsErc721ContractTxMgr) GetFromMethodParam(methodParam string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_param = ?", methodParam).Find(&results).Error

	return
}

// GetBatchFromMethodParam 批量查找 调用参数
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromMethodParam(methodParams []string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("method_param IN (?)", methodParams).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_PolygonsErc721ContractTxMgr) GetFromRemark(remark string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromRemark(remarks []string) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("remark IN (?)", remarks).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 更新时间
func (obj *_PolygonsErc721ContractTxMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 更新时间
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 创建时间
func (obj *_PolygonsErc721ContractTxMgr) GetFromCreatedTime(createdTime time.Time) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 创建时间
func (obj *_PolygonsErc721ContractTxMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_PolygonsErc721ContractTxMgr) FetchByPrimaryKey(id uint64) (result PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByTxidIndex primay or index 获取唯一内容
func (obj *_PolygonsErc721ContractTxMgr) FetchUniqueByTxidIndex(txid string) (result PolygonsErc721ContractTx, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("txid = ?", txid).Find(&result).Error

	return
}
