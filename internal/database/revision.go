package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/onestore-ai/onestore/pkg/onestore/types"
)

func (db *DB) ListRevision(ctx context.Context, groupName *string) ([]*types.Revision, error) {
	query := "SELECT * FROM feature_group_revision"
	var cond []interface{}
	if groupName != nil {
		query += " WHERE group_name = $1"
		cond = append(cond, *groupName)
	}
	revisions := make([]*types.Revision, 0)

	if err := db.SelectContext(ctx, &revisions, query, cond...); err != nil {
		return nil, err
	}
	return revisions, nil
}

func (db *DB) GetRevision(ctx context.Context, groupName string, revision int64) (*types.Revision, error) {
	query := "SELECT * FROM feature_group_revision WHERE group_name = $1 and revision = $2"
	var rs types.Revision
	if err := db.GetContext(ctx, rs, query, groupName, revision); err != nil {
		return nil, err
	}
	return &rs, nil
}

func InsertRevision(ctx context.Context, tx *sqlx.Tx, groupName string, revision int64, dataTable string, description string) error {
	cmd := "INSERT INTO feature_group_revision(group_name, revision, data_table, description) VALUES ($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, cmd, groupName, revision, dataTable, description)
	return err
}
