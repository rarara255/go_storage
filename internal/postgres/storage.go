package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"productStorage/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"

	"go_storage/internal/models"
	"go_storage/internal/storage"
)


type PGStorage struct {
	db *sql.DB
}

type Config struct {
	Host	string
	Port	string
	User	string
	Password	string
	DBName	string
}

func NewPGStorage(cfg Config) (storage.ProductStorage, error) {
	//Data Source Name - строка которая содержит 
	//все параметры для подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
						cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)	

	connConfig, err := pgx.ParseConfig(dsn) // вернёт или *pgx.ConnConfig или error
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	db := stdlib.OpenDB(*connConfig)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	if err := createTable(db); err != nil {
		return nil, fmt.Errorf("create table error: %w", err)
	}

	return &PGStorage{db:	db}, nil
}


func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS products (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	sbin TEXT NOT NULL,
	expiry_date DATE NOT NULL
	);
	`
	_, err := db.Exec(query)
	return err
}

func (storage *PGStorage) Add(product *models.Product) {
	query := `INSERT INTO products(name, sbin, expiry_date) VALUES
	($1, $2, $3)`

	_, err := storage.db.Exec(query,product.Name, product.SBIN, product.ExpiryDate)
	if err != nil {
		log.Printf("Add product error %s: %v", product.Name, err)
	}
}

func (storage *PGStorage) GetAll () []*models.Product {
	//rows - *sql.Rows итератор по строкам 
	// результата запроса
	rows, err := storage.db.Query(`SELECT name, sbin, expiry_date FROM products`)
	if err != nil {
		log.Printf("query error: %v", err)
		return nil
	}

	defer rows.Close()

	var products []*models.Product

	for rows.Next() {
		var name, sbin string
		var expiryDate time.Time

		if err := rows.Scan(&name, &sbin, &expiryDate); err != nil {
			log.Printf("scan error: %v", err)
			continue
		}

		products = append(products, &models.Product{
			Name: name,
			SBIN: sbin,
			ExpiryDate: expiryDate,
		})
	}

	return products
}