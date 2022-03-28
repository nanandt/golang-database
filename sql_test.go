package golangdatabase

import (
	"context"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	script := "insert into customer(id,name)values('2','bagus')"
	_, err := db.ExecContext(ctx, script) // function exec context tdk mengembalikan nilai
	if err != nil {
		panic(err)
	}
	fmt.Println("success insert new customer")
}