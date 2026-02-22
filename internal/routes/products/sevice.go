package products

import (
	"context"

	"github.com/arashn0uri/go-server/internal/form"
	"github.com/arashn0uri/go-server/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	Products(ctx context.Context) ([]repository.Product, error)
	ProductByID(ctx context.Context, id pgtype.UUID) (repository.Product, error)
	CreateProduct(ctx context.Context, data form.CreateProductRequest) error
	UpdateProduct(ctx context.Context, id pgtype.UUID, data form.UpdateProductRequest) error
	DeleteProduct(ctx context.Context, id pgtype.UUID) (int64, error)
}

type service struct {
	repository *repository.Queries
}

func NewService(db *repository.Queries) Service {
	return &service{
		repository: db,
	}
}

func (s *service) Products(ctx context.Context) ([]repository.Product, error) {
	products, err := s.repository.GetAllProducts(ctx)

	if err != nil {
		return []repository.Product{}, err
	}

	return products, nil
}

func (s *service) ProductByID(ctx context.Context, id pgtype.UUID) (repository.Product, error) {
	product, err := s.repository.GetProductByID(ctx, id)

	if err != nil {
		return repository.Product{}, err
	}
	return product, nil
}

func (s *service) CreateProduct(ctx context.Context, data form.CreateProductRequest) error {
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
