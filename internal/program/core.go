package program

import (
	fhttpRouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"metamask-auth/config"
	contractHttp "metamask-auth/internal/contract/delivery/http"
	contractUseCase "metamask-auth/internal/contract/usecase"
	"metamask-auth/internal/middleware"
	appCloser "metamask-auth/pkg/app_closer"
)

type coreProgram struct {
	cfg       *config.Config
	logger    *zap.Logger
	closer    appCloser.Closer
	appRouter *fhttpRouter.Router
}

func NewCoreProgram(cfg *config.Config, logger *zap.Logger, closer appCloser.Closer) Program {
	contractUC := contractUseCase.NewUseCase(logger, cfg)

	contractHandlers := contractHttp.NewHandlers(logger, contractUC)
	mw := middleware.NewMw(logger)

	appRouter := fhttpRouter.New()

	contractGroup := appRouter.Group("/contract")

	contractHttp.InitRouter(contractGroup, contractHandlers, mw)

	return &coreProgram{
		cfg:       cfg,
		logger:    logger,
		closer:    closer,
		appRouter: appRouter,
	}
}

func (core *coreProgram) Run() {
	core.logger.Info("starting coreProgram")
	server := fasthttp.Server{
		Handler: middleware.CORS(core.appRouter.Handler),
	}

	core.closer.AddCloser(func() {
		if err := server.Shutdown(); err != nil {
			core.logger.Panic("cannot stop coreProgram", zap.Error(err))
		}
	}, "coreProgram")

	listenAddr := core.cfg.System.Host + ":" + core.cfg.System.Port
	if err := server.ListenAndServe(listenAddr); err != nil {
		core.logger.Panic("Error in running server", zap.Error(err))
	}
}
