package db

// TODO: currently not useable
import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("teststin")
	enverr := godotenv.Load(".env")

	if enverr != nil{
  		log.Fatalf("Error loading .env file: %s", enverr)
 	}

	fmt.Println(os.Getenv("DATABASE_URL"))

	m, err := migrate.New(
		"file://migrations",
		os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}