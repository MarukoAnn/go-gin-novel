package common

import "gorm.io/gorm"

type NovelInfo struct {
	gorm.Model
	Art_title        string `json:"art_Title"`
	Art_author       string `json:"art_author"`
	Art_sub          string `json:"art_sub"`
	Art_type         string `json:"art_type"`
	Art_link         string `json:"art_link"`
	Art_update_time  string `json:"art_update_time"`
	Art_chapter_name string `json:"art_chapter_name"`
}
