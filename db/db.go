package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pismo/TransactionRoutineAPI/model"
)

func CreateDB() {
	url := model.Url()
	db := sqlx.MustConnect("mysql", url)

	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS transaction_db")
	if err != nil {
		panic(err)

	}

	_, err = db.Exec("USE transaction_db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS accounts(
		account_ID         INT(11) NOT NULL AUTO_INCREMENT
	   ,document_number   VARCHAR(200) NOT NULL
	   ,PRIMARY KEY (account_ID)
	 )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transactions(
		transaction_ID      INT(11) NOT NULL AUTO_INCREMENT
	   ,account_ID   	   	INT(11) NOT NULL
	   ,operationType_ID   	INT(11) NOT NULL
	   ,amount				DECIMAL(10,2) NOT NULL
	   ,eventDate 			DATETIME DEFAULT CURRENT_TIMESTAMP
	   ,PRIMARY KEY (transaction_ID)
	   ,FOREIGN KEY (account_ID) REFERENCES accounts (account_ID)
	 )ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;`)
	if err != nil {
		panic(err)
	}

}
