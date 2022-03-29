package golangdatabase

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// func TestExecSql(t *testing.T) {
// 	db := GetConnecttion()
// 	defer db.Close()

// 	ctx := context.Background()
// 	script := "insert into customer(id,name)values('10','gatau males')"
// 	_, err := db.ExecContext(ctx, script) // function exec context tdk mengembalikan nilai
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("success insert new customer")
// }

func TestQuerySql(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	script := "select id, name from customer"
	rows, err := db.QueryContext(ctx, script) // function query context mengembalikan nilai
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}

	defer rows.Close()
}

func TestInsertSql(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	script := "insert into customer(id,name,email,balance,rating,birth_date,married)values('2','gunawan','gunawan@gmail.com',9000,9.5,'1995-05-22',true)"
	_, err := db.ExecContext(ctx, script) // function exec context tdk mengembalikan nilai
	if err != nil {
		panic(err)
	}
	fmt.Println("success insert new customer")
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	script := "select id, name, email, balance, rating,created_at, birth_date,married from customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &createdAt, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		fmt.Println("email :", email)
		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		fmt.Println("birth_date :", birthDate)
		fmt.Println("created_at :", createdAt)
		fmt.Println("married :", married)
	}

	defer rows.Close()
}
