package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/qRe0/innowise-cart-api/internal/models"

	_ "github.com/lib/pq"
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository() *CartRepository {
	return &CartRepository{
		db: Init(),
	}
}

func Init() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=qre password=2411 dbname=cart_api sslmode=disable")
	if err != nil {
		e := errors.New("failed to connect to database")
		panic(e)
	}

	return db
}

func (r *CartRepository) CreateCart(cart models.Cart) error {
	_, err := r.db.Exec("INSERT INTO carts (id) VALUES ($1)", cart.ID)
	if err != nil {
		return errors.New("failed to create cart")
	}

	return nil
}

func (r *CartRepository) AddItemToCart(cart models.Cart, item models.CartItem) error {
	_, err := r.db.Exec("INSERT INTO items (id, cart_id, product, quantity) VALUES ($1, $2, $3, $4)", item.ID, cart.ID, item.Product, item.Quantity)
	if err != nil {
		return errors.New("error. cart not found in database")
	}

	return nil
}

func (r *CartRepository) RemoveItemFromCart(cart models.Cart, item models.CartItem) error {
	var itemCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM items WHERE id = $1 AND cart_id = $2", item.ID, cart.ID).Scan(&itemCount)
	if err != nil {
		return errors.New("error getting items count from database")
	}

	if itemCount == 0 {
		return errors.New("error. item not found in database")
	}

	_, err = r.db.Exec("DELETE FROM items WHERE id = $1 AND cart_id = $2", item.ID, cart.ID)
	if err != nil {
		return errors.New("failed to remove item from cart")
	}

	return nil
}

func (r *CartRepository) GetCart(cart models.Cart, item models.CartItem) error {
	return nil
}

func (r *CartRepository) GetLastID() (int, error) {
	var id int
	err := r.db.QueryRow("SELECT COUNT(*) FROM carts").Scan(&id)
	if err != nil {
		return 0, errors.New("error getting last cart id from database")
	}

	return id, nil
}

func (r *CartRepository) GetLastItemID() (int, error) {
	var id int
	err := r.db.QueryRow("SELECT MAX(id) FROM items").Scan(&id)
	if err != nil {
		return 0, errors.New("error getting last item id from database. database is empty. currentID = 0")
	}

	return id, nil
}
