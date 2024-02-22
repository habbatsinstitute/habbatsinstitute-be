package item

type Item struct {
	ID          int    `gorm:"type:int(11)"`
	Name        string `gorm:"type:varchar(255)"`
	Cat         string `gorm:"type:varchar(255)"`
	Price       int    `gorm:"type:int(11)"`
	Description string `gorm:"type:varchar(255)"`
	Images      string `gorm:"type:text"`
}
