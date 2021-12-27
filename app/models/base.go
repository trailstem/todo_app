package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo_app/config"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

//テーブル作成のコード書く
var Db *sql.DB

var err error

/*
//テーブル名宣言
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)
*/

//テーブル作成
func init() {

	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}
	/*
		Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
		if err != nil {
			log.Fatalln(err)
		}

		//Users tabele がなければ作成する
		//%sでテーブルネームの定数usersが入る
		cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			name STRING,
			email STRING,
			password STRING,
			created_at DATETIME)`, tableNameUser)
		//コマンド実行
		Db.Exec(cmdU)

		//Todoモデルのテーブル作成
		//%sでテーブルネームの定数todoが入る

		cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT,
			user_id INTEGER,
			created_at DATETIME)`, tableNameTodo)

		Db.Exec(cmdT)

		cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			uuid STRING NOT NULL UNIQUE,
			email STRING,
			user_id INTEGER,
			created_at DATETIME)`, tableNameSession)
		Db.Exec(cmdS)
	*/
}

//uuid.UUID パッケージのUUID型を使用している
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
