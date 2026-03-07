package wowgd

// ToyIndex contains an index of toys
type ToyIndex struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Toys []struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"toys"`
}

// Toy contains toy data
type Toy struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	ID   int `json:"id"`
	Item struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"item"`
	Source struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"source"`
	SourceDescription string `json:"source_description,omitempty"`
	Media             struct {
		Key struct {
			Href string `json:"href"`
		} `json:"key"`
		ID int `json:"id"`
	} `json:"media"`
}
