package client

import (
	"context"

	"github.com/joshdstockdale/go-truck-tracker/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
	GetInvoice(context.Context, int) (*types.Invoice, error)
}
