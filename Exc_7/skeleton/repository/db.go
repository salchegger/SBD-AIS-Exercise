package repository

import (
	"errors"
	"fmt"
	"log/slog"
	"ordersystem/model"
	"ordersystem/secrets"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseHandler struct {
	dbConn *gorm.DB
}

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
	return &DatabaseHandler{dbConn: dbConn}, nil
}

func getDsn() (string, error) {
	dbUser, err := secrets.LoadSecretOrEnv("POSTGRES_USER")
	if err != nil {
		return "", err
	}
	dbPw, err := secrets.LoadSecretOrEnv("POSTGRES_PASSWORD")
	if err != nil {
		return "", err
	}
	dbName, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return "", errors.New("environment variable 'POSTGRES_DB' is not set")
	}
	dbPort, ok := os.LookupEnv("PGPORT")
	if !ok {
		return "", errors.New("environment variable 'PGPORT' is not set")
	}
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return "", errors.New("environment variable 'DB_HOST' is not set")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		dbHost, dbUser, dbPw, dbName, dbPort)
	return dsn, nil
}

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

func (db *DatabaseHandler) GetOrder(id uint) (dbOrder *model.Order, err error) {
	err = db.dbConn.
		Where("id = ?", id).
		First(&dbOrder).Error
	if err != nil {
		return nil, err
	}
	return dbOrder, nil
}

const totalledStmt = `SELECT drink_id, SUM(amount) AS total_amount_ordered FROM orders WHERE deleted_at IS NULL GROUP BY drink_id ORDER BY drink_id;`

func (db *DatabaseHandler) GetTotalledOrders() (totals []model.DrinkOrderTotal, err error) {
	err = db.dbConn.Raw(totalledStmt).Scan(&totals).Error
	if err != nil {
		return nil, err
	}
	return totals, nil
}

func (db *DatabaseHandler) AddOrder(order *model.Order) (*model.Order, error) {
	err := db.dbConn.Create(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
