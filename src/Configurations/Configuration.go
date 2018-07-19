package Configurations

import (
	"io/ioutil"
	"github.com/json-iterator/go"
	"sync"
)

type configuration struct {
	RetargetingBidCap              float32 `json:"RETARGETING_BID_CAP"`
	MaxPhi                         float32 `json:"MAX_PHI"`
	MinPhi                         float32 `json:"MIN_PHI"`
	MaxPacing                      float32 `json:"MAX_PACING"`
	RedisHost                      string  `json:"REDIS_HOST"`
	RedisPort                      string  `json:"REDIS_PORT"`
	RedisPassword                  string  `json:"REDIS_PASSWORD"`
	ConversionProbabilityPrefix    string  `json:"CONVERSION_PROBABILITY_PREFIX"`
	ConversionDiscountFactor       float32 `json:"CONVERSION_DISCOUNT_FACTOR"`
	OptimzerInventorySessionPrefix string  `json:"OPTIMZER_INVENTORY_SESSION_PREFIX"`
	OptimizerTimeInterval          int32   `json:"OPTIMIZER_TIME_INTERVAL"`
}

var instance *configuration
var once sync.Once

func GetConfig() *configuration {
	if instance == nil {
		file, _ := ioutil.ReadFile("properties.json")
		once.Do(func() {
			instance = &configuration{}
		})
		jsoniter.Unmarshal(file, instance)
	}
	return instance
}
