package belajar_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	query := "INSERT into customer(id, name) values('budi', 'Budi')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert Customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	query := "select * from customer"
	data, err := db.QueryContext(ctx, query)
	defer data.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Select Customer")
	for data.Next() {
		var id, name string
		err := data.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id)
		fmt.Println("name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	query := "select id,name,email,balance,rating,created_at,birth_date,merried from customer"
	data, err := db.QueryContext(ctx, query)
	defer data.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Select Customer")
	for data.Next() {
		var id, name, email sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var created_at, birth_date sql.NullTime
		var merried sql.NullBool
		err := data.Scan(&id, &name, &email, &balance, &rating, &created_at, &birth_date, &merried)
		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id.String)
		fmt.Println("name:", name.String)
		if email.Valid {
			fmt.Println("email:", email.String)
		}
		fmt.Println("balance:", balance.Int32)
		fmt.Println("rating:", rating.Float64)
		fmt.Println("created_at:", created_at.Time)
		if birth_date.Valid {
			fmt.Println("birth_date:", birth_date.Time)
		}
		fmt.Println("merried:", merried.Bool)
		fmt.Println("========================================================")
	}
}

func TestSqlQueryParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "admin"
	password := "12345"
	query := "select username from user where username = ? and password = ? limit 1"
	data, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	if data.Next() {
		var username string
		err := data.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Berhasil Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSqlExecParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "admin2"
	password := "12345"
	query := "insert into user(username, password) values(?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert")
}