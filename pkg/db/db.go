package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"go-postgres/pkg/models"
)

type PostgresRepo struct {
	ConnString string
}

func (r PostgresRepo) createConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), r.ConnString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to db.  %v", err)
	}
	return conn, nil
}

func (r PostgresRepo) InsertUser(user models.User) (int, error) {
	db, err := r.createConnection()
	if err != nil {
		return 0, fmt.Errorf("unable to connect to db.  %v", err)
	}

	ctx := context.Background()
	defer db.Close(ctx)

	sqlStatement := `INSERT INTO users (name, rating) VALUES ($1, $2) RETURNING id`

	var id int

	// Scan function will save the insert id in the id
	err = db.QueryRow(ctx, sqlStatement, user.Name, user.Rating).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	return id, nil
}

func (r PostgresRepo) GetUser(id int) (models.User, error) {
	db, err := r.createConnection()
	if err != nil {
		return models.User{}, fmt.Errorf("unable to connect to db.  %v", err)
	}

	ctx := context.Background()
	defer db.Close(ctx)

	var user models.User

	sqlStatement := `SELECT * FROM users WHERE id=$1`

	row := db.QueryRow(ctx, sqlStatement, id)

	err = row.Scan(&user.ID, &user.Name, &user.Rating)

	if err != nil {
		return models.User{}, fmt.Errorf("unable to scan the row. %v", err)
	}

	return user, nil
}

func (r PostgresRepo) GetAllUsers() ([]models.User, error) {
	db, err := r.createConnection()
	if err != nil {
		return []models.User{}, fmt.Errorf("unable to connect to db.  %v", err)
	}

	ctx := context.Background()
	defer db.Close(ctx)

	var users []models.User

	sqlStatement := `SELECT * FROM users`

	rows, err := db.Query(ctx, sqlStatement)

	if err != nil {
		return []models.User{}, fmt.Errorf("unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Rating)

		if err != nil {
			return []models.User{}, fmt.Errorf("unable to scan the row. %v", err)
		}

		users = append(users, user)

	}

	return users, nil
}

func (r PostgresRepo) UpdateUser(id int, user models.User) (int64, error) {
	db, err := r.createConnection()
	if err != nil {
		return 0, fmt.Errorf("unable to connect to db.  %v", err)
	}

	ctx := context.Background()
	defer db.Close(ctx)

	sqlStatement := `UPDATE users SET name=$2, rating=$3 WHERE id=$1`

	res, err := db.Exec(ctx, sqlStatement, id, user.Name, user.Rating)

	if err != nil {
		return 0, fmt.Errorf("unable to execute the query. %v", err)
	}

	rowsAffected := res.RowsAffected()

	return rowsAffected, nil
}

func (r PostgresRepo) DeleteUser(id int) (int64, error) {
	db, err := r.createConnection()
	if err != nil {
		return 0, fmt.Errorf("unable to connect to db.  %v", err)
	}

	ctx := context.Background()
	defer db.Close(ctx)

	sqlStatement := `DELETE FROM users WHERE id=$1`

	res, err := db.Exec(ctx, sqlStatement, id)

	if err != nil {
		return 0, fmt.Errorf("unable to execute the query. %v", err)
	}

	rowsAffected := res.RowsAffected()

	if err != nil {
		return 0, fmt.Errorf("error while checking the affected rows. %v", err)
	}

	return rowsAffected, nil
}
