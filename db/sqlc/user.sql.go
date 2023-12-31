// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package simpleshop

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, 
  password, 
  salt, 
  role
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, username, password, salt, created_at, role, profile_id
`

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Role     string `json:"role"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.Password,
		arg.Salt,
		arg.Role,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Salt,
		&i.CreatedAt,
		&i.Role,
		&i.ProfileID,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, salt, created_at, role, profile_id FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Salt,
		&i.CreatedAt,
		&i.Role,
		&i.ProfileID,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, username, password, salt, created_at, role, profile_id FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUserParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUser(ctx context.Context, arg ListUserParams) ([]User, error) {
	rows, err := q.query(ctx, q.listUserStmt, listUser, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.Salt,
			&i.CreatedAt,
			&i.Role,
			&i.ProfileID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserRole = `-- name: UpdateUserRole :one
UPDATE users SET role = $2
WHERE id = $1 
RETURNING id, username, password, salt, created_at, role, profile_id
`

type UpdateUserRoleParams struct {
	ID   int64  `json:"id"`
	Role string `json:"role"`
}

func (q *Queries) UpdateUserRole(ctx context.Context, arg UpdateUserRoleParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserRoleStmt, updateUserRole, arg.ID, arg.Role)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Salt,
		&i.CreatedAt,
		&i.Role,
		&i.ProfileID,
	)
	return i, err
}
