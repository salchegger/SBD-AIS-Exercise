package repository

import (
	"errors"
	"fmt"
	"log/slog"
	"ordersystem/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Holds the connection to the PostgreSQL database via GORM
type DatabaseHandler struct {
	dbConn *gorm.DB
}

// Connects to the DB, auto-migrates tables (`Drink` & `Order`),
// and calls `prepopulate()` to insert test data if empty
func NewDatabaseHandler() (*DatabaseHandler, error) {
	slog.Info("Connecting to database")
	// connect to db
	dsn, err := getDsn()
	if err != nil {
		return nil, err
	}
	dbConn, err := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// create tables and migrate
	err = dbConn.AutoMigrate(&model.Drink{}, &model.Order{})
	if err != nil {
		return nil, err
	}
	// add test data to db
	err = prepopulate(dbConn)
	if err != nil {
		return nil, err
	}
	return &DatabaseHandler{dbConn: dbConn}, nil
}

// Builds the Postgres connection string from environment variables
func getDsn() (string, error) {
	dbUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return "", errors.New("environment variable 'POSTGRES_USER' is not set")
	}
	dbPw, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return "", errors.New("environment variable 'POSTGRES_PASSWORD' is not set")
	}
	dbName, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return "", errors.New("environment variable 'POSTGRES_DB' is not set")
	}
	dbPort, ok := os.LookupEnv("POSTGRES_TCP_PORT")
	if !ok {
		return "", errors.New("environment variable 'POSTGRES_TCP_PORT' is not set")
	}
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return "", errors.New("environment variable 'DB_HOST' is not set")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPw, dbName, dbPort)
	return dsn, nil
}

// Intended to insert initial drinks and orders
func prepopulate(dbConn *gorm.DB) error {
	// check if prepopulate has already run once
	var exists bool
	err := dbConn.Model(&model.Drink{}).
		Select("count(*) > 0").
		Find(&exists).
		Error
	if err != nil {
		return err
	}
	if exists {
		// don't prepopulate if has already run
		return nil
	}
	// create drink menu
	// todo create drinks
	drinks := []model.Drink{
		{Name: "Red Bull Peach", Price: 2.5, Description: "Peachy-flavored rocket fuel"},
		{Name: "Club-Mate", Price: 3.5, Description: "Hipster energy, slightly bitter"},
		{Name: "Espresso", Price: 2, Description: "Strong, bitter and dark"},
		{Name: "Pumpkin Spice Latte", Price: 4.5, Description: "Warm, cozy, sugary spice"},
		{Name: "Thai Iced Tea", Price: 5.0, Description: "Heavenly, icy refreshment bliss"},
	}
	if err := dbConn.Create(&drinks).Error; err != nil {
		return err
	}

	// todo create orders (only DrinkID + Amount)
	orders := []model.Order{
		{DrinkID: drinks[0].ID, Amount: 4},
		{DrinkID: drinks[1].ID, Amount: 2},
		{DrinkID: drinks[2].ID, Amount: 6},
		{DrinkID: drinks[3].ID, Amount: 3},
		{DrinkID: drinks[4].ID, Amount: 7},
	}
	if err := dbConn.Create(&orders).Error; err != nil {
		return err
	}

	// GORM documentation can be found here: https://gorm.io/docs/index.html

	return nil
}

// The following 3 functions:
// Retrieve drinks, orders, or summarized order totals from DB
func (db *DatabaseHandler) GetDrinks() (drinks []model.Drink, err error) {
	err = db.dbConn.Find(&drinks).Error
	if err != nil {
		return nil, err
	}
	return drinks, nil
}

func (db *DatabaseHandler) GetOrders() (orders []model.Order, err error) {
	err = db.dbConn.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

const totalledStmt = `SELECT drink_id, SUM(amount) AS total_amount_ordered FROM orders WHERE deleted_at IS NULL GROUP BY drink_id ORDER BY drink_id;`

func (db *DatabaseHandler) GetTotalledOrders() (totals []model.DrinkOrderTotal, err error) {
	err = db.dbConn.Raw(totalledStmt).Scan(&totals).Error
	if err != nil {
		return nil, err
	}
	return totals, nil
}

// Adds a new order to the DB
func (db *DatabaseHandler) AddOrder(order *model.Order) error {
	err := db.dbConn.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}
