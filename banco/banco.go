package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //drive sql
)

// funcao para conectar com o banco mysql
func Con() (*sql.DB, error) {

	users := "golang"                                  //usuario do banco
	pass := ":golang@"                                 // senha
	namedb := "/devbook"                               // nome do banco
	config := "?charset=utf8&parseTime=true&loc=Local" // configuracoes adicionais

	//"golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local
	stringCompleta := users + pass + namedb + config

	db, err := sql.Open("mysql", stringCompleta) //abrinco a conexao
	if err != nil {
		return nil, err //vendo se tem algum erro
	}

	if err := db.Ping(); err != nil { //vendo se nao passou a string de conexao errada
		return nil, err
	}

	return db, nil
}
