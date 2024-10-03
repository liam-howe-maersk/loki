package file

import "github.com/prometheus/client_golang/prometheus"

// Metrics hold the set of file-based metrics.
type Metrics struct {
	// Registerer used. May be nil.
	reg prometheus.Registerer

	// File-specific metrics
	readBytes                     *prometheus.GaugeVec
	totalBytes                    *prometheus.GaugeVec
	readLines                     *prometheus.CounterVec
	receivedLineChannel           *prometheus.CounterVec
	skippedLines                  *prometheus.CounterVec
	channelClosureCount           *prometheus.CounterVec
	readFilesGoRoutineExitCounter *prometheus.CounterVec
	encodingFailures              *prometheus.CounterVec
	filesActive                   prometheus.Gauge

	// Manager metrics
	failedTargets *prometheus.CounterVec
	targetsActive prometheus.Gauge
}

// NewMetrics creates a new set of file metrics. If reg is non-nil, the metrics
// will be registered.
func NewMetrics(reg prometheus.Registerer) *Metrics {
	var m Metrics
	m.reg = reg

	m.readBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "promtail",
		Name:      "read_bytes_total",
		Help:      "Number of bytes read.",
	}, []string{"path"})
	m.totalBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "promtail",
		Name:      "file_bytes_total",
		Help:      "Number of bytes total.",
	}, []string{"path"})
	m.readLines = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "promtail",
		Name:      "read_lines_total",
		Help:      "Number of lines read.",
	}, []string{"path", "contains_routing_queries_request"})
	m.receivedLineChannel = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "promtail",
		Name:      "received_lines_channel_total",
		Help:      "Number of lines received via channel.",
	}, []string{"path", "contains_routing_queries_request"})
	m.skippedLines = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "promtail",
		Name:      "skipped_lines_total",
		Help:      "Number of lines skipped.",
	}, []string{"path"})
	m.channelClosureCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "promtail",
		Name:      "channel_closure_total",
		Help:      "Number of times the channel was closed.",
	}, []string{"path"})
	m.readFilesGoRoutineExitCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "promtail",
		Name:      "read_files_go_routine_exit_total",
		Help:      "Number of times the read files go routine exited.",
	}, []string{"path"})

	m.filesActive = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "promtail",
		Name:      "files_active_total",
		Help:      "Number of active files.",
	})

	m.failedTargets = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "promtail",
		Name:      "targets_failed_total",
		Help:      "Number of failed targets.",
	}, []string{"reason"})
	m.targetsActive = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "promtail",
		Name:      "targets_active_total",
		Help:      "Number of active total.",
	})

	if reg != nil {
		reg.MustRegister(
			m.readBytes,
			m.totalBytes,
			m.readLines,
			m.receivedLineChannel,
			m.filesActive,
			m.failedTargets,
			m.targetsActive,
			m.channelClosureCount,
			m.readFilesGoRoutineExitCounter,
			m.skippedLines,
		)
	}

	return &m
}
