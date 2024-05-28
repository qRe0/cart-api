package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	myErrors "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/models"

	_ "github.com/lib/pq"
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
	db, err := sqlx.Open("postgres", "user=pgadmin password=24112004 dbname=cart_api host=database sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to DB successfully!")

	return db, nil
}

func (r *CartRepository) CreateCart() (*models.Cart, error) {
	_, err := r.db.Exec(`INSERT INTO carts DEFAULT VALUES `)
	if err != nil {
		return nil, myErrors.ErrCreatingCart
	}

	var id int
	err = r.db.QueryRow(`SELECT MAX(id) from carts`).Scan(&id)
	if err != nil {
		return nil, myErrors.ErrWrongCartID
	}

	cart := models.Cart{
		Entity: models.Entity{
			ID: &id,
		},
		Items: []models.CartItem{},
	}

	return &cart, nil
}

func (r *CartRepository) AddItemToCart(cartID int, item models.CartItem) (*models.CartItem, error) {
	var cartCount int
	_ = r.db.QueryRow(`SELECT COUNT(id) FROM carts WHERE id = $1`, cartID).Scan(&cartCount)
	if cartCount == 0 {
		return nil, myErrors.ErrCartNotFound
	}

	_, err := r.db.Exec(`
    INSERT INTO items (cart_id, product, quantity) 
    VALUES ($1, $2, $3)
    ON CONFLICT (cart_id, product) 
    DO UPDATE SET quantity = items.quantity + EXCLUDED.quantity`,
		cartID, item.Product, item.Quantity)
	if err != nil {
		return nil, myErrors.ErrAddItemToCart
	}

	id, err := r.GetLastItemID()
	if err != nil {
		return nil, myErrors.ErrGettingLastItemID
	}
	item = models.CartItem{
		Entity: models.Entity{
			ID: &id,
		},
		CartID:   cartID,
		Product:  item.Product,
		Quantity: item.Quantity,
	}

	return &item, nil
}

func (r *CartRepository) RemoveItemFromCart(cartID, itemID int) error {
	var itemCount, cartCount int
	_ = r.db.QueryRow(`SELECT COUNT(id) FROM carts WHERE id = $1`, cartID).Scan(&cartCount)
	if cartCount == 0 {
		return myErrors.ErrCartNotFound
	}

	_ = r.db.QueryRow(`SELECT COUNT(id) FROM items WHERE id = $1 AND cart_id = $2`, itemID, cartID).Scan(&itemCount)
	if itemCount == 0 {
		return myErrors.ErrItemNotFound
	}

	_, err := r.db.Exec(`DELETE FROM items WHERE id = $1 AND cart_id = $2`, itemID, cartID)
	if err != nil {
		return myErrors.ErrRemoveItemFromCart
	}

	return nil
}

func (r *CartRepository) GetCart(cartID int) (*models.Cart, error) {
	var cartCount int
	_ = r.db.QueryRow(`SELECT COUNT(*) FROM carts WHERE id = $1`, cartID).Scan(&cartCount)
	if cartCount == 0 {
		return nil, myErrors.ErrCartNotFound
	}

	rows, err := r.db.Query(`SELECT * FROM items WHERE cart_id = $1`, cartID)
	if err != nil {
		return nil, myErrors.ErrGetItems
	}
	defer rows.Close()

	item := models.CartItem{}
	cart := models.Cart{}
	cart.ID = &cartID
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.CartID, &item.Product, &item.Quantity)
		if err != nil {
			return nil, myErrors.ErrRowsScan
		}
		cart.Items = append(cart.Items, item)
	}

	return &cart, nil
}

func (r *CartRepository) GetLastItemID() (int, error) {
	var id int
	err := r.db.QueryRow(`SELECT MAX(id) FROM items`).Scan(&id)
	if err != nil {
		return 0, myErrors.ErrGettingLastItemID
	}

	return id, nil
}
