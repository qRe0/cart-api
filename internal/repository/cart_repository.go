package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/qRe0/innowise-cart-api/configs"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/models"
)

const (
	connStrTmpl = "user=%s password=%s dbname=%s host=%s sslmode=disable"

	createCartQuery = `INSERT INTO carts DEFAULT VALUES`
	maxCartIDQuery  = `SELECT MAX(id) FROM carts`
	insertItemQuery = `INSERT INTO items (cart_id, product, quantity) 
    	VALUES ($1, $2, $3)
    	ON CONFLICT (cart_id, product) 
    	DO UPDATE SET quantity = items.quantity + EXCLUDED.quantity RETURNING id`
	checkItemQuery  = `SELECT id FROM items WHERE id = $1 AND cart_id = $2`
	deleteItemQuery = `DELETE FROM items WHERE id = $1 AND cart_id = $2`
	selectItemQuery = `SELECT id, cart_id, product, quantity FROM items WHERE cart_id = $1`
	checkCartQuery  = `SELECT id FROM carts WHERE id = $1`
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func Init(cfg configs.DBConfig) (*sqlx.DB, error) {
	connStr := fmt.Sprintf(connStrTmpl, cfg.User, cfg.Password, cfg.DBName, cfg.Host)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Connected to DB successfully!")

	return db, nil
}

func (r *CartRepository) CreateCart(ctx context.Context) (*models.Cart, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errs.ErrStartTransaction
	}

	_, err = tx.ExecContext(ctx, createCartQuery)
	if err != nil {
		return nil, errs.ErrCreateCart
	}

	var id int
	err = tx.QueryRowContext(ctx, maxCartIDQuery).Scan(&id)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, errs.ErrCommitTransaction
	}

	cart := models.Cart{
		ID:    id,
		Items: []models.CartItem{},
	}

	return &cart, nil
}

func (r *CartRepository) AddItemToCart(ctx context.Context, item models.CartItem) (*models.CartItem, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errs.ErrStartTransaction
	}

	row := tx.QueryRowContext(ctx, checkCartQuery, item.CartID)
	err = row.Scan(&item.CartID)
	if err != nil {
		tx.Rollback()
		return nil, errs.ErrCartNotFound
	}

	var id int
	err = tx.QueryRowContext(ctx, insertItemQuery, item.CartID, item.Product, item.Quantity).Scan(&id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, errs.ErrCommitTransaction
	}

	item.ID = id
	return &item, nil
}

func (r *CartRepository) RemoveItemFromCart(item *models.CartItem) error {
	row := r.db.QueryRow(checkCartQuery, item.CartID)
	err := row.Scan(&item.CartID)
	if err != nil {
		return errs.ErrCartNotFound
	}

	row = r.db.QueryRow(checkItemQuery, item.ID, item.CartID)
	err = row.Scan(&item.ID)
	if err != nil {
		return errs.ErrItemNotFound
	}

	_, err = r.db.Exec(deleteItemQuery, item.ID, item.CartID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CartRepository) GetCart(ctx context.Context, cart *models.Cart) (*models.Cart, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errs.ErrStartTransaction
	}

	row := tx.QueryRowContext(ctx, checkCartQuery, cart.ID)
	err = row.Scan(&cart.ID)
	if err != nil {
		tx.Rollback()
		return nil, errs.ErrCartNotFound
	}

	rows, err := tx.QueryContext(ctx, selectItemQuery, cart.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	item := models.CartItem{}
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.CartID, &item.Product, &item.Quantity)
		if err != nil {
			return nil, err
		}
		cart.Items = append(cart.Items, item)
	}

	err = tx.Commit()
	if err != nil {
		return nil, errs.ErrCommitTransaction
	}

	return cart, nil
}
