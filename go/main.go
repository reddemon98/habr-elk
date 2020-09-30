package main

import (
	"go.uber.org/zap/zapcore"
	"math/rand"
	"os"
	"strconv"
	"time"

	"go.uber.org/zap"
)

const logFile = "/logs/app.log"

var requests = []map[string]interface{}{
	{"method": "get", "url": "/price"},
	{"method": "get", "url": "/orders"},
	{"method": "post", "url": "/order"},
}

func main() {
	_, err := os.OpenFile(logFile, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "@timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.OutputPaths = []string{"stdout", logFile}
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()

	i := 0
	for {
		i++
		time.Sleep(time.Second * 3)

		v, status := getRequest()
		if status < 400 {
			sugar.Infow("request #"+strconv.Itoa(i), v...)
		} else {
			sugar.Errorw("request #"+strconv.Itoa(i), v...)
		}
	}
}

func getRequest() ([]interface{}, int) {
	requestNum := rand.Intn(len(requests))
	request := requests[requestNum]

	status := rand.Intn(5)
	if status == 0 {
		request["statusCode"] = 500
		request["message"] = "Internal server error"
	} else if status == 1 {
		request["statusCode"] = 400
		request["message"] = "Bad request"
	} else {
		request["statusCode"] = 200
		request["message"] = "Successfully"
	}

	var result []interface{}

	for k, v := range request {
		result = append(result, k, v)
	}

	return result, request["statusCode"].(int)
}
