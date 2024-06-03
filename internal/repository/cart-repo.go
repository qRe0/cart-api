package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/qRe0/innowise-cart-api/configs"
	"github.com/qRe0/innowise-cart-api/internal/models"
)

const (
	connStrTmpl = "user=%s password=%s dbname=%s host=%s sslmode=disable"

	createCartQuery = `INSERT INTO carts DEFAULT VALUES`
	maxCartIDQuery  = `SELECT MAX(id) FROM carts`
	cartCountQuery  = `SELECT COUNT(id) FROM carts WHERE id = $1`
	insertItemQuery = `INSERT INTO items (cart_id, product, quantity) 
    	VALUES ($1, $2, $3)
    	ON CONFLICT (cart_id, product) 
    	DO UPDATE SET quantity = items.quantity + EXCLUDED.quantity RETURNING id`
	itemCountQuery  = `SELECT COUNT(id) FROM items WHERE id = $1 AND cart_id = $2`
	deleteItemQuery = `DELETE FROM items WHERE id = $1 AND cart_id = $2`
	selectItemQuery = `SELECT id, cart_id, product, quantity FROM items WHERE cart_id = $1`
	selectCartQuery = `SELECT id FROM carts WHERE id = $1`
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

func (r *CartRepository) CreateCart() (*models.Cart, error) {
	_, err := r.db.Exec(createCartQuery)
	if err != nil {
		return nil, err
	}

	var id int
	err = r.db.QueryRow(maxCartIDQuery).Scan(&id)
	if err != nil {
		return nil, err
	}

	cart := models.Cart{
		ID:    id,
		Items: []models.CartItem{},
	}

	return &cart, nil
}

func (r *CartRepository) AddItemToCart(item models.CartItem) (*models.CartItem, error) {
	var id int
	err := r.db.QueryRow(insertItemQuery, item.CartID, item.Product, item.Quantity).Scan(&id)
	if err != nil {
		return nil, err
	}

	item.ID = id
	return &item, nil
}

func (r *CartRepository) RemoveItemFromCart(cartID, itemID int) error {
	_, err := r.db.Exec(deleteItemQuery, itemID, cartID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CartRepository) GetCart(cart *models.Cart) (*models.Cart, error) {
	row := r.db.QueryRow(selectCartQuery, cart.ID)
	err := row.Scan(&cart.ID)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(selectItemQuery, cart.ID)
	if err != nil {
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

	return cart, nil
}
