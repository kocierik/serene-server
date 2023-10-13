package models

type Music struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Desc     string `json:"desc"`
	Format   string `json:"format"`
	Duration int    `json:"duration"`
	Path     string `json:"path"`
}
