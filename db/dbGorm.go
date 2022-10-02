package dbSql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "root:@tcp(localhost:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"
var DbGorm = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en coxion a base de datos", err)
		panic(err)
	} else {
		fmt.Println("Conexion a base de datos OK")
		return db
	}
}()

const url = "root:@tcp(localhost:3306)/go_db"

var dbSql *sql.DB

func Connect() {
	// dbGorm.
	conx, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	// fmt.Println("conexion a base de datos exitosa")
	dbSql = conx
}

func Close() {
	dbSql.Close()
}

// VERIFICA LA CONEXION A BASE DE DATOS
func VerifyPing() {
	if err := dbSql.Ping(); err != nil {
		panic(err)
	}
}

func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		println("Error:", err)
		return false
	}
	exist := rows.Next()
	if exist {
		println("Ya existe la tabla:", tableName)
	}
	return exist

}

// CREA UNA TABLA
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		ok, err := Exec(schema)
		if err != nil {
			panic(err)
		}
		fmt.Println(ok)
	}
}

func Exec(query string, arg ...interface{}) (sql.Result, error) {
	Connect()
	res, err := dbSql.Exec(query, arg...)
	Close()

	if err != nil {
		fmt.Println(err)
	}
	return res, err
}

func Query(query string, arg ...interface{}) (*sql.Rows, error) {
	Connect()
	res, err := dbSql.Query(query, arg...)
	Close()
	if err != nil {

		fmt.Println(err)
	}
	return res, err
}
