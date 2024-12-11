package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SupportedTokensMgr struct {
	*_BaseMgr
}

// SupportedTokensMgr open func
func SupportedTokensMgr(db *gorm.DB) *_SupportedTokensMgr {
	if db == nil {
		panic(fmt.Errorf("SupportedTokensMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SupportedTokensMgr{_BaseMgr: &_BaseMgr{DB: db.Table("supported_tokens"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SupportedTokensMgr) GetTableName() string {
	return "supported_tokens"
}

// Get 获取
func (obj *_SupportedTokensMgr) Get() (result SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SupportedTokensMgr) Gets() (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一标识每个代币记录的主键
func (obj *_SupportedTokensMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTokenName token_name获取 代币名称，如 USDC, WETH, DAI 等
func (obj *_SupportedTokensMgr) WithTokenName(tokenName string) Option {
	return optionFunc(func(o *options) { o.query["token_name"] = tokenName })
}

// WithSymbol symbol获取 代币符号，如 USDC, WETH, DAI 等
func (obj *_SupportedTokensMgr) WithSymbol(symbol string) Option {
	return optionFunc(func(o *options) { o.query["symbol"] = symbol })
}

// WithContractAddress contract_address获取 代币的合约地址，标准以太坊地址格式
func (obj *_SupportedTokensMgr) WithContractAddress(contractAddress string) Option {
	return optionFunc(func(o *options) { o.query["contract_address"] = contractAddress })
}

// WithIcon icon获取 代币图标
func (obj *_SupportedTokensMgr) WithIcon(icon string) Option {
	return optionFunc(func(o *options) { o.query["icon"] = icon })
}

// WithDecimals decimals获取 代币的小数位数，通常为 18
func (obj *_SupportedTokensMgr) WithDecimals(decimals uint8) Option {
	return optionFunc(func(o *options) { o.query["decimals"] = decimals })
}

// WithChainType chain_type获取 代币所在的链类型（1: 以太坊, 2: Polygon, ...）
func (obj *_SupportedTokensMgr) WithChainType(chainType uint8) Option {
	return optionFunc(func(o *options) { o.query["chain_type"] = chainType })
}

// WithRate rate获取 代币兑换美元的汇率
func (obj *_SupportedTokensMgr) WithRate(rate float64) Option {
	return optionFunc(func(o *options) { o.query["rate"] = rate })
}

// WithIsActive is_active获取 代币是否可用，1 表示可用，0 表示不可用
func (obj *_SupportedTokensMgr) WithIsActive(isActive bool) Option {
	return optionFunc(func(o *options) { o.query["is_active"] = isActive })
}

// WithBinanceSymbol binance_symbol获取 Binance API 中用于查询价格的代币符号
func (obj *_SupportedTokensMgr) WithBinanceSymbol(binanceSymbol string) Option {
	return optionFunc(func(o *options) { o.query["binance_symbol"] = binanceSymbol })
}

// WithCreatedTime created_time获取 记录的创建时间
func (obj *_SupportedTokensMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// WithUpdatedTime updated_time获取 记录的最后更新时间
func (obj *_SupportedTokensMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// GetByOption 功能选项模式获取
func (obj *_SupportedTokensMgr) GetByOption(opts ...Option) (result SupportedTokens, err error) {
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
func (obj *_SupportedTokensMgr) GetByOptions(opts ...Option) (results []*SupportedTokens, err error) {
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

// GetFromID 通过id获取内容 唯一标识每个代币记录的主键
func (obj *_SupportedTokensMgr) GetFromID(id uint64) (result SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一标识每个代币记录的主键
func (obj *_SupportedTokensMgr) GetBatchFromID(ids []uint64) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromTokenName 通过token_name获取内容 代币名称，如 USDC, WETH, DAI 等
func (obj *_SupportedTokensMgr) GetFromTokenName(tokenName string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_name = ?", tokenName).Find(&results).Error

	return
}

// GetBatchFromTokenName 批量查找 代币名称，如 USDC, WETH, DAI 等
func (obj *_SupportedTokensMgr) GetBatchFromTokenName(tokenNames []string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_name IN (?)", tokenNames).Find(&results).Error

	return
}

// GetFromSymbol 通过symbol获取内容 代币符号，如 USDC, WETH, DAI 等
func (obj *_SupportedTokensMgr) GetFromSymbol(symbol string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol = ?", symbol).Find(&results).Error

	return
}

// GetBatchFromSymbol 批量查找 代币符号，如 USDC, WETH, DAI 等
func (obj *_SupportedTokensMgr) GetBatchFromSymbol(symbols []string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("symbol IN (?)", symbols).Find(&results).Error

	return
}

// GetFromContractAddress 通过contract_address获取内容 代币的合约地址，标准以太坊地址格式
func (obj *_SupportedTokensMgr) GetFromContractAddress(contractAddress string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_address = ?", contractAddress).Find(&results).Error

	return
}

// GetBatchFromContractAddress 批量查找 代币的合约地址，标准以太坊地址格式
func (obj *_SupportedTokensMgr) GetBatchFromContractAddress(contractAddresss []string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("contract_address IN (?)", contractAddresss).Find(&results).Error

	return
}

// GetFromIcon 通过icon获取内容 代币图标
func (obj *_SupportedTokensMgr) GetFromIcon(icon string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon = ?", icon).Find(&results).Error

	return
}

// GetBatchFromIcon 批量查找 代币图标
func (obj *_SupportedTokensMgr) GetBatchFromIcon(icons []string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("icon IN (?)", icons).Find(&results).Error

	return
}

// GetFromDecimals 通过decimals获取内容 代币的小数位数，通常为 18
func (obj *_SupportedTokensMgr) GetFromDecimals(decimals uint8) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("decimals = ?", decimals).Find(&results).Error

	return
}

// GetBatchFromDecimals 批量查找 代币的小数位数，通常为 18
func (obj *_SupportedTokensMgr) GetBatchFromDecimals(decimalss []uint8) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("decimals IN (?)", decimalss).Find(&results).Error

	return
}

// GetFromChainType 通过chain_type获取内容 代币所在的链类型（1: 以太坊, 2: Polygon, ...）
func (obj *_SupportedTokensMgr) GetFromChainType(chainType uint8) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type = ?", chainType).Find(&results).Error

	return
}

// GetBatchFromChainType 批量查找 代币所在的链类型（1: 以太坊, 2: Polygon, ...）
func (obj *_SupportedTokensMgr) GetBatchFromChainType(chainTypes []uint8) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type IN (?)", chainTypes).Find(&results).Error

	return
}

// GetFromRate 通过rate获取内容 代币兑换美元的汇率
func (obj *_SupportedTokensMgr) GetFromRate(rate float64) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate = ?", rate).Find(&results).Error

	return
}

// GetBatchFromRate 批量查找 代币兑换美元的汇率
func (obj *_SupportedTokensMgr) GetBatchFromRate(rates []float64) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("rate IN (?)", rates).Find(&results).Error

	return
}

// GetFromIsActive 通过is_active获取内容 代币是否可用，1 表示可用，0 表示不可用
func (obj *_SupportedTokensMgr) GetFromIsActive(isActive bool) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_active = ?", isActive).Find(&results).Error

	return
}

// GetBatchFromIsActive 批量查找 代币是否可用，1 表示可用，0 表示不可用
func (obj *_SupportedTokensMgr) GetBatchFromIsActive(isActives []bool) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_active IN (?)", isActives).Find(&results).Error

	return
}

// GetFromBinanceSymbol 通过binance_symbol获取内容 Binance API 中用于查询价格的代币符号
func (obj *_SupportedTokensMgr) GetFromBinanceSymbol(binanceSymbol string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("binance_symbol = ?", binanceSymbol).Find(&results).Error

	return
}

// GetBatchFromBinanceSymbol 批量查找 Binance API 中用于查询价格的代币符号
func (obj *_SupportedTokensMgr) GetBatchFromBinanceSymbol(binanceSymbols []string) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("binance_symbol IN (?)", binanceSymbols).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 记录的创建时间
func (obj *_SupportedTokensMgr) GetFromCreatedTime(createdTime time.Time) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 记录的创建时间
func (obj *_SupportedTokensMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 记录的最后更新时间
func (obj *_SupportedTokensMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 记录的最后更新时间
func (obj *_SupportedTokensMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_SupportedTokensMgr) FetchByPrimaryKey(id uint64) (result SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByUniqueToken primay or index 获取唯一内容
func (obj *_SupportedTokensMgr) FetchUniqueIndexByUniqueToken(tokenName string, contractAddress string, chainType uint8) (result SupportedTokens, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_name = ? AND contract_address = ? AND chain_type = ?", tokenName, contractAddress, chainType).Find(&result).Error

	return
}
