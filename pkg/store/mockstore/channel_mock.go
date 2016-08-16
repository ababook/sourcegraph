// generated by gen-mocks; DO NOT EDIT

package mockstore

import (
	"context"

	"sourcegraph.com/sourcegraph/sourcegraph/pkg/store"
)

type Channel struct {
	Listen_ func(ctx context.Context, channel string) (ch <-chan store.ChannelNotification, unlisten func(), err error)
	Notify_ func(ctx context.Context, channel, payload string) error
}

func (s *Channel) Listen(ctx context.Context, channel string) (ch <-chan store.ChannelNotification, unlisten func(), err error) {
	return s.Listen_(ctx, channel)
}

func (s *Channel) Notify(ctx context.Context, channel, payload string) error {
	return s.Notify_(ctx, channel, payload)
}

var _ store.Channel = (*Channel)(nil)
