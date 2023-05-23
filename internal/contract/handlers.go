package contract

import "github.com/valyala/fasthttp"

type Handlers interface {
	StartAuction() fasthttp.RequestHandler
	EndAuction() fasthttp.RequestHandler
	MakeBid() fasthttp.RequestHandler
	BidWin() fasthttp.RequestHandler
}
