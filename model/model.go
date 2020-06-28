package model

import "sync"

type Covid19Stat struct {  //一个省份的数据
	StateName     string `json:"stateName"`
	ConfirmedCase int    `json:"confirmed"`
	ActiveCase    int    `json:"-"`
	Cured         int    `json:"recovered"`
	Death         int    `json:"death"`
}

type Covid19StatMap map[string]Covid19Stat  //多个省份的数据

type Covid19StatMapDateWise map[string]Covid19StatMap  //多个日期多个省份的数据

type dataCacheStruct struct {
	sync.Mutex
	covid19StatsMapDateWise Covid19StatMapDateWise
}

func (d *dataCacheStruct) UpdateCache(data Covid19StatMapDateWise) {
	d.Lock()
	d.covid19StatsMapDateWise = data
	d.Unlock()
}

func (d *dataCacheStruct) GetCache() Covid19StatMapDateWise {
	return d.covid19StatsMapDateWise
}

var DataCache dataCacheStruct = dataCacheStruct{}

type FormatedStatResult struct {
	Data []FormatedStatData `json:"data"`
}

type FormatedStatData struct {
	Date      string                      `json:"date"`
	Confirmed int                         `json:"confirmed"`
	Recovered int                         `json:"recovered"`
	Death     int                         `json:"death"`
	Active    int                         `json:"active"`
	StateWise []StateWiseFormatedStatData `json:"stateWise"`
}

type StateWiseFormatedStatData struct {
	StateName string `json:"stateName"`
	Confirmed int    `json:"confirmed"`
	Recovered int    `json:"recovered"`
	Death     int    `json:"death"`
	Active    int    `json:"active"`
}

const DateFormatPattern = "02-01-2006"
