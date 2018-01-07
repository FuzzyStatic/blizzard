/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-07 12:38:35
 * @Last Modified by:   FuzzyStatic
 * @Last Modified time: 2018-01-07 12:38:35
 */

package diablo3

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

// SeasonAchievementPoints structure
type SeasonAchievementPoints struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Row []struct {
		Row
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
	LastUpdateTime    string `json:"last_update_time"`
	GeneratedBy       string `json:"generated_by"`
	AchievementPoints bool   `json:"achievement_points"`
	Season            int    `json:"season"`
}

// SeasonRift structure
type SeasonRift struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Row []struct {
		Row
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
	Hardcore             bool   `json:"hardcore"`
	GreaterRift          bool   `json:"greater_rift"`
	GreaterRiftSoloClass string `json:"greater_rift_solo_class"`
	Season               int    `json:"season"`
}

// Row player, hero, and rank data for leaderboard
type Row struct {
	Player []struct {
		Player
	} `json:"player"`
	Order int `json:"order"`
	Data  []struct {
		Data
	} `json:"data"`
}

// Player rank data for leaderboard player
type Player struct {
	Key       string `json:"key"`
	AccountID int    `json:"accountId"`
	Data      []struct {
		Data
	} `json:"data"`
}

// Data rank data for leaderboard player
type Data struct {
	ID        string `json:"id"`
	Number    int    `json:"number,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
	String    string `json:"string,omitempty"`
}

// JSON2Struct creates SeasonAchievementPoints structure from JSON byte array
func (s *SeasonAchievementPoints) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, s)
}

// JSON2Struct creates SeasonRift structure from JSON byte array
func (s *SeasonRift) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, s)
}

// GetSeasonRiftRowFromRank get player information from season leaderboard via rank
func (s *SeasonRift) GetSeasonRiftRowFromRank(rankNum int) (*Row, error) {
	for i := range s.Row {
		v := reflect.ValueOf(s.Row[i])

		for j := 0; j < v.NumField(); j++ {
			row := v.Field(j).Interface().(Row)
			data, _ := row.GetSeasonRowData()

			if data.ID == "Rank" && data.Number == rankNum {
				return &row, nil
			}
		}
	}

	return nil, errors.New("no player at rank " + strconv.Itoa(rankNum) + " found")
}

// GetSeasonRowData get data for row from season leaderboard
func (r *Row) GetSeasonRowData() (*Data, error) {
	for i := range (*r).Data {
		v := reflect.ValueOf((*r).Data[i])

		for j := 0; j < v.NumField(); j++ {
			data := v.Field(j).Interface().(Data)
			return &data, nil
		}
	}

	return nil, errors.New("no row data found")
}

// GetSeasonRowPlayer get player for row from season leaderboard
func (r *Row) GetSeasonRowPlayer() (*Player, error) {
	for i := range (*r).Player {
		v := reflect.ValueOf((*r).Player[i])

		for j := 0; j < v.NumField(); j++ {
			player := v.Field(j).Interface().(Player)
			return &player, nil
		}
	}

	return nil, errors.New("no row player found")
}

// GetHeroBattleTagFromSeasonRowPlayer get hero battle tag from row player from season
// leaderboard
func (p *Player) GetHeroBattleTagFromSeasonRowPlayer() (*string, error) {
	for _, v := range (*p).Data {
		if v.ID == "HeroBattleTag" {
			return &v.String, nil
		}
	}

	return nil, errors.New("no HeroBattleTag found")
}

// GetHeroIDFromSeasonRowPlayer get hero ID from row player from season leaderboard
func (p *Player) GetHeroIDFromSeasonRowPlayer() (*int, error) {
	for _, v := range (*p).Data {
		if v.ID == "HeroId" {
			return &v.Number, nil
		}
	}

	return nil, errors.New("no HeroId found")
}

// GetHeroBattleTagAndIDFromSeasonRow get hero battle tag ID from row from season
// leaderboard
func (r *Row) GetHeroBattleTagAndIDFromSeasonRow() (*string, *int, error) {
	var (
		player    *Player
		battleTag *string
		heroID    *int
		err       error
	)

	if player, err = r.GetSeasonRowPlayer(); err != nil {
		return nil, nil, err
	}

	if battleTag, err = player.GetHeroBattleTagFromSeasonRowPlayer(); err != nil {
		return nil, nil, err
	}

	if heroID, err = player.GetHeroIDFromSeasonRowPlayer(); err != nil {
		return nil, nil, err
	}

	return battleTag, heroID, nil
}
