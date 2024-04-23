package strava

import "time"

type Map struct {
	ID              string `json:"id"`
	SummaryPolyline string `json:"summary_polyline"`
	ResourceState   int    `json:"resource_state"`
}

type StravaOAuthResponse struct {
	TokenType    string  `json:"token_type"`
	ExpiresIn    int64   `json:"expires_in"`
	ExpiresAt    int64   `json:"expires_at"`
	RefreshToken string  `json:"refresh_token"`
	AccessToken  string  `json:"access_token"`
	Athlete      Athlete `json:"athlete"`
}

type Athlete struct {
	ID            int64     `json:"id"`
	Username      string    `json:"username"`
	ResourceState int       `json:"resource_state"`
	Firstname     string    `json:"firstname"`
	Lastname      string    `json:"lastname"`
	Bio           string    `json:"bio"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	Sex           string    `json:"sex"`
	Premium       bool      `json:"premium"`
	Summit        bool      `json:"summit"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	BadgeTypeID   int       `json:"badge_type_id"`
	Weight        any       `json:"weight"`
	ProfileMedium string    `json:"profile_medium"`
	Profile       string    `json:"profile"`
	Friend        any       `json:"friend"`
	Follower      any       `json:"follower"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Activity struct {
	ResourceState              int       `json:"resource_state"`
	Athlete                    Athlete   `json:"athlete"`
	Name                       string    `json:"name"`
	Distance                   float64   `json:"distance"`
	MovingTime                 int       `json:"moving_time"`
	ElapsedTime                int       `json:"elapsed_time"`
	TotalElevationGain         float64   `json:"total_elevation_gain"`
	Type                       string    `json:"type"`
	SportType                  string    `json:"sport_type"`
	WorkoutType                int       `json:"workout_type"`
	ID                         int64     `json:"id"`
	StartDate                  time.Time `json:"start_date"`
	StartDateLocal             time.Time `json:"start_date_local"`
	Timezone                   string    `json:"timezone"`
	UtcOffset                  float64   `json:"utc_offset"`
	LocationCity               any       `json:"location_city"`
	LocationState              any       `json:"location_state"`
	LocationCountry            any       `json:"location_country"`
	AchievementCount           int       `json:"achievement_count"`
	KudosCount                 int       `json:"kudos_count"`
	CommentCount               int       `json:"comment_count"`
	AthleteCount               int       `json:"athlete_count"`
	PhotoCount                 int       `json:"photo_count"`
	Map                        Map       `json:"map"`
	Trainer                    bool      `json:"trainer"`
	Commute                    bool      `json:"commute"`
	Manual                     bool      `json:"manual"`
	Private                    bool      `json:"private"`
	Visibility                 string    `json:"visibility"`
	Flagged                    bool      `json:"flagged"`
	GearID                     any       `json:"gear_id"`
	StartLatlng                []LatLng  `json:"start_latlng"`
	EndLatlng                  []LatLng  `json:"end_latlng"`
	AverageSpeed               float64   `json:"average_speed"`
	MaxSpeed                   float64   `json:"max_speed"`
	AverageCadence             float64   `json:"average_cadence"`
	AverageWatts               float64   `json:"average_watts"`
	MaxWatts                   int       `json:"max_watts"`
	WeightedAverageWatts       int       `json:"weighted_average_watts"`
	Kilojoules                 float64   `json:"kilojoules"`
	DeviceWatts                bool      `json:"device_watts"`
	HasHeartrate               bool      `json:"has_heartrate"`
	AverageHeartrate           float64   `json:"average_heartrate"`
	MaxHeartrate               float64   `json:"max_heartrate"`
	HeartrateOptOut            bool      `json:"heartrate_opt_out"`
	DisplayHideHeartrateOption bool      `json:"display_hide_heartrate_option"`
	ElevHigh                   float64   `json:"elev_high"`
	ElevLow                    float64   `json:"elev_low"`
	UploadID                   int64     `json:"upload_id"`
	UploadIDStr                string    `json:"upload_id_str"`
	ExternalID                 string    `json:"external_id"`
	FromAcceptedTag            bool      `json:"from_accepted_tag"`
	PrCount                    int       `json:"pr_count"`
	TotalPhotoCount            int       `json:"total_photo_count"`
	HasKudoed                  bool      `json:"has_kudoed"`
	SufferScore                float64   `json:"suffer_score"`
}
