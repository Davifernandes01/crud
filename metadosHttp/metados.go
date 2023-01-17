package metado

import (
	
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// struct para a criação do usuario no banco de dados
type Users struct {
	Id    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// inserir

//lendo a requisição
func Post(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("falha ao ler a requisição"))
		return
	}

	var user Users

	//convertendo o json em struct
	if err = json.Unmarshal(request, &user); err != nil {
		w.Write([]byte("erro ao converter users Json para struct"))

		return
	}

	//conectando ao bd
	db, err := banco.Con()
	if err != nil {
		w.Write([]byte("erro ao conectar com o bd"))

		return
	}

	defer db.Close()

	//criando o statement
	statement, err := db.Prepare("insert into users (nome, email) values (?,?)")
	if err != nil {
		w.Write([]byte("erro ao criar o users no banco de dados"))

		return
	}

	//fechando o statment
	defer statement.Close()

	insert, err := statement.Exec(user.Nome, user.Email)
	if err != nil {
		w.Write([]byte("erro ao executar o banco de dados"))

		return
	}

	//pegando o id do users
	id, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("erro ao pagar o id  no banco de dados"))

		return
	}

	//retornando o status code do http
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("users inserido com sucesso!! id: %d", id)))
}
