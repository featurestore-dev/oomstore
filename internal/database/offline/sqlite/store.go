package sqlite

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/oom-ai/oomstore/internal/database/dbutil"
	"github.com/oom-ai/oomstore/internal/database/offline"
	"github.com/oom-ai/oomstore/internal/database/offline/sqlutil"
	"github.com/oom-ai/oomstore/pkg/oomstore/types"
)

const SQLiteBatchSize = 20

var _ offline.Store = &DB{}

type DB struct {
	*sqlx.DB
}

func (db *DB) Ping(ctx context.Context) error {
	return db.PingContext(ctx)
}

func Open(option *types.SQLiteOpt) (*DB, error) {
	db, err := dbutil.OpenSQLite(option.DBFile)
	return &DB{db}, err
}

func (db *DB) Import(ctx context.Context, opt offline.ImportOpt) (int64, error) {
	return sqlutil.Import(ctx, db.DB, opt, dbutil.LoadDataFromSource(types.SQLite, SQLiteBatchSize), types.SQLite)
}

func (db *DB) Export(ctx context.Context, opt offline.ExportOpt) (<-chan types.ExportRecord, <-chan error) {
	return sqlutil.Export(ctx, db.DB, opt, types.SQLite)
}

func (db *DB) Join(ctx context.Context, opt offline.JoinOpt) (*types.JoinResult, error) {
	return sqlutil.Join(ctx, db.DB, opt, types.SQLite)
}

func (db *DB) TypeTag(dbType string) (string, error) {
	return TypeTag(dbType)
}

func (db *DB) TableSchema(ctx context.Context, tableName string) (*types.DataTableSchema, error) {
	return nil, fmt.Errorf("not implemented")
}
