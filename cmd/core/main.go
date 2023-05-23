package main

import (
	"context"
	"metamask-auth/config"
	"metamask-auth/internal/program"
	appCloser "metamask-auth/pkg/app_closer"
	"metamask-auth/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	cfg := config.CreateConfig()
	mainlog := logger.FromEnv("[main]")
	closer := appCloser.InitCloser(mainlog)

	coreProgram := program.NewCoreProgram(cfg, logger.FromEnv("[coreProgram]"), closer)
	go coreProgram.Run()

	select {
	case <-ctx.Done():
		go func() {
			closer.CloseAll()
		}()
		stop()
		os.Exit(0)
	}
}
