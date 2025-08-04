package main

import (
	"context"
	"fmt"

	"github.com/tf63/go-gorm-tx-sample/internal/anti-pattern/application"
	"github.com/tf63/go-gorm-tx-sample/internal/anti-pattern/infrastracture"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// DB setup
	db, err := gorm.Open(mysql.Open("root:rootpassword@tcp(127.0.0.1:3306)/sample"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// DI
	ar := infrastracture.NewAccountRepository(db)
	au := application.NewAccountUsecase(ar)

	// Request
	ctx := context.Background()
	err = au.Transfer(ctx, "1", "2", 1000)
	if err != nil {
		fmt.Println("Transfer failed:", err)
		return
	}

	// Output
	fmt.Println("Transfer completed successfully", "from_account_id", "to_account_id", 1000)
}
