package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MintingNftMgr struct {
	*_BaseMgr
}

// MintingNftMgr open func
func MintingNftMgr(db *gorm.DB) *_MintingNftMgr {
	if db == nil {
		panic(fmt.Errorf("MintingNftMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MintingNftMgr{_BaseMgr: &_BaseMgr{DB: db.Table("minting_nft"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MintingNftMgr) GetTableName() string {
	return "minting_nft"
}

// Get 获取
func (obj *_MintingNftMgr) Get() (result MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MintingNftMgr) Gets() (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 铸造NFT
func (obj *_MintingNftMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithChainType chain_type获取 链类型：1-Ethereum 2-Polygons
func (obj *_MintingNftMgr) WithChainType(chainType int) Option {
	return optionFunc(func(o *options) { o.query["chain_type"] = chainType })
}

// WithHeight height获取 区块高度
func (obj *_MintingNftMgr) WithHeight(height int64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithMintAddress mint_address获取 铸币地址
func (obj *_MintingNftMgr) WithMintAddress(mintAddress string) Option {
	return optionFunc(func(o *options) { o.query["mint_address"] = mintAddress })
}

// WithNft721Address nft721_address获取 合约地址
func (obj *_MintingNftMgr) WithNft721Address(nft721Address string) Option {
	return optionFunc(func(o *options) { o.query["nft721_address"] = nft721Address })
}

// WithTxID tx_id获取 交易Hash
func (obj *_MintingNftMgr) WithTxID(txID string) Option {
	return optionFunc(func(o *options) { o.query["tx_id"] = txID })
}

// WithTokenID token_id获取 NFT市场token ID
func (obj *_MintingNftMgr) WithTokenID(tokenID int) Option {
	return optionFunc(func(o *options) { o.query["token_id"] = tokenID })
}

// WithNftName nft_name获取 铸币名称
func (obj *_MintingNftMgr) WithNftName(nftName string) Option {
	return optionFunc(func(o *options) { o.query["nft_name"] = nftName })
}

// WithMinerFee miner_fee获取 矿工费
func (obj *_MintingNftMgr) WithMinerFee(minerFee float64) Option {
	return optionFunc(func(o *options) { o.query["miner_fee"] = minerFee })
}

// WithNoticeStatus notice_status获取 通知状态(0--未通知,1--已通知)
func (obj *_MintingNftMgr) WithNoticeStatus(noticeStatus int) Option {
	return optionFunc(func(o *options) { o.query["notice_status"] = noticeStatus })
}

// WithTradeStatus trade_status获取 交易状态(0--失败，1--成功)
func (obj *_MintingNftMgr) WithTradeStatus(tradeStatus int) Option {
	return optionFunc(func(o *options) { o.query["trade_status"] = tradeStatus })
}

// WithTimestamp timestamp获取 链上确认时间
func (obj *_MintingNftMgr) WithTimestamp(timestamp int64) Option {
	return optionFunc(func(o *options) { o.query["timestamp"] = timestamp })
}

// WithCreatedTime created_time获取 创建时间
func (obj *_MintingNftMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// WithUpdatedTime updated_time获取 更新时间
func (obj *_MintingNftMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// GetByOption 功能选项模式获取
func (obj *_MintingNftMgr) GetByOption(opts ...Option) (result MintingNft, err error) {
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
func (obj *_MintingNftMgr) GetByOptions(opts ...Option) (results []*MintingNft, err error) {
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

// GetFromID 通过id获取内容 铸造NFT
func (obj *_MintingNftMgr) GetFromID(id int64) (result MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 铸造NFT
func (obj *_MintingNftMgr) GetBatchFromID(ids []int64) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromChainType 通过chain_type获取内容 链类型：1-Ethereum 2-Polygons
func (obj *_MintingNftMgr) GetFromChainType(chainType int) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type = ?", chainType).Find(&results).Error

	return
}

// GetBatchFromChainType 批量查找 链类型：1-Ethereum 2-Polygons
func (obj *_MintingNftMgr) GetBatchFromChainType(chainTypes []int) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type IN (?)", chainTypes).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容 区块高度
func (obj *_MintingNftMgr) GetFromHeight(height int64) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量查找 区块高度
func (obj *_MintingNftMgr) GetBatchFromHeight(heights []int64) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("height IN (?)", heights).Find(&results).Error

	return
}

// GetFromMintAddress 通过mint_address获取内容 铸币地址
func (obj *_MintingNftMgr) GetFromMintAddress(mintAddress string) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mint_address = ?", mintAddress).Find(&results).Error

	return
}

// GetBatchFromMintAddress 批量查找 铸币地址
func (obj *_MintingNftMgr) GetBatchFromMintAddress(mintAddresss []string) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mint_address IN (?)", mintAddresss).Find(&results).Error

	return
}

// GetFromNft721Address 通过nft721_address获取内容 合约地址
func (obj *_MintingNftMgr) GetFromNft721Address(nft721Address string) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft721_address = ?", nft721Address).Find(&results).Error

	return
}

// GetBatchFromNft721Address 批量查找 合约地址
func (obj *_MintingNftMgr) GetBatchFromNft721Address(nft721Addresss []string) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft721_address IN (?)", nft721Addresss).Find(&results).Error

	return
}

// GetFromTxID 通过tx_id获取内容 交易Hash
func (obj *_MintingNftMgr) GetFromTxID(txID string) (result MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tx_id = ?", txID).Find(&result).Error

	return
}

// GetBatchFromTxID 批量查找 交易Hash
func (obj *_MintingNftMgr) GetBatchFromTxID(txIDs []string) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tx_id IN (?)", txIDs).Find(&results).Error

	return
}

// GetFromTokenID 通过token_id获取内容 NFT市场token ID
func (obj *_MintingNftMgr) GetFromTokenID(tokenID int) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_id = ?", tokenID).Find(&results).Error

	return
}

// GetBatchFromTokenID 批量查找 NFT市场token ID
func (obj *_MintingNftMgr) GetBatchFromTokenID(tokenIDs []int) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_id IN (?)", tokenIDs).Find(&results).Error

	return
}

// GetFromNftName 通过nft_name获取内容 铸币名称
func (obj *_MintingNftMgr) GetFromNftName(nftName string) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_name = ?", nftName).Find(&results).Error

	return
}

// GetBatchFromNftName 批量查找 铸币名称
func (obj *_MintingNftMgr) GetBatchFromNftName(nftNames []string) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_name IN (?)", nftNames).Find(&results).Error

	return
}

// GetFromMinerFee 通过miner_fee获取内容 矿工费
func (obj *_MintingNftMgr) GetFromMinerFee(minerFee float64) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("miner_fee = ?", minerFee).Find(&results).Error

	return
}

// GetBatchFromMinerFee 批量查找 矿工费
func (obj *_MintingNftMgr) GetBatchFromMinerFee(minerFees []float64) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("miner_fee IN (?)", minerFees).Find(&results).Error

	return
}

// GetFromNoticeStatus 通过notice_status获取内容 通知状态(0--未通知,1--已通知)
func (obj *_MintingNftMgr) GetFromNoticeStatus(noticeStatus int) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_status = ?", noticeStatus).Find(&results).Error

	return
}

// GetBatchFromNoticeStatus 批量查找 通知状态(0--未通知,1--已通知)
func (obj *_MintingNftMgr) GetBatchFromNoticeStatus(noticeStatuss []int) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("notice_status IN (?)", noticeStatuss).Find(&results).Error

	return
}

// GetFromTradeStatus 通过trade_status获取内容 交易状态(0--失败，1--成功)
func (obj *_MintingNftMgr) GetFromTradeStatus(tradeStatus int) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("trade_status = ?", tradeStatus).Find(&results).Error

	return
}

// GetBatchFromTradeStatus 批量查找 交易状态(0--失败，1--成功)
func (obj *_MintingNftMgr) GetBatchFromTradeStatus(tradeStatuss []int) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("trade_status IN (?)", tradeStatuss).Find(&results).Error

	return
}

// GetFromTimestamp 通过timestamp获取内容 链上确认时间
func (obj *_MintingNftMgr) GetFromTimestamp(timestamp int64) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("timestamp = ?", timestamp).Find(&results).Error

	return
}

// GetBatchFromTimestamp 批量查找 链上确认时间
func (obj *_MintingNftMgr) GetBatchFromTimestamp(timestamps []int64) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("timestamp IN (?)", timestamps).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 创建时间
func (obj *_MintingNftMgr) GetFromCreatedTime(createdTime time.Time) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 创建时间
func (obj *_MintingNftMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 更新时间
func (obj *_MintingNftMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 更新时间
func (obj *_MintingNftMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_MintingNftMgr) FetchByPrimaryKey(id int64) (result MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByTxIDIndex primay or index 获取唯一内容
func (obj *_MintingNftMgr) FetchUniqueByTxIDIndex(txID string) (result MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("tx_id = ?", txID).Find(&result).Error

	return
}

// FetchIndexByTokenID  获取多个内容
func (obj *_MintingNftMgr) FetchIndexByTokenID(tokenID int) (results []*MintingNft, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_id = ?", tokenID).Find(&results).Error

	return
}
