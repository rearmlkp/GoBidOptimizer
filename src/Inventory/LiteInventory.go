package Inventory

import (
	"GoBidOptimizer/src/IO"
	"strconv"
	"GoBidOptimizer/src/Configurations"
)

type LiteInventory struct {
	Id                string
	CampaignID        string
	AdgroupID         string
	Domain            string
	PublisherId       string
	StartTime         int32
	EndTime           int32
	FixedPrice        float32
	ReferenceCPM      float32
	LastIntervalIndex int32
	TimeInterval      int32
}

func New(ortb IO.OpenRTBRequest) *LiteInventory {
	return &LiteInventory{
		CampaignID:        ortb.Ext.CampaignID,
		AdgroupID:         ortb.Ext.AdgroupID,
		PublisherId:       ortb.SourceString(),
		StartTime:         ortb.Ext.CampaignStartDate,
		EndTime:           ortb.Ext.CampaignEndDate,
		FixedPrice:        ortb.Ext.ReferenceCPM,
		ReferenceCPM:      ortb.Ext.ReferenceCPM,
		LastIntervalIndex: -1,
		Id:                GenerateSessionID(ortb),
		TimeInterval:      Configurations.GetConfig().OptimizerTimeInterval * 60,
	}
}

func NewFromRedis(data map[string]string) *LiteInventory {
	startT, _ := strconv.ParseInt(data["startTime"], 10, 32)
	endT, _ := strconv.ParseInt(data["endTime"], 10, 32)
	fixPrice, _ := strconv.ParseFloat(data["fixedPrice"], 32)
	refC, _ := strconv.ParseFloat(data["refCpm"], 32)
	lastIntervalIndex, _ := strconv.ParseInt(data["lastIntervalIndex"], 10, 32)

	return &LiteInventory{
		Id:                data["Id"],
		CampaignID:        data["campaignId"],
		AdgroupID:         data["adGroupId"],
		PublisherId:       data["publisherId"],
		StartTime:         int32(startT),
		EndTime:           int32(endT),
		FixedPrice:        float32(fixPrice),
		ReferenceCPM:      float32(refC),
		LastIntervalIndex: int32(lastIntervalIndex),
		TimeInterval:      Configurations.GetConfig().OptimizerTimeInterval * 60,
	}
}

func (l *LiteInventory) toMap() *map[string]interface{} {
	result := make(map[string]interface{})
	result["Id"] = l.Id
	result["campaignId"] = l.CampaignID
	result["adGroupId"] = l.AdgroupID
	result["publisherId"] = l.PublisherId
	result["startTime"] = l.StartTime
	result["endTime"] = l.EndTime
	result["fixedPrice"] = l.FixedPrice
	result["refCpm"] = l.ReferenceCPM
	result["lastIntervalIndex"] = l.LastIntervalIndex
	return &result
}
