package mysql

import (
	entity "desafio-itens-app/internal/domain/item"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type MySQLItemRepository struct {
	db *gorm.DB
}

func NewMySQLItemRepository(db *gorm.DB) *MySQLItemRepository {
	return &MySQLItemRepository{db: db}
}

func (r *MySQLItemRepository) GetItem(id int) (*entity.Item, error) {
	var model ItemModel

	err := r.db.First(&model, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Item não encontrado")
		}
		return nil, fmt.Errorf("Erro ao buscar item: %w", err)
	}

	item := model.ToEntity()
	return &item, nil
}

func (r *MySQLItemRepository) GetItens() ([]entity.Item, error) {
	var models []ItemModel

	err := r.db.Find(&models).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar itens: %w", err)
	}

	var itens []entity.Item
	for _, model := range models {
		itens = append(itens, model.ToEntity())
	}

	return itens, nil
}

func (r *MySQLItemRepository) GetItensFiltrados(status *entity.Status, limit int) ([]entity.Item, error) {

	var models []ItemModel
	query := r.db.Model(&ItemModel{})

	if status != nil {
		query = query.Where("status = ?", string(*status))
	}

	err := query.Limit(limit).Find(&models).Error
	if err != nil {
		return nil, fmt.Errorf("Erro ao buscar itens filtrados: %w", err)
	}

	var itens []entity.Item
	for _, model := range models {
		itens = append(itens, model.ToEntity())
	}

	return itens, nil
}

func (r *MySQLItemRepository) GetItensPaginados(offset, limit int) ([]entity.Item, int, error) {
	var models []ItemModel
	var totalCount int64

	err := r.db.Model(&ItemModel{}).Count(&totalCount).Error
	if err != nil {
		return nil, 0, fmt.Errorf("Erro ao contar itens: %w", err)
	}

	err = r.db.Offset(offset).Limit(limit).Order("Created_at DESC").Find(&models).Error
	if err != nil {
		return nil, 0, fmt.Errorf("Erro ao buscar itens paginados: %w", err)
	}

	var itens []entity.Item
	for _, model := range models {
		itens = append(itens, model.ToEntity())
	}

	return itens, int(totalCount), nil
}

func (r *MySQLItemRepository) GetItensFiltradosPaginados(status *entity.Status, offset, limit int) ([]entity.Item, int, error) {

	var models []ItemModel
	var totalCount int64

	query := r.db.Model(&ItemModel{})
	if status != nil {
		query = query.Where("status = ?", string(*status))
	}

	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, fmt.Errorf("Erro ao contar itens filtrados: %w", err)
	}

	err = query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&models).Error
	if err != nil {
		return nil, 0, fmt.Errorf("Erro ao buscar itens filtrados paginados: %w", err)
	}

	var itens []entity.Item
	for _, model := range models {
		itens = append(itens, model.ToEntity())
	}
	return itens, int(totalCount), nil
}

func (r *MySQLItemRepository) CountItens(status *entity.Status) (int, error) {
	var count int64
	query := r.db.Model(&ItemModel{})

	if status != nil {
		query = query.Where("status = ?", string(*status))
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, fmt.Errorf("Erro ao contar itens: %w", err)
	}
	return int(count), nil
}

func (r *MySQLItemRepository) CodeExists(code string) (bool, error) {
	var count int64

	err := r.db.Model(&ItemModel{}).Where("code = ?", code).Count(&count).Error
	if err != nil {
		return false, fmt.Errorf("Erro ao verificar código: %w", err)
	}

	return count > 0, nil
}

func (r *MySQLItemRepository) AddItem(item entity.Item) (entity.Item, error) {

	model := FromEntity(item)

	err := r.db.Create(&model).Error
	if err != nil {
		return entity.Item{}, fmt.Errorf("Erro ao criar item: %w", err)
	}

	return model.ToEntity(), nil

}

func (r *MySQLItemRepository) UpdateItem(item entity.Item) error {
	model := FromEntity(item)

	err := r.db.Save(&model).Error
	if err != nil {
		return fmt.Errorf("Erro ao atualiazar item :%w", err)
	}
	return nil
}

func (r *MySQLItemRepository) DeleteItem(id int) error {
	err := r.db.Delete(&ItemModel{}, id).Error
	if err != nil {
		return fmt.Errorf("Erro ao deletar item: %w", err)
	}
	return nil
}
