package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _NftCollectionsMgr struct {
	*_BaseMgr
}

// NftCollectionsMgr open func
func NftCollectionsMgr(db *gorm.DB) *_NftCollectionsMgr {
	if db == nil {
		panic(fmt.Errorf("NftCollectionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NftCollectionsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("nft_collections"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NftCollectionsMgr) GetTableName() string {
	return "nft_collections"
}

// Get 获取
func (obj *_NftCollectionsMgr) Get() (result NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NftCollectionsMgr) Gets() (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一标识每个NFT合集的主键
func (obj *_NftCollectionsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithChainType chain_type获取 链类型 (1: 以太坊, 2: Polygon, ...)
func (obj *_NftCollectionsMgr) WithChainType(chainType int) Option {
	return optionFunc(func(o *options) { o.query["chain_type"] = chainType })
}

// WithCollectionName collection_name获取 NFT合集名称
func (obj *_NftCollectionsMgr) WithCollectionName(collectionName string) Option {
	return optionFunc(func(o *options) { o.query["collection_name"] = collectionName })
}

// WithNftAddress nft_address获取 NFT合集的合约地址，标准以太坊地址
func (obj *_NftCollectionsMgr) WithNftAddress(nftAddress string) Option {
	return optionFunc(func(o *options) { o.query["nft_address"] = nftAddress })
}

// WithMintStatus mint_status获取 合约的铸造状态（0-未添加到预言机订阅，1-已添加到预言机订阅）
func (obj *_NftCollectionsMgr) WithMintStatus(mintStatus int) Option {
	return optionFunc(func(o *options) { o.query["mint_status"] = mintStatus })
}

// WithCoverImageURL cover_image_url获取 NFT合集的封面图片URL地址
func (obj *_NftCollectionsMgr) WithCoverImageURL(coverImageURL string) Option {
	return optionFunc(func(o *options) { o.query["cover_image_url"] = coverImageURL })
}

// WithAvatarFrameURL avatar_frame_url获取 头像地址
func (obj *_NftCollectionsMgr) WithAvatarFrameURL(avatarFrameURL string) Option {
	return optionFunc(func(o *options) { o.query["avatar_frame_url"] = avatarFrameURL })
}

// WithMintedNftCount minted_nft_count获取 已铸造的NFT数量
func (obj *_NftCollectionsMgr) WithMintedNftCount(mintedNftCount int) Option {
	return optionFunc(func(o *options) { o.query["minted_nft_count"] = mintedNftCount })
}

// WithListedNftCount listed_nft_count获取 正在出售的NFT数量
func (obj *_NftCollectionsMgr) WithListedNftCount(listedNftCount int) Option {
	return optionFunc(func(o *options) { o.query["listed_nft_count"] = listedNftCount })
}

// WithFloorPrice floor_price获取 当前地板价，单位为美元
func (obj *_NftCollectionsMgr) WithFloorPrice(floorPrice float64) Option {
	return optionFunc(func(o *options) { o.query["floor_price"] = floorPrice })
}

// WithTotalVolume total_volume获取 总交易额，单位为美元
func (obj *_NftCollectionsMgr) WithTotalVolume(totalVolume float64) Option {
	return optionFunc(func(o *options) { o.query["total_volume"] = totalVolume })
}

// WithCreatedTime created_time获取 记录的创建时间
func (obj *_NftCollectionsMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// WithUpdatedTime updated_time获取 记录的最后更新时间
func (obj *_NftCollectionsMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// GetByOption 功能选项模式获取
func (obj *_NftCollectionsMgr) GetByOption(opts ...Option) (result NftCollections, err error) {
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
func (obj *_NftCollectionsMgr) GetByOptions(opts ...Option) (results []*NftCollections, err error) {
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

// GetFromID 通过id获取内容 唯一标识每个NFT合集的主键
func (obj *_NftCollectionsMgr) GetFromID(id int64) (result NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一标识每个NFT合集的主键
func (obj *_NftCollectionsMgr) GetBatchFromID(ids []int64) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromChainType 通过chain_type获取内容 链类型 (1: 以太坊, 2: Polygon, ...)
func (obj *_NftCollectionsMgr) GetFromChainType(chainType int) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type = ?", chainType).Find(&results).Error

	return
}

// GetBatchFromChainType 批量查找 链类型 (1: 以太坊, 2: Polygon, ...)
func (obj *_NftCollectionsMgr) GetBatchFromChainType(chainTypes []int) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type IN (?)", chainTypes).Find(&results).Error

	return
}

// GetFromCollectionName 通过collection_name获取内容 NFT合集名称
func (obj *_NftCollectionsMgr) GetFromCollectionName(collectionName string) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("collection_name = ?", collectionName).Find(&results).Error

	return
}

// GetBatchFromCollectionName 批量查找 NFT合集名称
func (obj *_NftCollectionsMgr) GetBatchFromCollectionName(collectionNames []string) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("collection_name IN (?)", collectionNames).Find(&results).Error

	return
}

// GetFromNftAddress 通过nft_address获取内容 NFT合集的合约地址，标准以太坊地址
func (obj *_NftCollectionsMgr) GetFromNftAddress(nftAddress string) (result NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_address = ?", nftAddress).Find(&result).Error

	return
}

// GetBatchFromNftAddress 批量查找 NFT合集的合约地址，标准以太坊地址
func (obj *_NftCollectionsMgr) GetBatchFromNftAddress(nftAddresss []string) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_address IN (?)", nftAddresss).Find(&results).Error

	return
}

// GetFromMintStatus 通过mint_status获取内容 合约的铸造状态（0-未添加到预言机订阅，1-已添加到预言机订阅）
func (obj *_NftCollectionsMgr) GetFromMintStatus(mintStatus int) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mint_status = ?", mintStatus).Find(&results).Error

	return
}

// GetBatchFromMintStatus 批量查找 合约的铸造状态（0-未添加到预言机订阅，1-已添加到预言机订阅）
func (obj *_NftCollectionsMgr) GetBatchFromMintStatus(mintStatuss []int) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("mint_status IN (?)", mintStatuss).Find(&results).Error

	return
}

// GetFromCoverImageURL 通过cover_image_url获取内容 NFT合集的封面图片URL地址
func (obj *_NftCollectionsMgr) GetFromCoverImageURL(coverImageURL string) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover_image_url = ?", coverImageURL).Find(&results).Error

	return
}

// GetBatchFromCoverImageURL 批量查找 NFT合集的封面图片URL地址
func (obj *_NftCollectionsMgr) GetBatchFromCoverImageURL(coverImageURLs []string) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("cover_image_url IN (?)", coverImageURLs).Find(&results).Error

	return
}

// GetFromAvatarFrameURL 通过avatar_frame_url获取内容 头像地址
func (obj *_NftCollectionsMgr) GetFromAvatarFrameURL(avatarFrameURL string) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_frame_url = ?", avatarFrameURL).Find(&results).Error

	return
}

// GetBatchFromAvatarFrameURL 批量查找 头像地址
func (obj *_NftCollectionsMgr) GetBatchFromAvatarFrameURL(avatarFrameURLs []string) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("avatar_frame_url IN (?)", avatarFrameURLs).Find(&results).Error

	return
}

// GetFromMintedNftCount 通过minted_nft_count获取内容 已铸造的NFT数量
func (obj *_NftCollectionsMgr) GetFromMintedNftCount(mintedNftCount int) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minted_nft_count = ?", mintedNftCount).Find(&results).Error

	return
}

// GetBatchFromMintedNftCount 批量查找 已铸造的NFT数量
func (obj *_NftCollectionsMgr) GetBatchFromMintedNftCount(mintedNftCounts []int) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("minted_nft_count IN (?)", mintedNftCounts).Find(&results).Error

	return
}

// GetFromListedNftCount 通过listed_nft_count获取内容 正在出售的NFT数量
func (obj *_NftCollectionsMgr) GetFromListedNftCount(listedNftCount int) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("listed_nft_count = ?", listedNftCount).Find(&results).Error

	return
}

// GetBatchFromListedNftCount 批量查找 正在出售的NFT数量
func (obj *_NftCollectionsMgr) GetBatchFromListedNftCount(listedNftCounts []int) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("listed_nft_count IN (?)", listedNftCounts).Find(&results).Error

	return
}

// GetFromFloorPrice 通过floor_price获取内容 当前地板价，单位为美元
func (obj *_NftCollectionsMgr) GetFromFloorPrice(floorPrice float64) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("floor_price = ?", floorPrice).Find(&results).Error

	return
}

// GetBatchFromFloorPrice 批量查找 当前地板价，单位为美元
func (obj *_NftCollectionsMgr) GetBatchFromFloorPrice(floorPrices []float64) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("floor_price IN (?)", floorPrices).Find(&results).Error

	return
}

// GetFromTotalVolume 通过total_volume获取内容 总交易额，单位为美元
func (obj *_NftCollectionsMgr) GetFromTotalVolume(totalVolume float64) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_volume = ?", totalVolume).Find(&results).Error

	return
}

// GetBatchFromTotalVolume 批量查找 总交易额，单位为美元
func (obj *_NftCollectionsMgr) GetBatchFromTotalVolume(totalVolumes []float64) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("total_volume IN (?)", totalVolumes).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 记录的创建时间
func (obj *_NftCollectionsMgr) GetFromCreatedTime(createdTime time.Time) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 记录的创建时间
func (obj *_NftCollectionsMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容 记录的最后更新时间
func (obj *_NftCollectionsMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找 记录的最后更新时间
func (obj *_NftCollectionsMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_NftCollectionsMgr) FetchByPrimaryKey(id int64) (result NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueByNftAddressIndex primay or index 获取唯一内容
func (obj *_NftCollectionsMgr) FetchUniqueByNftAddressIndex(nftAddress string) (result NftCollections, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_address = ?", nftAddress).Find(&result).Error

	return
}
