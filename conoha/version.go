package conoha

type Version struct {
	Status     string      `json:"status"`
	Updated    string      `json:"updated"`
	MediaTypes []MediaType `json:"media-types"`
	Id         string      `json:"id"`
	Links      []Link      `json:"links"`
}
