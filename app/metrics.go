package app

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type MetricsHandler struct {
	logger      *logrus.Logger
	namespace   string
	serviceName string
}

func CreateMetricsHandler(logger *logrus.Logger, namespace string, serviceName string) (*MetricsHandler, error) {
	return &MetricsHandler{
		logger:    logger,
		namespace: namespace,
		serviceName: serviceName,
	}, nil
}

func (m *MetricsHandler) CreateRequestsMetric() (*prometheus.Counter, error) {
	labels := prometheus.Labels{
		"namespace": m.namespace,
		"service_name": m.serviceName,
	}
	requestsCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "num_of_requests",
		Help: "Total number of requests forwarded.",
		ConstLabels: labels,
	})

	if err := prometheus.Register(requestsCounter); err != nil {
		logrus.WithError(err).Error("Metric num_of_requests failed to register")
		return nil, err
	} else {
		logrus.Info("Metric num_of_requests registered successfully")
	}

	return &requestsCounter, nil
}
