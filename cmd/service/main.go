package main

import (
	"flag"
	"fmt"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/nnkolesnikov/transport-learning/pkg/service"
	"github.com/nnkolesnikov/transport-learning/pkg/service/httpserver"
)

const (
	serviceVersion = "dev"
)

type configuration struct {
	Port               string `envconfig:"PORT" required:"true" default:"8080"`
	MaxRequestBodySize int    `envconfig:"MAX_REQUEST_BODY_SIZE" default:"10485760"` // 10 MB
	Debug              bool   `envconfig:"DEBUG" default:"false"`

	ReadTimeout time.Duration `envconnfig:"READ_TIMEOUT" default:"1s"`
}

func main() {
	printVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *printVersion {
		fmt.Println(serviceVersion)
		os.Exit(0)
	}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	_ = level.Info(logger).Log("msg", "initializing", "version", serviceVersion)

	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		_ = level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}
	if !cfg.Debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	svc, err := service.NewService()
	if err != nil {
		level.Error(logger).Log("msg", "failed to create service", "err", err)
		os.Exit(1)
	}
	svc = service.NewLoggingMiddleware(logger, svc)

	router := httpserver.NewPreparedServer(svc)
	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
	fasthttpServer := &fasthttp.Server{
		Handler:            router.Handler,
		MaxRequestBodySize: cfg.MaxRequestBodySize,
		ReadTimeout:        cfg.ReadTimeout,
	}

	go func() {
		_ = level.Info(logger).Log("msg", "starting http server", "port", cfg.Port)
		if err := fasthttpServer.ListenAndServe(":" + cfg.Port); err != nil {
			_ = level.Error(logger).Log("msg", "server run failure", "err", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		_ = level.Info(logger).Log("msg", "received signal, exiting", "signal", sig)

		if err := fasthttpServer.Shutdown(); err != nil {
			_ = level.Error(logger).Log("msg", "server shutdown failure", "err", err)
		}

		_ = level.Info(logger).Log("msg", "goodbye")
	}(<-c)
}
