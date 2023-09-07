package controllers

import (
	"fmt"
	"server/configs"
	"server/models"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

var validate = validator.New()

func ApplicantCreate(c *fiber.Ctx) error {
	var applicant models.Applicant
	applicant.Submitted_At = time.Now()

	// parse request body
	if err := c.BodyParser(&applicant); err != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "无效的请求体"})
	}

	if err := validate.Struct(&applicant); err != nil {
		return c.Status(400).JSON(models.Response{Success: false, Message: "表单内容不合法", Data: err.Error()})
	}

	options := &models.Applicant{Student_ID: applicant.Student_ID}
	result := configs.Db.Where(options).Assign(applicant).FirstOrCreate(&applicant)

	if result.Error != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: result.Error.Error()})
	} else if result.RowsAffected > 0 {
		return c.Status(201).JSON(models.Response{Success: true, Message: "表单提交成功", Data: applicant})
	} else {
		return c.Status(200).JSON(models.Response{Success: false, Message: "你的提交没有变化", Data: applicant})
	}
}

func ApplicantQuery(c *fiber.Ctx) error {
	var applicant models.Applicant
	studentId := c.Query("student_id", "")

	if studentId == "" {
		return c.Status(400).JSON(models.Response{Success: false, Message: "无效的请求体"})
	} else if err := configs.Db.Where(&models.Applicant{Student_ID: studentId}).First(&applicant).Error; err != nil {
		return c.Status(404).JSON(models.Response{Success: false, Message: "找不到相关记录", Data: err.Error()})
	} else {
		return c.Status(200).JSON(models.Response{Success: true, Message: "数据读取成功", Data: applicant})
	}
}

func ApplicantsQuery(c *fiber.Ctx) error {
	var applicants []models.Applicant
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "20"))

	var total int64
	configs.Db.Model(&models.Applicant{}).Count(&total)

	result := configs.Db.Order("submitted_at desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&applicants)
	if result.Error != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: result.Error.Error()})
	} else {
		return c.Status(200).JSON(models.Response{Success: true, Message: "数据读取成功", Data: applicants, Pagination: models.Pagination{Page: page, Total: total, Page_Size: pageSize}})
	}
}

func ApplicantsDownload(c *fiber.Ctx) error {
	// query databse
	var applicants []models.Applicant
	if err := configs.Db.Find(&applicants).Error; err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	}

	// create a sheet file
	file := excelize.NewFile()
	defer file.Close()

	// update sheet name
	name := "Mars报名表单"
	file.SetSheetName("Sheet1", name)

	// setup sheet header
	file.SetCellValue(name, "A1", "序号")
	file.SetCellValue(name, "B1", "姓名")
	file.SetCellValue(name, "C1", "性别")
	file.SetCellValue(name, "D1", "邮箱")
	file.SetCellValue(name, "E1", "QQ")
	file.SetCellValue(name, "F1", "学号")
	file.SetCellValue(name, "G1", "学院")
	file.SetCellValue(name, "H1", "专业")
	file.SetCellValue(name, "I1", "提交时间")
	file.SetCellValue(name, "J1", "自我介绍")

	// write data to sheet
	for i, d := range applicants {
		row := i + 2
		file.SetCellValue(name, "A"+fmt.Sprintf("%d", row), i+1)
		file.SetCellValue(name, "B"+fmt.Sprintf("%d", row), d.Name)
		file.SetCellValue(name, "C"+fmt.Sprintf("%d", row), d.Gender)
		file.SetCellValue(name, "D"+fmt.Sprintf("%d", row), d.Email)
		file.SetCellValue(name, "E"+fmt.Sprintf("%d", row), d.QQ)
		file.SetCellValue(name, "F"+fmt.Sprintf("%d", row), d.Student_ID)
		file.SetCellValue(name, "G"+fmt.Sprintf("%d", row), d.College)
		file.SetCellValue(name, "H"+fmt.Sprintf("%d", row), d.Major)
		file.SetCellValue(name, "I"+fmt.Sprintf("%d", row), d.Submitted_At.Format("2006-01-02 15:04"))
		file.SetCellValue(name, "J"+fmt.Sprintf("%d", row), d.Introduction)
	}

	// setup collum with
	file.SetColWidth(name, "D", "I", 17)
	file.SetColWidth(name, "J", "J", 100)

	// setup sheet style
	wrapStyle, _ := file.NewStyle(&excelize.Style{Alignment: &excelize.Alignment{WrapText: true}})
	alignStyle, _ := file.NewStyle(&excelize.Style{Alignment: &excelize.Alignment{Vertical: "top", Horizontal: "left"}})
	file.SetColStyle(name, "J", wrapStyle)
	file.SetColStyle(name, "A:I", alignStyle)

	// Set the appropriate headers for the response
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", "attachment; filename="+name+".xlsx")

	// Write the Excel data to the client
	if err := file.Write(c.Response().BodyWriter()); err != nil {
		return c.Status(500).JSON(models.Response{Success: false, Message: "服务器内部错误", Data: err.Error()})
	}

	return nil
}
