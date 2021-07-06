package dto

type RedisObject struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
