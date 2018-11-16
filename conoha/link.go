package conoha

// Link represents the information of the link.
type Link struct {
	Href string `json:"href"`
	Type string `json:"type"`
	Rel  string `json:"rel"`
}
