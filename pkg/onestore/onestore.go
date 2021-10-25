package onestore

import (
	"context"
	"fmt"
	"io"

	"github.com/onestore-ai/onestore/internal/database/metadata"
	"github.com/onestore-ai/onestore/internal/database/offline"
	"github.com/onestore-ai/onestore/internal/database/online"
	"github.com/onestore-ai/onestore/pkg/onestore/types"
)

type OneStore struct {
	online   online.Store
	offline  offline.Store
	metadata metadata.Store
}

func Open(ctx context.Context, opt types.OneStoreOpt) (*OneStore, error) {
	optV2 := opt.ToOneStoreOptV2()

	onlineStore, err := online.Open(optV2.OnlineStoreOpt)
	if err != nil {
		return nil, err
	}
	offlineStore, err := offline.Open(optV2.OfflineStoreOpt)
	if err != nil {
		return nil, err
	}
	metadataStore, err := metadata.Open(optV2.MetaStoreOpt)
	if err != nil {
		return nil, err
	}

	return &OneStore{
		online:   onlineStore,
		offline:  offlineStore,
		metadata: metadataStore,
	}, nil
}

func Create(ctx context.Context, opt types.OneStoreOpt) (*OneStore, error) {
	optV2 := opt.ToOneStoreOptV2()
	if err := metadata.CreateDatabase(ctx, optV2.MetaStoreOpt); err != nil {
		return nil, err
	}

	return Open(ctx, opt)
}

func (s *OneStore) Close() error {
	errs := []error{}

	for _, closer := range []io.Closer{s.online, s.offline, s.metadata} {
		if err := closer.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return fmt.Errorf("failed closing store: %v", errs)
	}
	return nil
}
