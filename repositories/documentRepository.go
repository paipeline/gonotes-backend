// repositories/documentRepository.go
package repositories

import (
	"goauth/models"

	"gorm.io/gorm"
)

// DocumentRepository 接口定义了所有文档相关的数据库操作
type DocumentRepository interface {
	GetDocument(id string) (*models.Document, error)
	CreateDocument(document *models.Document) error
	UpdateDocument(document *models.Document) error
	DeleteDocument(id string) error
	GetSidebarDocuments(userId string, parentDocument *string) ([]models.Document, error)
	GetChildDocuments(parentId string, userId string) ([]models.Document, error)
	// 可以根据需要添加其他方法
}

// documentRepository 是 DocumentRepository 接口的具体实现
type documentRepository struct {
	DB *gorm.DB
}

// NewDocumentRepository 创建一个新的 DocumentRepository 实例
func NewDocumentRepository(db *gorm.DB) DocumentRepository {
	return &documentRepository{DB: db}
}

func (r *documentRepository) GetDocument(id string) (*models.Document, error) {
	var document models.Document
	err := r.DB.First(&document, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &document, nil
}

func (r *documentRepository) CreateDocument(document *models.Document) error {
	return r.DB.Create(document).Error
}

func (r *documentRepository) UpdateDocument(document *models.Document) error {
	return r.DB.Save(document).Error
}

func (r *documentRepository) DeleteDocument(id string) error {
	return r.DB.Delete(&models.Document{}, "id = ?", id).Error
}

func (r *documentRepository) GetSidebarDocuments(userId string, parentDocument *string) ([]models.Document, error) {
	var documents []models.Document
	query := r.DB.Where("user_id = ? AND is_archived = ?", userId, false)

	if parentDocument != nil {
		query = query.Where("parent_document = ?", *parentDocument)
	} else {
		query = query.Where("parent_document IS NULL")
	}

	err := query.Order("created_at desc").Find(&documents).Error
	return documents, err
}

func (r *documentRepository) GetChildDocuments(parentId string, userId string) ([]models.Document, error) {
	var documents []models.Document
	err := r.DB.Where("parent_document = ? AND user_id = ?", parentId, userId).Find(&documents).Error
	return documents, err
}
