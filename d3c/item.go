package d3c

// Item structure
type Item struct {
	ID             string `json:"id"`
	Slug           string `json:"slug"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	TooltipParams  string `json:"tooltipParams"`
	RequiredLevel  int    `json:"requiredLevel"`
	StackSizeMax   int    `json:"stackSizeMax"`
	AccountBound   bool   `json:"accountBound"`
	FlavorText     string `json:"flavorText"`
	FlavorTextHTML string `json:"flavorTextHtml"`
	TypeName       string `json:"typeName"`
	Type           struct {
		TwoHanded bool   `json:"twoHanded"`
		ID        string `json:"id"`
	} `json:"type"`
	Damage                 string   `json:"damage"`
	Dps                    string   `json:"dps"`
	DamageHTML             string   `json:"damageHtml"`
	Color                  string   `json:"color"`
	IsSeasonRequiredToDrop bool     `json:"isSeasonRequiredToDrop"`
	SeasonRequiredToDrop   int      `json:"seasonRequiredToDrop"`
	Slots                  []string `json:"slots"`
	Attributes             struct {
		Primary []struct {
			TextHTML string `json:"textHtml"`
			Text     string `json:"text"`
		} `json:"primary"`
		Secondary []struct {
			TextHTML string `json:"textHtml"`
			Text     string `json:"text"`
		} `json:"secondary"`
		Other []interface{} `json:"other"`
	} `json:"attributes"`
	RandomAffixes []struct {
		OneOf []struct {
			TextHTML string `json:"textHtml"`
			Text     string `json:"text"`
		} `json:"oneOf"`
	} `json:"randomAffixes"`
	SetItems []string `json:"setItems"`
}
