package controllers

import (
	"time"

	db "example.com/go-fiber/config"
	"example.com/go-fiber/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var cashiers []models.Cashier
var cashier models.Cashier

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"success": "false",
				"message": "Invalid request body",
			},
		)
	}

	if data["name"] == "" || data["passcode"] == "" {
		return c.Status(400).JSON(
			fiber.Map{
				"success": "false",
				"message": "Name and passcode should be given",
			},
		)
	}

	now := time.Now()
	cashier := models.Cashier{
		Name:      data["name"],
		Passcode:  data["passcode"],
		CreatedAt: now,
		UpdatedAt: now,
	}
	db.DB.Create(&cashier)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "cashier created",
		"name":    cashier.Name,
	})
}

func CashiersList(c *fiber.Ctx) error {
	// limit, _ := strconv.Atoi(c.Query("limit", "1"))
	// skip, _ := strconv.Atoi(c.Query("skip", "10"))
	// var count int64

	// db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashiers)
	db.DB.Find(&cashiers)

	return c.Status(200).JSON(fiber.Map{"status": "success", "results": len(cashiers), "notes": cashiers})
}

func GetCashierDetails(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	result := db.DB.Select("id, name, created_at, updated_at").Where("id = ?", cashierId).First(&cashier)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No item with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	cashierData := make(map[string]interface{})
	cashierData["cashierId"] = cashier.Id
	cashierData["name"] = cashier.Name
	cashierData["createdAt"] = cashier.CreatedAt
	cashierData["updatedAt"] = cashier.UpdatedAt

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   cashierData,
	})
}

func UpdateCashier(c *fiber.Ctx) error {
	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	return nil
}
