package sqlite

import (
	"context"

	"github.com/oom-ai/oomstore/pkg/errdefs"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"

	"github.com/oom-ai/oomstore/internal/database/metadata"
)

func createRevision(ctx context.Context, sqlxCtx metadata.SqlxContext, opt metadata.CreateRevisionOpt) (int, error) {
	var snapshotTable, cdcTable string
	if opt.SnapshotTable != nil {
		snapshotTable = *opt.SnapshotTable
	}
	if opt.CdcTable != nil {
		cdcTable = *opt.CdcTable
	}

	insertQuery := "INSERT INTO feature_group_revision(group_id, revision, snapshot_table, cdc_table, anchored, description) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := sqlxCtx.ExecContext(ctx, sqlxCtx.Rebind(insertQuery), opt.GroupID, opt.Revision, snapshotTable, cdcTable, opt.Anchored, opt.Description)
	if err != nil {
		if sqliteErr, ok := err.(*sqlite.Error); ok {
			if sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
				return 0, errdefs.Errorf("revision already exists: groupID=%d, revision=%d", opt.GroupID, opt.Revision)
			}
		}
		return 0, errdefs.WithStack(err)
	}
	revisionID, err := res.LastInsertId()
	if err != nil {
		return 0, errdefs.WithStack(err)
	}

	return int(revisionID), nil
}
