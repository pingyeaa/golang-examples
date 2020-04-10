package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.WithFields(logrus.Fields{
		"name":   "平也",
		"gender": "man",
		"age":    "26",
	}).Info("记录内容")

	_, err := os.Create("runtime/abc.txt")
	logrus.WithError(err).Info("hehe")
}
