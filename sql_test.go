package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSqlAja(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	script := "insert into customer(id,name)values('3','dimas subekti')"
	_, err := db.ExecContext(ctx, script) // function exec context tdk mengembalikan nilai
	if err != nil {
		panic(err)
	}
	fmt.Println("success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	script := "select id, name from customer"
	rows, err := db.QueryContext(ctx, script) // function query context mengembalikan nilai
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}

}

func TestInsertSql(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	script := "insert into customer(id,name,email,balance,rating,birth_date,married)values('1','rizky','rizky@gmail.com' ,9000,9.5,'1997-03-22' ,true)"
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

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString // data yg bisa null
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &createdAt, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid {
			fmt.Println("email :", email.String)
		}

		fmt.Println("balance :", balance)
		fmt.Println("rating :", rating)
		if birthDate.Valid {
			fmt.Println("birth_date :", birthDate.Time)
		}
		fmt.Println("created_at :", createdAt)
		fmt.Println("married :", married)
	}

}

func TestSqlInjection(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()

	//username := "admin'; #"
	//password := "salah"
	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("succes login", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("succes login", username)
	} else {
		fmt.Println("gagal login")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	username := "wahyu'; DROP TABLE USER; #"
	password := "wahyu"

	ctx := context.Background()
	script := "INSERT INTO user (username,password) VALUES (?,?)"
	_, err := db.ExecContext(ctx, script, username, password) // function exec context tdk mengembalikan nilai
	if err != nil {
		panic(err)
	}
	fmt.Println("success insert new user")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	email := "wahyu2@gmail.com"
	comment := "test ke DUA"

	ctx := context.Background()
	script := "INSERT INTO comments (email,comment) VALUES (?,?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last Insert ID : ", id)
}

func TestQueryAutoIncrement(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id,email,comment FROM comments"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var email, comment string
		err = rows.Scan(&id, &email, &comment)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("Id :", id)
		fmt.Println("Email :", email)
		fmt.Println("Comment :", comment)
	}
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments (email,comment) VALUES (?,?)"
	stmt, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 30; i++ {
		email := "rizky" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)
		execContext, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := execContext.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id ", insertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnecttion()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	// do transaction
	script := "INSERT INTO comments (email,comment) VALUES (?,?)"
	for i := 0; i < 10; i++ {
		email := "rizky" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)
		execContext, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}
		insertId, err := execContext.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment Id ", insertId)
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
