package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _NftPreSalesMgr struct {
	*_BaseMgr
}

// NftPreSalesMgr open func
func NftPreSalesMgr(db *gorm.DB) *_NftPreSalesMgr {
	if db == nil {
		panic(fmt.Errorf("NftPreSalesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NftPreSalesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("nft_pre_sales"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NftPreSalesMgr) GetTableName() string {
	return "nft_pre_sales"
}

// Get 获取
func (obj *_NftPreSalesMgr) Get() (result NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NftPreSalesMgr) Gets() (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一标识每个预售记录的主键
func (obj *_NftPreSalesMgr) WithID(id uint64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithChainType chain_type获取 链类型 (1: 以太坊, 2: Polygon, ...)',
func (obj *_NftPreSalesMgr) WithChainType(chainType int) Option {
	return optionFunc(func(o *options) { o.query["chain_type"] = chainType })
}

// WithNftAddress nft_address获取 NFT 合约地址，标准以太坊地址
func (obj *_NftPreSalesMgr) WithNftAddress(nftAddress string) Option {
	return optionFunc(func(o *options) { o.query["nft_address"] = nftAddress })
}

// WithNftName nft_name获取 NFT名称
func (obj *_NftPreSalesMgr) WithNftName(nftName string) Option {
	return optionFunc(func(o *options) { o.query["nft_name"] = nftName })
}

// WithTotalSupply total_supply获取 发售的NFT总数量
func (obj *_NftPreSalesMgr) WithTotalSupply(totalSupply int) Option {
	return optionFunc(func(o *options) { o.query["total_supply"] = totalSupply })
}

// WithNftImageURL nft_image_url获取 NFT展示图像的URL地址
func (obj *_NftPreSalesMgr) WithNftImageURL(nftImageURL string) Option {
	return optionFunc(func(o *options) { o.query["nft_image_url"] = nftImageURL })
}

// WithAvatarFrameURL avatar_frame_url获取 NFT头像框的URL地址
func (obj *_NftPreSalesMgr) WithAvatarFrameURL(avatarFrameURL string) Option {
	return optionFunc(func(o *options) { o.query["avatar_frame_url"] = avatarFrameURL })
}

// WithStartTime start_time获取 预售开始时间，使用时间戳格式（秒）
func (obj *_NftPreSalesMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithSoldCount sold_count获取 已售出的NFT数量
func (obj *_NftPreSalesMgr) WithSoldCount(soldCount int) Option {
	return optionFunc(func(o *options) { o.query["sold_count"] = soldCount })
}

// WithIsVisible is_visible获取 预售区的可见状态，0 表示不可见，1 表示可见
func (obj *_NftPreSalesMgr) WithIsVisible(isVisible int) Option {
	return optionFunc(func(o *options) { o.query["is_visible"] = isVisible })
}

// WithIsActive is_active获取 预售是否进行中，0 表示已结束，1 表示进行中
func (obj *_NftPreSalesMgr) WithIsActive(isActive int) Option {
	return optionFunc(func(o *options) { o.query["is_active"] = isActive })
}

// WithCreatedTime created_time获取 记录的创建时间
func (obj *_NftPreSalesMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// WithUpdatedTime updated_time获取 记录的最后更新时间
func (obj *_NftPreSalesMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// GetByOption 功能选项模式获取
func (obj *_NftPreSalesMgr) GetByOption(opts ...Option) (result NftPreSales, err error) {
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
func (obj *_NftPreSalesMgr) GetByOptions(opts ...Option) (results []*NftPreSales, err error) {
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

// GetFromID 通过id获取内容 唯一标识每个预售记录的主键
func (obj *_NftPreSalesMgr) GetFromID(id uint64) (result NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一标识每个预售记录的主键
func (obj *_NftPreSalesMgr) GetBatchFromID(ids []uint64) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromChainType 通过chain_type获取内容 链类型 (1: 以太坊, 2: Polygon, ...)',
func (obj *_NftPreSalesMgr) GetFromChainType(chainType int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type = ?", chainType).Find(&results).Error

	return
}

// GetBatchFromChainType 批量查找 链类型 (1: 以太坊, 2: Polygon, ...)',
func (obj *_NftPreSalesMgr) GetBatchFromChainType(chainTypes []int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type IN (?)", chainTypes).Find(&results).Error

	return
}

// GetFromNftAddress 通过nft_address获取内容 NFT 合约地址，标准以太坊地址
func (obj *_NftPreSalesMgr) GetFromNftAddress(nftAddress string) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_address = ?", nftAddress).Find(&results).Error

	return
}

// GetBatchFromNftAddress 批量查找 NFT 合约地址，标准以太坊地址
func (obj *_NftPreSalesMgr) GetBatchFromNftAddress(nftAddresss []string) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_address IN (?)", nftAddresss).Find(&results).Error

	return
}

// GetFromNftName 通过nft_name获取内容 NFT名称
func (obj *_NftPreSalesMgr) GetFromNftName(nftName string) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_name = ?", nftName).Find(&results).Error

	return
}

// GetBatchFromNftName 批量查找 NFT名称
func (obj *_NftPreSalesMgr) GetBatchFromNftName(nftNames []string) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_name IN (?)", nftNames).Find(&results).Error

	return
}

// GetFromTotalSupply 通过total_supply获取内容 发售的NFT总数量
func (obj *_NftPreSalesMgr) GetFromTotalSupply(totalSupply int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_supply = ?", totalSupply).Find(&results).Error

	return
}

// GetBatchFromTotalSupply 批量查找 发售的NFT总数量
func (obj *_NftPreSalesMgr) GetBatchFromTotalSupply(totalSupplys []int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_supply IN (?)", totalSupplys).Find(&results).Error

	return
}

// GetFromNftImageURL 通过nft_image_url获取内容 NFT展示图像的URL地址
func (obj *_NftPreSalesMgr) GetFromNftImageURL(nftImageURL string) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_image_url = ?", nftImageURL).Find(&results).Error

	return
}

// GetBatchFromNftImageURL 批量查找 NFT展示图像的URL地址
func (obj *_NftPreSalesMgr) GetBatchFromNftImageURL(nftImageURLs []string) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_image_url IN (?)", nftImageURLs).Find(&results).Error

	return
}

// GetFromAvatarFrameURL 通过avatar_frame_url获取内容 NFT头像框的URL地址
func (obj *_NftPreSalesMgr) GetFromAvatarFrameURL(avatarFrameURL string) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_frame_url = ?", avatarFrameURL).Find(&results).Error

	return
}

// GetBatchFromAvatarFrameURL 批量查找 NFT头像框的URL地址
func (obj *_NftPreSalesMgr) GetBatchFromAvatarFrameURL(avatarFrameURLs []string) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_frame_url IN (?)", avatarFrameURLs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 预售开始时间，使用时间戳格式（秒）
func (obj *_NftPreSalesMgr) GetFromStartTime(startTime int64) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_time = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 预售开始时间，使用时间戳格式（秒）
func (obj *_NftPreSalesMgr) GetBatchFromStartTime(startTimes []int64) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("start_time IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromSoldCount 通过sold_count获取内容 已售出的NFT数量
func (obj *_NftPreSalesMgr) GetFromSoldCount(soldCount int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sold_count = ?", soldCount).Find(&results).Error

	return
}

// GetBatchFromSoldCount 批量查找 已售出的NFT数量
func (obj *_NftPreSalesMgr) GetBatchFromSoldCount(soldCounts []int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("sold_count IN (?)", soldCounts).Find(&results).Error

	return
}

// GetFromIsVisible 通过is_visible获取内容 预售区的可见状态，0 表示不可见，1 表示可见
func (obj *_NftPreSalesMgr) GetFromIsVisible(isVisible int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_visible = ?", isVisible).Find(&results).Error

	return
}

// GetBatchFromIsVisible 批量查找 预售区的可见状态，0 表示不可见，1 表示可见
func (obj *_NftPreSalesMgr) GetBatchFromIsVisible(isVisibles []int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_visible IN (?)", isVisibles).Find(&results).Error

	return
}

// GetFromIsActive 通过is_active获取内容 预售是否进行中，0 表示已结束，1 表示进行中
func (obj *_NftPreSalesMgr) GetFromIsActive(isActive int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_active = ?", isActive).Find(&results).Error

	return
}

// GetBatchFromIsActive 批量查找 预售是否进行中，0 表示已结束，1 表示进行中
func (obj *_NftPreSalesMgr) GetBatchFromIsActive(isActives []int) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("is_active IN (?)", isActives).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 记录的创建时间
func (obj *_NftPreSalesMgr) GetFromCreatedTime(createdTime time.Time) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 记录的创建时间
func (obj *_NftPreSalesMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 记录的最后更新时间
func (obj *_NftPreSalesMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 记录的最后更新时间
func (obj *_NftPreSalesMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_NftPreSalesMgr) FetchByPrimaryKey(id uint64) (result NftPreSales, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}
