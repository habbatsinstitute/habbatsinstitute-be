package realtimechat

type Message struct {
	UserID   int    `gorm:"column:user_id"`
	UserTo   int    `gorm:"column:user_to"`
	RoomID   int    `gorm:"column:room_id"`
	Messages string `gorm:"column:messages"`
	Images   string `gorm:"column:images"`
}