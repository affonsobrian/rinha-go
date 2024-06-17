package models

import "time"

type Cliente struct {
	ID         uint        `json:"-" gorm:"primaryKey"`
	Limite     int         `json:"limite"`
	Saldo      int         `json:"saldo"`
	Transacoes []Transacao `json:"-" gorm:"foreignKey:ClienteID"`
}

type Transacao struct {
	ID          uint      `json:"-" gorm:"primaryKey"`
	ClienteID   uint      `json:"-"`
	Cliente     Cliente   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Valor       int       `json:"valor"`
	Tipo        string    `json:"tipo" gorm:"size:1"`
	Descricao   string    `json:"descricao"`
	RealizadoEm time.Time `json:"realizada_em"`
}

type Saldo struct {
	Limite      int       `json:"limite"`
	Total       int       `json:"total"`
	DataExtrato time.Time `json:"data_extrato"`
}

type Extrato struct {
	Saldo             Saldo       `json:"saldo"`
	UltimasTransacoes []Transacao `json:"ultimas_transacoes"`
}
