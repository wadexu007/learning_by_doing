package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"instrumented-app/controllers"
	"instrumented-app/models"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	serviceName        = os.Getenv("SERVICE_NAME")
	serviceEnvironment = os.Getenv("SERVICE_ENVIRONMENT")
	otelHttpEndpoint   = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
)

func initTracer(rs *resource.Resource) func(context.Context) error {

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithInsecure(),
			otlptracehttp.WithEndpoint(otelHttpEndpoint),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(rs),
		),
	)
	return exporter.Shutdown
}

func initMeter(rs *resource.Resource) func(context.Context) error {

	// init otlpmetrichttp exporter
	meterExp, err := otlpmetrichttp.New(context.Background(),
		otlpmetrichttp.WithInsecure(),
		otlpmetrichttp.WithEndpoint(otelHttpEndpoint),
		// otlpmetrichttp.WithHeaders(map[string]string{
		// 	"X-Scope-OrgID": "tenant-demo",
		// }),
	)
	if err != nil {
		log.Fatal(err)
	}

	// init stdoutmetric exporter just for test
	// stdoutExp, _ := stdoutmetric.New()

	// init prometheus exporter just for test
	// promeExp, err := promexporter.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// init meter provider
	meterProvider := metric.NewMeterProvider(
		metric.WithResource(rs),
		// init provider for otlpmetrichttp exporter
		metric.WithReader(metric.NewPeriodicReader(meterExp, metric.WithInterval(10*time.Second))),
		// init stdoutmetric exporter
		// metric.WithReader(metric.NewPeriodicReader(stdoutExp, metric.WithInterval(10*time.Second))),
		// init provider for prometheus exporter
		// metric.WithReader(promeExp),
	)

	global.SetMeterProvider(meterProvider)

	return meterExp.Shutdown
}

func timeDuration() func(ctx *gin.Context) {
	meter := global.Meter("bookstore-latency-meter")
	httpDurationsHistogram, _ := meter.SyncFloat64().Histogram(
		"bookstore_http_durations_histogram_seconds",
		instrument.WithDescription("Http latency distributions."),
	)

	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		status := strconv.Itoa(ctx.Writer.Status())
		method := ctx.Request.Method
		elapsed := float64(time.Since(start)) / float64(time.Second)
		httpDurationsHistogram.Record(
			ctx,
			elapsed,
			attribute.String("method", method),
			attribute.String("status", status),
		)
	}
}

func main() {

	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("service.environment", serviceEnvironment),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Println("Could not set resources: ", err)
	}

	cleanupTracer := initTracer(resources)
	defer cleanupTracer(context.Background())

	cleanupMeter := initMeter(resources)
	defer cleanupMeter(context.Background())

	r := gin.Default()
	r.Use(otelgin.Middleware(serviceName), timeDuration())
	// Connect to database
	models.ConnectDatabase()

	// register prometheus handler
	// promHandler := promhttp.HandlerFor(
	// 	prometheus.DefaultGatherer,
	// 	promhttp.HandlerOpts{
	// 		EnableOpenMetrics: true,
	// 	},
	// )
	// r.GET("/metrics", func(ctx *gin.Context) {
	// 	promHandler.ServeHTTP(ctx.Writer, ctx.Request)
	// })

	// Routes
	r.GET("/ping", controllers.Status)
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Run the server
	r.Run(":8080")
}
