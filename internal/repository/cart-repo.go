package repository

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
	"github.com/qRe0/innowise-cart-api/internal/models"
)

const (
	createCartQuery = `INSERT INTO carts DEFAULT VALUES`
	maxCartIDQuery  = `SELECT MAX(id) FROM carts`
	cartCountQuery  = `SELECT COUNT(id) FROM carts WHERE id = $1`
	insertItemQuery = `INSERT INTO items (cart_id, product, quantity) 
    	VALUES ($1, $2, $3)
    	ON CONFLICT (cart_id, product) 
    	DO UPDATE SET quantity = items.quantity + EXCLUDED.quantity`
	itemCountQuery  = `SELECT COUNT(id) FROM items WHERE id = $1 AND cart_id = $2`
	deleteItemQuery = `DELETE FROM items WHERE id = $1 AND cart_id = $2`
	selectItemQuery = `SELECT * FROM items WHERE cart_id = $1`
	maxItemIDQuery  = `SELECT MAX(id) FROM items`
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
		return nil, err
	}

	connStr := "user=" + os.Getenv("DATABASE_USER") +
		" password=" + os.Getenv("DATABASE_PASSWORD") +
		" dbname=" + os.Getenv("DATABASE_NAME") +
		" host=" + os.Getenv("DATABASE_HOST") +
		" sslmode=disable"

	db, err := sqlx.Open("postgres", connStr)
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
	_, err := r.db.Exec(createCartQuery)
	if err != nil {
		return nil, errs.ErrCreatingCart
	}

	var id int
	err = r.db.QueryRow(maxCartIDQuery).Scan(&id)
	if err != nil {
		return nil, errs.ErrWrongCartID
	}

	cart := models.Cart{
		ID:    id,
		Items: []models.CartItem{},
	}

	return &cart, nil
}

func (r *CartRepository) AddItemToCart(cartID int, item models.CartItem) (*models.CartItem, error) {
	var cartCount int
	_ = r.db.QueryRow(cartCountQuery, cartID).Scan(&cartCount)
	if cartCount == 0 {
		return nil, errs.ErrCartNotFound
	}

	_, err := r.db.Exec(insertItemQuery, cartID, item.Product, item.Quantity)
	if err != nil {
		return nil, errs.ErrAddItemToCart
	}

	id, err := r.GetLastItemID()
	if err != nil {
		return nil, errs.ErrGettingLastItemID
	}
	item = models.CartItem{
		ID:       id,
		CartID:   cartID,
		Product:  item.Product,
		Quantity: item.Quantity,
	}

	return &item, nil
}

func (r *CartRepository) RemoveItemFromCart(cartID, itemID int) error {
	var itemCount, cartCount int
	_ = r.db.QueryRow(cartCountQuery, cartID).Scan(&cartCount)
	if cartCount == 0 {
		return errs.ErrCartNotFound
	}

	_ = r.db.QueryRow(itemCountQuery, itemID, cartID).Scan(&itemCount)
	if itemCount == 0 {
		return errs.ErrItemNotFound
	}

	_, err := r.db.Exec(deleteItemQuery, itemID, cartID)
	if err != nil {
		return errs.ErrRemoveItemFromCart
	}

	return nil
}

func (r *CartRepository) GetCart(cartID int) (*models.Cart, error) {
	var cartCount int
	_ = r.db.QueryRow(cartCountQuery, cartID).Scan(&cartCount)
	if cartCount == 0 {
		return nil, errs.ErrCartNotFound
	}

	rows, err := r.db.Query(selectItemQuery, cartID)
	if err != nil {
		return nil, errs.ErrItemNotFound
	}
	defer rows.Close()

	item := models.CartItem{}
	cart := models.Cart{}
	cart.ID = cartID
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.CartID, &item.Product, &item.Quantity)
		if err != nil {
			return nil, errs.ErrRowsScan
		}
		cart.Items = append(cart.Items, item)
	}

	return &cart, nil
}

func (r *CartRepository) GetLastItemID() (int, error) {
	var id int
	err := r.db.QueryRow(maxItemIDQuery).Scan(&id)
	if err != nil {
		return 0, errs.ErrGettingLastItemID
	}

	return id, nil
}
