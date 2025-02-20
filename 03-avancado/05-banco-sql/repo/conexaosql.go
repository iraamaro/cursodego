package repo

import (
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	/*
		github.com/go-sql-driver/mysql não é usado diretamente pela aplicação
	*/
	_ "github.com/go-sql-driver/mysql"
)

// variavel singleton que armazena a conexao
var db *sqlx.DB

// AbreConexaoComBancoDeDadosSQL funcao que abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (db *sqlx.DB, err error) {
	envPath := "./.env"
	absPath, err := filepath.Abs(envPath)
	if err != nil {
		log.Fatalf("Erro ao obter caminho absoluto: %v", err)
	}

	// Carrega as variáveis do arquivo .env
	err = godotenv.Load(absPath)
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
	// Obtém as credenciais do banco de dados
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	err = nil
	dsn := dbUser + ":" + dbPass +
		"@tcp(" + dbHost + ":" + dbPort + ")/" + dbName +
		"?charset=utf8&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro na conexao: ", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		log.Println("[AbreConexaoComBancoDeDadosSQL] Erro no ping na conexao: ", err.Error())
		return
	}
	return
}

// GetDBConnection Obtem a conexao com o banco de dados
func GetDBConnection() (localdb *sqlx.DB, err error) {
	if db == nil {
		db, err = AbreConexaoComBancoDeDadosSQL()
		if err != nil {
			log.Println("[GetDBConnection] Erro na conexao: ", err.Error())
			return
		}
	}
	err = db.Ping()
	if err != nil {
		log.Println("[GetDBConnection] Erro no ping na conexao: ", err.Error())
		return
	}
	localdb = db
	return
}
