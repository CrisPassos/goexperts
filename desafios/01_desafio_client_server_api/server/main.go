package main

import (
	"1_Client_Server_API/shared"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	//Cria o servidor Mux para usar a /cotacao
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", CotacaoHandler)

	http.ListenAndServe(":8080", mux)
}

func CotacaoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/desafios")
	HandleError(err, "Não foi possível conectar com a base de dados")

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS cotacao (id varchar(255) PRIMARY KEY, bid varchar(255))")
	HandleError(err, "Erro ao criar tabela:")

	cotacao, error := GetCotacao(db)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(cotacao)
}

func GetCotacao(db *sql.DB) (*shared.Cotacao, error) {
	//cria o contexto
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	//faz a requisição
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//convert para JSON
	var cotacao shared.Cotacao
	err = json.NewDecoder(resp.Body).Decode(&cotacao)

	if err != nil {
		return nil, err
	}

	err = insertBID(ctx, db, cotacao.USDBRL.Bid)

	HandleError(err, "Erro ao salvar cotação no banco de dados:")

	fmt.Println("Cotação salva com sucesso:", cotacao.USDBRL.Bid)

	return &cotacao, nil
}

func insertBID(ctx context.Context, db *sql.DB, bid string) error {
	//insere o BID na base de dados
	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	stmt, err := db.PrepareContext(ctxDB, "INSERT INTO cotacao (id, bid) VALUES (?,?)")
	HandleError(err, "Erro no insert:")

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, uuid.New().String(), bid)

	return err
}

func HandleError(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		panic(err)
	}
}
