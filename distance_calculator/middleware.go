package main

import (
	"time"

	"github.com/joshdstockdale/go-truck-tracker/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next CalcServicer
}

func NewLogMiddleware(next CalcServicer) CalcServicer {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": dist,
		}).Info("Calculate distance:")
	}(time.Now())
	dist, err = m.next.CalculateDistance(data)
	return
}
