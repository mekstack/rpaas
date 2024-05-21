package main

import (
	"context"
	"flag"
	"log/slog"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"sync"
	"syscall"
	xdsapp "xds_server/internal/app"
)

var (
	nodeID           string
	isDebug          bool
	configPath       string
	cpuProfile       bool
	memProfile       bool
	goroutineProfile bool
)

func init() {
	flag.BoolVar(&isDebug, "debug", false, "Enable xDS server debug logging")

	// Tell Envoy to use this Node ID
	flag.StringVar(&nodeID, "nodeID", "test-id", "Node ID")

	// Path to config file
	flag.StringVar(&configPath, "cfgPath", "config/config.yaml", "path to config file")

	// Flag to CPU profile
	flag.BoolVar(&cpuProfile, "cpu-profile", false, "flag to cpu profile")

	// Flag to MEM profile
	flag.BoolVar(&memProfile, "mem-profile", false, "flag to mem profile")

	// Flag to GOROUTINE profile
	flag.BoolVar(&goroutineProfile, "goroutine-profile", false, "flag to mem profile")
}

func main() {
	flag.Parse()

	//MEM
	if memProfile {
		defer func() {
			f, _ := os.Create("mem.pb.gz")
			runtime.GC()
			_ = pprof.WriteHeapProfile(f)
			_ = f.Close()
		}()
	}

	//CPU
	if cpuProfile {
		f, _ := os.Create("cpu.pb.gz")
		_ = pprof.StartCPUProfile(f)
		defer func() {
			pprof.StopCPUProfile()
			_ = f.Close()
		}()
	}

	//GOROUTINE
	if goroutineProfile {
		go func() {
			fg, _ := os.Create("goroutine.pb.gz")
			_ = pprof.Lookup("goroutine").WriteTo(fg, 0)
		}()
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	logger := Logger()

	a, err := xdsapp.New(ctx, configPath, logger, nodeID)
	if err != nil {
		logger.Error("failed to initialize xds app", "error", err.Error())
		stop()
		os.Exit(1)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		if err = a.Run(ctx); err != nil {
			logger.Error("failed to run or an error occurred  application", slog.String("error", err.Error()))
			stop()
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	logger.Info("received signal to shut down the application")

	a.Stop()

	wg.Wait()
}

func Logger() *slog.Logger {
	var h slog.Handler
	switch isDebug {
	case true:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})
	default:
		h = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelWarn,
			AddSource: true,
		})
	}
	logger := slog.New(h)
	return logger
}
