package golangdatabase

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	d, err := sql.Open("mysql", "gatau:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	// gunakan db
	defer d.Close()
}
