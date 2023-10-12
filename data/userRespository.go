package data

import (
	"yakki/blogo/models"
)

type UserRepositoryInterface interface {
	GetUsers() (*models.UserList, error)
	// GetUser(id int) (*models.User, error)
	CreateUser(user models.NewUser) error
}

type UserRepository struct {
	ds Datasource
}

func (repo *UserRepository) GetUsers() (*models.UserList, error) {
	users := &models.UserList{}
	sql := "select * from b_user"

	rows, err := repo.ds.GetDB().Query(sql)

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var item models.User

		err := rows.Scan(&item.Id, &item.Name, &item.Username, &item.Email)

		if err != nil {
			return users, err
		}

		users.Users = append(users.Users, item)

	}

	return users, nil
}

func (repo *UserRepository) CreateUser(user models.NewUser) error {
	sql := "insert into b_user (name, username, email) values ($1, $2, $3) returning id"
	id := 0
	err := repo.ds.GetDB().QueryRow(sql, user.Name, user.Username, user.Email).Scan(&id)
	println(id)
	if err != nil {
		return err
	}

	return nil
}

func ProvideUserRepository(ds Datasource) UserRepositoryInterface {
	return &UserRepository{ds}
}
