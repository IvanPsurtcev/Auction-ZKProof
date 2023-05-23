package contractHttp

import (
	"context"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"metamask-auth/internal/contract"
	"metamask-auth/pkg/fhttp"
)

type gameHandlers struct {
	logger     *zap.Logger
	contractUC contract.UseCase
}

func NewHandlers(logger *zap.Logger, contractUC contract.UseCase) contract.Handlers {
	return &gameHandlers{
		logger:     logger,
		contractUC: contractUC,
	}
}

func (h *gameHandlers) StartAuction() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		err := h.contractUC.StartAuction(context.TODO())
		if err != nil {
			h.logger.Error("Start err:", zap.Error(err))
		}
	}
}

func (h *gameHandlers) EndAuction() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		err := h.contractUC.EndAuction(context.TODO())
		if err != nil {
			h.logger.Error("End err:", zap.Error(err))
		}
	}
}

func (h *gameHandlers) MakeBid() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		var params contract.Bid
		if err := json.Unmarshal(ctx.PostBody(), &params); err != nil {
			h.logger.Error("CreateGame err:", zap.Error(err))
			fhttp.ErrProcessRequest(ctx)
			return
		}
		err := h.contractUC.MakeBid(context.TODO(), &params)
		if err != nil {
			h.logger.Error("MakeBid err:", zap.Error(err))
		}
	}
}

func (h *gameHandlers) BidWin() fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		err := h.contractUC.BidWin(context.TODO())
		if err != nil {
			h.logger.Error("BidWin err:", zap.Error(err))
		}
	}
}
