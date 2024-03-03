package repository

const (
	createTask = `INSERT INTO tasks (list_id, name, description, status, tag, priority, due_date, created_at)
					VALUES ($1, $2, $3, $4, $5, $6, $7, now())
					RETURNING *`
	updateTask = `UPDATE tasks
					SET list_id=COALESCE(NULLIF($1, ''), list_id),
						name=COALESCE(NULLIF($2, ''), name),
						description=COALESCE(NULLIF($3, ''), description),
						status=COALESCE(NULLIF($4, ''), status),
						tag=COALESCE(NULLIF($5, ''), tag),
						priority=COALESCE(NULLIF($6, ''), priority),
						due_date=COALESCE(NULLIF($7, ''), due_date),
						updated_at=now()
					WHERE id=$8
					RETURNING *`
	getTask    = `SELECT * FROM tasks WHERE id=$1`
	deleteTask = `DELETE FROM tasks where id=$1`
	allTask    = `SELECT * FROM tasks
					ORDER BY created_at, updated_at
					LIMIT $1
					OFFSET $2`
	getTotal = `SELECT COUNT(id) FROM tasks`
)
