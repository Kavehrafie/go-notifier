package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{})

	// ✅ load config
	//_, err := config.Load()
	//if err != nil {
	//	logger.Fatal(err)
	//}

}
