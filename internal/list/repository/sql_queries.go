package repository

const (
	createList = `INSERT INTO lists (title, type, created_at) 
					VALUES ($1, $2, now()) 
					RETURNING *`
	allLists = `SELECT * FROM lists 
					ORDER BY created_at, updated_at 
					LIMIT $1 
					OFFSET $2`
	getList    = `SELECT * FROM lists WHERE id=$1`
	updateList = `UPDATE lists 
					SET title=COALESCE(NULLIF($1, ''), title), 
						type=COALESCE(NULLIF($2, ''), type),
						updated_at=now() 
					WHERE id=$3
					RETURNING *`
	deleteList = `DELETE FROM lists WHERE id=$1`
)
