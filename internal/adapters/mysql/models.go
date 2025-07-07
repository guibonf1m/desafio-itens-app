package mysql

import (
	entity "desafio-itens-app/internal/domain/item"
	"gorm.io/gorm"
	"time"
)

// ✅ Model específico para GORM
type ItemModel struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	Code      string         `gorm:"uniqueIndex;size:50;not null"`
	Nome      string         `gorm:"size:100;not null"`
	Descricao string         `gorm:"size:500"`
	Preco     float64        `gorm:"type:decimal(10,2);not null"`
	Estoque   int            `gorm:"default:0;not null"`
	Status    string         `gorm:"type:enum('active','inactive');default:'active'"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ItemModel) TableName() string {
	return "itens"
}

// 🔄 CONVERSÕES Domain ↔ Model
func (m *ItemModel) ToEntity() entity.Item {
	return entity.Item{
		ID:        m.ID,
		Code:      m.Code,
		Nome:      m.Nome,
		Descricao: m.Descricao,
		Preco:     m.Preco,
		Estoque:   m.Estoque,
		Status:    entity.Status(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func FromEntity(item entity.Item) ItemModel {
	return ItemModel{
		ID:        item.ID,
		Code:      item.Code,
		Nome:      item.Nome,
		Descricao: item.Descricao,
		Preco:     item.Preco,
		Estoque:   item.Estoque,
		Status:    string(item.Status),
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}
