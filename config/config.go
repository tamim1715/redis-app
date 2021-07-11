package config

import (
	"errors"
	"os"
	"strconv"
)

// ServerPort and RedisPort are embedded into actual code, as they are part of the codebase
//If you run Redis in a different port, please change the default value
//If you want to use a different port for the application, change the server port. Please also make changes to docker file (expose the same port)
const ServerPort = "4040"
const RedisPort = "6379"

//envVars
var RedisPassword string
var RedisMasterEndpoint string
var RedisSlaveEndpoints [100]string
var RedisSlaveCount int
var slaveCountTemp string

//temp variable
var boolVal bool

func InitEnvironmentVariables() error {

	//DB CREDENTIALS + Cluster Endpoint + Others
	RedisPassword, boolVal = os.LookupEnv("REDIS_PASSWORD")
	if boolVal == false {
		return errors.New("REDIS_PASSWORD not found in envVars")
	}
	RedisMasterEndpoint, boolVal = os.LookupEnv("MASTER_ENDPOINT")
	if boolVal == false {
		return errors.New("MASTER_ENDPOINT not found in envVars")
	}
	err := initSlaveEndpoints()
	if err != nil {
		return err
	}
	return nil
}
func initSlaveEndpoints() error {
	slaveCountTemp, boolVal = os.LookupEnv("REDIS_SLAVE_COUNT")
	if boolVal == true {
		RedisSlaveCount, err := strconv.Atoi(slaveCountTemp)
		if err != nil {
			return err
		}
		if RedisSlaveCount < 0 || RedisSlaveCount > 50 {
			return errors.New("invalid slave number: " + slaveCountTemp)
		}
	} else {
		RedisSlaveCount = 0
		return nil
	}
	for i := 0; i < RedisSlaveCount; i++ {
		RedisSlaveEndpoints[i], boolVal = os.LookupEnv("SLAVE_ENDPOINT_" + strconv.Itoa(i))
		if boolVal == false {
			return errors.New("SLAVE_ENDPOINT_" + strconv.Itoa(i) + " not found in envVars")
		}
	}
	return nil
}
