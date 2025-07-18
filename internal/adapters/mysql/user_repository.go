package mysql

import (
	"context"
	"desafio-itens-app/internal/application/ports/repositories"
	userDomain "desafio-itens-app/internal/domain/user"
	"fmt"
	"gorm.io/gorm"
)

type MySQLUserRepository struct {
	db *gorm.DB
}

var _ repositories.UserRepository = (*MySQLUserRepository)(nil)

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

func (r *MySQLUserRepository) List(ctx context.Context, limit, offset int) ([]*userDomain.User, int64, error) {
	var models []UserModel
	var totalCount int64

	err := r.db.WithContext(ctx).Model(&UserModel{}).Count(&totalCount).Error
	if err != nil {
		return nil, 0, fmt.Errorf("erro ao contar os usuários: %w", err)
	}

	err = r.db.WithContext(ctx).Limit(limit).Offset(offset).Order("created_at DESC").Find(&models).Error
	if err != nil {
		return nil, 0, fmt.Errorf("erro ao buscar usuários: %w", err)
	}

	var users []*userDomain.User
	for _, model := range models {
		user := model.toEntity()
		users = append(users, &user)
	}

	return users, totalCount, nil
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

func (r *MySQLUserRepository) GetByEmail(email string) (*userDomain.User, error) {
	var model UserModel

	err := r.db.Where("email = ?", email).First(&model).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("usuário com email %s não encontrado", email)
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
		return fmt.Errorf("usuário com ID %d não encontrado", id)
	}
	return nil
}

func (r *MySQLUserRepository) UserNameExists(username string) (bool, error) {
	var count int64

	err := r.db.Model(&UserModel{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("erro ao verificar username: %w", err)
	}

	return count > 0, nil
}

func (r *MySQLUserRepository) EmailExists(email string) (bool, error) {
	var count int64

	err := r.db.Model(&UserModel{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("erro ao verificar email: %w", err)
	}
	return count > 0, nil
}
