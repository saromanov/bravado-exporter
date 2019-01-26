package exporter

// Exporter defines interface for
type Exporter interface {
	Update(*Metric) error
}

// Metric provides definition of the metric
type Metric struct {
}
