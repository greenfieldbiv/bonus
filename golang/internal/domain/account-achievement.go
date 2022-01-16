package domain

type AccountAchievement struct {
	Id            *int64 `json:"id" db:"id"`
	AccountId     *int64 `json:"accountid" db:"accountid"`
	AchievementId *int64 `json:"achievementid" db:"achievementid"`
}
