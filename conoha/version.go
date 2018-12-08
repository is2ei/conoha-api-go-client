package conoha

// Version represents API version for each endpoint.
type Version struct {
	Status     string       `json:"status,omitempty"`
	Updated    string       `json:"updated,omitempty"`
	MediaTypes []*mediaType `json:"media-types,omitempty"`
	ID         string       `json:"id,omitempty"`
	Links      []*Link      `json:"links,omitempty"`
}
