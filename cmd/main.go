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

	fromID := 2
	toID := 3

	account_1, _ := ar.FindByID(ctx, fromID)
	account_2, _ := ar.FindByID(ctx, toID)

	fmt.Printf("Account %d Balance: %d\n", fromID, account_1.Balance)
	fmt.Printf("Account %d Balance: %d\n", toID, account_2.Balance)

	err = au.Transfer(ctx, fromID, toID, 100)
	if err != nil {
		fmt.Println("Transfer failed:", err)
		return
	}

	// Output
	fmt.Println("Transfer completed successfully")

	account_1, _ = ar.FindByID(ctx, fromID)
	account_2, _ = ar.FindByID(ctx, toID)

	fmt.Printf("Account %d Balance: %d\n", fromID, account_1.Balance)
	fmt.Printf("Account %d Balance: %d\n", toID, account_2.Balance)
}
