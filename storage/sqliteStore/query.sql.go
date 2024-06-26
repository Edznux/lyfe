// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqliteStore

import (
	"context"
	"database/sql"
)

const createAuthToken = `-- name: CreateAuthToken :exec
INSERT INTO strava_auth_token (athlete_id, firstname, lastname, refresh_token, access_token, expires_in, expires_at) VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateAuthTokenParams struct {
	AthleteID    int64
	Firstname    string
	Lastname     sql.NullString
	RefreshToken string
	AccessToken  string
	ExpiresIn    int64
	ExpiresAt    int64
}

func (q *Queries) CreateAuthToken(ctx context.Context, arg CreateAuthTokenParams) error {
	_, err := q.db.ExecContext(ctx, createAuthToken,
		arg.AthleteID,
		arg.Firstname,
		arg.Lastname,
		arg.RefreshToken,
		arg.AccessToken,
		arg.ExpiresIn,
		arg.ExpiresAt,
	)
	return err
}

const getAuthToken = `-- name: GetAuthToken :one
SELECT id, athlete_id, firstname, lastname, refresh_token, access_token, expires_in, expires_at FROM strava_auth_token
WHERE athlete_id = ? LIMIT 1
`

func (q *Queries) GetAuthToken(ctx context.Context, athleteID int64) (StravaAuthToken, error) {
	row := q.db.QueryRowContext(ctx, getAuthToken, athleteID)
	var i StravaAuthToken
	err := row.Scan(
		&i.ID,
		&i.AthleteID,
		&i.Firstname,
		&i.Lastname,
		&i.RefreshToken,
		&i.AccessToken,
		&i.ExpiresIn,
		&i.ExpiresAt,
	)
	return i, err
}

const getAuthTokenForUsername = `-- name: GetAuthTokenForUsername :one
SELECT strava_auth_token.id, strava_auth_token.athlete_id, strava_auth_token.firstname, strava_auth_token.lastname, strava_auth_token.refresh_token, strava_auth_token.access_token, strava_auth_token.expires_in, strava_auth_token.expires_at FROM strava_auth_token, user
WHERE user.strava_athlete_id = strava_auth_token.athlete_id
AND user.username = ? LIMIT 1
`

func (q *Queries) GetAuthTokenForUsername(ctx context.Context, username string) (StravaAuthToken, error) {
	row := q.db.QueryRowContext(ctx, getAuthTokenForUsername, username)
	var i StravaAuthToken
	err := row.Scan(
		&i.ID,
		&i.AthleteID,
		&i.Firstname,
		&i.Lastname,
		&i.RefreshToken,
		&i.AccessToken,
		&i.ExpiresIn,
		&i.ExpiresAt,
	)
	return i, err
}
