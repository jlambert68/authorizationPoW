package userAuthorizationEngine

import (
	"github.com/sirupsen/logrus"
	"jlambert/authorizationPoW/common_config"
	"log"
	"os"
	"time"
)

/****************************************************/
//  Set up logging specified in common config file
func (userAuthorizationEngineServerObject *userAuthorizationEngineServerObjectStruct) InitLogger(filename string) {
	userAuthorizationEngineServerObject.logger = logrus.StandardLogger()

	switch common_config.LoggingLevel {

	case logrus.DebugLevel:
		log.Println("Using logging level: ", common_config.LoggingLevel)

	case logrus.InfoLevel:
		log.Println("Using logging level: ", common_config.LoggingLevel)

	case logrus.WarnLevel:
		log.Println("Using logging level: ", common_config.LoggingLevel)

	default:
		log.Println("No correct value for debugging-level, this was used: ", common_config.LoggingLevel)
		os.Exit(0)

	}

	logrus.SetLevel(common_config.LoggingLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})

	//If no file then set standard out
	if filename == "" {
		userAuthorizationEngineServerObject.logger.Out = os.Stdout

	} else {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			userAuthorizationEngineServerObject.logger.Out = file
		} else {
			log.Println("Failed to log to file, using default stderr")
		}
	}
}
