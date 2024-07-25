package user

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/client/db"
	"github.com/DenisCom3/m-auth/internal/model"
	"github.com/DenisCom3/m-auth/internal/repository"
	"github.com/DenisCom3/m-auth/internal/repository/user/converter"
	modelRepo "github.com/DenisCom3/m-auth/internal/repository/user/model"
	sq "github.com/Masterminds/squirrel"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func New(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := psql.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "userRepo.create",
		QueryRaw: query,
	}

	var user modelRepo.User

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Create(ctx context.Context, user model.CreateUser) (int64, error) {
	query, args, err := psql.Insert(tableName).
		Columns(nameColumn, passwordColumn, emailColumn, roleColumn).
		Values(user.Info.Name, user.Password, user.Info.Email, user.Info.Role).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "userRepo.Create",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Update(ctx context.Context, user model.UpdateUser) error {

	query, args, err := psql.Update(tableName).
		Set("name", user.Info.Name).
		Set("email", user.Info.Email).
		Set("role", user.Info.Role).
		Where(sq.Eq{"id": user.ID}).
		ToSql()

	if err != nil {
		return err
	}

	q := db.Query{Name: "userRepo.Update", QueryRaw: query}

	err = r.db.DB().QueryRowContext(ctx, q, args).Scan()

	if err != nil {
		return err
	}

	return nil

}

func (r *repo) Delete(ctx context.Context, id int64) error {

	sql, args, err := psql.Delete(tableName).Where(sq.Eq{"id": id}).ToSql()

	if err != nil {
		return err
	}

	q := db.Query{Name: "userRepo.Delete", QueryRaw: sql}

	err = r.db.DB().QueryRowContext(ctx, q, args).Scan()

	if err != nil {
		return err
	}

	return nil

}
