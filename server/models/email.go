package models

import "time"

type EmailRecord struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Sent_At    time.Time `json:"sent_at"`
	Status     string    `json:"status"`
	Email      string    `json:"email"`
	Student_ID string    `json:"student_id"`
}
