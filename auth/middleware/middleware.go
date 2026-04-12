package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	grpcRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_grpc_requests_total",
			Help: "Total number of gRPC requests",
		},
		[]string{"method", "code"},
	)
	grpcDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "auth_grpc_request_duration_seconds",
			Help:    "Duration of gRPC requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)
	grpcInflight = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "auth_grpc_inflight_requests",
			Help: "In-flight gRPC requests",
		},
	)
)

func init() {
	prometheus.MustRegister(grpcRequests, grpcDuration, grpcInflight)
}

// UnaryLoggingInterceptor перехватывает все unary RPC запросы и логирует их
func UnaryMetricsInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		grpcInflight.Inc()
		start := time.Now()

		resp, err := handler(ctx, req)

		duration := time.Since(start).Seconds()
		grpcInflight.Dec()

		code := "OK"
		if err != nil {
			if st, ok := status.FromError(err); ok {
				code = st.Code().String()
			} else {
				code = "Unknown"
			}
			logger.Warn("rpc error",
				zap.String("method", info.FullMethod),
				zap.String("code", code),
				zap.Error(err),
			)
		} else {
			logger.Info("rpc completed",
				zap.String("method", info.FullMethod),
				zap.Duration("duration", time.Since(start)),
			)
		}

		grpcDuration.WithLabelValues(info.FullMethod).Observe(duration)
		grpcRequests.WithLabelValues(info.FullMethod, code).Inc()

		return resp, err
	}
}

// StartMetricsServer starts HTTP server exposing /metrics on given addr (e.g. :9090).
func StartMetricsServer(addr string, logger *zap.Logger) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	go func() {
		logger.Info("starting metrics HTTP server", zap.String("addr", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("metrics server failed", zap.Error(err))
		}
	}()
}
