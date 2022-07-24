package users

import (
	"aero-internship/internal/domain/entity/tokens"
	"aero-internship/internal/domain/entity/users"
	"aero-internship/pkg/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Users struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewUsers(cfg *config.Config, db *sqlx.DB) *Users {
	return &Users{cfg: cfg, db: db}
}

func (u *Users) MakeRegistrationTxn(cfg *config.Config, userDTO users.UserDTO) error {

	//начало транзакции
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}

	q := `
		insert into Users (name,email,passwordHash,isadmin,registeredFrom) 
		values ($1,$2,$3,$4,$5)
	`
	_, err = tx.Exec(
		q,
		userDTO.Name,
		userDTO.Email,
		userDTO.Password,
		userDTO.IsAdmin,
		time.Now().Unix(),
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (u *Users) GetTokenDTOFromUserDTO(cfg *config.Config, userDTO *users.UserDTO) (*tokens.TokenDTO, error) {

	//начало транзакции
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	var tokenDTO tokens.TokenDTO

	log.Printf("going to the DB and scan data for tokenDTO\n")

	selectUserByEmail := fmt.Sprintf("select id, isadmin from Users where email=$1")
	row := tx.QueryRow(selectUserByEmail, userDTO.Email)
	if err := row.Scan(&tokenDTO.UserId, &tokenDTO.IsAdmin); err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &tokenDTO, nil
}

func (u *Users) GetUserDTObyEmail(cfg *config.Config, email string) (*users.User, error) {
	//начало транзакции
	tx, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	var user users.User

	selectUserByEmail := fmt.Sprintf("select name,email,passwordHash,isadmin from Users where email=$1")
	row := tx.QueryRow(selectUserByEmail, email)
	if err := row.Scan(&user.Name, &user.Email, &user.Password, &user.IsAdmin); err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &user, nil
}
