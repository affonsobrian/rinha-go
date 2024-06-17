package main

import (
	"github.com/affonsobrian/rinha-go/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/clientes", handlers.CreateCliente)
	app.Get("/clientes", handlers.GetAllClientes)
	app.Post("/clientes/:id/transacoes", handlers.CreateTransacao)
	app.Get("/clientes/:id/transacoes", handlers.GetAllTransacoes)
	app.Get("/clientes/:id/extrato", handlers.GetExtrato)
}