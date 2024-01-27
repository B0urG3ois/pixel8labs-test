package db

const (
	// Users Table.
	queryCreateUsersTable = `
		CREATE TABLE IF NOT EXISTS pixel_mst_users (
		    id INTEGER PRIMARY KEY,
		    avatar_url VARCHAR NOT NULL,
		    name VARCHAR NOT NULL,
		    username VARCHAR NOT NULL,
		    email VARCHAR NOT NULL,
		    bio VARCHAR,
		    followers INTEGER NOT NULL DEFAULT 0,
		    following INTEGER NOT NULL DEFAULT 0,
			followers_url VARCHAR,
			following_url VARCHAR,
			visitors INTEGER NOT NULL DEFAULT 0,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	queryCreateIndexUsersID = `
		CREATE INDEX IF NOT EXISTS pixel_mst_users_idx_id ON pixel_mst_users (id);
    `

	queryCreateIndexUsersEmail = `
		CREATE INDEX IF NOT EXISTS pixel_mst_users_idx_email ON pixel_mst_users (email);
    `

	queryCreateIndexUsersName = `
		CREATE INDEX IF NOT EXISTS pixel_mst_users_idx_name ON pixel_mst_users (name);
    `

	queryCreateIndexUsersUsername = `
		CREATE INDEX IF NOT EXISTS pixel_mst_users_idx_username ON pixel_mst_users (username);
    `
)
