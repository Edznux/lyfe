-- name: GetAuthToken :one
SELECT * FROM strava_auth_token
WHERE athlete_id = ? LIMIT 1;

-- name: GetAuthTokenForUsername :one
SELECT strava_auth_token.* FROM strava_auth_token, user
WHERE user.strava_athlete_id = strava_auth_token.athlete_id
AND user.username = ? LIMIT 1;

-- name: CreateAuthToken :exec
INSERT INTO strava_auth_token (athlete_id, firstname, lastname, refresh_token, access_token, expires_in, expires_at) VALUES (?, ?, ?, ?, ?, ?, ?);