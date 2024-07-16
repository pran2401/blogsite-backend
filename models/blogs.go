package models

type Blog struct {
	Id          uint   `json:id`
	Title       string `json:title`
	Description string `json:description`
	Image       string `json:image`
	Userid      string `json:userid`
	User        User   `gorm:"foreignKey:Userid"`
}
