package mysql

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/oom-ai/oomstore/internal/database/dbutil"
	"github.com/oom-ai/oomstore/internal/database/metadata"
)

func createRevision(ctx context.Context, sqlxCtx metadata.SqlxContext, opt metadata.CreateRevisionOpt) (int, string, error) {
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
		if er, ok := err.(*mysql.MySQLError); ok {
			if er.Number == ER_DUP_ENTRY {
				return 0, "", errors.Errorf("revision already exists: groupID=%d, revision=%d", opt.GroupID, opt.Revision)
			}
		}
		return 0, "", errors.WithStack(err)
	}
	revisionID, err := res.LastInsertId()
	if err != nil {
		return 0, "", errors.WithStack(err)
	}

	if opt.SnapshotTable == nil {
		updateQuery := "UPDATE feature_group_revision SET snapshot_table = ? WHERE id = ?"
		snapshotTable = dbutil.OfflineBatchTableName(opt.GroupID, revisionID)
		result, err := sqlxCtx.ExecContext(ctx, sqlxCtx.Rebind(updateQuery), snapshotTable, revisionID)
		if err != nil {
			return 0, "", errors.WithStack(err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return 0, "", errors.WithStack(err)
		}
		if rowsAffected != 1 {
			return 0, "", errors.Errorf("failed to update revision %d: revision not found", revisionID)
		}
	}
	return int(revisionID), snapshotTable, nil
}
