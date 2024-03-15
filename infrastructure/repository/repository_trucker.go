package repository

import (
	"context"
	"gobrax-api/domain/entities"
)

func (repo *mysqlRepo) CreateTrucker(ctx context.Context, trucker entities.Trucker) (err error) {
	repo.logger.Info("repository-create-trucker")

	query := `INSERT INTO public.trucker (name, age, nationality) 
    VALUES ($1, $2, $3)`
	_, err = repo.DB.Exec(query, trucker.Name, trucker.Age, trucker.Nationality)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	return nil
}

func (repo *mysqlRepo) GetTrucker(ctx context.Context) (truckers []entities.Trucker, err error) {
	repo.logger.Info("repository-get-trucker")

	query := "SELECT * from public.trucker"
	stmt, err := repo.DB.PrepareContext(ctx, query)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	for rows.Next() {
		var trucker entities.Trucker

		err = rows.Scan(&trucker.ID, &trucker.Name, &trucker.Age, &trucker.Nationality)
		if err != nil {
			repo.logger.Error("repository-failed:", err.Error())
			return
		}

		truckers = append(truckers, trucker)
	}

	return
}

func (repo *mysqlRepo) GetTruckerByID(ctx context.Context, id int) (trucker entities.Trucker, err error) {
	repo.logger.Info("repository-get-trucker-by-id")

	query := "SELECT * from public.trucker WHERE id = $1"
	stmt, err := repo.DB.PrepareContext(ctx, query)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	if rows.Next() {

		err = rows.Scan(&trucker.ID, &trucker.Name, &trucker.Age, &trucker.Nationality)
		if err != nil {
			repo.logger.Error("repository-failed:", err.Error())
			return
		}
	}

	return
}

func (repo *mysqlRepo) UpdateTrucker(ctx context.Context, trucker entities.Trucker, id int) (err error) {
	repo.logger.Info("repository-update-trucker")

	query := "UPDATE public.trucker SET name = $1, age = $2, nationality = $3 WHERE id = $4"

	stmt, err := repo.DB.PrepareContext(ctx, query)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	_, err = stmt.ExecContext(ctx, trucker.Name, trucker.Age, trucker.Nationality, id)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	return nil
}

func (repo *mysqlRepo) DeleteTrucker(ctx context.Context, id int) (err error) {
	repo.logger.Info("repository-delete-trucker")

	query := "DELETE FROM public.trucker WHERE id = $1"

	stmt, err := repo.DB.PrepareContext(ctx, query)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	return nil
}
