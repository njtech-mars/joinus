package models

import "time"

type EmailRecord struct {
	Sent_At time.Time `json:"sent_at"`
	Status  string    `json:"status"`
}

type Applicant struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Submitted_At  time.Time `json:"submitted_at"`
	Email_Status  string    `json:"email_status"`
	Name          string    `json:"name" validate:"required,min=2,max=20"`
	Gender        string    `json:"gender" validate:"required,oneof=男 女"`
	Email         string    `json:"email" validate:"required,email,max=320"`
	QQ            string    `json:"qq" validate:"required,min=5,max=11,number"`
	Student_ID    string    `json:"student_id" validate:"required,len=12,number"`
	College       string    `json:"college" validate:"required,max=50"`
	Major         string    `json:"major" validate:"required,max=50"`
	First_Choice  string    `json:"first_choice" validate:"required,len=3"`
	Second_Choice string    `json:"second_choice" validate:"required,len=3,nefield=First_Choice"`
	Introduction  string    `json:"introduction" validate:"required,min=4,max=500"`
}
