package Inventory

import "GoBidOptimizer/src/Configurations"

type LiteInventoryPipeline struct {
}

func GenerateConversionProbKey(exchangeName string, exchangeUserId string) string {
	if exchangeName == "" {
		return Configurations.GetConfig().ConversionProbabilityPrefix + ":" + exchangeUserId
	}
	return Configurations.GetConfig().ConversionProbabilityPrefix + ":" + exchangeName + ":" + exchangeUserId
}
