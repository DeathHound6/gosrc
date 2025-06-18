package v2

type ChallengeModerator struct {
	ChallengeID string             `json:"challengeId"`
	UserID      string             `json:"userId"`
	Level       GameModeratorLevel `json:"level"`
}
