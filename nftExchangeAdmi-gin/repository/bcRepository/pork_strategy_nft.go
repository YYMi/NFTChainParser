package bcRepository

import (
	"time"
)

// BcConfig [...]
type BcConfig struct {
	ID          uint32    `gorm:"primaryKey" json:"-"`
	Module      string    `json:"module"`                                    // 模块名称
	Name        string    `gorm:"uniqueIndex:name_value_index" json:"name"`  // 配置变量名
	Value       string    `gorm:"uniqueIndex:name_value_index" json:"value"` // 值
	Description string    `json:"description"`                               // 描述
	UpdatedTime time.Time `json:"updatedTime"`                               // 更新时间
	CreatedTime time.Time `json:"createdTime"`                               // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *BcConfig) TableName() string {
	return "bc_config"
}

// BcConfigColumns get sql column name.获取数据库列名
var BcConfigColumns = struct {
	ID          string
	Module      string
	Name        string
	Value       string
	Description string
	UpdatedTime string
	CreatedTime string
}{
	ID:          "id",
	Module:      "module",
	Name:        "name",
	Value:       "value",
	Description: "description",
	UpdatedTime: "updated_time",
	CreatedTime: "created_time",
}

// BcPolygonsBlock [...]
type BcPolygonsBlock struct {
	Height              uint64    `gorm:"primaryKey;unique" json:"-"`           // 高度
	ParentHash          string    `json:"parentHash"`                           // 父hash
	BlockHash           string    `json:"blockHash"`                            // 区块hash
	Transaction         int       `gorm:"default:0" json:"transaction"`         // 交易
	ContractTransaction int       `gorm:"default:0" json:"contractTransaction"` // 合约内部交易
	BlockTimeStamp      int64     `json:"blockTimeStamp"`                       // 出块时间
	CreatedTime         time.Time `json:"createdTime"`                          // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *BcPolygonsBlock) TableName() string {
	return "bc_polygons_block"
}

// BcPolygonsBlockColumns get sql column name.获取数据库列名
var BcPolygonsBlockColumns = struct {
	Height              string
	ParentHash          string
	BlockHash           string
	Transaction         string
	ContractTransaction string
	BlockTimeStamp      string
	CreatedTime         string
}{
	Height:              "height",
	ParentHash:          "parent_hash",
	BlockHash:           "block_hash",
	Transaction:         "transaction",
	ContractTransaction: "contract_transaction",
	BlockTimeStamp:      "block_time_stamp",
	CreatedTime:         "created_time",
}

// MintingNft 铸造NFT交易
type MintingNft struct {
	ID            int64     `gorm:"primaryKey" json:"-"`                          // 铸造NFT
	ChainType     int       `json:"chainType"`                                    // 链类型：1-Ethereum 2-Polygons
	Height        int64     `json:"height"`                                       // 区块高度
	MintAddress   string    `json:"mintAddress"`                                  // 铸币地址
	Nft721Address string    `json:"nft721Address"`                                // 合约地址
	TxID          string    `gorm:"unique" json:"txId"`                           // 交易Hash
	TokenID       int       `gorm:"index:token_id" json:"tokenId"`                // NFT市场token ID
	NftName       string    `json:"nftName"`                                      // 铸币名称
	MinerFee      float64   `json:"minerFee"`                                     // 矿工费
	NoticeStatus  int       `gorm:"default:0" json:"noticeStatus"`                // 通知状态(0--未通知,1--已通知)
	TradeStatus   int       `json:"tradeStatus"`                                  // 交易状态(0--失败，1--成功)
	Timestamp     int64     `json:"timestamp"`                                    // 链上确认时间
	CreatedTime   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"` // 创建时间
	UpdatedTime   time.Time `json:"updatedTime"`                                  // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *MintingNft) TableName() string {
	return "minting_nft"
}

// MintingNftColumns get sql column name.获取数据库列名
var MintingNftColumns = struct {
	ID            string
	ChainType     string
	Height        string
	MintAddress   string
	Nft721Address string
	TxID          string
	TokenID       string
	NftName       string
	MinerFee      string
	NoticeStatus  string
	TradeStatus   string
	Timestamp     string
	CreatedTime   string
	UpdatedTime   string
}{
	ID:            "id",
	ChainType:     "chain_type",
	Height:        "height",
	MintAddress:   "mint_address",
	Nft721Address: "nft721_address",
	TxID:          "tx_id",
	TokenID:       "token_id",
	NftName:       "nft_name",
	MinerFee:      "miner_fee",
	NoticeStatus:  "notice_status",
	TradeStatus:   "trade_status",
	Timestamp:     "timestamp",
	CreatedTime:   "created_time",
	UpdatedTime:   "updated_time",
}

// NftCollections NFT合集管理表
type NftCollections struct {
	ID             int64     `gorm:"primaryKey" json:"-"`                          // 唯一标识每个NFT合集的主键
	ChainType      int       `json:"chainType"`                                    // 链类型 (1: 以太坊, 2: Polygon, ...)
	CollectionName string    `json:"collectionName"`                               // NFT合集名称
	NftAddress     string    `gorm:"unique" json:"nftAddress"`                     // NFT合集的合约地址，标准以太坊地址
	MintStatus     int       `gorm:"default:0" json:"mintStatus"`                  // 合约的铸造状态（0-未添加到预言机订阅，1-已添加到预言机订阅）
	CoverImageURL  string    `json:"coverImageUrl"`                                // NFT合集的封面图片URL地址
	AvatarFrameURL string    `json:"avatarFrameUrl"`                               // 头像地址
	MintedNftCount int       `gorm:"default:0" json:"mintedNftCount"`              // 已铸造的NFT数量
	ListedNftCount int       `gorm:"default:0" json:"listedNftCount"`              // 正在出售的NFT数量
	FloorPrice     float64   `gorm:"default:0.00" json:"floorPrice"`               // 当前地板价，单位为美元
	TotalVolume    float64   `gorm:"default:0.00" json:"totalVolume"`              // 总交易额，单位为美元
	CreatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"` // 记录的创建时间
	UpdatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"` // 记录的最后更新时间
}

// TableName get sql table name.获取数据库表名
func (m *NftCollections) TableName() string {
	return "nft_collections"
}

// NftCollectionsColumns get sql column name.获取数据库列名
var NftCollectionsColumns = struct {
	ID             string
	ChainType      string
	CollectionName string
	NftAddress     string
	MintStatus     string
	CoverImageURL  string
	AvatarFrameURL string
	MintedNftCount string
	ListedNftCount string
	FloorPrice     string
	TotalVolume    string
	CreatedTime    string
	UpdatedTime    string
}{
	ID:             "id",
	ChainType:      "chain_type",
	CollectionName: "collection_name",
	NftAddress:     "nft_address",
	MintStatus:     "mint_status",
	CoverImageURL:  "cover_image_url",
	AvatarFrameURL: "avatar_frame_url",
	MintedNftCount: "minted_nft_count",
	ListedNftCount: "listed_nft_count",
	FloorPrice:     "floor_price",
	TotalVolume:    "total_volume",
	CreatedTime:    "created_time",
	UpdatedTime:    "updated_time",
}

// NftDetails NFT详情信息表
type NftDetails struct {
	ID                int64     `gorm:"primaryKey" json:"-"`                                                     // 唯一标识每个NFT详情的主键
	OrderID           string    `json:"orderId"`                                                                 // 关联的NFT订单表的ID
	NftName           string    `json:"nftName"`                                                                 // NFT的名称
	CreatorName       string    `json:"creatorName"`                                                             // NFT创建者的名称
	NftLikes          int       `json:"nftLikes"`                                                                // NFT的点赞数
	CreatorAvatarURL  string    `json:"creatorAvatarUrl"`                                                        // NFT创建者的头像URL
	CollectionAddress string    `gorm:"uniqueIndex:collection_addressAnToken_id_index" json:"collectionAddress"` // NFT的合集合约地址，标准以太坊地址
	TokenID           int       `gorm:"uniqueIndex:collection_addressAnToken_id_index" json:"tokenId"`           // NFT的Token ID
	CreatorAddress    string    `json:"creatorAddress"`                                                          // NFT创建者的地址，标准以太坊地址
	OwnerAddress      string    `json:"ownerAddress"`                                                            // NFT所有者的地址，标准以太坊地址
	TokenStandard     string    `json:"tokenStandard"`                                                           // NFT的代币标准（如ERC721、ERC1155）
	ChainType         int       `json:"chainType"`                                                               // NFT所在的链类型（1: 以太坊, 2: Polygon, ...）
	CreatedTime       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`                            // NFT的创建时间
	NftImageURL       string    `json:"nftImageUrl"`                                                             // NFT的图片URL
	NftMetadata       string    `json:"nftMetadata"`                                                             // NFT的元数据，存储图片元素和其他信息的JSON格式
	UpdatedTime       time.Time `json:"updatedTime"`                                                             //  更新时间
}

// TableName get sql table name.获取数据库表名
func (m *NftDetails) TableName() string {
	return "nft_details"
}

// NftDetailsColumns get sql column name.获取数据库列名
var NftDetailsColumns = struct {
	ID                string
	OrderID           string
	NftName           string
	CreatorName       string
	NftLikes          string
	CreatorAvatarURL  string
	CollectionAddress string
	TokenID           string
	CreatorAddress    string
	OwnerAddress      string
	TokenStandard     string
	ChainType         string
	CreatedTime       string
	NftImageURL       string
	NftMetadata       string
	UpdatedTime       string
}{
	ID:                "id",
	OrderID:           "order_id",
	NftName:           "nft_name",
	CreatorName:       "creator_name",
	NftLikes:          "nft_likes",
	CreatorAvatarURL:  "creator_avatar_url",
	CollectionAddress: "collection_address",
	TokenID:           "token_id",
	CreatorAddress:    "creator_address",
	OwnerAddress:      "owner_address",
	TokenStandard:     "token_standard",
	ChainType:         "chain_type",
	CreatedTime:       "created_time",
	NftImageURL:       "nft_image_url",
	NftMetadata:       "nft_metadata",
	UpdatedTime:       "updated_time",
}

// NftOrders NFT订单信息表
type NftOrders struct {
	ID                  int64     `gorm:"primaryKey" json:"-"`                          // 唯一标识每个NFT订单的主键
	OrderID             string    `gorm:"unique" json:"orderId"`                        // 订单号
	NftDetailsID        int64     `json:"nftDetailsId"`                                 // 关联的NFT详情表的ID
	SellerAddress       string    `json:"sellerAddress"`                                // 卖家的地址，标准以太坊地址
	BuyerAddress        string    `json:"buyerAddress"`                                 // 买家的地址，标准以太坊地址
	Price               float64   `json:"price"`                                        // 交易价格，单位为支付货币的最小单位（例如ETH的wei）
	Fee                 float64   `json:"fee"`                                          // 手续费
	PaymentTokenAddress string    `json:"paymentTokenAddress"`                          // 支付时使用的货币token地址
	Symbol              string    `json:"symbol"`                                       // 支付货币符号
	OrderStatus         int       `json:"orderStatus"`                                  // 订单状态，1 表示进行中，2 表示已完成，3 表示已取消
	ChainTime           int64     `json:"chainTime"`                                    // 链确认时间
	CreatedTime         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"` // 订单创建时间
	UpdatedTime         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"` // 订单的最后更新时间
}

// TableName get sql table name.获取数据库表名
func (m *NftOrders) TableName() string {
	return "nft_orders"
}

// NftOrdersColumns get sql column name.获取数据库列名
var NftOrdersColumns = struct {
	ID                  string
	OrderID             string
	NftDetailsID        string
	SellerAddress       string
	BuyerAddress        string
	Price               string
	Fee                 string
	PaymentTokenAddress string
	Symbol              string
	OrderStatus         string
	ChainTime           string
	CreatedTime         string
	UpdatedTime         string
}{
	ID:                  "id",
	OrderID:             "order_id",
	NftDetailsID:        "nft_details_id",
	SellerAddress:       "seller_address",
	BuyerAddress:        "buyer_address",
	Price:               "price",
	Fee:                 "fee",
	PaymentTokenAddress: "payment_token_address",
	Symbol:              "symbol",
	OrderStatus:         "order_status",
	ChainTime:           "chain_time",
	CreatedTime:         "created_time",
	UpdatedTime:         "updated_time",
}

// NftPreSales NFT管理预售表
type NftPreSales struct {
	ID             uint64    `gorm:"primaryKey" json:"-"`                          // 唯一标识每个预售记录的主键
	ChainType      int       `json:"chainType"`                                    // 链类型 (1: 以太坊, 2: Polygon, ...)',
	NftAddress     string    `json:"nftAddress"`                                   // NFT 合约地址，标准以太坊地址
	NftName        string    `json:"nftName"`                                      // NFT名称
	TotalSupply    int       `json:"totalSupply"`                                  // 发售的NFT总数量
	NftImageURL    string    `json:"nftImageUrl"`                                  // NFT展示图像的URL地址
	AvatarFrameURL string    `json:"avatarFrameUrl"`                               // NFT头像框的URL地址
	StartTime      int64     `json:"startTime"`                                    // 预售开始时间，使用时间戳格式（秒）
	SoldCount      int       `gorm:"default:0" json:"soldCount"`                   // 已售出的NFT数量
	IsVisible      int       `gorm:"default:1" json:"isVisible"`                   // 预售区的可见状态，0 表示不可见，1 表示可见
	IsActive       int       `gorm:"default:0" json:"isActive"`                    // 预售是否进行中，0 表示已结束，1 表示进行中
	CreatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"` // 记录的创建时间
	UpdatedTime    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"` // 记录的最后更新时间
}

// TableName get sql table name.获取数据库表名
func (m *NftPreSales) TableName() string {
	return "nft_pre_sales"
}

// NftPreSalesColumns get sql column name.获取数据库列名
var NftPreSalesColumns = struct {
	ID             string
	ChainType      string
	NftAddress     string
	NftName        string
	TotalSupply    string
	NftImageURL    string
	AvatarFrameURL string
	StartTime      string
	SoldCount      string
	IsVisible      string
	IsActive       string
	CreatedTime    string
	UpdatedTime    string
}{
	ID:             "id",
	ChainType:      "chain_type",
	NftAddress:     "nft_address",
	NftName:        "nft_name",
	TotalSupply:    "total_supply",
	NftImageURL:    "nft_image_url",
	AvatarFrameURL: "avatar_frame_url",
	StartTime:      "start_time",
	SoldCount:      "sold_count",
	IsVisible:      "is_visible",
	IsActive:       "is_active",
	CreatedTime:    "created_time",
	UpdatedTime:    "updated_time",
}

// PolygonChainlinkContractTx [...]
type PolygonChainlinkContractTx struct {
	ID          uint64    `gorm:"primaryKey" json:"-"`
	Height      int64     `json:"height"`             // 区块高度
	MethodID    string    `json:"methodId"`           // 交易方法id
	ContractAdd string    `json:"contractAdd"`        // 合约地址
	FromAddr    string    `json:"fromAddr"`           // from地址
	TxID        string    `gorm:"unique" json:"txId"` // 交易hash
	Status      int       `json:"status"`             // 交易状态（0-等待链确认 1-到达安全高度 2-交易失败）
	Fee         float64   `json:"fee"`                // 交易收费
	MethodParam string    `json:"methodParam"`        // 调用参数
	Remark      string    `json:"remark"`             // 备注
	UpdatedTime time.Time `json:"updatedTime"`        // 更新时间
	CreatadTime time.Time `json:"creatadTime"`        // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *PolygonChainlinkContractTx) TableName() string {
	return "polygon_chainlink_contract_tx"
}

// PolygonChainlinkContractTxColumns get sql column name.获取数据库列名
var PolygonChainlinkContractTxColumns = struct {
	ID          string
	Height      string
	MethodID    string
	ContractAdd string
	FromAddr    string
	TxID        string
	Status      string
	Fee         string
	MethodParam string
	Remark      string
	UpdatedTime string
	CreatadTime string
}{
	ID:          "id",
	Height:      "height",
	MethodID:    "method_id",
	ContractAdd: "contract_add",
	FromAddr:    "from_addr",
	TxID:        "tx_id",
	Status:      "status",
	Fee:         "fee",
	MethodParam: "method_param",
	Remark:      "remark",
	UpdatedTime: "updated_time",
	CreatadTime: "creatad_time",
}

// PolygonMarketplaceOrderTx NFT交易所-挂单合约交易详情
type PolygonMarketplaceOrderTx struct {
	ID          uint64    `gorm:"primaryKey" json:"-"`
	Height      int64     `json:"height"`             // 区块高度
	MethodID    string    `json:"methodId"`           // 交易方法id
	OrderID     string    `json:"orderId"`            // 交易订单号
	ContractAdd string    `json:"contractAdd"`        // 合约地址
	FromAddr    string    `json:"fromAddr"`           // from地址
	TxID        string    `gorm:"unique" json:"txId"` // 交易hash
	Status      int       `json:"status"`             // 交易状态（0-等待链确认 1-到达安全高度 2-交易失败）
	Fee         float64   `json:"fee"`                // 交易收费
	MethodParam string    `json:"methodParam"`        // 调用参数
	Remark      string    `json:"remark"`             // 备注
	TimeStamp   int64     `json:"timeStamp"`          // 链确认时间
	UpdatedTime time.Time `json:"updatedTime"`        // 更新时间
	CreatadTime time.Time `json:"creatadTime"`        // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *PolygonMarketplaceOrderTx) TableName() string {
	return "polygon_marketplace_order_tx"
}

// PolygonMarketplaceOrderTxColumns get sql column name.获取数据库列名
var PolygonMarketplaceOrderTxColumns = struct {
	ID          string
	Height      string
	MethodID    string
	OrderID     string
	ContractAdd string
	FromAddr    string
	TxID        string
	Status      string
	Fee         string
	MethodParam string
	Remark      string
	TimeStamp   string
	UpdatedTime string
	CreatadTime string
}{
	ID:          "id",
	Height:      "height",
	MethodID:    "method_id",
	OrderID:     "order_id",
	ContractAdd: "contract_add",
	FromAddr:    "from_addr",
	TxID:        "tx_id",
	Status:      "status",
	Fee:         "fee",
	MethodParam: "method_param",
	Remark:      "remark",
	TimeStamp:   "time_stamp",
	UpdatedTime: "updated_time",
	CreatadTime: "creatad_time",
}

// PolygonsAuthorizedContractTx 多边形-授权合约交易
type PolygonsAuthorizedContractTx struct {
	ID           uint64    `gorm:"primaryKey" json:"-"`
	Height       int64     `json:"height"`             // 区块高度
	MethodID     string    `json:"methodId"`           // 方法ID：,081812fc : getApproved(uint256) ,42842e0e : safeTransferFrom(address,address,uint256) ,4b0795a8 : s_lastError() ,e05ce1bb : upDateMintStartTime(uint256) ,03b4f88b : authorizedUpdateTokenURI(uint256) ,d5abeb01 : maxSupply() ,6c83bbfa : frontCoverUrl() ,6352211e : ownerOf(uint256) ,46559e2f : authorizedAddr(address) ,0ca76175 : handleOracleFulfillment(bytes32,bytes,bytes) ,09c1ba2e : subscriptionId() ,c0f69491 : updataSubscriptionId(uint64) ,e985e9c5 : isApprovedForAll(address,address) ,06fdde03 : name() ,095ea7b3 : approve(address,uint256) ,69e15404 : feeAmount() ,9a4c483d : upDateFrontCoverUrl(string) ,8d91b620 : avatarFrame() ,8da5cb5b : owner() ,14f710fe : mintNFT() ,846c67ec : requestIdByAdd(bytes32) ,01ffc9a7 : supportsInterface(bytes4) ,8e218783 : upDateAavatarFrame(string) ,12065fe0 : getBalance() ,23b872dd : transferFrom(address,address,uint256) ,9ea55bb0 : updateFeeAmount(uint256) ,70a08231 : balanceOf(address) ,c87b56dd : tokenURI(uint256) ,b88d4fde : safeTransferFr
	ContractAddr string    `json:"contractAddr"`       // 合约地址
	FromAddr     string    `json:"fromAddr"`           // from地址
	Txid         string    `gorm:"unique" json:"txid"` // 交易HASH
	Status       int       `json:"status"`             // 交易状态：0-等待链确认 1-达到安全高度 2-交易失败
	Fee          float64   `json:"fee"`                // 手续费（单位ETH）
	MethodParam  string    `json:"methodParam"`        // 调用参数
	Remark       string    `json:"remark"`             // 备注
	UpdatedTime  time.Time `json:"updatedTime"`        // 更新时间
	CreatedTime  time.Time `json:"createdTime"`        // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *PolygonsAuthorizedContractTx) TableName() string {
	return "polygons_authorized_contract_tx"
}

// PolygonsAuthorizedContractTxColumns get sql column name.获取数据库列名
var PolygonsAuthorizedContractTxColumns = struct {
	ID           string
	Height       string
	MethodID     string
	ContractAddr string
	FromAddr     string
	Txid         string
	Status       string
	Fee          string
	MethodParam  string
	Remark       string
	UpdatedTime  string
	CreatedTime  string
}{
	ID:           "id",
	Height:       "height",
	MethodID:     "method_id",
	ContractAddr: "contract_addr",
	FromAddr:     "from_addr",
	Txid:         "txid",
	Status:       "status",
	Fee:          "fee",
	MethodParam:  "method_param",
	Remark:       "remark",
	UpdatedTime:  "updated_time",
	CreatedTime:  "created_time",
}

// PolygonsContract [...]
type PolygonsContract struct {
	ID              uint64    `gorm:"primaryKey" json:"-"`
	Type            int       `json:"type"`                          // 合约类型：0-预言机订阅合约 1-721合约 2-挂单交易所合约 3-竞拍交易所合约 4-授权合约
	ContractAddress string    `gorm:"unique" json:"contractAddress"` // 合约地址
	Name            string    `json:"name"`                          // 合约名称
	OwnerAddress    string    `json:"ownerAddress"`                  // 合约所有者地址
	Status          int       `gorm:"default:0" json:"status"`       // 状态：0-停用 1-启用
	PropID          int64     `json:"propId"`                        // 合约绑定的道具ID：-1 没有道具
	CreatedTime     time.Time `json:"createdTime"`                   // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *PolygonsContract) TableName() string {
	return "polygons_contract"
}

// PolygonsContractColumns get sql column name.获取数据库列名
var PolygonsContractColumns = struct {
	ID              string
	Type            string
	ContractAddress string
	Name            string
	OwnerAddress    string
	Status          string
	PropID          string
	CreatedTime     string
}{
	ID:              "id",
	Type:            "type",
	ContractAddress: "contract_address",
	Name:            "name",
	OwnerAddress:    "owner_address",
	Status:          "status",
	PropID:          "prop_id",
	CreatedTime:     "created_time",
}

// PolygonsErc721ContractTx 多边形-erc721合约交易
type PolygonsErc721ContractTx struct {
	ID           uint64    `gorm:"primaryKey" json:"-"`
	Height       int64     `json:"height"`             // 区块高度
	MethodID     string    `json:"methodId"`           // 方法ID：,081812fc : getApproved(uint256) ,42842e0e : safeTransferFrom(address,address,uint256) ,4b0795a8 : s_lastError() ,e05ce1bb : upDateMintStartTime(uint256) ,03b4f88b : authorizedUpdateTokenURI(uint256) ,d5abeb01 : maxSupply() ,6c83bbfa : frontCoverUrl() ,6352211e : ownerOf(uint256) ,46559e2f : authorizedAddr(address) ,0ca76175 : handleOracleFulfillment(bytes32,bytes,bytes) ,09c1ba2e : subscriptionId() ,c0f69491 : updataSubscriptionId(uint64) ,e985e9c5 : isApprovedForAll(address,address) ,06fdde03 : name() ,095ea7b3 : approve(address,uint256) ,69e15404 : feeAmount() ,9a4c483d : upDateFrontCoverUrl(string) ,8d91b620 : avatarFrame() ,8da5cb5b : owner() ,14f710fe : mintNFT() ,846c67ec : requestIdByAdd(bytes32) ,01ffc9a7 : supportsInterface(bytes4) ,8e218783 : upDateAavatarFrame(string) ,12065fe0 : getBalance() ,23b872dd : transferFrom(address,address,uint256) ,9ea55bb0 : updateFeeAmount(uint256) ,70a08231 : balanceOf(address) ,c87b56dd : tokenURI(uint256) ,b88d4fde : safeTransferFr
	ContractAddr string    `json:"contractAddr"`       // 合约地址
	FromAddr     string    `json:"fromAddr"`           // from地址
	Txid         string    `gorm:"unique" json:"txid"` // 交易HASH
	Status       int       `json:"status"`             // 交易状态：0-等待链确认 1-达到安全高度 2-交易失败
	Fee          float64   `json:"fee"`                // 手续费（单位ETH）
	MethodParam  string    `json:"methodParam"`        // 调用参数
	Remark       string    `json:"remark"`             // 备注
	UpdatedTime  time.Time `json:"updatedTime"`        // 更新时间
	CreatedTime  time.Time `json:"createdTime"`        // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *PolygonsErc721ContractTx) TableName() string {
	return "polygons_erc721_contract_tx"
}

// PolygonsErc721ContractTxColumns get sql column name.获取数据库列名
var PolygonsErc721ContractTxColumns = struct {
	ID           string
	Height       string
	MethodID     string
	ContractAddr string
	FromAddr     string
	Txid         string
	Status       string
	Fee          string
	MethodParam  string
	Remark       string
	UpdatedTime  string
	CreatedTime  string
}{
	ID:           "id",
	Height:       "height",
	MethodID:     "method_id",
	ContractAddr: "contract_addr",
	FromAddr:     "from_addr",
	Txid:         "txid",
	Status:       "status",
	Fee:          "fee",
	MethodParam:  "method_param",
	Remark:       "remark",
	UpdatedTime:  "updated_time",
	CreatedTime:  "created_time",
}

// PolygonsMianCurrencyTx 多边形主币交易
type PolygonsMianCurrencyTx struct {
	ID          uint64    `gorm:"primaryKey" json:"-"`
	Heigth      int64     `json:"heigth"`             // 区块高度
	Type        int       `json:"type"`               // 交易类型
	Status      int       `json:"status"`             // 交易状态:0-等待链确认 1-达到安全高度 2-交易失败
	Txid        string    `gorm:"unique" json:"txid"` // 交易HASH
	FromAddress string    `json:"fromAddress"`        // from地址
	ToAddress   string    `json:"toAddress"`          // to地址
	Amount      float64   `json:"amount"`             // 交易金额
	Fee         float64   `json:"fee"`                // 手续费
	Remark      string    `json:"remark"`             // 备注
	UpdatedTime time.Time `json:"updatedTime"`        // 更新时间
	CreatedTime time.Time `json:"createdTime"`        // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *PolygonsMianCurrencyTx) TableName() string {
	return "polygons_mian_currency_tx"
}

// PolygonsMianCurrencyTxColumns get sql column name.获取数据库列名
var PolygonsMianCurrencyTxColumns = struct {
	ID          string
	Heigth      string
	Type        string
	Status      string
	Txid        string
	FromAddress string
	ToAddress   string
	Amount      string
	Fee         string
	Remark      string
	UpdatedTime string
	CreatedTime string
}{
	ID:          "id",
	Heigth:      "heigth",
	Type:        "type",
	Status:      "status",
	Txid:        "txid",
	FromAddress: "from_address",
	ToAddress:   "to_address",
	Amount:      "amount",
	Fee:         "fee",
	Remark:      "remark",
	UpdatedTime: "updated_time",
	CreatedTime: "created_time",
}

// SupportedTokens 支持的NFT交易所代币表
type SupportedTokens struct {
	ID              uint64    `gorm:"primaryKey" json:"-"`                             // 唯一标识每个代币记录的主键
	TokenName       string    `gorm:"uniqueIndex:unique_token" json:"tokenName"`       // 代币名称，如 USDC, WETH, DAI 等
	Symbol          string    `json:"symbol"`                                          // 代币符号，如 USDC, WETH, DAI 等
	ContractAddress string    `gorm:"uniqueIndex:unique_token" json:"contractAddress"` // 代币的合约地址，标准以太坊地址格式
	Icon            string    `json:"icon"`                                            // 代币图标
	Decimals        uint8     `gorm:"default:18" json:"decimals"`                      // 代币的小数位数，通常为 18
	ChainType       uint8     `gorm:"uniqueIndex:unique_token" json:"chainType"`       // 代币所在的链类型（1: 以太坊, 2: Polygon, ...）
	Rate            float64   `gorm:"default:1.00000000" json:"rate"`                  // 代币兑换美元的汇率
	IsActive        bool      `gorm:"default:1" json:"isActive"`                       // 代币是否可用，1 表示可用，0 表示不可用
	BinanceSymbol   string    `json:"binanceSymbol"`                                   // Binance API 中用于查询价格的代币符号
	CreatedTime     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdTime"`    // 记录的创建时间
	UpdatedTime     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedTime"`    // 记录的最后更新时间
}

// TableName get sql table name.获取数据库表名
func (m *SupportedTokens) TableName() string {
	return "supported_tokens"
}

// SupportedTokensColumns get sql column name.获取数据库列名
var SupportedTokensColumns = struct {
	ID              string
	TokenName       string
	Symbol          string
	ContractAddress string
	Icon            string
	Decimals        string
	ChainType       string
	Rate            string
	IsActive        string
	BinanceSymbol   string
	CreatedTime     string
	UpdatedTime     string
}{
	ID:              "id",
	TokenName:       "token_name",
	Symbol:          "symbol",
	ContractAddress: "contract_address",
	Icon:            "icon",
	Decimals:        "decimals",
	ChainType:       "chain_type",
	Rate:            "rate",
	IsActive:        "is_active",
	BinanceSymbol:   "binance_symbol",
	CreatedTime:     "created_time",
	UpdatedTime:     "updated_time",
}
