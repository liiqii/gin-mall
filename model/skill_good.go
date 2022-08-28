package model

type SkillGoods struct {
	Id         uint `gorm:"primarykey"`
	ProductId  uint `gorm:"not null"`
	BossId     uint `gorm:"not null"`
	Title      string
	Money      float64
	Num        int `gorm:"not null"`
	CustomId   uint
	CustomName string
}

type SkillGood2MQ struct {
	ProductId uint    `json:"product_id"`
	BossId    uint    `json:"boss_id"`
	UserId    uint    `json:"user_id"`
	Money     float64 `json:"money"`
	Key       string  `json:"key"`
}
