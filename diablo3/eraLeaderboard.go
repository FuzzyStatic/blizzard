/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:39:09
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-07 13:01:07
 */

package diablo3

import "encoding/json"

// EraLeaderboard structure
type EraLeaderboard struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Row []struct {
		Player []struct {
			Key       string `json:"key"`
			AccountID int    `json:"accountId"`
			Data      []struct {
				ID     string `json:"id"`
				String string `json:"string,omitempty"`
				Number int    `json:"number,omitempty"`
			} `json:"data"`
		} `json:"player"`
		Order int `json:"order"`
		Data  []struct {
			ID        string `json:"id"`
			Number    int    `json:"number,omitempty"`
			Timestamp int    `json:"timestamp,omitempty"`
			String    string `json:"string,omitempty"`
		} `json:"data"`
	} `json:"row"`
	Key   string `json:"key"`
	Title struct {
		EnUS string `json:"en_US"`
		EsMX string `json:"es_MX"`
		PtBR string `json:"pt_BR"`
		DeDE string `json:"de_DE"`
		EnGB string `json:"en_GB"`
		EsES string `json:"es_ES"`
		FrFR string `json:"fr_FR"`
		ItIT string `json:"it_IT"`
		PlPL string `json:"pl_PL"`
		PTPT string `json:"PT_PT"`
		RuRU string `json:"ru_RU"`
		KoKR string `json:"ko_KR"`
		ZhTW string `json:"zh_TW"`
		ZhCN string `json:"zh_CN"`
	} `json:"title"`
	Column []struct {
		ID     string `json:"id"`
		Hidden bool   `json:"hidden"`
		Order  int    `json:"order,omitempty"`
		Label  struct {
			EnUS string `json:"en_US"`
			EsMX string `json:"es_MX"`
			PtBR string `json:"pt_BR"`
			DeDE string `json:"de_DE"`
			EnGB string `json:"en_GB"`
			EsES string `json:"es_ES"`
			FrFR string `json:"fr_FR"`
			ItIT string `json:"it_IT"`
			PlPL string `json:"pl_PL"`
			PTPT string `json:"PT_PT"`
			RuRU string `json:"ru_RU"`
			KoKR string `json:"ko_KR"`
			ZhTW string `json:"zh_TW"`
			ZhCN string `json:"zh_CN"`
		} `json:"label"`
		Type string `json:"type"`
	} `json:"column"`
	LastUpdateTime       string `json:"last_update_time"`
	GeneratedBy          string `json:"generated_by"`
	GreaterRift          bool   `json:"greater_rift"`
	GreaterRiftSoloClass string `json:"greater_rift_solo_class"`
	Era                  int    `json:"era"`
}

// JSON2Struct creates EraLeaderboard structure from JSON byte array
func (e *EraLeaderboard) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, e)
}
