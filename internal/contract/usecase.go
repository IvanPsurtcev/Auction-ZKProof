package contract

import "context"

type UseCase interface {
	StartAuction(ctx context.Context) error
	EndAuction(ctx context.Context) error
	MakeBid(ctx context.Context, bid *Bid) error
	BidWin(ctx context.Context) error
	// SetStatus(ctx context.Context) error
}
