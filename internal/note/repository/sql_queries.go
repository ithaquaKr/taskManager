package repository

const (
	createNote = `INSERT INTO notes (list_id, name, content, created_at)
					VALUE ($1, $2, $3, now())
					RETURNING *`
	allNotes = `SELECT * FROM notes 
					ORDER BY created_at, updated_at
					LIMIT $1
					OFFSET $2`
	getNote    = `SELECT * FROM notes WHERE id=$1`
	updateNote = `UPDATE notes
					SET list_id=COALESCE(NULLIF($1, ''), list_id),
						name=COALESCE(NULLIF($2, ''), name),
						content=COALESCE(NULLIF($3, ''), content),
						updated_at=now()
					WHERE id=$4
					RETURNING *`
	deleteNote = `DELETE FROM notes where id=$1`
)
