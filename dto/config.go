package dto

type Configuration struct {
	Mongo MongoConfiguration `json:"mongo"`
}

type MongoConfiguration struct {
	Url          string `json:"url"`
	UserName     string `json:"userName"`
	Password     string `json:"password"`
	AuthDatabase string `json:"authDatabase"`
	Database     string `json:"database"`
}
