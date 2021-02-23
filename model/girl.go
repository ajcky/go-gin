package model

type Girl struct {
	ID       int64  `gorm:"column:id" json:"id" form:"id"`
	Nickname string `gorm:"column:nickname" json:"nickname" form:"nickname"`
}

func (g *Girl) TableName() string {
	return "girl"
}
