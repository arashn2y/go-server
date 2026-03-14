package products

import (
	"context"

	"github.com/arashn0uri/go-server/internal/constants"
	"github.com/arashn0uri/go-server/internal/form"
	"github.com/arashn0uri/go-server/internal/models"
	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/arashn0uri/go-server/internal/routes/permissions"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	Products(ctx context.Context) (*[]models.Product, error)
	ProductByID(ctx context.Context, id pgtype.UUID) (*models.Product, error)
	CreateProduct(ctx context.Context, data form.CreateProductRequest) error
	UpdateProduct(ctx context.Context, id pgtype.UUID, data form.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, id pgtype.UUID) (int64, error)
}

type service struct {
	repository *repository.Queries
	permission permissions.Service
}

func NewService(db *repository.Queries, permission permissions.Service) Service {
	return &service{
		repository: db,
		permission: permission,
	}
}

func (s *service) Products(ctx context.Context) (*[]models.Product, error) {
	products, err := s.repository.GetAllProducts(ctx)

	if err != nil {
		return nil, err
	}

	result := make([]models.Product, len(products))

	for i, product := range products {
		result[i] = models.Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description.String,
			Price:       float64(product.PriceInCents) / 100,
		}
	}

	return &result, nil
}

func (s *service) ProductByID(ctx context.Context, id pgtype.UUID) (*models.Product, error) {
	product, err := s.repository.GetProductByID(ctx, id)

	if err != nil {
		return nil, err
	}

	result := &models.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description.String,
		Price:       float64(product.PriceInCents) / 100,
	}

	return result, nil
}

func (s *service) CreateProduct(ctx context.Context, data form.CreateProductRequest) error {
	permissionErr := s.permission.CheckPermission(ctx, string(constants.ResourceProducts), string(constants.PermissionCreate))
	if permissionErr != nil {
		return permissionErr
	}
	err := s.repository.CreateProduct(ctx, repository.CreateProductParams{
		Name:         data.Name,
		Description:  pgtype.Text{String: data.Description, Valid: true},
		PriceInCents: data.PriceInCents,
	})

	return err
}

func (s *service) UpdateProduct(ctx context.Context, id pgtype.UUID, data form.UpdateProductRequest) error {
	err := s.repository.UpdateProduct(ctx, repository.UpdateProductParams{
		ID:           id,
		Name:         data.Name,
		Description:  pgtype.Text{String: data.Description, Valid: true},
		PriceInCents: data.PriceInCents,
	})

	return err
}

func (s *service) DeleteProduct(ctx context.Context, id pgtype.UUID) (int64, error) {
	rows, err := s.repository.DeleteProduct(ctx, id)

	return rows, err
}
