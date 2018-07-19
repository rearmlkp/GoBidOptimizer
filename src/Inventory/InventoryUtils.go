package Inventory

import (
	"GoBidOptimizer/src/IO"
	"GoBidOptimizer/src/Configurations"
	"GoBidOptimizer/src/Managers"
	"fmt"
)

func GenerateSessionID(o IO.OpenRTBRequest) string {
	prefix := Configurations.GetConfig().OptimzerInventorySessionPrefix
	return prefix + ":" + o.Ext.CampaignID + ":" + o.Ext.AdgroupID
}

func GetInventory(o IO.OpenRTBRequest) *LiteInventory {
	id := GenerateSessionID(o)
	if data, err := Managers.GetRedisClient().HGetAll(id).Result(); err != nil {
		fmt.Println(err)
	} else {
		if len(data) > 0 {
			return NewFromRedis(data)
		}
	}
	return New(o)
}

func UpdateMap(id string, data *LiteInventory) {
	Managers.AutoHmSet(id, *data.toMap())
}
