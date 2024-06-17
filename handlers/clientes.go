package handlers

import (
	"strconv"
	"time"

	"github.com/affonsobrian/rinha-go/database"
	"github.com/affonsobrian/rinha-go/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllClientes(c *fiber.Ctx) error {
	var clientes []models.Cliente
	database.DB.Db.Find(&clientes)
	return c.Status(fiber.StatusOK).JSON(clientes)
}

func CreateCliente(c *fiber.Ctx) error {
	cliente := new(models.Cliente)
	if err := c.BodyParser(cliente); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&cliente)

	return c.Status(fiber.StatusCreated).JSON(cliente)
}

func GetAllTransacoes(c *fiber.Ctx) error {
	var transacoes []models.Transacao
	ClienteID := c.Params("id")
	database.DB.Db.Preload("Cliente").Where("cliente_id = ?", ClienteID).Find(&transacoes)
	return c.Status(fiber.StatusOK).JSON(transacoes)
}

func CreateTransacao(c *fiber.Ctx) error {
	transacao := new(models.Transacao)
	cliente := new(models.Cliente)
	if err := c.BodyParser(transacao); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if clientId, err := strconv.ParseUint(c.Params("id"), 10, 32); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	} else {
		transacao.ClienteID = uint(clientId)
	}

	db_transaction := database.DB.Db.Begin()
	if err := database.DB.Db.Where("id = ?", transacao.ClienteID).Find(&cliente).Error; err != nil {
		db_transaction.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if cliente.ID == 0 {
		db_transaction.Rollback()
		return c.SendStatus(fiber.StatusNotFound)
	}
	if transacao.Tipo == "d" {
		cliente.Saldo -= transacao.Valor
		if cliente.Limite < -cliente.Saldo {
			db_transaction.Rollback()
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}
	} else {
		cliente.Saldo += transacao.Valor
	}
	if err := database.DB.Db.Create(&transacao).Error; err != nil {
		db_transaction.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := database.DB.Db.Save(&cliente).Error; err != nil {
		db_transaction.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	db_transaction.Commit()

	return c.Status(fiber.StatusOK).JSON(cliente)
}

func GetExtrato(c *fiber.Ctx) error {
	clientId, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	cliente := new(models.Cliente)
	extrato := new(models.Extrato)
	var transacoes []models.Transacao

	if err := database.DB.Db.Where("id = ?", clientId).Find(&cliente).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if cliente.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}
	extrato.Saldo.Limite = cliente.Limite
	extrato.Saldo.Total = cliente.Saldo
	extrato.Saldo.DataExtrato = time.Now()

	if err := database.DB.Db.Where("cliente_id = ?", clientId).Order("realizado_em DESC").Limit(10).Find(&transacoes).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	extrato.UltimasTransacoes = transacoes
	return c.Status(fiber.StatusOK).JSON(extrato)
}
