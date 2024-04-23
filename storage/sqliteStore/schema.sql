CREATE TABLE IF NOT EXISTS strava_auth_token (
    id INTEGER PRIMARY KEY,
    athlete_id int NOT NULL,
    firstname text NOT NULL,
    lastname  text,
    refresh_token text NOT NULL,
    access_token text NOT NULL,
    expires_in int NOT NULL,
    expires_at int NOT NULL
);

CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY,
    username text NOT NULL,
    strava_athlete_id int NOT NULL
);

CREATE TABLE IF NOT EXISTS strava_activities (
    id INTEGER PRIMARY KEY,
    activity_id int NOT NULL,
    user_id int NOT NULL,
    activity_name text NOT NULL,
    distance float NOT NULL,
    moving_time int NOT NULL,
    elapsed_time int NOT NULL,
    total_elevation_gain float NOT NULL,
    activity_type text NOT NULL,
    start_date text NOT NULL,
    start_date_local text NOT NULL,
    timezone text NOT NULL,
    achievement_count int NOT NULL,
    kudos_count int NOT NULL,
    comment_count int NOT NULL,
    athlete_count int NOT NULL,
    photo_count int NOT NULL,
    is_trainer boolean NOT NULL,
    is_commute boolean NOT NULL,
    is_manual boolean NOT NULL,
    is_private boolean NOT NULL,
    is_flagged boolean NOT NULL,
    gear_id INTEGER,
    from_strava boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS strava_gear (
    id INTEGER PRIMARY KEY,
    gear_id text NOT NULL,
    user_id int NOT NULL,
    name text NOT NULL,
    distance float NOT NULL,
    is_primary boolean NOT NULL
);

INSERT OR IGNORE INTO user (id, username, strava_athlete_id) VALUES (1, 'edznux', 111727886);