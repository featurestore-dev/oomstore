package onestore

import (
	"context"
	"fmt"
	"io"

	"github.com/onestore-ai/onestore/internal/database"
	"github.com/onestore-ai/onestore/internal/database/metadata"
	"github.com/onestore-ai/onestore/internal/database/offline"
	"github.com/onestore-ai/onestore/internal/database/online"
	"github.com/onestore-ai/onestore/pkg/onestore/types"
)

type OneStore struct {
	db *database.DB

	online   online.Store
	offline  offline.Store
	metadata metadata.Store
}

func Open(ctx context.Context, opt types.OneStoreOpt) (*OneStore, error) {
	optV2 := opt.ToOneStoreOptV2()

	db, err := database.Open(toDatabaseOption(&opt))
	if err != nil {
		return nil, err
	}

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
		db:       db,
		online:   onlineStore,
		offline:  offlineStore,
		metadata: metadataStore,
	}, nil
}

func Create(ctx context.Context, opt types.OneStoreOpt) (*OneStore, error) {
	if err := database.CreateDatabase(ctx, toDatabaseOption(&opt)); err != nil {
		return nil, err
	}

	return Open(ctx, opt)
}

func (s *OneStore) Close() error {
	errs := []error{}

	for _, closer := range []io.Closer{s.db, s.online, s.offline, s.metadata} {
		if err := closer.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return fmt.Errorf("failed closing store: %v", errs)
	}
	return nil
}

func toDatabaseOption(opt *types.OneStoreOpt) database.Option {
	return database.Option{
		Host:   opt.Host,
		Port:   opt.Port,
		User:   opt.User,
		Pass:   opt.Pass,
		DbName: opt.Workspace,
	}
}
