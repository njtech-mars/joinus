package controllers

import (
	"server/configs"
	"server/lib"
	"server/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getRecords(studentId string, options interface{}) []models.EmailRecord {
	var records []models.EmailRecord
	result := configs.Db.Where(&options).Find(&records)
	if result.Error != nil {
		return nil
	} else {
		return records
	}
}

func EmailQuery(c *fiber.Ctx) error {
	studentId := c.Query("student_id", "")
	if studentId == "" {
		return c.Status(400).JSON(models.Response{Success: false, Message: "无效的请求体"})
	} else {
		records := getRecords(studentId, models.EmailRecord{Student_ID: studentId})
		if records == nil {
			return c.Status(404).JSON(models.Response{Success: false, Message: "找不到相关记录"})
		} else {
			return c.Status(200).JSON(models.Response{Success: true, Message: "数据获取成功", Data: records})
		}
	}
}

func EmailSend(c *fiber.Ctx) error {
	var applicant models.Applicant
	once := c.QueryBool("once", true)
	studentId := c.Query("student_id", "")

	if studentId == "" {
		return c.Status(400).JSON(models.Response{Success: false, Message: "无效的请求体"})
	} else if err := configs.Db.Where(&models.Applicant{Student_ID: studentId}).First(&applicant).Error; err != nil {
		return c.Status(404).JSON(models.Response{Success: false, Message: "找不到相关记录", Data: err.Error()})
	} else {
		if once && len(getRecords(studentId, models.EmailRecord{Student_ID: studentId, Status: "success"})) > 0 {
			return c.Status(400).JSON(models.Response{Success: false, Message: "邮件已经成功发送过啦"})
		}

		if err := lib.SendEmail(applicant.Email, applicant.Name); err != nil {
			var record = models.EmailRecord{Email: applicant.Email, Student_ID: studentId, Sent_At: time.Now(), Status: "failed"}
			configs.Db.Create(&record)
			return c.Status(500).JSON(models.Response{Success: false, Message: "邮件发送失败", Data: err.Error()})
		} else {
			var record = models.EmailRecord{Email: applicant.Email, Student_ID: studentId, Sent_At: time.Now(), Status: "success"}
			configs.Db.Create(&record)
			return c.Status(200).JSON(models.Response{Success: true, Message: "邮件发送成功", Data: record})
		}
	}
}
