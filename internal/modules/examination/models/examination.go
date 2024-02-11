package models

import (
	"github.com/Surdy-A/amis_portal/internal/modules/user/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

var Events = map[int]string{
	1: "Qur'an",
	2: "QUIZ",
	3: "KHUTBAH",
	4: "EXHIBITION",
	5: "PICK & TALK",
	6: "CREATIVE WRITING",
	7: "CALLIGRAPHY",
	8: "DEBATE",
	9: "ESSAY WRITING",
}

type Participant struct {
	FirstName  string         `form:"first_name" json:"first_name" binding:"required"`
	LastName   string         `form:"last_name" json:"last_name" binding:"required"`
	MiddleName string         `form:"middle_name" json:"middle_name"`
	Class      string         `form:"class" json:"class" binding:"required"`
	Image      string         `form:"image" json:"image" binding:"required"`
	Events     pq.StringArray `form:"events" json:"events" gorm:"type:text[]"`
}

type Examination struct {
	gorm.Model
	UserID uint
	User   models.User
	//SChool Relationship
	Address            string         `form:"address" json:"address" binding:"required,min=3,max=100"`
	Zone               string         `form:"zone"  json:"zone" binding:"required,min=3,max=100"`
	SchoolPhoneNumber  string         `form:"school_phone_number" json:"school_phone_number"`
	TeacherName        string         `form:"teacher_name" json:"teacher_name" binding:"required,min=3,max=100"`
	TeacherPhoneNumber string         `form:"teacher_phone_number" json:"teacher_phone_number"`
	HeadTeacherName    string         `form:"head_teacher_name" json:"HeadTeacherName"`
	LGA                string         `form:"lga" json:"lga"`
	Paid               bool           `form:"other_association" json:"paid"`
	Participants       Participant    `form:"participants" json:"participants" gorm:"type:text[]"`
	Category           pq.StringArray `form:"category" json:"category" gorm:"type:text[]"`
}
