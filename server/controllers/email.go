package controllers

import (
	"server/configs"
	"server/lib"
	"server/models"

	"github.com/gofiber/fiber/v2"
)

func EmailSend(c *fiber.Ctx) error {
	once := c.QueryBool("once", true)
	studentId := c.Query("student_id", "")

	if studentId == "" {
		return c.Status(400).JSON(models.Response{Success: false, Message: "无效的请求体"})
	} else {
		var applicant models.Applicant
		result := configs.Db.Where(&models.Applicant{Student_ID: studentId}).Limit(1).Find(&applicant)

		if err := result.Error; err != nil {
			return c.Status(404).JSON(models.Response{Success: false, Message: "找不到相关记录", Data: err.Error()})
		} else {
			if once && applicant.Email_Status == "success" {
				return c.Status(400).JSON(models.Response{Success: false, Message: "邮件已经成功发送过啦"})
			}

			if err := lib.SendEmail(applicant.Email, applicant.Name); err != nil {
				applicant.Email_Status = "failed"
				configs.Db.Save(applicant)
				return c.Status(500).JSON(models.Response{Success: false, Message: "邮件发送失败", Data: err.Error()})
			} else {
				applicant.Email_Status = "success"
				configs.Db.Save(applicant)
				return c.Status(200).JSON(models.Response{Success: true, Message: "邮件发送成功"})
			}
		}
	}
}
