package bcRepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _NftDetailsMgr struct {
	*_BaseMgr
}

// NftDetailsMgr open func
func NftDetailsMgr(db *gorm.DB) *_NftDetailsMgr {
	if db == nil {
		panic(fmt.Errorf("NftDetailsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NftDetailsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("nft_details"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NftDetailsMgr) GetTableName() string {
	return "nft_details"
}

// Get 获取
func (obj *_NftDetailsMgr) Get() (result NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NftDetailsMgr) Gets() (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Find(&results).Error

	return
}

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 唯一标识每个NFT详情的主键
func (obj *_NftDetailsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取 关联的NFT订单表的ID
func (obj *_NftDetailsMgr) WithOrderID(orderID string) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithNftName nft_name获取 NFT的名称
func (obj *_NftDetailsMgr) WithNftName(nftName string) Option {
	return optionFunc(func(o *options) { o.query["nft_name"] = nftName })
}

// WithCreatorName creator_name获取 NFT创建者的名称
func (obj *_NftDetailsMgr) WithCreatorName(creatorName string) Option {
	return optionFunc(func(o *options) { o.query["creator_name"] = creatorName })
}

// WithNftLikes nft_likes获取 NFT的点赞数
func (obj *_NftDetailsMgr) WithNftLikes(nftLikes int) Option {
	return optionFunc(func(o *options) { o.query["nft_likes"] = nftLikes })
}

// WithCreatorAvatarURL creator_avatar_url获取 NFT创建者的头像URL
func (obj *_NftDetailsMgr) WithCreatorAvatarURL(creatorAvatarURL string) Option {
	return optionFunc(func(o *options) { o.query["creator_avatar_url"] = creatorAvatarURL })
}

// WithCollectionAddress collection_address获取 NFT的合集合约地址，标准以太坊地址
func (obj *_NftDetailsMgr) WithCollectionAddress(collectionAddress string) Option {
	return optionFunc(func(o *options) { o.query["collection_address"] = collectionAddress })
}

// WithTokenID token_id获取 NFT的Token ID
func (obj *_NftDetailsMgr) WithTokenID(tokenID int) Option {
	return optionFunc(func(o *options) { o.query["token_id"] = tokenID })
}

// WithCreatorAddress creator_address获取 NFT创建者的地址，标准以太坊地址
func (obj *_NftDetailsMgr) WithCreatorAddress(creatorAddress string) Option {
	return optionFunc(func(o *options) { o.query["creator_address"] = creatorAddress })
}

// WithOwnerAddress owner_address获取 NFT所有者的地址，标准以太坊地址
func (obj *_NftDetailsMgr) WithOwnerAddress(ownerAddress string) Option {
	return optionFunc(func(o *options) { o.query["owner_address"] = ownerAddress })
}

// WithTokenStandard token_standard获取 NFT的代币标准（如ERC721、ERC1155）
func (obj *_NftDetailsMgr) WithTokenStandard(tokenStandard string) Option {
	return optionFunc(func(o *options) { o.query["token_standard"] = tokenStandard })
}

// WithChainType chain_type获取 NFT所在的链类型（1: 以太坊, 2: Polygon, ...）
func (obj *_NftDetailsMgr) WithChainType(chainType int) Option {
	return optionFunc(func(o *options) { o.query["chain_type"] = chainType })
}

// WithCreatedTime created_time获取 NFT的创建时间
func (obj *_NftDetailsMgr) WithCreatedTime(createdTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["created_time"] = createdTime })
}

// WithNftImageURL nft_image_url获取 NFT的图片URL
func (obj *_NftDetailsMgr) WithNftImageURL(nftImageURL string) Option {
	return optionFunc(func(o *options) { o.query["nft_image_url"] = nftImageURL })
}

// WithNftMetadata nft_metadata获取 NFT的元数据，存储图片元素和其他信息的JSON格式
func (obj *_NftDetailsMgr) WithNftMetadata(nftMetadata string) Option {
	return optionFunc(func(o *options) { o.query["nft_metadata"] = nftMetadata })
}

// WithUpdatedTime updated_time获取  更新时间
func (obj *_NftDetailsMgr) WithUpdatedTime(updatedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated_time"] = updatedTime })
}

// GetByOption 功能选项模式获取
func (obj *_NftDetailsMgr) GetByOption(opts ...Option) (result NftDetails, err error) {
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
func (obj *_NftDetailsMgr) GetByOptions(opts ...Option) (results []*NftDetails, err error) {
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

// GetFromID 通过id获取内容 唯一标识每个NFT详情的主键
func (obj *_NftDetailsMgr) GetFromID(id int64) (result NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 唯一标识每个NFT详情的主键
func (obj *_NftDetailsMgr) GetBatchFromID(ids []int64) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 关联的NFT订单表的ID
func (obj *_NftDetailsMgr) GetFromOrderID(orderID string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_id = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 关联的NFT订单表的ID
func (obj *_NftDetailsMgr) GetBatchFromOrderID(orderIDs []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("order_id IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromNftName 通过nft_name获取内容 NFT的名称
func (obj *_NftDetailsMgr) GetFromNftName(nftName string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_name = ?", nftName).Find(&results).Error

	return
}

// GetBatchFromNftName 批量查找 NFT的名称
func (obj *_NftDetailsMgr) GetBatchFromNftName(nftNames []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_name IN (?)", nftNames).Find(&results).Error

	return
}

// GetFromCreatorName 通过creator_name获取内容 NFT创建者的名称
func (obj *_NftDetailsMgr) GetFromCreatorName(creatorName string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creator_name = ?", creatorName).Find(&results).Error

	return
}

// GetBatchFromCreatorName 批量查找 NFT创建者的名称
func (obj *_NftDetailsMgr) GetBatchFromCreatorName(creatorNames []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creator_name IN (?)", creatorNames).Find(&results).Error

	return
}

// GetFromNftLikes 通过nft_likes获取内容 NFT的点赞数
func (obj *_NftDetailsMgr) GetFromNftLikes(nftLikes int) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_likes = ?", nftLikes).Find(&results).Error

	return
}

// GetBatchFromNftLikes 批量查找 NFT的点赞数
func (obj *_NftDetailsMgr) GetBatchFromNftLikes(nftLikess []int) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_likes IN (?)", nftLikess).Find(&results).Error

	return
}

// GetFromCreatorAvatarURL 通过creator_avatar_url获取内容 NFT创建者的头像URL
func (obj *_NftDetailsMgr) GetFromCreatorAvatarURL(creatorAvatarURL string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creator_avatar_url = ?", creatorAvatarURL).Find(&results).Error

	return
}

// GetBatchFromCreatorAvatarURL 批量查找 NFT创建者的头像URL
func (obj *_NftDetailsMgr) GetBatchFromCreatorAvatarURL(creatorAvatarURLs []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creator_avatar_url IN (?)", creatorAvatarURLs).Find(&results).Error

	return
}

// GetFromCollectionAddress 通过collection_address获取内容 NFT的合集合约地址，标准以太坊地址
func (obj *_NftDetailsMgr) GetFromCollectionAddress(collectionAddress string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("collection_address = ?", collectionAddress).Find(&results).Error

	return
}

// GetBatchFromCollectionAddress 批量查找 NFT的合集合约地址，标准以太坊地址
func (obj *_NftDetailsMgr) GetBatchFromCollectionAddress(collectionAddresss []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("collection_address IN (?)", collectionAddresss).Find(&results).Error

	return
}

// GetFromTokenID 通过token_id获取内容 NFT的Token ID
func (obj *_NftDetailsMgr) GetFromTokenID(tokenID int) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_id = ?", tokenID).Find(&results).Error

	return
}

// GetBatchFromTokenID 批量查找 NFT的Token ID
func (obj *_NftDetailsMgr) GetBatchFromTokenID(tokenIDs []int) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_id IN (?)", tokenIDs).Find(&results).Error

	return
}

// GetFromCreatorAddress 通过creator_address获取内容 NFT创建者的地址，标准以太坊地址
func (obj *_NftDetailsMgr) GetFromCreatorAddress(creatorAddress string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creator_address = ?", creatorAddress).Find(&results).Error

	return
}

// GetBatchFromCreatorAddress 批量查找 NFT创建者的地址，标准以太坊地址
func (obj *_NftDetailsMgr) GetBatchFromCreatorAddress(creatorAddresss []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("creator_address IN (?)", creatorAddresss).Find(&results).Error

	return
}

// GetFromOwnerAddress 通过owner_address获取内容 NFT所有者的地址，标准以太坊地址
func (obj *_NftDetailsMgr) GetFromOwnerAddress(ownerAddress string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("owner_address = ?", ownerAddress).Find(&results).Error

	return
}

// GetBatchFromOwnerAddress 批量查找 NFT所有者的地址，标准以太坊地址
func (obj *_NftDetailsMgr) GetBatchFromOwnerAddress(ownerAddresss []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("owner_address IN (?)", ownerAddresss).Find(&results).Error

	return
}

// GetFromTokenStandard 通过token_standard获取内容 NFT的代币标准（如ERC721、ERC1155）
func (obj *_NftDetailsMgr) GetFromTokenStandard(tokenStandard string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_standard = ?", tokenStandard).Find(&results).Error

	return
}

// GetBatchFromTokenStandard 批量查找 NFT的代币标准（如ERC721、ERC1155）
func (obj *_NftDetailsMgr) GetBatchFromTokenStandard(tokenStandards []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("token_standard IN (?)", tokenStandards).Find(&results).Error

	return
}

// GetFromChainType 通过chain_type获取内容 NFT所在的链类型（1: 以太坊, 2: Polygon, ...）
func (obj *_NftDetailsMgr) GetFromChainType(chainType int) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type = ?", chainType).Find(&results).Error

	return
}

// GetBatchFromChainType 批量查找 NFT所在的链类型（1: 以太坊, 2: Polygon, ...）
func (obj *_NftDetailsMgr) GetBatchFromChainType(chainTypes []int) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("chain_type IN (?)", chainTypes).Find(&results).Error

	return
}

// GetFromCreatedTime 通过created_time获取内容 NFT的创建时间
func (obj *_NftDetailsMgr) GetFromCreatedTime(createdTime time.Time) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time = ?", createdTime).Find(&results).Error

	return
}

// GetBatchFromCreatedTime 批量查找 NFT的创建时间
func (obj *_NftDetailsMgr) GetBatchFromCreatedTime(createdTimes []time.Time) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("created_time IN (?)", createdTimes).Find(&results).Error

	return
}

// GetFromNftImageURL 通过nft_image_url获取内容 NFT的图片URL
func (obj *_NftDetailsMgr) GetFromNftImageURL(nftImageURL string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_image_url = ?", nftImageURL).Find(&results).Error

	return
}

// GetBatchFromNftImageURL 批量查找 NFT的图片URL
func (obj *_NftDetailsMgr) GetBatchFromNftImageURL(nftImageURLs []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_image_url IN (?)", nftImageURLs).Find(&results).Error

	return
}

// GetFromNftMetadata 通过nft_metadata获取内容 NFT的元数据，存储图片元素和其他信息的JSON格式
func (obj *_NftDetailsMgr) GetFromNftMetadata(nftMetadata string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_metadata = ?", nftMetadata).Find(&results).Error

	return
}

// GetBatchFromNftMetadata 批量查找 NFT的元数据，存储图片元素和其他信息的JSON格式
func (obj *_NftDetailsMgr) GetBatchFromNftMetadata(nftMetadatas []string) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("nft_metadata IN (?)", nftMetadatas).Find(&results).Error

	return
}

// GetFromUpdatedTime 通过updated_time获取内容  更新时间
func (obj *_NftDetailsMgr) GetFromUpdatedTime(updatedTime time.Time) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time = ?", updatedTime).Find(&results).Error

	return
}

// GetBatchFromUpdatedTime 批量查找  更新时间
func (obj *_NftDetailsMgr) GetBatchFromUpdatedTime(updatedTimes []time.Time) (results []*NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("updated_time IN (?)", updatedTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primay or index 获取唯一内容
func (obj *_NftDetailsMgr) FetchByPrimaryKey(id int64) (result NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("id = ?", id).Find(&result).Error

	return
}

// FetchUniqueIndexByCollectionAddressAnTokenIDIndex primay or index 获取唯一内容
func (obj *_NftDetailsMgr) FetchUniqueIndexByCollectionAddressAnTokenIDIndex(collectionAddress string, tokenID int) (result NftDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Table(obj.GetTableName()).Where("collection_address = ? AND token_id = ?", collectionAddress, tokenID).Find(&result).Error

	return
}
