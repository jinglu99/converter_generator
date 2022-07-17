package thrifts

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

const (
	UNITFEN = "fen"

	UNITSECOND = "second"

	UNITPEOPLE = "people"

	UNITNUMBER = "number"
)

type TargetType int64

const (
	TargetType_Live                        TargetType = 1
	TargetType_Watch                       TargetType = 2
	TargetType_Revenue                     TargetType = 3
	TargetType_PayUV                       TargetType = 4
	TargetType_Live10KFans                 TargetType = 5
	TargetType_Follow                      TargetType = 6
	TargetType_EcommerceGMV                TargetType = 7
	TargetType_EcommerceSaleUV             TargetType = 8
	TargetType_Ecommerce                   TargetType = 9
	TargetType_Interaction                 TargetType = 10
	TargetType_CommercialSale              TargetType = 11
	TargetType_AuthorSupply                TargetType = 12
	TargetType_ContentConsumption          TargetType = 13
	TargetType_NicheScenarioCognition      TargetType = 14
	TargetType_LiveEarnings                TargetType = 15
	TargetType_AuthorExperience            TargetType = 16
	TargetType_CommercialCash              TargetType = 17
	TargetType_GameJointOperations         TargetType = 18
	TargetType_InnerExposureIncome         TargetType = 19
	TargetType_HotSpotIncome               TargetType = 20
	TargetType_CreatorIncome               TargetType = 21
	TargetType_OuterInfluenceIncome        TargetType = 22
	TargetType_PoliticMediaIncome          TargetType = 23
	TargetType_ContributeIncome            TargetType = 24
	TargetType_LongTermGameJointOperations TargetType = 25
	TargetType_UserActiveIncome            TargetType = 26
	TargetType_MusicCopyrightIncome        TargetType = 27
	TargetType_MovieTicketIncome           TargetType = 28
	TargetType_FunctionalPenetrationIncome TargetType = 29
	TargetType_BrandPerceptionIncome       TargetType = 30
	TargetType_Others                      TargetType = 255
)

func (p TargetType) String() string {
	switch p {
	case TargetType_Live:
		return "Live"
	case TargetType_Watch:
		return "Watch"
	case TargetType_Revenue:
		return "Revenue"
	case TargetType_PayUV:
		return "PayUV"
	case TargetType_Live10KFans:
		return "Live10KFans"
	case TargetType_Follow:
		return "Follow"
	case TargetType_EcommerceGMV:
		return "EcommerceGMV"
	case TargetType_EcommerceSaleUV:
		return "EcommerceSaleUV"
	case TargetType_Ecommerce:
		return "Ecommerce"
	case TargetType_Interaction:
		return "Interaction"
	case TargetType_CommercialSale:
		return "CommercialSale"
	case TargetType_AuthorSupply:
		return "AuthorSupply"
	case TargetType_ContentConsumption:
		return "ContentConsumption"
	case TargetType_NicheScenarioCognition:
		return "NicheScenarioCognition"
	case TargetType_LiveEarnings:
		return "LiveEarnings"
	case TargetType_AuthorExperience:
		return "AuthorExperience"
	case TargetType_CommercialCash:
		return "CommercialCash"
	case TargetType_GameJointOperations:
		return "GameJointOperations"
	case TargetType_InnerExposureIncome:
		return "InnerExposureIncome"
	case TargetType_HotSpotIncome:
		return "HotSpotIncome"
	case TargetType_CreatorIncome:
		return "CreatorIncome"
	case TargetType_OuterInfluenceIncome:
		return "OuterInfluenceIncome"
	case TargetType_PoliticMediaIncome:
		return "PoliticMediaIncome"
	case TargetType_ContributeIncome:
		return "ContributeIncome"
	case TargetType_LongTermGameJointOperations:
		return "LongTermGameJointOperations"
	case TargetType_UserActiveIncome:
		return "UserActiveIncome"
	case TargetType_MusicCopyrightIncome:
		return "MusicCopyrightIncome"
	case TargetType_MovieTicketIncome:
		return "MovieTicketIncome"
	case TargetType_FunctionalPenetrationIncome:
		return "FunctionalPenetrationIncome"
	case TargetType_BrandPerceptionIncome:
		return "BrandPerceptionIncome"
	case TargetType_Others:
		return "Others"
	}
	return "<UNSET>"
}

func TargetTypeFromString(s string) (TargetType, error) {
	switch s {
	case "Live":
		return TargetType_Live, nil
	case "Watch":
		return TargetType_Watch, nil
	case "Revenue":
		return TargetType_Revenue, nil
	case "PayUV":
		return TargetType_PayUV, nil
	case "Live10KFans":
		return TargetType_Live10KFans, nil
	case "Follow":
		return TargetType_Follow, nil
	case "EcommerceGMV":
		return TargetType_EcommerceGMV, nil
	case "EcommerceSaleUV":
		return TargetType_EcommerceSaleUV, nil
	case "Ecommerce":
		return TargetType_Ecommerce, nil
	case "Interaction":
		return TargetType_Interaction, nil
	case "CommercialSale":
		return TargetType_CommercialSale, nil
	case "AuthorSupply":
		return TargetType_AuthorSupply, nil
	case "ContentConsumption":
		return TargetType_ContentConsumption, nil
	case "NicheScenarioCognition":
		return TargetType_NicheScenarioCognition, nil
	case "LiveEarnings":
		return TargetType_LiveEarnings, nil
	case "AuthorExperience":
		return TargetType_AuthorExperience, nil
	case "CommercialCash":
		return TargetType_CommercialCash, nil
	case "GameJointOperations":
		return TargetType_GameJointOperations, nil
	case "InnerExposureIncome":
		return TargetType_InnerExposureIncome, nil
	case "HotSpotIncome":
		return TargetType_HotSpotIncome, nil
	case "CreatorIncome":
		return TargetType_CreatorIncome, nil
	case "OuterInfluenceIncome":
		return TargetType_OuterInfluenceIncome, nil
	case "PoliticMediaIncome":
		return TargetType_PoliticMediaIncome, nil
	case "ContributeIncome":
		return TargetType_ContributeIncome, nil
	case "LongTermGameJointOperations":
		return TargetType_LongTermGameJointOperations, nil
	case "UserActiveIncome":
		return TargetType_UserActiveIncome, nil
	case "MusicCopyrightIncome":
		return TargetType_MusicCopyrightIncome, nil
	case "MovieTicketIncome":
		return TargetType_MovieTicketIncome, nil
	case "FunctionalPenetrationIncome":
		return TargetType_FunctionalPenetrationIncome, nil
	case "BrandPerceptionIncome":
		return TargetType_BrandPerceptionIncome, nil
	case "Others":
		return TargetType_Others, nil
	}
	return TargetType(0), fmt.Errorf("not a valid TargetType string")
}

func TargetTypePtr(v TargetType) *TargetType { return &v }

func (p *TargetType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = TargetType(result.Int64)
	return
}

func (p *TargetType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type ProfitTplType int64

const (
	ProfitTplType_Customized   ProfitTplType = 0
	ProfitTplType_Standardized ProfitTplType = 1
)

func (p ProfitTplType) String() string {
	switch p {
	case ProfitTplType_Customized:
		return "Customized"
	case ProfitTplType_Standardized:
		return "Standardized"
	}
	return "<UNSET>"
}

func ProfitTplTypeFromString(s string) (ProfitTplType, error) {
	switch s {
	case "Customized":
		return ProfitTplType_Customized, nil
	case "Standardized":
		return ProfitTplType_Standardized, nil
	}
	return ProfitTplType(0), fmt.Errorf("not a valid ProfitTplType string")
}

func ProfitTplTypePtr(v ProfitTplType) *ProfitTplType { return &v }

func (p *ProfitTplType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ProfitTplType(result.Int64)
	return
}

func (p *ProfitTplType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type VerticalType int64

const (
	VerticalType_Reading              VerticalType = 1
	VerticalType_PanHealth            VerticalType = 2
	VerticalType_FamilyEducation      VerticalType = 3
	VerticalType_Examination          VerticalType = 4
	VerticalType_MediaPerson          VerticalType = 5
	VerticalType_BeautyIndustry       VerticalType = 6
	VerticalType_EmotionalCounseling  VerticalType = 7
	VerticalType_Handicraft           VerticalType = 8
	VerticalType_PaintingCalligraphy  VerticalType = 9
	VerticalType_PregnantTreatment    VerticalType = 10
	VerticalType_Treatment            VerticalType = 11
	VerticalType_MusicEducation       VerticalType = 12
	VerticalType_SportsFitness        VerticalType = 13
	VerticalType_LanguageTraining     VerticalType = 14
	VerticalType_NurseryEducation     VerticalType = 15
	VerticalType_VocationalEducation  VerticalType = 16
	VerticalType_ChatRoom             VerticalType = 17
	VerticalType_KTV                  VerticalType = 18
	VerticalType_Game                 VerticalType = 19
	VerticalType_Star                 VerticalType = 20
	VerticalType_Composer             VerticalType = 21
	VerticalType_LiveKnowledgePayment VerticalType = 22
	VerticalType_PayConsultation      VerticalType = 23
	VerticalType_TravelOutdoor        VerticalType = 24
	VerticalType_Car                  VerticalType = 25
	VerticalType_HumanityArt          VerticalType = 26
	VerticalType_Media                VerticalType = 27
	VerticalType_Government           VerticalType = 28
	VerticalType_Finance              VerticalType = 29
	VerticalType_Other                VerticalType = 255
)

func (p VerticalType) String() string {
	switch p {
	case VerticalType_Reading:
		return "Reading"
	case VerticalType_PanHealth:
		return "PanHealth"
	case VerticalType_FamilyEducation:
		return "FamilyEducation"
	case VerticalType_Examination:
		return "Examination"
	case VerticalType_MediaPerson:
		return "MediaPerson"
	case VerticalType_BeautyIndustry:
		return "BeautyIndustry"
	case VerticalType_EmotionalCounseling:
		return "EmotionalCounseling"
	case VerticalType_Handicraft:
		return "Handicraft"
	case VerticalType_PaintingCalligraphy:
		return "PaintingCalligraphy"
	case VerticalType_PregnantTreatment:
		return "PregnantTreatment"
	case VerticalType_Treatment:
		return "Treatment"
	case VerticalType_MusicEducation:
		return "MusicEducation"
	case VerticalType_SportsFitness:
		return "SportsFitness"
	case VerticalType_LanguageTraining:
		return "LanguageTraining"
	case VerticalType_NurseryEducation:
		return "NurseryEducation"
	case VerticalType_VocationalEducation:
		return "VocationalEducation"
	case VerticalType_ChatRoom:
		return "ChatRoom"
	case VerticalType_KTV:
		return "KTV"
	case VerticalType_Game:
		return "Game"
	case VerticalType_Star:
		return "Star"
	case VerticalType_Composer:
		return "Composer"
	case VerticalType_LiveKnowledgePayment:
		return "LiveKnowledgePayment"
	case VerticalType_PayConsultation:
		return "PayConsultation"
	case VerticalType_TravelOutdoor:
		return "TravelOutdoor"
	case VerticalType_Car:
		return "Car"
	case VerticalType_HumanityArt:
		return "HumanityArt"
	case VerticalType_Media:
		return "Media"
	case VerticalType_Government:
		return "Government"
	case VerticalType_Finance:
		return "Finance"
	case VerticalType_Other:
		return "Other"
	}
	return "<UNSET>"
}

func VerticalTypeFromString(s string) (VerticalType, error) {
	switch s {
	case "Reading":
		return VerticalType_Reading, nil
	case "PanHealth":
		return VerticalType_PanHealth, nil
	case "FamilyEducation":
		return VerticalType_FamilyEducation, nil
	case "Examination":
		return VerticalType_Examination, nil
	case "MediaPerson":
		return VerticalType_MediaPerson, nil
	case "BeautyIndustry":
		return VerticalType_BeautyIndustry, nil
	case "EmotionalCounseling":
		return VerticalType_EmotionalCounseling, nil
	case "Handicraft":
		return VerticalType_Handicraft, nil
	case "PaintingCalligraphy":
		return VerticalType_PaintingCalligraphy, nil
	case "PregnantTreatment":
		return VerticalType_PregnantTreatment, nil
	case "Treatment":
		return VerticalType_Treatment, nil
	case "MusicEducation":
		return VerticalType_MusicEducation, nil
	case "SportsFitness":
		return VerticalType_SportsFitness, nil
	case "LanguageTraining":
		return VerticalType_LanguageTraining, nil
	case "NurseryEducation":
		return VerticalType_NurseryEducation, nil
	case "VocationalEducation":
		return VerticalType_VocationalEducation, nil
	case "ChatRoom":
		return VerticalType_ChatRoom, nil
	case "KTV":
		return VerticalType_KTV, nil
	case "Game":
		return VerticalType_Game, nil
	case "Star":
		return VerticalType_Star, nil
	case "Composer":
		return VerticalType_Composer, nil
	case "LiveKnowledgePayment":
		return VerticalType_LiveKnowledgePayment, nil
	case "PayConsultation":
		return VerticalType_PayConsultation, nil
	case "TravelOutdoor":
		return VerticalType_TravelOutdoor, nil
	case "Car":
		return VerticalType_Car, nil
	case "HumanityArt":
		return VerticalType_HumanityArt, nil
	case "Media":
		return VerticalType_Media, nil
	case "Government":
		return VerticalType_Government, nil
	case "Finance":
		return VerticalType_Finance, nil
	case "Other":
		return VerticalType_Other, nil
	}
	return VerticalType(0), fmt.Errorf("not a valid VerticalType string")
}

func VerticalTypePtr(v VerticalType) *VerticalType { return &v }

func (p *VerticalType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = VerticalType(result.Int64)
	return
}

func (p *VerticalType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type ProfitIndex int64

const (
	// 提开播
	ProfitIndex_TotalLiveDurationIncrement ProfitIndex = 1
	ProfitIndex_PerLiveDurationIncrement   ProfitIndex = 2
	ProfitIndex_LiveUVIncrement            ProfitIndex = 3
	// 提看播
	ProfitIndex_TotalWatchDurationForAnchorIncrement ProfitIndex = 11
	ProfitIndex_TotalWatchDurationIncrement          ProfitIndex = 12
	ProfitIndex_PerWatchDurationForAnchorIncrement   ProfitIndex = 13
	ProfitIndex_PerWatchDurationIncrement            ProfitIndex = 14
	ProfitIndex_WatchUVForAnchorIncrement            ProfitIndex = 15
	ProfitIndex_WatchUVIncrement                     ProfitIndex = 16
	// 提营收
	ProfitIndex_TotalRevenueForAnchorIncrement ProfitIndex = 21
	ProfitIndex_TotalRevenueIncrement          ProfitIndex = 22
	// 提关注
	ProfitIndex_TotalFollowerForAnchorIncrement ProfitIndex = 31
	ProfitIndex_TotalFollowerIncrement          ProfitIndex = 32
	// 提互动
	ProfitIndex_TotalPairEnterForAnchorIncrement ProfitIndex = 41
	ProfitIndex_TotalLinkUVForAnchorIncrement    ProfitIndex = 42
	// 提付费UV
	ProfitIndex_TotalPayUVForAnchorIncrement ProfitIndex = 51
	ProfitIndex_TotalPayUVIncrement          ProfitIndex = 52
	// 提电商GMV
	ProfitIndex_TotalGMVForAnchorIncrement ProfitIndex = 61
	ProfitIndex_TotalGMVIncrement          ProfitIndex = 62
	// 提动销UV
	ProfitIndex_TotalEcommerceSaleUVIncrement ProfitIndex = 71
	// 提游戏联运
	ProfitIndex_TotalGameJointOperationsRevenueForAnchorIncrement ProfitIndex = 81
	ProfitIndex_TotalActivityGameIncrement                        ProfitIndex = 82
	ProfitIndex_TotalInnerAdvertisementIncrement                  ProfitIndex = 83
	ProfitIndex_TotalNewActivationUserNumber                      ProfitIndex = 84
	// 站内曝光收益
	ProfitIndex_ActivityTopicVVIncrement ProfitIndex = 91
	// 热点收益
	ProfitIndex_HotSpotVV     ProfitIndex = 101
	ProfitIndex_HotSpotNumber ProfitIndex = 102
	// 创作者收益
	ProfitIndex_AuthorAddFans ProfitIndex = 103
	// 站外影响力收益
	ProfitIndex_RpNumber    ProfitIndex = 111
	ProfitIndex_BaiduMark   ProfitIndex = 112
	ProfitIndex_XinlangMark ProfitIndex = 113
	ProfitIndex_WeiXinMark  ProfitIndex = 114
	ProfitIndex_Other       ProfitIndex = 119
	// 投稿收益
	ProfitIndex_ContributeIncome ProfitIndex = 120
	// 拉新拉活
	ProfitIndex_TotalNewExternalUserNumber ProfitIndex = 130
	ProfitIndex_TotalNewActiveUserNumber   ProfitIndex = 131
	// 端内版权
	ProfitIndex_TotalMusicCVCopyrightIncome ProfitIndex = 140
	ProfitIndex_TotalMovieTicketIncome      ProfitIndex = 141
	// 功能渗透-指标项
	ProfitIndex_FunctionalPenetrationAARatioIncrement    ProfitIndex = 150
	ProfitIndex_FunctionalPenetrationABRatioIncrement    ProfitIndex = 151
	ProfitIndex_FunctionalPenetrationOtherRatioIncrement ProfitIndex = 152
	// 品牌感知-指标项
	ProfitIndex_OnSiteActivityPerceptionRatioIncrement ProfitIndex = 160
	ProfitIndex_BrandReputationIncrement               ProfitIndex = 161
	ProfitIndex_BrandPerceptionOtherIncrement          ProfitIndex = 162
	ProfitIndex_Others                                 ProfitIndex = 255
)

func (p ProfitIndex) String() string {
	switch p {
	case ProfitIndex_TotalLiveDurationIncrement:
		return "TotalLiveDurationIncrement"
	case ProfitIndex_PerLiveDurationIncrement:
		return "PerLiveDurationIncrement"
	case ProfitIndex_LiveUVIncrement:
		return "LiveUVIncrement"
	case ProfitIndex_TotalWatchDurationForAnchorIncrement:
		return "TotalWatchDurationForAnchorIncrement"
	case ProfitIndex_TotalWatchDurationIncrement:
		return "TotalWatchDurationIncrement"
	case ProfitIndex_PerWatchDurationForAnchorIncrement:
		return "PerWatchDurationForAnchorIncrement"
	case ProfitIndex_PerWatchDurationIncrement:
		return "PerWatchDurationIncrement"
	case ProfitIndex_WatchUVForAnchorIncrement:
		return "WatchUVForAnchorIncrement"
	case ProfitIndex_WatchUVIncrement:
		return "WatchUVIncrement"
	case ProfitIndex_TotalRevenueForAnchorIncrement:
		return "TotalRevenueForAnchorIncrement"
	case ProfitIndex_TotalRevenueIncrement:
		return "TotalRevenueIncrement"
	case ProfitIndex_TotalFollowerForAnchorIncrement:
		return "TotalFollowerForAnchorIncrement"
	case ProfitIndex_TotalFollowerIncrement:
		return "TotalFollowerIncrement"
	case ProfitIndex_TotalPairEnterForAnchorIncrement:
		return "TotalPairEnterForAnchorIncrement"
	case ProfitIndex_TotalLinkUVForAnchorIncrement:
		return "TotalLinkUVForAnchorIncrement"
	case ProfitIndex_TotalPayUVForAnchorIncrement:
		return "TotalPayUVForAnchorIncrement"
	case ProfitIndex_TotalPayUVIncrement:
		return "TotalPayUVIncrement"
	case ProfitIndex_TotalGMVForAnchorIncrement:
		return "TotalGMVForAnchorIncrement"
	case ProfitIndex_TotalGMVIncrement:
		return "TotalGMVIncrement"
	case ProfitIndex_TotalEcommerceSaleUVIncrement:
		return "TotalEcommerceSaleUVIncrement"
	case ProfitIndex_TotalGameJointOperationsRevenueForAnchorIncrement:
		return "TotalGameJointOperationsRevenueForAnchorIncrement"
	case ProfitIndex_TotalActivityGameIncrement:
		return "TotalActivityGameIncrement"
	case ProfitIndex_TotalInnerAdvertisementIncrement:
		return "TotalInnerAdvertisementIncrement"
	case ProfitIndex_TotalNewActivationUserNumber:
		return "TotalNewActivationUserNumber"
	case ProfitIndex_ActivityTopicVVIncrement:
		return "ActivityTopicVVIncrement"
	case ProfitIndex_HotSpotVV:
		return "HotSpotVV"
	case ProfitIndex_HotSpotNumber:
		return "HotSpotNumber"
	case ProfitIndex_AuthorAddFans:
		return "AuthorAddFans"
	case ProfitIndex_RpNumber:
		return "RpNumber"
	case ProfitIndex_BaiduMark:
		return "BaiduMark"
	case ProfitIndex_XinlangMark:
		return "XinlangMark"
	case ProfitIndex_WeiXinMark:
		return "WeiXinMark"
	case ProfitIndex_Other:
		return "Other"
	case ProfitIndex_ContributeIncome:
		return "ContributeIncome"
	case ProfitIndex_TotalNewExternalUserNumber:
		return "TotalNewExternalUserNumber"
	case ProfitIndex_TotalNewActiveUserNumber:
		return "TotalNewActiveUserNumber"
	case ProfitIndex_TotalMusicCVCopyrightIncome:
		return "TotalMusicCVCopyrightIncome"
	case ProfitIndex_TotalMovieTicketIncome:
		return "TotalMovieTicketIncome"
	case ProfitIndex_FunctionalPenetrationAARatioIncrement:
		return "FunctionalPenetrationAARatioIncrement"
	case ProfitIndex_FunctionalPenetrationABRatioIncrement:
		return "FunctionalPenetrationABRatioIncrement"
	case ProfitIndex_FunctionalPenetrationOtherRatioIncrement:
		return "FunctionalPenetrationOtherRatioIncrement"
	case ProfitIndex_OnSiteActivityPerceptionRatioIncrement:
		return "OnSiteActivityPerceptionRatioIncrement"
	case ProfitIndex_BrandReputationIncrement:
		return "BrandReputationIncrement"
	case ProfitIndex_BrandPerceptionOtherIncrement:
		return "BrandPerceptionOtherIncrement"
	case ProfitIndex_Others:
		return "Others"
	}
	return "<UNSET>"
}

func ProfitIndexFromString(s string) (ProfitIndex, error) {
	switch s {
	case "TotalLiveDurationIncrement":
		return ProfitIndex_TotalLiveDurationIncrement, nil
	case "PerLiveDurationIncrement":
		return ProfitIndex_PerLiveDurationIncrement, nil
	case "LiveUVIncrement":
		return ProfitIndex_LiveUVIncrement, nil
	case "TotalWatchDurationForAnchorIncrement":
		return ProfitIndex_TotalWatchDurationForAnchorIncrement, nil
	case "TotalWatchDurationIncrement":
		return ProfitIndex_TotalWatchDurationIncrement, nil
	case "PerWatchDurationForAnchorIncrement":
		return ProfitIndex_PerWatchDurationForAnchorIncrement, nil
	case "PerWatchDurationIncrement":
		return ProfitIndex_PerWatchDurationIncrement, nil
	case "WatchUVForAnchorIncrement":
		return ProfitIndex_WatchUVForAnchorIncrement, nil
	case "WatchUVIncrement":
		return ProfitIndex_WatchUVIncrement, nil
	case "TotalRevenueForAnchorIncrement":
		return ProfitIndex_TotalRevenueForAnchorIncrement, nil
	case "TotalRevenueIncrement":
		return ProfitIndex_TotalRevenueIncrement, nil
	case "TotalFollowerForAnchorIncrement":
		return ProfitIndex_TotalFollowerForAnchorIncrement, nil
	case "TotalFollowerIncrement":
		return ProfitIndex_TotalFollowerIncrement, nil
	case "TotalPairEnterForAnchorIncrement":
		return ProfitIndex_TotalPairEnterForAnchorIncrement, nil
	case "TotalLinkUVForAnchorIncrement":
		return ProfitIndex_TotalLinkUVForAnchorIncrement, nil
	case "TotalPayUVForAnchorIncrement":
		return ProfitIndex_TotalPayUVForAnchorIncrement, nil
	case "TotalPayUVIncrement":
		return ProfitIndex_TotalPayUVIncrement, nil
	case "TotalGMVForAnchorIncrement":
		return ProfitIndex_TotalGMVForAnchorIncrement, nil
	case "TotalGMVIncrement":
		return ProfitIndex_TotalGMVIncrement, nil
	case "TotalEcommerceSaleUVIncrement":
		return ProfitIndex_TotalEcommerceSaleUVIncrement, nil
	case "TotalGameJointOperationsRevenueForAnchorIncrement":
		return ProfitIndex_TotalGameJointOperationsRevenueForAnchorIncrement, nil
	case "TotalActivityGameIncrement":
		return ProfitIndex_TotalActivityGameIncrement, nil
	case "TotalInnerAdvertisementIncrement":
		return ProfitIndex_TotalInnerAdvertisementIncrement, nil
	case "TotalNewActivationUserNumber":
		return ProfitIndex_TotalNewActivationUserNumber, nil
	case "ActivityTopicVVIncrement":
		return ProfitIndex_ActivityTopicVVIncrement, nil
	case "HotSpotVV":
		return ProfitIndex_HotSpotVV, nil
	case "HotSpotNumber":
		return ProfitIndex_HotSpotNumber, nil
	case "AuthorAddFans":
		return ProfitIndex_AuthorAddFans, nil
	case "RpNumber":
		return ProfitIndex_RpNumber, nil
	case "BaiduMark":
		return ProfitIndex_BaiduMark, nil
	case "XinlangMark":
		return ProfitIndex_XinlangMark, nil
	case "WeiXinMark":
		return ProfitIndex_WeiXinMark, nil
	case "Other":
		return ProfitIndex_Other, nil
	case "ContributeIncome":
		return ProfitIndex_ContributeIncome, nil
	case "TotalNewExternalUserNumber":
		return ProfitIndex_TotalNewExternalUserNumber, nil
	case "TotalNewActiveUserNumber":
		return ProfitIndex_TotalNewActiveUserNumber, nil
	case "TotalMusicCVCopyrightIncome":
		return ProfitIndex_TotalMusicCVCopyrightIncome, nil
	case "TotalMovieTicketIncome":
		return ProfitIndex_TotalMovieTicketIncome, nil
	case "FunctionalPenetrationAARatioIncrement":
		return ProfitIndex_FunctionalPenetrationAARatioIncrement, nil
	case "FunctionalPenetrationABRatioIncrement":
		return ProfitIndex_FunctionalPenetrationABRatioIncrement, nil
	case "FunctionalPenetrationOtherRatioIncrement":
		return ProfitIndex_FunctionalPenetrationOtherRatioIncrement, nil
	case "OnSiteActivityPerceptionRatioIncrement":
		return ProfitIndex_OnSiteActivityPerceptionRatioIncrement, nil
	case "BrandReputationIncrement":
		return ProfitIndex_BrandReputationIncrement, nil
	case "BrandPerceptionOtherIncrement":
		return ProfitIndex_BrandPerceptionOtherIncrement, nil
	case "Others":
		return ProfitIndex_Others, nil
	}
	return ProfitIndex(0), fmt.Errorf("not a valid ProfitIndex string")
}

func ProfitIndexPtr(v ProfitIndex) *ProfitIndex { return &v }

func (p *ProfitIndex) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ProfitIndex(result.Int64)
	return
}

func (p *ProfitIndex) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type CrowdChooseType int64

const (
	CrowdChooseType_All       CrowdChooseType = 1
	CrowdChooseType_AllowList CrowdChooseType = 2
	CrowdChooseType_Custom    CrowdChooseType = 3
)

func (p CrowdChooseType) String() string {
	switch p {
	case CrowdChooseType_All:
		return "All"
	case CrowdChooseType_AllowList:
		return "AllowList"
	case CrowdChooseType_Custom:
		return "Custom"
	}
	return "<UNSET>"
}

func CrowdChooseTypeFromString(s string) (CrowdChooseType, error) {
	switch s {
	case "All":
		return CrowdChooseType_All, nil
	case "AllowList":
		return CrowdChooseType_AllowList, nil
	case "Custom":
		return CrowdChooseType_Custom, nil
	}
	return CrowdChooseType(0), fmt.Errorf("not a valid CrowdChooseType string")
}

func CrowdChooseTypePtr(v CrowdChooseType) *CrowdChooseType { return &v }

func (p *CrowdChooseType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = CrowdChooseType(result.Int64)
	return
}

func (p *CrowdChooseType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type CrowdType int64

const (
	CrowdType_Anchor CrowdType = 1
	CrowdType_User   CrowdType = 2
)

func (p CrowdType) String() string {
	switch p {
	case CrowdType_Anchor:
		return "Anchor"
	case CrowdType_User:
		return "User"
	}
	return "<UNSET>"
}

func CrowdTypeFromString(s string) (CrowdType, error) {
	switch s {
	case "Anchor":
		return CrowdType_Anchor, nil
	case "User":
		return CrowdType_User, nil
	}
	return CrowdType(0), fmt.Errorf("not a valid CrowdType string")
}

func CrowdTypePtr(v CrowdType) *CrowdType { return &v }

func (p *CrowdType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = CrowdType(result.Int64)
	return
}

func (p *CrowdType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type MarkFansLevel int64

const (
	MarkFansLevel_Normal          MarkFansLevel = 1
	MarkFansLevel_TenThonsand     MarkFansLevel = 2
	MarkFansLevel_HundredThousand MarkFansLevel = 3
	MarkFansLevel_Millon          MarkFansLevel = 4
)

func (p MarkFansLevel) String() string {
	switch p {
	case MarkFansLevel_Normal:
		return "Normal"
	case MarkFansLevel_TenThonsand:
		return "TenThonsand"
	case MarkFansLevel_HundredThousand:
		return "HundredThousand"
	case MarkFansLevel_Millon:
		return "Millon"
	}
	return "<UNSET>"
}

func MarkFansLevelFromString(s string) (MarkFansLevel, error) {
	switch s {
	case "Normal":
		return MarkFansLevel_Normal, nil
	case "TenThonsand":
		return MarkFansLevel_TenThonsand, nil
	case "HundredThousand":
		return MarkFansLevel_HundredThousand, nil
	case "Millon":
		return MarkFansLevel_Millon, nil
	}
	return MarkFansLevel(0), fmt.Errorf("not a valid MarkFansLevel string")
}

func MarkFansLevelPtr(v MarkFansLevel) *MarkFansLevel { return &v }

func (p *MarkFansLevel) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = MarkFansLevel(result.Int64)
	return
}

func (p *MarkFansLevel) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type GameType int64

const (
	GameType_BringGame    GameType = 1
	GameType_LiveRoomGame GameType = 2
)

func (p GameType) String() string {
	switch p {
	case GameType_BringGame:
		return "BringGame"
	case GameType_LiveRoomGame:
		return "LiveRoomGame"
	}
	return "<UNSET>"
}

func GameTypeFromString(s string) (GameType, error) {
	switch s {
	case "BringGame":
		return GameType_BringGame, nil
	case "LiveRoomGame":
		return GameType_LiveRoomGame, nil
	}
	return GameType(0), fmt.Errorf("not a valid GameType string")
}

func GameTypePtr(v GameType) *GameType { return &v }

func (p *GameType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = GameType(result.Int64)
	return
}

func (p *GameType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type ROIType int64

const (
	ROIType_Customized            ROIType = 0
	ROIType_ShortTermROI          ROIType = 1
	ROIType_BusinessInvestmentROI ROIType = 2
	ROIType_LongTermROI           ROIType = 3
	ROIType_OtherROI              ROIType = 4
)

func (p ROIType) String() string {
	switch p {
	case ROIType_Customized:
		return "Customized"
	case ROIType_ShortTermROI:
		return "ShortTermROI"
	case ROIType_BusinessInvestmentROI:
		return "BusinessInvestmentROI"
	case ROIType_LongTermROI:
		return "LongTermROI"
	case ROIType_OtherROI:
		return "OtherROI"
	}
	return "<UNSET>"
}

func ROITypeFromString(s string) (ROIType, error) {
	switch s {
	case "Customized":
		return ROIType_Customized, nil
	case "ShortTermROI":
		return ROIType_ShortTermROI, nil
	case "BusinessInvestmentROI":
		return ROIType_BusinessInvestmentROI, nil
	case "LongTermROI":
		return ROIType_LongTermROI, nil
	case "OtherROI":
		return ROIType_OtherROI, nil
	}
	return ROIType(0), fmt.Errorf("not a valid ROIType string")
}

func ROITypePtr(v ROIType) *ROIType { return &v }

func (p *ROIType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ROIType(result.Int64)
	return
}

func (p *ROIType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type ProjectType int64

const (
	ProjectType_Default ProjectType = 1
	ProjectType_Main    ProjectType = 2
	ProjectType_Sub     ProjectType = 3
)

func (p ProjectType) String() string {
	switch p {
	case ProjectType_Default:
		return "Default"
	case ProjectType_Main:
		return "Main"
	case ProjectType_Sub:
		return "Sub"
	}
	return "<UNSET>"
}

func ProjectTypeFromString(s string) (ProjectType, error) {
	switch s {
	case "Default":
		return ProjectType_Default, nil
	case "Main":
		return ProjectType_Main, nil
	case "Sub":
		return ProjectType_Sub, nil
	}
	return ProjectType(0), fmt.Errorf("not a valid ProjectType string")
}

func ProjectTypePtr(v ProjectType) *ProjectType { return &v }

func (p *ProjectType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ProjectType(result.Int64)
	return
}

func (p *ProjectType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type ProjectSourceType int64

const (
	ProjectSourceType_Gaiya       ProjectSourceType = 0
	ProjectSourceType_Chengjiao   ProjectSourceType = 1
	ProjectSourceType_LocalLife   ProjectSourceType = 2
	ProjectSourceType_LocalLifeV2 ProjectSourceType = 3
)

func (p ProjectSourceType) String() string {
	switch p {
	case ProjectSourceType_Gaiya:
		return "Gaiya"
	case ProjectSourceType_Chengjiao:
		return "Chengjiao"
	case ProjectSourceType_LocalLife:
		return "LocalLife"
	case ProjectSourceType_LocalLifeV2:
		return "LocalLifeV2"
	}
	return "<UNSET>"
}

func ProjectSourceTypeFromString(s string) (ProjectSourceType, error) {
	switch s {
	case "Gaiya":
		return ProjectSourceType_Gaiya, nil
	case "Chengjiao":
		return ProjectSourceType_Chengjiao, nil
	case "LocalLife":
		return ProjectSourceType_LocalLife, nil
	case "LocalLifeV2":
		return ProjectSourceType_LocalLifeV2, nil
	}
	return ProjectSourceType(0), fmt.Errorf("not a valid ProjectSourceType string")
}

func ProjectSourceTypePtr(v ProjectSourceType) *ProjectSourceType { return &v }

func (p *ProjectSourceType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ProjectSourceType(result.Int64)
	return
}

func (p *ProjectSourceType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type ProjectReplayStatus int64

const (
	ProjectReplayStatus_None                ProjectReplayStatus = 0
	ProjectReplayStatus_RecordingProfit     ProjectReplayStatus = 1
	ProjectReplayStatus_CreatorConfirm      ProjectReplayStatus = 2
	ProjectReplayStatus_FinanceConfirm      ProjectReplayStatus = 3
	ProjectReplayStatus_Confirmed           ProjectReplayStatus = 4
	ProjectReplayStatus_CannotBeEstimated   ProjectReplayStatus = 5
	ProjectReplayStatus_ReplayInfoError     ProjectReplayStatus = 6
	ProjectReplayStatus_LongTermDataPrepare ProjectReplayStatus = 7
)

func (p ProjectReplayStatus) String() string {
	switch p {
	case ProjectReplayStatus_None:
		return "None"
	case ProjectReplayStatus_RecordingProfit:
		return "RecordingProfit"
	case ProjectReplayStatus_CreatorConfirm:
		return "CreatorConfirm"
	case ProjectReplayStatus_FinanceConfirm:
		return "FinanceConfirm"
	case ProjectReplayStatus_Confirmed:
		return "Confirmed"
	case ProjectReplayStatus_CannotBeEstimated:
		return "CannotBeEstimated"
	case ProjectReplayStatus_ReplayInfoError:
		return "ReplayInfoError"
	case ProjectReplayStatus_LongTermDataPrepare:
		return "LongTermDataPrepare"
	}
	return "<UNSET>"
}

func ProjectReplayStatusFromString(s string) (ProjectReplayStatus, error) {
	switch s {
	case "None":
		return ProjectReplayStatus_None, nil
	case "RecordingProfit":
		return ProjectReplayStatus_RecordingProfit, nil
	case "CreatorConfirm":
		return ProjectReplayStatus_CreatorConfirm, nil
	case "FinanceConfirm":
		return ProjectReplayStatus_FinanceConfirm, nil
	case "Confirmed":
		return ProjectReplayStatus_Confirmed, nil
	case "CannotBeEstimated":
		return ProjectReplayStatus_CannotBeEstimated, nil
	case "ReplayInfoError":
		return ProjectReplayStatus_ReplayInfoError, nil
	case "LongTermDataPrepare":
		return ProjectReplayStatus_LongTermDataPrepare, nil
	}
	return ProjectReplayStatus(0), fmt.Errorf("not a valid ProjectReplayStatus string")
}

func ProjectReplayStatusPtr(v ProjectReplayStatus) *ProjectReplayStatus { return &v }

func (p *ProjectReplayStatus) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ProjectReplayStatus(result.Int64)
	return
}

func (p *ProjectReplayStatus) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type ProjectReplayFlowNodeStatus int64

const (
	ProjectReplayFlowNodeStatus_NotStart     ProjectReplayFlowNodeStatus = 0
	ProjectReplayFlowNodeStatus_InProcessing ProjectReplayFlowNodeStatus = 1
	ProjectReplayFlowNodeStatus_Finished     ProjectReplayFlowNodeStatus = 2
)

func (p ProjectReplayFlowNodeStatus) String() string {
	switch p {
	case ProjectReplayFlowNodeStatus_NotStart:
		return "NotStart"
	case ProjectReplayFlowNodeStatus_InProcessing:
		return "InProcessing"
	case ProjectReplayFlowNodeStatus_Finished:
		return "Finished"
	}
	return "<UNSET>"
}

func ProjectReplayFlowNodeStatusFromString(s string) (ProjectReplayFlowNodeStatus, error) {
	switch s {
	case "NotStart":
		return ProjectReplayFlowNodeStatus_NotStart, nil
	case "InProcessing":
		return ProjectReplayFlowNodeStatus_InProcessing, nil
	case "Finished":
		return ProjectReplayFlowNodeStatus_Finished, nil
	}
	return ProjectReplayFlowNodeStatus(0), fmt.Errorf("not a valid ProjectReplayFlowNodeStatus string")
}

func ProjectReplayFlowNodeStatusPtr(v ProjectReplayFlowNodeStatus) *ProjectReplayFlowNodeStatus {
	return &v
}

func (p *ProjectReplayFlowNodeStatus) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ProjectReplayFlowNodeStatus(result.Int64)
	return
}

func (p *ProjectReplayFlowNodeStatus) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type ProjectReplayAuthType int64

const (
	ProjectReplayAuthType_ProfitRecorder       ProjectReplayAuthType = 1
	ProjectReplayAuthType_ProfitRecorderLeader ProjectReplayAuthType = 2
	ProjectReplayAuthType_ProjectCreator       ProjectReplayAuthType = 3
	ProjectReplayAuthType_ProjectCreatorLeader ProjectReplayAuthType = 4
	ProjectReplayAuthType_Finance              ProjectReplayAuthType = 5
	ProjectReplayAuthType_FinanceLeader        ProjectReplayAuthType = 6
	ProjectReplayAuthType_ROIAdmin             ProjectReplayAuthType = 7
	ProjectReplayAuthType_Cooperator           ProjectReplayAuthType = 8
)

func (p ProjectReplayAuthType) String() string {
	switch p {
	case ProjectReplayAuthType_ProfitRecorder:
		return "ProfitRecorder"
	case ProjectReplayAuthType_ProfitRecorderLeader:
		return "ProfitRecorderLeader"
	case ProjectReplayAuthType_ProjectCreator:
		return "ProjectCreator"
	case ProjectReplayAuthType_ProjectCreatorLeader:
		return "ProjectCreatorLeader"
	case ProjectReplayAuthType_Finance:
		return "Finance"
	case ProjectReplayAuthType_FinanceLeader:
		return "FinanceLeader"
	case ProjectReplayAuthType_ROIAdmin:
		return "ROIAdmin"
	case ProjectReplayAuthType_Cooperator:
		return "Cooperator"
	}
	return "<UNSET>"
}

func ProjectReplayAuthTypeFromString(s string) (ProjectReplayAuthType, error) {
	switch s {
	case "ProfitRecorder":
		return ProjectReplayAuthType_ProfitRecorder, nil
	case "ProfitRecorderLeader":
		return ProjectReplayAuthType_ProfitRecorderLeader, nil
	case "ProjectCreator":
		return ProjectReplayAuthType_ProjectCreator, nil
	case "ProjectCreatorLeader":
		return ProjectReplayAuthType_ProjectCreatorLeader, nil
	case "Finance":
		return ProjectReplayAuthType_Finance, nil
	case "FinanceLeader":
		return ProjectReplayAuthType_FinanceLeader, nil
	case "ROIAdmin":
		return ProjectReplayAuthType_ROIAdmin, nil
	case "Cooperator":
		return ProjectReplayAuthType_Cooperator, nil
	}
	return ProjectReplayAuthType(0), fmt.Errorf("not a valid ProjectReplayAuthType string")
}

func ProjectReplayAuthTypePtr(v ProjectReplayAuthType) *ProjectReplayAuthType { return &v }

func (p *ProjectReplayAuthType) Scan(value interface{}) (err error) {
	var result sql.NullInt64
	err = result.Scan(value)
	*p = ProjectReplayAuthType(result.Int64)
	return
}

func (p *ProjectReplayAuthType) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return int64(*p), nil
}

type Objective struct {
	Population               *string                    `thrift:"Population,1" json:"Population,omitempty"`
	Targets                  []TargetType               `thrift:"Targets,2" json:"Targets,omitempty"`
	Profits                  []*ProjectProfit           `thrift:"Profits,3" json:"Profits,omitempty"`
	ROI                      *ROI                       `thrift:"ROI,4" json:"ROI,omitempty"`
	ParentBudgetDepartmentID *string                    `thrift:"ParentBudgetDepartmentID,5" json:"ParentBudgetDepartmentID,omitempty"`
	ROIs                     []*ROI                     `thrift:"ROIs,6" json:"ROIs,omitempty"`
	MainROIType              *ROIType                   `thrift:"MainROIType,7" json:"MainROIType,omitempty"`
	LocalLifeBusinessTargets []*LocalLifeBusinessTarget `thrift:"LocalLifeBusinessTargets,8" json:"LocalLifeBusinessTargets,omitempty"`
}

type ProjectProfit struct {
	ProfitType         string              `thrift:"ProfitType,1,required" json:"ProfitType,omitempty"`
	Value              string              `thrift:"Value,2,required" json:"Value,omitempty"`
	ProfitTplType      *ProfitTplType      `thrift:"ProfitTplType,3" json:"ProfitTplType,omitempty"`
	StandardizedProfit *StandardizedProfit `thrift:"StandardizedProfit,4" json:"StandardizedProfit,omitempty"`
	Target             *TargetType         `thrift:"Target,5" json:"Target,omitempty"`
}

type StandardizedProfit struct {
	TargetType             *TargetType              `thrift:"TargetType,1" json:"TargetType,omitempty"`
	ProfitIndex            *ProfitIndex             `thrift:"ProfitIndex,2" json:"ProfitIndex,omitempty"`
	IncrementValue         *float64                 `thrift:"IncrementValue,3" json:"IncrementValue,omitempty"`
	VerticalTypes          []VerticalType           `thrift:"VerticalTypes,4" json:"VerticalTypes,omitempty"`
	ExpectRate             *float64                 `thrift:"ExpectRate,5" json:"ExpectRate,omitempty"`
	CrowdID                *string                  `thrift:"CrowdID,6" json:"CrowdID,omitempty"`
	CrowdEstimatedNum      *float64                 `thrift:"CrowdEstimatedNum,7" json:"CrowdEstimatedNum,omitempty"`
	CrowdConfigJSON        *string                  `thrift:"CrowdConfigJson,8" json:"CrowdConfigJson,omitempty"`
	CrowdType              *int32                   `thrift:"CrowdType,9" json:"CrowdType,omitempty"`
	TargetCustomName       *string                  `thrift:"TargetCustomName,10" json:"TargetCustomName,omitempty"`
	ProfitIndexName        *string                  `thrift:"ProfitIndexName,11" json:"ProfitIndexName,omitempty"`
	CrowdUpdateTime        *int32                   `thrift:"CrowdUpdateTime,12" json:"CrowdUpdateTime,omitempty"`
	GameIDs                []string                 `thrift:"GameIDs,13" json:"GameIDs,omitempty"`
	CustomValue            *string                  `thrift:"CustomValue,14" json:"CustomValue,omitempty"`
	TopicIDs               []string                 `thrift:"TopicIDs,15" json:"TopicIDs,omitempty"`
	MulitProfits           []*MulitProfits          `thrift:"MulitProfits,16" json:"MulitProfits,omitempty"`
	AdditionalRemarks      *string                  `thrift:"AdditionalRemarks,17" json:"AdditionalRemarks,omitempty"`
	ContributeLevelIncomes []*ContributeLevelIncome `thrift:"ContributeLevelIncomes,18" json:"ContributeLevelIncomes,omitempty"`
	ActivityHotSpots       []string                 `thrift:"ActivityHotSpots,19" json:"ActivityHotSpots,omitempty"`
	GameMulitProfits       []*GameMulitProfits      `thrift:"GameMulitProfits,20" json:"GameMulitProfits,omitempty"`
	ProcessingMessage      *ProcessingMessage       `thrift:"ProcessingMessage,21" json:"ProcessingMessage,omitempty"`
	MusicMulitProfits      []*MusicMulitProfit      `thrift:"MusicMulitProfits,22" json:"MusicMulitProfits,omitempty"`
	MovieProfit            *MovieProfit             `thrift:"MovieProfit,23" json:"MovieProfit,omitempty"`
	ActivityHotSpotIDs     []string                 `thrift:"ActivityHotSpotIDs,24" json:"ActivityHotSpotIDs,omitempty"`
	MusicProfitInfos       []*MusicProfitInfo       `thrift:"MusicProfitInfos,25" json:"MusicProfitInfos,omitempty"`
	AbilityName            *string                  `thrift:"AbilityName,26" json:"AbilityName,omitempty"`
	LibraURL               *string                  `thrift:"LibraURL,27" json:"LibraURL,omitempty"`
	ChannelInfoList        []*ChannelInfo           `thrift:"ChannelInfoList,40" json:"ChannelInfoList,omitempty"`
}

type MulitProfits struct {
	ProfitIndex      *ProfitIndex `thrift:"ProfitIndex,1" json:"ProfitIndex,omitempty"`
	IncrementValue   *float64     `thrift:"IncrementValue,2" json:"IncrementValue,omitempty"`
	CustomValue      *string      `thrift:"CustomValue,3" json:"CustomValue,omitempty"`
	ExpectRate       *float64     `thrift:"ExpectRate,4" json:"ExpectRate,omitempty"`
	ProfitIndexName  *string      `thrift:"ProfitIndexName,5" json:"ProfitIndexName,omitempty"`
	TargetCustomName *string      `thrift:"TargetCustomName,6" json:"TargetCustomName,omitempty"`
}

type ContributeLevelIncome struct {
	MarkFansLevel  MarkFansLevel `thrift:"MarkFansLevel,1,required" json:"MarkFansLevel"`
	IncrementValue float64       `thrift:"IncrementValue,2,required" json:"IncrementValue"`
}

type GameMulitProfits struct {
	GameID       *string         `thrift:"GameID,1" json:"GameID,omitempty"`
	GameName     *string         `thrift:"GameName,2" json:"GameName,omitempty"`
	MulitProfits []*MulitProfits `thrift:"MulitProfits,3" json:"MulitProfits,omitempty"`
	LTV          *float64        `thrift:"LTV,4" json:"LTV,omitempty"`
	GameType     *GameType       `thrift:"GameType,5" json:"GameType,omitempty"`
}

type ProcessingMessage struct {
	GameID          *string   `thrift:"GameID,1" json:"GameID,omitempty"`
	GameName        *string   `thrift:"GameName,2" json:"GameName,omitempty"`
	LTV             *float64  `thrift:"LTV,3" json:"LTV,omitempty"`
	GameType        *GameType `thrift:"GameType,4" json:"GameType,omitempty"`
	GameRolePercent *float64  `thrift:"GameRolePercent,5" json:"GameRolePercent,omitempty"`
}

type MusicMulitProfit struct {
	MusicCopyrightType *string  `thrift:"MusicCopyrightType,1" json:"MusicCopyrightType,omitempty"`
	IncrementValue     *float64 `thrift:"IncrementValue,2" json:"IncrementValue,omitempty"`
	GroupID            *string  `thrift:"GroupID,3" json:"GroupID,omitempty"`
	ClipIDs            []string `thrift:"ClipIDs,4" json:"ClipIDs,omitempty"`
	CopyrightIncome    *float64 `thrift:"CopyrightIncome,5" json:"CopyrightIncome,omitempty"`
}

type MusicProfitInfo struct {
	MusicCopyrightType *string  `thrift:"MusicCopyrightType,1" json:"MusicCopyrightType,omitempty"`
	Name               *string  `thrift:"Name,2" json:"Name,omitempty"`
	GroupID            *string  `thrift:"GroupID,3" json:"GroupID,omitempty"`
	ClipIDs            []string `thrift:"ClipIDs,4" json:"ClipIDs,omitempty"`
}

func NewMusicProfitInfo() *MusicProfitInfo {
	return &MusicProfitInfo{}
}

type MovieProfit struct {
	MovieIDs       []string `thrift:"MovieIDs,1" json:"MovieIDs,omitempty"`
	PlatformIncome *float64 `thrift:"PlatformIncome,2" json:"PlatformIncome,omitempty"`
	MovieNames     []string `thrift:"MovieNames,3" json:"MovieNames,omitempty"`
}

type ChannelInfo struct {
	ChannelID   *int32  `thrift:"ChannelID,1" json:"ChannelID,omitempty"`
	ChannelName *string `thrift:"ChannelName,2" json:"ChannelName,omitempty"`
}

type ReturnFormulaIndex struct {
	Key   *string  `thrift:"Key,1" json:"Key,omitempty"`
	Value *float64 `thrift:"Value,2" json:"Value,omitempty"`
	Desc  *string  `thrift:"Desc,3" json:"Desc,omitempty"`
}

type ShortTermROI struct {
	ReturnFormulaIndexs map[string]*ReturnFormulaIndex `thrift:"ReturnFormulaIndexs,1" json:"ReturnFormulaIndexs,omitempty"`
	Desc                *string                        `thrift:"Desc,2" json:"Desc,omitempty"`
	Roi                 *string                        `thrift:"Roi,3" json:"Roi,omitempty"`
	I                   *int64                         `thrift:"I,4" json:"I,omitempty"`
}

type BusinessInvestmentROI struct {
	ReturnFormulaIndexs map[string]*ReturnFormulaIndex `thrift:"ReturnFormulaIndexs,1" json:"ReturnFormulaIndexs,omitempty"`
	Desc                *string                        `thrift:"Desc,2" json:"Desc,omitempty"`
	Roi                 *string                        `thrift:"Roi,3" json:"Roi,omitempty"`
	I                   *int64                         `thrift:"I,4" json:"I,omitempty"`
}

type LongTermROI struct {
	ReturnFormulaIndexs map[string]*ReturnFormulaIndex `thrift:"ReturnFormulaIndexs,1" json:"ReturnFormulaIndexs,omitempty"`
	Desc                *string                        `thrift:"Desc,2" json:"Desc,omitempty"`
	Roi                 *string                        `thrift:"Roi,3" json:"Roi,omitempty"`
	I                   *int64                         `thrift:"I,4" json:"I,omitempty"`
}

type OtherROI struct {
	Desc        *string  `thrift:"Desc,1" json:"Desc,omitempty"`
	ROIValue    *string  `thrift:"ROIValue,2" json:"ROIValue,omitempty"`
	CostValue   *float64 `thrift:"CostValue,3" json:"CostValue,omitempty"`
	ReturnValue *float64 `thrift:"ReturnValue,4" json:"ReturnValue,omitempty"`
}

type ROI struct {
	Value                 *string                `thrift:"Value,1" json:"Value,omitempty"`
	Desc                  *string                `thrift:"Desc,2" json:"Desc,omitempty"`
	ROIType               *ROIType               `thrift:"ROIType,3" json:"ROIType,omitempty"`
	ShortTermROI          *ShortTermROI          `thrift:"ShortTermROI,4" json:"ShortTermROI,omitempty"`
	BusinessInvestmentROI *BusinessInvestmentROI `thrift:"BusinessInvestmentROI,5" json:"BusinessInvestmentROI,omitempty"`
	LongTermROI           *LongTermROI           `thrift:"LongTermROI,6" json:"LongTermROI,omitempty"`
	OtherROI              *OtherROI              `thrift:"OtherROI,7" json:"OtherROI,omitempty"`
}

type LocalLifeBusinessTarget struct {
	TargetID    *int64   `thrift:"TargetID,1" json:"TargetID,omitempty"`
	TargetVlue  *float64 `thrift:"TargetVlue,2" json:"TargetVlue,omitempty"`
	CustomValue *string  `thrift:"CustomValue,3" json:"CustomValue,omitempty"`
}

type ProjectBudget struct {
	TotalCost    *int64 `thrift:"TotalCost,1" json:"TotalCost,omitempty"`
	TotalCostFin *int64 `thrift:"TotalCostFin,2" json:"TotalCostFin,omitempty"`
}

type ProjectRelation struct {
	MainProjectID *int64  `thrift:"MainProjectID,1" json:"MainProjectID,omitempty"`
	SubProjectIDs []int64 `thrift:"SubProjectIDs,2" json:"SubProjectIDs,omitempty"`
}

type ProjectReplayAuthPeople struct {
	Email           string  `thrift:"Email,1,required" json:"Email"`
	LeaderEmail     *string `thrift:"LeaderEmail,2" json:"LeaderEmail,omitempty"`
	ProfitIndexList []int64 `thrift:"ProfitIndexList,3" json:"ProfitIndexList,omitempty"`
}

type ProjectReplayFlowNode struct {
	NodeType     ProjectReplayStatus         `thrift:"NodeType,1,required" json:"NodeType"`
	OperatorList []*ProjectReplayAuthPeople  `thrift:"OperatorList,2,required" json:"OperatorList"`
	NodeStatus   ProjectReplayFlowNodeStatus `thrift:"NodeStatus,3,required" json:"NodeStatus"`
	OperateTime  *int64                      `thrift:"OperateTime,4" json:"OperateTime,omitempty"`
}

type ProjectReplayAuth struct {
	AuthType  ProjectReplayAuthType      `thrift:"AuthType,1,required" json:"AuthType"`
	StaffList []*ProjectReplayAuthPeople `thrift:"StaffList,2,required" json:"StaffList"`
}

type InputCost struct {
	CostName string  `thrift:"CostName,1,required" json:"CostName"`
	Cost     float64 `thrift:"Cost,2,required" json:"Cost"`
}

type ProjectReplayInputCost struct {
	InputCostList   []*InputCost `thrift:"InputCostList,1,required" json:"InputCostList"`
	InputCostReason string       `thrift:"InputCostReason,2,required" json:"InputCostReason"`
}

type CommonInputProfit struct {
	TargetType         TargetType  `thrift:"TargetType,1,required" json:"TargetType"`
	ProfitIndex        ProfitIndex `thrift:"ProfitIndex,2,required" json:"ProfitIndex"`
	Profit             *float64    `thrift:"Profit,3" json:"Profit,omitempty"`
	IncrementValue     *float64    `thrift:"IncrementValue,4" json:"IncrementValue,omitempty"`
	Remark             *string     `thrift:"Remark,5" json:"Remark,omitempty"`
	ProfitUnit         *string     `thrift:"ProfitUnit,6" json:"ProfitUnit,omitempty"`
	IncrementValueUnit *string     `thrift:"IncrementValueUnit,7" json:"IncrementValueUnit,omitempty"`
	CustomProfit       *string     `thrift:"CustomProfit,8" json:"CustomProfit,omitempty"`
}

type GameInputProfit struct {
	TargetType         TargetType  `thrift:"TargetType,1,required" json:"TargetType"`
	ProfitIndex        ProfitIndex `thrift:"ProfitIndex,2,required" json:"ProfitIndex"`
	GameID             string      `thrift:"GameID,3,required" json:"GameID"`
	Profit             *float64    `thrift:"Profit,4" json:"Profit,omitempty"`
	IncrementValue     *float64    `thrift:"IncrementValue,5" json:"IncrementValue,omitempty"`
	Remark             *string     `thrift:"Remark,6" json:"Remark,omitempty"`
	ProfitUnit         *string     `thrift:"ProfitUnit,7" json:"ProfitUnit,omitempty"`
	IncrementValueUnit *string     `thrift:"IncrementValueUnit,8" json:"IncrementValueUnit,omitempty"`
	CustomProfit       *string     `thrift:"CustomProfit,9" json:"CustomProfit,omitempty"`
}

type ContributeLevelInputProfit struct {
	TargetType         TargetType    `thrift:"TargetType,1,required" json:"TargetType"`
	ProfitIndex        ProfitIndex   `thrift:"ProfitIndex,2,required" json:"ProfitIndex"`
	MarkFansLevel      MarkFansLevel `thrift:"MarkFansLevel,3,required" json:"MarkFansLevel"`
	Profit             *float64      `thrift:"Profit,4" json:"Profit,omitempty"`
	IncrementValue     *float64      `thrift:"IncrementValue,5" json:"IncrementValue,omitempty"`
	Remark             *string       `thrift:"Remark,6" json:"Remark,omitempty"`
	ProfitUnit         *string       `thrift:"ProfitUnit,7" json:"ProfitUnit,omitempty"`
	IncrementValueUnit *string       `thrift:"IncrementValueUnit,8" json:"IncrementValueUnit,omitempty"`
	CustomProfit       *string       `thrift:"CustomProfit,9" json:"CustomProfit,omitempty"`
}

type MultiInputProfit struct {
	TargetType         TargetType `thrift:"TargetType,1,required" json:"TargetType"`
	ProfitIndexName    *string    `thrift:"ProfitIndexName,2" json:"ProfitIndexName,omitempty"`
	Profit             *float64   `thrift:"Profit,3" json:"Profit,omitempty"`
	IncrementValue     *float64   `thrift:"IncrementValue,4" json:"IncrementValue,omitempty"`
	Remark             *string    `thrift:"Remark,5" json:"Remark,omitempty"`
	ProfitUnit         *string    `thrift:"ProfitUnit,6" json:"ProfitUnit,omitempty"`
	IncrementValueUnit *string    `thrift:"IncrementValueUnit,7" json:"IncrementValueUnit,omitempty"`
	CustomProfit       *string    `thrift:"CustomProfit,8" json:"CustomProfit,omitempty"`
}

type MusicInputProfit struct {
	TargetType         TargetType  `thrift:"TargetType,1,required" json:"TargetType"`
	ProfitIndex        ProfitIndex `thrift:"ProfitIndex,2,required" json:"ProfitIndex"`
	MusicCopyrightType string      `thrift:"MusicCopyrightType,3,required" json:"MusicCopyrightType"`
	Profit             *float64    `thrift:"Profit,4" json:"Profit,omitempty"`
	IncrementValue     *float64    `thrift:"IncrementValue,5" json:"IncrementValue,omitempty"`
	Remark             *string     `thrift:"Remark,6" json:"Remark,omitempty"`
	ProfitUnit         *string     `thrift:"ProfitUnit,7" json:"ProfitUnit,omitempty"`
	IncrementValueUnit *string     `thrift:"IncrementValueUnit,8" json:"IncrementValueUnit,omitempty"`
	CustomProfit       *string     `thrift:"CustomProfit,9" json:"CustomProfit,omitempty"`
}

type InputProfit struct {
	CommonInputProfit          *CommonInputProfit          `thrift:"CommonInputProfit,1" json:"CommonInputProfit,omitempty"`
	GameInputProfit            *GameInputProfit            `thrift:"GameInputProfit,2" json:"GameInputProfit,omitempty"`
	ContributeLevelInputProfit *ContributeLevelInputProfit `thrift:"ContributeLevelInputProfit,3" json:"ContributeLevelInputProfit,omitempty"`
	MultiInputProfitList       []*MultiInputProfit         `thrift:"MultiInputProfitList,4" json:"MultiInputProfitList,omitempty"`
	MusicInputProfit           *MusicInputProfit           `thrift:"MusicInputProfit,5" json:"MusicInputProfit,omitempty"`
}

type ProjectReplayInputProfit struct {
	CommonInputProfitList          []*CommonInputProfit          `thrift:"CommonInputProfitList,1" json:"CommonInputProfitList,omitempty"`
	GameInputProfitList            []*GameInputProfit            `thrift:"GameInputProfitList,2" json:"GameInputProfitList,omitempty"`
	ContributeLevelInputProfitList []*ContributeLevelInputProfit `thrift:"ContributeLevelInputProfitList,3" json:"ContributeLevelInputProfitList,omitempty"`
	MultiInputProfitList           []*MultiInputProfit           `thrift:"MultiInputProfitList,4" json:"MultiInputProfitList,omitempty"`
	MusicInputProfitList           []*MusicInputProfit           `thrift:"MusicInputProfitList,5" json:"MusicInputProfitList,omitempty"`
}

type ProjectReplayDocInfo struct {
	DocLink string  `thrift:"DocLink,1,required" json:"DocLink"`
	DocName *string `thrift:"DocName,2" json:"DocName,omitempty"`
}

type ReplayOtherInfo struct {
	CannotBeEstimatedReason *string `thrift:"CannotBeEstimatedReason,1" json:"CannotBeEstimatedReason,omitempty"`
	ReplayInfoErrorReason   *string `thrift:"ReplayInfoErrorReason,2" json:"ReplayInfoErrorReason,omitempty"`
}

//
//type ProjectReplayInfo struct {
//	ProjectID                int64                     `thrift:"ProjectID,1,required" json:"ProjectID"`
//	Status                   *ProjectReplayStatus      `thrift:"Status,2" json:"Status,omitempty"`
//	ProjectReplayInputCost   *ProjectReplayInputCost   `thrift:"ProjectReplayInputCost,3" json:"ProjectReplayInputCost,omitempty"`
//	ProjectReplayInputProfit *ProjectReplayInputProfit `thrift:"ProjectReplayInputProfit,4" json:"ProjectReplayInputProfit,omitempty"`
//	ReplayDoc                *ProjectReplayDoc         `thrift:"ReplayDoc,5" json:"ReplayDoc,omitempty"`
//	OtherInfo                *ReplayOtherInfo          `thrift:"OtherInfo,6" json:"OtherInfo,omitempty"`
//}

type ObjectiveModel struct {
	ROIs        []*ROI           `thrift:"ROIs,1" json:"ROIs,omitempty"`
	MainROIType *ROIType         `thrift:"MainROIType,2" json:"MainROIType,omitempty"`
	Profits     []*ProjectProfit `thrift:"Profits,3" json:"Profits,omitempty"`
}
