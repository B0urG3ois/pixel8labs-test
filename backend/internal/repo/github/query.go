package github

const (
	tableDef = `pixel_mst_users`

	querySelectFromBaseRepo = `SELECT
								id,
									avatar_url,
									name,
									username,
									email,
									bio,
									followers,
									following,
									followers_url,
									following_url,
									visitors,
									created_at,
									updated_at FROM ` + tableDef

	querySelectByID    = querySelectFromBaseRepo + ` WHERE id = $1`
	querySelectByEmail = querySelectFromBaseRepo + ` WHERE email = $1`

	queryInsertFromBaseRepo = ` INSERT INTO ` + tableDef + `
								(
									id,
									avatar_url,
									name,
									username,
									email,
									bio,
									followers,
									following,
									followers_url,
									following_url,
									visitors,
									created_at,
									updated_at
								) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`

	queryUpdateVisitorsByID = `UPDATE ` + tableDef + ` 
								SET visitors = $1 WHERE id = $2`
)
