package repository

import (
	"github.com/jmoiron/sqlx"
	myErrors "github.com/qRe0/innowise-cart-api/internal/errors"
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
		e := myErrors.ErrConnectingToDB
		panic(e)
	}

	return db
}

func (r *CartRepository) CreateCart(cart models.Cart) error {
	_, err := r.db.Exec("INSERT INTO carts (id) VALUES ($1)", cart.ID)
	if err != nil {
		return myErrors.ErrCreatingCart
	}

	return nil
}

func (r *CartRepository) AddItemToCart(cart models.Cart, item models.CartItem) error {
	_, err := r.db.Exec("INSERT INTO items (id, cart_id, product, quantity) VALUES ($1, $2, $3, $4)", item.ID, cart.ID, item.Product, item.Quantity)
	if err != nil {
		return myErrors.ErrAddItemToCart
	}

	return nil
}

func (r *CartRepository) RemoveItemFromCart(cart models.Cart, item models.CartItem) error {
	var itemCount, cartCount int
	_ = r.db.QueryRow("SELECT COUNT(id) FROM carts WHERE id = $1", cart.ID).Scan(&cartCount)
	if cartCount == 0 {
		return myErrors.ErrCartNotFound
	}

	_ = r.db.QueryRow("SELECT COUNT(id) FROM items WHERE id = $1 AND cart_id = $2", item.ID, cart.ID).Scan(&itemCount)
	if itemCount == 0 {
		return myErrors.ErrItemNotFound
	}

	_, err := r.db.Exec("DELETE FROM items WHERE id = $1 AND cart_id = $2", item.ID, cart.ID)
	if err != nil {
		return myErrors.ErrRemoveItemFromCart
	}

	return nil
}

func (r *CartRepository) GetCart(cart *models.Cart, item models.CartItem) error {
	var cartCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM carts WHERE id = $1", cart.ID).Scan(&cartCount)
	if err != nil {
		return myErrors.ErrGettingItemsCount
	}

	if cartCount == 0 {
		return myErrors.ErrCartNotFound
	}

	rows, err := r.db.Query("SELECT * FROM items WHERE cart_id = $1", cart.ID)
	if err != nil {
		return myErrors.ErrGetItems
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&item.ID, &item.CartID, &item.Product, &item.Quantity)
		if err != nil {
			return myErrors.ErrRowsScan
		}
		cart.Items = append(cart.Items, item)
	}

	return nil
}

func (r *CartRepository) GetLastCartID() (int, error) {
	var id int
	err := r.db.QueryRow("SELECT COUNT(*) FROM carts").Scan(&id)
	if err != nil {
		return 0, myErrors.ErrGettingLastCartID
	}

	return id, nil
}

func (r *CartRepository) GetLastItemID() (int, error) {
	var id int
	err := r.db.QueryRow("SELECT MAX(id) FROM items").Scan(&id)
	if err != nil {
		return 0, myErrors.ErrGettingLastItemID
	}

	return id, nil
}
