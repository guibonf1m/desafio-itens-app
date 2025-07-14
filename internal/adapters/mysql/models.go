package mysql

import (
	entity "desafio-itens-app/internal/domain/item"
	userEntity "desafio-itens-app/internal/domain/user"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID        int            `gorm:"primaryKey;autoIncrement"`
	Username  string         `gorm:"uniqueIndex;size:50;not null"`
	Email     string         `gorm:"uniqueIndex;size:255;not null"`
	Password  string         `gorm:"size:255;not null"`
	Role      string         `gorm:"type:enum('admin','user');default:'user';not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (UserModel) TableName() string {
	return "users"
}

func (m *UserModel) toEntity() userEntity.User {
	return userEntity.User{
		ID:        m.ID,
		Username:  m.Username,
		Email:     m.Email,
		Password:  m.Password,
		Role:      userEntity.Role(m.Role),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func fromUserEntity(user userEntity.User) UserModel {
	return UserModel{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Role:      string(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type ItemModel struct {
	ID            int            `gorm:"primaryKey;autoIncrement"`
	Code          string         `gorm:"uniqueIndex;size:50;not null"`
	Nome          string         `gorm:"size:100;not null"`
	Descricao     string         `gorm:"size:500"`
	Preco         float64        `gorm:"type:decimal(10,2);not null"`
	Estoque       int            `gorm:"default:0;not null"`
	Status        string         `gorm:"type:enum('active','inactive');default:'active'"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	CreatedBy     *int           `gorm:"column:created_by;index"`
	UpdatedBy     *int           `gorm:"column:updated_by;index"`
	CreatedByUser *UserModel     `gorm:"foreignKey:CreatedBy;references:ID"`
	UpdatedByUser *UserModel     `gorm:"foreignKey:UpdatedBy;references:ID"`
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
		CreatedBy: m.CreatedBy,
		UpdateBy:  m.UpdatedBy,
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
		CreatedBy: item.CreatedBy,
		UpdatedBy: item.UpdateBy,
	}
}
