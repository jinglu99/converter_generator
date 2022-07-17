package data


type Objective struct {
	Population               *string                    `json:"Population,omitempty"`
	Targets                  []int64                    `json:"Targets,omitempty"`
	Profits                  []*ProjectProfit           `json:"Profits,omitempty"`
	ROI                      *ROI                       `json:"ROI,omitempty"`
	ParentBudgetDepartmentID *string                    `json:"ParentBudgetDepartmentID,omitempty"`
	ROIs                     []*ROI                     `json:"ROIs,omitempty"`
	MainROIType              *ROIType                   `json:"MainROIType,omitempty"`
	LocalLifeBusinessTargets []*LocalLifeBusinessTarget `json:"LocalLifeBusinessTargets,omitempty"`
}

type LocalLifeBusinessTarget struct {
	TargetID    *int64   `json:"TargetID,omitempty"`
	TargetVlue  *float64 `json:"TargetVlue,omitempty"`
	CustomValue *string  `json:"CustomValue,omitempty"`
}

type ROI struct {
	Value                 *string                `json:"Value,omitempty"`
	Desc                  *string                `json:"Desc,omitempty"`
	ROIType               *ROIType               `json:"ROIType,omitempty"`
	ShortTermROI          *ShortTermROI          `json:"ShortTermROI,omitempty"`
	BusinessInvestmentROI *BusinessInvestmentROI `json:"BusinessInvestmentROI,omitempty"`
	LongTermROI           *LongTermROI           `json:"LongTermROI,omitempty"`
	OtherROI              *OtherROI              `json:"OtherROI,omitempty"`
}

type OtherROI struct {
	Desc        *string  `json:"Desc,omitempty"`
	ROIValue    *string  `json:"ROIValue,omitempty"`
	CostValue   *float64 `json:"CostValue,omitempty"`
	ReturnValue *float64 `json:"ReturnValue,omitempty"`
}

type LongTermROI struct {
	ReturnFormulaIndexs map[string]*ReturnFormulaIndex `json:"ReturnFormulaIndexs,omitempty"`
	Desc                *string                        `json:"Desc,omitempty"`
	Roi                 *string                        `json:"Roi,omitempty"`
	I                   *int64                         `json:"I,omitempty"`
}

type BusinessInvestmentROI struct {
	ReturnFormulaIndexs map[string]*ReturnFormulaIndex `json:"ReturnFormulaIndexs,omitempty"`
	Desc                *string                        `json:"Desc,omitempty"`
	Roi                 *string                        `json:"Roi,omitempty"`
	I                   *int64                         `json:"I,omitempty"`
}

type ReturnFormulaIndex struct {
	Key   *string  `json:"Key,omitempty"`
	Value *float64 `json:"Value,omitempty"`
	Desc  *string  `json:"Desc,omitempty"`
}

type ShortTermROI struct {
	ReturnFormulaIndexs map[string]*ReturnFormulaIndex `json:"ReturnFormulaIndexs,omitempty"`
	Desc                *string                        `json:"Desc,omitempty"`
	Roi                 *string                        `json:"Roi,omitempty"`
	I                   *int64                         `json:"I,omitempty"`
}

type ProjectProfit struct {
	ProfitType         string              `json:"ProfitType,omitempty"`
	Value              string              `json:"Value,omitempty"`
	ProfitTplType      *ProfitTplType      `json:"ProfitTplType,omitempty"`
	StandardizedProfit *StandardizedProfit `json:"StandardizedProfit,omitempty"`
	Target             *int64              `json:"Target,omitempty"`
}

type StandardizedProfit struct {
	TargetType             *int64                   `json:"TargetType,omitempty"`
	ProfitIndex            *int64                   `json:"ProfitIndex,omitempty"`
	IncrementValue         *float64                 `json:"IncrementValue,omitempty"`
	VerticalTypes          []int64                  `json:"VerticalTypes,omitempty"`
	ExpectRate             *float64                 `json:"ExpectRate,omitempty"`
	CrowdID                *string                  `json:"CrowdID,omitempty"`
	CrowdEstimatedNum      *float64                 `json:"CrowdEstimatedNum,omitempty"`
	CrowdConfigJSON        *string                  `json:"CrowdConfigJson,omitempty"`
	CrowdType              *int32                   `json:"CrowdType,omitempty"`
	TargetCustomName       *string                  `json:"TargetCustomName,omitempty"`
	ProfitIndexName        *string                  `json:"ProfitIndexName,omitempty"`
	CrowdUpdateTime        *int32                   `json:"CrowdUpdateTime,omitempty"`
	GameIDs                []string                 `json:"GameIDs,omitempty"`
	CustomValue            *string                  `json:"CustomValue,omitempty"`
	TopicIDs               []string                 `json:"TopicIDs,omitempty"`
	MulitProfits           []*MulitProfits          `json:"MulitProfits,omitempty"`
	AdditionalRemarks      *string                  `json:"AdditionalRemarks,omitempty"`
	ContributeLevelIncomes []*ContributeLevelIncome `json:"ContributeLevelIncomes,omitempty"`
	ActivityHotSpots       []string                 `json:"ActivityHotSpots,omitempty"`
	GameMulitProfits       []*GameMulitProfits      `json:"GameMulitProfits,omitempty"`
	ProcessingMessage      *ProcessingMessage       `json:"ProcessingMessage,omitempty"`
	MusicMulitProfits      []*MusicMulitProfit      `json:"MusicMulitProfits,omitempty"`
	MovieProfit            *MovieProfit             `json:"MovieProfit,omitempty"`
	ActivityHotSpotIDs     []string                 `json:"ActivityHotSpotIDs,omitempty"`
	MusicProfitInfos       []*MusicProfitInfo       `json:"MusicProfitInfos,omitempty"`
	AbilityName            *string                  `json:"AbilityName,omitempty"`
	LibraURL               *string                  `json:"LibraURL,omitempty"`
	ChannelInfoList        []*ChannelInfo           `json:"ChannelInfoList,omitempty"`
}

type ChannelInfo struct {
	ChannelID   *int32  `json:"ChannelID,omitempty"`
	ChannelName *string `json:"ChannelName,omitempty"`
}

type MusicProfitInfo struct {
	MusicCopyrightType *string  `json:"MusicCopyrightType,omitempty"`
	Name               *string  `json:"Name,omitempty"`
	GroupID            *string  `json:"GroupID,omitempty"`
	ClipIDs            []string `json:"ClipIDs,omitempty"`
}

type MovieProfit struct {
	MovieIDs       []string `json:"MovieIDs,omitempty"`
	PlatformIncome *float64 `json:"PlatformIncome,omitempty"`
	MovieNames     []string `json:"MovieNames,omitempty"`
}

type MusicMulitProfit struct {
	MusicCopyrightType *string  `json:"MusicCopyrightType,omitempty"`
	IncrementValue     *float64 `json:"IncrementValue,omitempty"`
	GroupID            *string  `json:"GroupID,omitempty"`
	ClipIDs            []string `json:"ClipIDs,omitempty"`
	CopyrightIncome    *float64 `json:"CopyrightIncome,omitempty"`
}

type ProcessingMessage struct {
	GameID          *string   `json:"GameID,omitempty"`
	GameName        *string   `json:"GameName,omitempty"`
	LTV             *float64  `json:"LTV,omitempty"`
	GameType        *GameType `json:"GameType,omitempty"`
	GameRolePercent *float64  `json:"GameRolePercent,omitempty"`
}

type GameMulitProfits struct {
	GameID       *string         `json:"GameID,omitempty"`
	GameName     *string         `json:"GameName,omitempty"`
	MulitProfits []*MulitProfits `json:"MulitProfits,omitempty"`
	LTV          *float64        `json:"LTV,omitempty"`
	GameType     *GameType       `json:"GameType,omitempty"`
}

type MulitProfits struct {
	ProfitIndex      *int64   `json:"ProfitIndex,omitempty"`
	IncrementValue   *float64 `json:"IncrementValue,omitempty"`
	CustomValue      *string  `json:"CustomValue,omitempty"`
	ExpectRate       *float64 `json:"ExpectRate,omitempty"`
	ProfitIndexName  *string  `json:"ProfitIndexName,omitempty"`
	TargetCustomName *string  `json:"TargetCustomName,omitempty"`
}

type ContributeLevelIncome struct {
	MarkFansLevel  int64   `json:"MarkFansLevel"`
	IncrementValue float64 `json:"IncrementValue"`
}

//type MarkFansLevel int64
//
//const (
//	MarkFansLevel_Normal          MarkFansLevel = 1
//	MarkFansLevel_TenThonsand     MarkFansLevel = 2
//	MarkFansLevel_HundredThousand MarkFansLevel = 3
//	MarkFansLevel_Millon          MarkFansLevel = 4
//)
//
//type TargetType int64
//
//const (
//	TargetType_Live                        TargetType = 1
//	TargetType_Watch                       TargetType = 2
//	TargetType_Revenue                     TargetType = 3
//	TargetType_PayUV                       TargetType = 4
//	TargetType_Live10KFans                 TargetType = 5
//	TargetType_Follow                      TargetType = 6
//	TargetType_EcommerceGMV                TargetType = 7
//	TargetType_EcommerceSaleUV             TargetType = 8
//	TargetType_Ecommerce                   TargetType = 9
//	TargetType_Interaction                 TargetType = 10
//	TargetType_CommercialSale              TargetType = 11
//	TargetType_AuthorSupply                TargetType = 12
//	TargetType_ContentConsumption          TargetType = 13
//	TargetType_NicheScenarioCognition      TargetType = 14
//	TargetType_LiveEarnings                TargetType = 15
//	TargetType_AuthorExperience            TargetType = 16
//	TargetType_CommercialCash              TargetType = 17
//	TargetType_GameJointOperations         TargetType = 18
//	TargetType_InnerExposureIncome         TargetType = 19
//	TargetType_HotSpotIncome               TargetType = 20
//	TargetType_CreatorIncome               TargetType = 21
//	TargetType_OuterInfluenceIncome        TargetType = 22
//	TargetType_PoliticMediaIncome          TargetType = 23
//	TargetType_ContributeIncome            TargetType = 24
//	TargetType_LongTermGameJointOperations TargetType = 25
//	TargetType_UserActiveIncome            TargetType = 26
//	TargetType_MusicCopyrightIncome        TargetType = 27
//	TargetType_MovieTicketIncome           TargetType = 28
//	TargetType_FunctionalPenetrationIncome TargetType = 29
//	TargetType_BrandPerceptionIncome       TargetType = 30
//	TargetType_Others                      TargetType = 255
//)

type ROIType int64

const (
	ROIType_Customized            ROIType = 0
	ROIType_ShortTermROI          ROIType = 1
	ROIType_BusinessInvestmentROI ROIType = 2
	ROIType_LongTermROI           ROIType = 3
	ROIType_OtherROI              ROIType = 4
)

type ProfitTplType int64

const (
	ProfitTplType_Customized   ProfitTplType = 0
	ProfitTplType_Standardized ProfitTplType = 1
)

type GameType int64

const (
	GameType_BringGame    GameType = 1
	GameType_LiveRoomGame GameType = 2
)

//
//type ProfitIndex int64
//
//const (
//	// 提开播
//	ProfitIndex_TotalLiveDurationIncrement ProfitIndex = 1
//	ProfitIndex_PerLiveDurationIncrement   ProfitIndex = 2
//	ProfitIndex_LiveUVIncrement            ProfitIndex = 3
//	// 提看播
//	ProfitIndex_TotalWatchDurationForAnchorIncrement ProfitIndex = 11
//	ProfitIndex_TotalWatchDurationIncrement          ProfitIndex = 12
//	ProfitIndex_PerWatchDurationForAnchorIncrement   ProfitIndex = 13
//	ProfitIndex_PerWatchDurationIncrement            ProfitIndex = 14
//	ProfitIndex_WatchUVForAnchorIncrement            ProfitIndex = 15
//	ProfitIndex_WatchUVIncrement                     ProfitIndex = 16
//	// 提营收
//	ProfitIndex_TotalRevenueForAnchorIncrement ProfitIndex = 21
//	ProfitIndex_TotalRevenueIncrement          ProfitIndex = 22
//	// 提关注
//	ProfitIndex_TotalFollowerForAnchorIncrement ProfitIndex = 31
//	ProfitIndex_TotalFollowerIncrement          ProfitIndex = 32
//	// 提互动
//	ProfitIndex_TotalPairEnterForAnchorIncrement ProfitIndex = 41
//	ProfitIndex_TotalLinkUVForAnchorIncrement    ProfitIndex = 42
//	// 提付费UV
//	ProfitIndex_TotalPayUVForAnchorIncrement ProfitIndex = 51
//	ProfitIndex_TotalPayUVIncrement          ProfitIndex = 52
//	// 提电商GMV
//	ProfitIndex_TotalGMVForAnchorIncrement ProfitIndex = 61
//	ProfitIndex_TotalGMVIncrement          ProfitIndex = 62
//	// 提动销UV
//	ProfitIndex_TotalEcommerceSaleUVIncrement ProfitIndex = 71
//	// 提游戏联运
//	ProfitIndex_TotalGameJointOperationsRevenueForAnchorIncrement ProfitIndex = 81
//	ProfitIndex_TotalActivityGameIncrement                        ProfitIndex = 82
//	ProfitIndex_TotalInnerAdvertisementIncrement                  ProfitIndex = 83
//	ProfitIndex_TotalNewActivationUserNumber                      ProfitIndex = 84
//	// 站内曝光收益
//	ProfitIndex_ActivityTopicVVIncrement ProfitIndex = 91
//	// 热点收益
//	ProfitIndex_HotSpotVV     ProfitIndex = 101
//	ProfitIndex_HotSpotNumber ProfitIndex = 102
//	// 创作者收益
//	ProfitIndex_AuthorAddFans ProfitIndex = 103
//	// 站外影响力收益
//	ProfitIndex_RpNumber    ProfitIndex = 111
//	ProfitIndex_BaiduMark   ProfitIndex = 112
//	ProfitIndex_XinlangMark ProfitIndex = 113
//	ProfitIndex_WeiXinMark  ProfitIndex = 114
//	ProfitIndex_Other       ProfitIndex = 119
//	// 投稿收益
//	ProfitIndex_ContributeIncome ProfitIndex = 120
//	// 拉新拉活
//	ProfitIndex_TotalNewExternalUserNumber ProfitIndex = 130
//	ProfitIndex_TotalNewActiveUserNumber   ProfitIndex = 131
//	// 端内版权
//	ProfitIndex_TotalMusicCVCopyrightIncome ProfitIndex = 140
//	ProfitIndex_TotalMovieTicketIncome      ProfitIndex = 141
//	// 功能渗透-指标项
//	ProfitIndex_FunctionalPenetrationAARatioIncrement    ProfitIndex = 150
//	ProfitIndex_FunctionalPenetrationABRatioIncrement    ProfitIndex = 151
//	ProfitIndex_FunctionalPenetrationOtherRatioIncrement ProfitIndex = 152
//	// 品牌感知-指标项
//	ProfitIndex_OnSiteActivityPerceptionRatioIncrement ProfitIndex = 160
//	ProfitIndex_BrandReputationIncrement               ProfitIndex = 161
//	ProfitIndex_BrandPerceptionOtherIncrement          ProfitIndex = 162
//	ProfitIndex_Others                                 ProfitIndex = 255
//)
//
//type VerticalType int64
//
//const (
//	VerticalType_Reading              VerticalType = 1
//	VerticalType_PanHealth            VerticalType = 2
//	VerticalType_FamilyEducation      VerticalType = 3
//	VerticalType_Examination          VerticalType = 4
//	VerticalType_MediaPerson          VerticalType = 5
//	VerticalType_BeautyIndustry       VerticalType = 6
//	VerticalType_EmotionalCounseling  VerticalType = 7
//	VerticalType_Handicraft           VerticalType = 8
//	VerticalType_PaintingCalligraphy  VerticalType = 9
//	VerticalType_PregnantTreatment    VerticalType = 10
//	VerticalType_Treatment            VerticalType = 11
//	VerticalType_MusicEducation       VerticalType = 12
//	VerticalType_SportsFitness        VerticalType = 13
//	VerticalType_LanguageTraining     VerticalType = 14
//	VerticalType_NurseryEducation     VerticalType = 15
//	VerticalType_VocationalEducation  VerticalType = 16
//	VerticalType_ChatRoom             VerticalType = 17
//	VerticalType_KTV                  VerticalType = 18
//	VerticalType_Game                 VerticalType = 19
//	VerticalType_Star                 VerticalType = 20
//	VerticalType_Composer             VerticalType = 21
//	VerticalType_LiveKnowledgePayment VerticalType = 22
//	VerticalType_PayConsultation      VerticalType = 23
//	VerticalType_TravelOutdoor        VerticalType = 24
//	VerticalType_Car                  VerticalType = 25
//	VerticalType_HumanityArt          VerticalType = 26
//	VerticalType_Media                VerticalType = 27
//	VerticalType_Government           VerticalType = 28
//	VerticalType_Finance              VerticalType = 29
//	VerticalType_Other                VerticalType = 255
//)
