package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	databases "mobee-test/db"
	router "mobee-test/libs/routers"
	"mobee-test/usecases"

	c "mobee-test/controllers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"mobee-test/repositories"
)

var (
	httpRouter router.Router          = router.NewChiRouter()
	dbRepoConn databases.DatabaseRepo = databases.NewPostgresRepo()
)

func main() {
	// Load envornment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	// Read environment variables
	port := os.Getenv("PORT")
	appName := os.Getenv("APP_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	timeoutContext := time.Duration(5 * time.Second)
	httpResult := router.NewResultset()

	portDB, _ := strconv.Atoi(dbPort)
	portConnect, _ := strconv.Atoi(port)

	db, err := dbRepoConn.Connect(dbHost, portDB, dbUser, dbPassword, dbName, sslMode)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repositories.NewUserRepository(db)
	walletRepo := repositories.NewWalletRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	userUC := usecases.NewUserUseCase(timeoutContext, userRepo, db)
	walletUC := usecases.NewWalletUseCase(timeoutContext, walletRepo, db)
	transactioUC := usecases.NewTransactionUseCase(timeoutContext, transactionRepo, db)

	user := c.NewUserController(userUC, httpResult)
	waller := c.NewWalletController(walletUC, httpResult)
	transaction := c.NewTransactionController(transactioUC, httpResult)

	// Define a simple GET route
	httpRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Welcome to %s!", appName)
	})

	httpRouter.Post("/users", user.CreateUser)
	httpRouter.Post("/wallets", waller.CreateWallet)
	httpRouter.Post("/deposit", transaction.Deposit)

	// Start the server
	fmt.Printf("Starting %s on port %d\n", appName, portConnect)

	httpRouter.Run(portConnect, appName)
}
