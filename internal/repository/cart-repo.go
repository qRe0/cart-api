package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/qRe0/innowise-cart-api/internal/models"
)

const (
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
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func Init() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load enviromental variables: %w", err)
	}

	connStr := "user=" + os.Getenv("DATABASE_USER") +
		" password=" + os.Getenv("DATABASE_PASSWORD") +
		" dbname=" + os.Getenv("DATABASE_NAME") +
		" host=" + os.Getenv("DATABASE_HOST") +
		" sslmode=disable"

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

func (r *CartRepository) IsItemExist(itemID, cartID int) (bool, error) {
	var count int
	err := r.db.QueryRow(itemCountQuery, itemID, cartID).Scan(&count)
	if count == 0 {
		return false, err
	}

	return true, nil
}

func (r *CartRepository) IsCartExist(cartID int) (bool, error) {
	var count int
	err := r.db.QueryRow(cartCountQuery, cartID).Scan(&count)
	if count == 0 {
		return false, err
	}

	return true, nil
}

func (r *CartRepository) GetCart(cartID int) (*models.Cart, error) {
	rows, err := r.db.Query(selectItemQuery, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	item := models.CartItem{}
	cart := models.Cart{}
	cart.ID = cartID
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.CartID, &item.Product, &item.Quantity)
		if err != nil {
			return nil, err
		}
		cart.Items = append(cart.Items, item)
	}

	return &cart, nil
}
