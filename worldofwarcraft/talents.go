/*
 * @Author: Allen Flickinger (allen.flickinger@gmail.com)
 * @Date: 2018-01-15 18:54:28
 * @Last Modified by: FuzzyStatic
 * @Last Modified time: 2018-01-15 18:55:07
 */

package worldofwarcraft

import "encoding/json"

// Talents structure
type Talents struct {
	Num1 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"1"`
	Num2 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				Range       string `json:"range"`
				PowerCost   string `json:"powerCost"`
				CastTime    string `json:"castTime"`
				Cooldown    string `json:"cooldown"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"2"`
	Num3 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		PetSpecs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"petSpecs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"3"`
	Num4 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"4"`
	Num5 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"5"`
	Num6 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"6"`
	Num7 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				Range       string `json:"range"`
				CastTime    string `json:"castTime"`
				Cooldown    string `json:"cooldown"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"7"`
	Num8 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"8"`
	Num9 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"9"`
	Num10 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				Range       string `json:"range"`
				CastTime    string `json:"castTime"`
				Cooldown    string `json:"cooldown"`
			} `json:"spell"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"10"`
	Num11 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"11"`
	Num12 struct {
		Specs []struct {
			Name            string `json:"name"`
			Role            string `json:"role"`
			BackgroundImage string `json:"backgroundImage"`
			Icon            string `json:"icon"`
			Description     string `json:"description"`
			Order           int    `json:"order"`
		} `json:"specs"`
		Talents [][][]struct {
			Tier   int `json:"tier"`
			Column int `json:"column"`
			Spell  struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Icon        string `json:"icon"`
				Description string `json:"description"`
				CastTime    string `json:"castTime"`
			} `json:"spell"`
			Spec struct {
				Name            string `json:"name"`
				Role            string `json:"role"`
				BackgroundImage string `json:"backgroundImage"`
				Icon            string `json:"icon"`
				Description     string `json:"description"`
				Order           int    `json:"order"`
			} `json:"spec"`
		} `json:"talents"`
		Class string `json:"class"`
	} `json:"12"`
}

// JSON2Struct creates Talents structure from JSON byte array
func (t *Talents) JSON2Struct(b *[]byte) error {
	return json.Unmarshal(*b, t)
}
