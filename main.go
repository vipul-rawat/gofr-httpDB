package httpDB

import "gofr.dev/pkg/gofr/datasource"

type httpDB struct {
	logger  Logger
	metrics Metrics
}

var DB = func(conf datasource.Config, logger datasource.Logger, metrics datasource.Metrics) datasource.HttpDb {
	return &httpDB{
		logger:  logger,
		metrics: metrics,
	}
}

func (h *httpDB) Get(key string) string {
	return ""
}

func (h *httpDB) Set(key, value string) {
	return
}
