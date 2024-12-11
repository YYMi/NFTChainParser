package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _NftOrdersMgr struct {
	*_BaseMgr
}

// NftOrdersMgr open func
func NftOrdersMgr(db *gorm.DB) *_NftOrdersMgr {
	if db == nil {
		panic(fmt.Errorf("NftOrdersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NftOrdersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("nft_orders"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NftOrdersMgr) GetTableName() string {
	return "nft_orders"
}

// Get 获取
func (obj *_NftOrdersMgr) Get() (result NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NftOrdersMgr) Gets() (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一标识每个NFT订单的主键
func (obj *_NftOrdersMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取 订单号
func (obj *_NftOrdersMgr) WithOrderID(orderID string) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithNftDetailsID nft_details_id获取 关联的NFT详情表的ID
func (obj *_NftOrdersMgr) WithNftDetailsID(nftDetailsID int64) Option {
	return optionFunc(func(o *options) { o.query["nft_details_id"] = nftDetailsID })
}

// WithSellerAddress seller_address获取 卖家的地址，标准以太坊地址
func (obj *_NftOrdersMgr) WithSellerAddress(sellerAddress string) Option {
	return optionFunc(func(o *options) { o.query["seller_address"] = sellerAddress })
}

// WithBuyerAddress buyer_address获取 买家的地址，标准以太坊地址
func (obj *_NftOrdersMgr) WithBuyerAddress(buyerAddress string) Option {
	return optionFunc(func(o *options) { o.query["buyer_address"] = buyerAddress })
}

// WithPrice price获取 交易价格，单位为支付货币的最小单位（例如ETH的wei）
func (obj *_NftOrdersMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithFee fee获取 手续费
func (obj *_NftOrdersMgr) WithFee(fee float64) Option {
	return optionFunc(func(o *options) { o.query["fee"] = fee })
}

// WithPaymentTokenAddress payment_token_address获取 支付时使用的货币token地址
func (obj *_NftOrdersMgr) WithPaymentTokenAddress(paymentTokenAddress string) Option {
	return optionFunc(func(o *options) { o.query["payment_token_address"] = paymentTokenAddress })
}

// WithSymbol symbol获取 支付货币符号
func (obj *_NftOrdersMgr) WithSymbol(symbol string) Option {
	return optionFunc(func(o *options) { o.query["symbol"] = symbol })
}

// WithOrderStatus order_status获取 订单状态，1 表示进行中，2 表示已完成，3 表示已取消
func (obj *_NftOrdersMgr) WithOrderStatus(orderStatus int) Option {
	return optionFunc(func(o *options) { o.query["order_status"] = orderStatus })
}

// WithChainTime chain_time获取 链确认时间
func (obj *_NftOrdersMgr) WithChainTime(chainTime int64) Option {
	return optionFunc(func(o *options) { o.query["chain_time"] = chainTime })
}

// WithCreatedTime created_time获取 订单创建时间
func (obj *_NftOrdersMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// WithUpdatedTime updated_time获取 订单的最后更新时间
func (obj *_NftOrdersMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// GetByOption 功能选项模式获取
func (obj *_NftOrdersMgr) GetByOption(opts ...Option) (result NftOrders, err error) {
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
func (obj *_NftOrdersMgr) GetByOptions(opts ...Option) (results []*NftOrders, err error) {
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

// GetFromID 通过id获取内容 唯一标识每个NFT订单的主键
func (obj *_NftOrdersMgr) GetFromID(id int64) (result NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一标识每个NFT订单的主键
func (obj *_NftOrdersMgr) GetBatchFromID(ids []int64) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 订单号
func (obj *_NftOrdersMgr) GetFromOrderID(orderID string) (result NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_id = ?", orderID).Find(&result).Error

	return
}

// GetBatchFromOrderID 批量查找 订单号
func (obj *_NftOrdersMgr) GetBatchFromOrderID(orderIDs []string) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_id IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromNftDetailsID 通过nft_details_id获取内容 关联的NFT详情表的ID
func (obj *_NftOrdersMgr) GetFromNftDetailsID(nftDetailsID int64) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_details_id = ?", nftDetailsID).Find(&results).Error

	return
}

// GetBatchFromNftDetailsID 批量查找 关联的NFT详情表的ID
func (obj *_NftOrdersMgr) GetBatchFromNftDetailsID(nftDetailsIDs []int64) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_details_id IN (?)", nftDetailsIDs).Find(&results).Error

	return
}

// GetFromSellerAddress 通过seller_address获取内容 卖家的地址，标准以太坊地址
func (obj *_NftOrdersMgr) GetFromSellerAddress(sellerAddress string) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("seller_address = ?", sellerAddress).Find(&results).Error

	return
}

// GetBatchFromSellerAddress 批量查找 卖家的地址，标准以太坊地址
func (obj *_NftOrdersMgr) GetBatchFromSellerAddress(sellerAddresss []string) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("seller_address IN (?)", sellerAddresss).Find(&results).Error

	return
}

// GetFromBuyerAddress 通过buyer_address获取内容 买家的地址，标准以太坊地址
func (obj *_NftOrdersMgr) GetFromBuyerAddress(buyerAddress string) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("buyer_address = ?", buyerAddress).Find(&results).Error

	return
}

// GetBatchFromBuyerAddress 批量查找 买家的地址，标准以太坊地址
func (obj *_NftOrdersMgr) GetBatchFromBuyerAddress(buyerAddresss []string) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("buyer_address IN (?)", buyerAddresss).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 交易价格，单位为支付货币的最小单位（例如ETH的wei）
func (obj *_NftOrdersMgr) GetFromPrice(price float64) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 交易价格，单位为支付货币的最小单位（例如ETH的wei）
func (obj *_NftOrdersMgr) GetBatchFromPrice(prices []float64) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("price IN (?)", prices).Find(&results).Error

	return
}

// GetFromFee 通过fee获取内容 手续费
func (obj *_NftOrdersMgr) GetFromFee(fee float64) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee = ?", fee).Find(&results).Error

	return
}

// GetBatchFromFee 批量查找 手续费
func (obj *_NftOrdersMgr) GetBatchFromFee(fees []float64) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("fee IN (?)", fees).Find(&results).Error

	return
}

// GetFromPaymentTokenAddress 通过payment_token_address获取内容 支付时使用的货币token地址
func (obj *_NftOrdersMgr) GetFromPaymentTokenAddress(paymentTokenAddress string) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment_token_address = ?", paymentTokenAddress).Find(&results).Error

	return
}

// GetBatchFromPaymentTokenAddress 批量查找 支付时使用的货币token地址
func (obj *_NftOrdersMgr) GetBatchFromPaymentTokenAddress(paymentTokenAddresss []string) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("payment_token_address IN (?)", paymentTokenAddresss).Find(&results).Error

	return
}

// GetFromSymbol 通过symbol获取内容 支付货币符号
func (obj *_NftOrdersMgr) GetFromSymbol(symbol string) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol = ?", symbol).Find(&results).Error

	return
}

// GetBatchFromSymbol 批量查找 支付货币符号
func (obj *_NftOrdersMgr) GetBatchFromSymbol(symbols []string) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol IN (?)", symbols).Find(&results).Error

	return
}

// GetFromOrderStatus 通过order_status获取内容 订单状态，1 表示进行中，2 表示已完成，3 表示已取消
func (obj *_NftOrdersMgr) GetFromOrderStatus(orderStatus int) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_status = ?", orderStatus).Find(&results).Error

	return
}

// GetBatchFromOrderStatus 批量查找 订单状态，1 表示进行中，2 表示已完成，3 表示已取消
func (obj *_NftOrdersMgr) GetBatchFromOrderStatus(orderStatuss []int) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_status IN (?)", orderStatuss).Find(&results).Error

	return
}

// GetFromChainTime 通过chain_time获取内容 链确认时间
func (obj *_NftOrdersMgr) GetFromChainTime(chainTime int64) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_time = ?", chainTime).Find(&results).Error

	return
}

// GetBatchFromChainTime 批量查找 链确认时间
func (obj *_NftOrdersMgr) GetBatchFromChainTime(chainTimes []int64) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_time IN (?)", chainTimes).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 订单创建时间
func (obj *_NftOrdersMgr) GetFromCreatedTime(createdTime time.Time) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 订单创建时间
func (obj *_NftOrdersMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 订单的最后更新时间
func (obj *_NftOrdersMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 订单的最后更新时间
func (obj *_NftOrdersMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_NftOrdersMgr) FetchByPrimaryKey(id int64) (result NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByOrderIDIndex primay or index 获取唯一内容
func (obj *_NftOrdersMgr) FetchUniqueByOrderIDIndex(orderID string) (result NftOrders, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_id = ?", orderID).Find(&result).Error

	return
}
