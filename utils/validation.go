package utils

import "github.com/khan1507017/redis-app/dto"

func DtoValidation(object *dto.RedisObject) bool {
	if object.Key == "" || object.Value == "" {
		return false
	}
	return true
}
