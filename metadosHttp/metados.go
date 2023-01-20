package metado

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// struct para a criação do usuario no banco de dados
type Users struct {
	Id    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// inserir

//lendo a requisição
//adicionar
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

	//fechando a conexao com o bd
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

//buscar
func Get(w http.ResponseWriter, r *http.Request){

	//abrindo a conexeao
	db, err := banco.Con()
	if err != nil {
		w.Write([]byte("erro ao conectar com o banco"))
		return
	}

	//fechando
	defer db.Close()

	//exexutando uma query que devolve as linhas do banmco de dados
	rows, err:= db.Query("select * from users")
	if err != nil {
		w.Write([]byte("erro ao buscar users"))

		return
	}

	//fechando as rows
	defer rows.Close()

	//criando um slice para armazenar os users
	var users []Users

	for rows.Next(){ //"proxima linha"
	//para cada linha no banco, ira se criar um novo users
		var user Users

		if err := rows.Scan(&user.Id, &user.Nome, &user.Email);err != nil {
			w.Write([]byte("erro ao escanear o users"))

			return
	    }

		//assim que se criar um user, ira coloca-lo dentro do slice
		users = append(users, user)
	}


	//status 200
	w.WriteHeader(http.StatusOK)

	//convertendo nosso slice para json
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("erro ao converter  struct para json"))

		return
	}

}


//buscar por id
func GetById(w http.ResponseWriter, r *http.Request){

	parametro:= mux.Vars(r)

	// id, err:= strconv.ParseUint(parametro["id"],10,32)
	// if err != nil{
	// 	w.Write([]byte("erro ao converter o parametro para inteiro"))

	// 	return
	// }

	ID, erro := strconv.ParseUint(parametro["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, err:= banco.Con()
	if err != nil {
		w.Write([]byte("erro ao se conecetar com o banco"))

		return
	}

	defer db.Close()


	row, err:= db.Query("select * from users where id = ?", ID)
	if err != nil{
		w.Write([]byte("erro ao buscar usuario"))
		return
	}

	var user Users

	if row.Next(){

		if err:= row.Scan(&user.Id, &user.Nome, &user.Email); err != nil{
			w.Write([]byte("erro ao escanear usuario"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err !=nil{

		w.Write([]byte("erro ao converter o usuario para json"))
		return
	}

}
