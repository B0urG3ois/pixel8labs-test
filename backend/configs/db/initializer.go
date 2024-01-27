package db

import "context"

func (q *Queries) InitUsersTable(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, queryCreateUsersTable)
	return err
}

func (q *Queries) Index(ctx context.Context) error {
	var err error

	// Users.
	if _, err = q.db.QueryContext(ctx, queryCreateIndexUsersID); err != nil {
		return err
	}
	if _, err = q.db.QueryContext(ctx, queryCreateIndexUsersEmail); err != nil {
		return err
	}
	if _, err = q.db.QueryContext(ctx, queryCreateIndexUsersName); err != nil {
		return err
	}
	if _, err = q.db.QueryContext(ctx, queryCreateIndexUsersUsername); err != nil {
		return err
	}

	return nil
}
