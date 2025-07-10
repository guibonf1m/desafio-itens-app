package mysql

import (
	userDomain "desafio-itens-app/internal/domain/user"
	"fmt"
	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

var _ userDomain.UserRepository = (*MySQLUserRepository)(nil)

func NewMySQLUserRepository(db *gorm.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(user userDomain.User) (userDomain.User, error) {
	model := fromUserEntity(user)

	err := r.db.Create(&model).Error
	if err != nil {
		return userDomain.User{}, fmt.Errorf("erro ao criar usuário: %w", err)
	}

	return model.toEntity(), nil
}

func (r *MySQLUserRepository) GetById(id int) (*userDomain.User, error) {
	var model UserModel

	err := r.db.First(&model, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("usuário com ID %d não encontrado", id)
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}
	user := model.toEntity()
	return &user, nil
}

func (r *MySQLUserRepository) GetByUsername(username string) (*userDomain.User, error) {
	var model UserModel
	err := r.db.Where("username = ?", username).First(&model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("usuário %s não encontrado", username)
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}
	user := model.toEntity()
	return &user, nil
}

func (r *MySQLUserRepository) Update(user userDomain.User) error {
	model := fromUserEntity(user)

	err := r.db.Save(&model).Error
	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %w", err)
	}
	return nil
}

func (r *MySQLUserRepository) Delete(id int) error {
	result := r.db.Delete(&UserModel{}, id)

	if result.Error != nil {
		return fmt.Errorf("erro ao deletar usuário: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("usuário com ID %d não encontrado,", id)
	}
	return nil
}

func (r *MySQLUserRepository) UserNameExists(username string) (bool, error) {
	var count int64

	err := r.db.Model(&UserModel{}).Where("username = ?", username).Count(&count)
	if err != nil {
		return false, fmt.Errorf("erro ao verificar username: %w", err)
	}

	return count > 0, nil
}
