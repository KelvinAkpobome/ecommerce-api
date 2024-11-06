package config

import (
	"ecommerce-api/domain"
	"ecommerce-api/utils"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=ecomm password=123456 dbname=mynewdatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!", err)
	}
	// AutoMigrate the schema
	err = db.AutoMigrate(&domain.User{}, &domain.Product{}, &domain.Order{}, &domain.OrderProduct{})
	if err != nil {
		log.Fatal("Failed to migrate the database schema:", err)
	}

	// Seed the users table
	users := []domain.User{
		{Email: "John Doe", Password: "123456", IsAdmin: true},
		{Email: "Steve Jobs", Password: "123456", IsAdmin: false},
		{Email: "Bill Gates", Password: "123456", IsAdmin: false},
	}

	db.Create(&users)

	// Seed the products table
	products := []domain.Product{
		{Name: "Product A", Price: 10.00},
		{Name: "Product B", Price: 15.50},
		{Name: "Product C", Price: 7.25},
	}
	db.Create(&products)
    

    // Seed an order
    order := domain.Order{
        UserID: users[0].ID,
        Status: utils.Pending,
    }
    db.Create(&order)

    // Associate the order with products through the OrderProduct table
    for _, product := range products {
        orderProduct := domain.OrderProduct{
            OrderID:   order.ID,
            ProductID: product.ID,
        }
        db.Create(&orderProduct)
    }

	log.Println("Database seeding completed successfully!")
	return db
}
