package repository

import (
	"context"
	"gobrax-api/domain/entities"
)

func (repo *mysqlRepo) CreateVehicle(ctx context.Context, vehicle entities.Vehicle) (err error) {
	repo.logger.Info("repository-create-vehicle")

	query := `INSERT INTO public.vehicle (name, brand, color, year, plate) 
    VALUES ($1, $2, $3, $4, $5)`

	_, err = repo.DB.Exec(query, vehicle.Name, vehicle.Brand, vehicle.Color, vehicle.Year, vehicle.Plate)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	return nil
}

func (repo *mysqlRepo) GetVehicle(ctx context.Context) (vehicles []entities.Vehicle, err error) {
	repo.logger.Info("repository-get-vehicle")

	query := "SELECT * from public.vehicle"
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
		var vehicle entities.Vehicle

		err = rows.Scan(&vehicle.ID, &vehicle.Name, &vehicle.Brand, &vehicle.Color, &vehicle.Year, &vehicle.Plate)
		if err != nil {
			repo.logger.Error("repository-failed:", err.Error())
			return
		}

		vehicles = append(vehicles, vehicle)
	}

	return
}

func (repo *mysqlRepo) GetVehicleByID(ctx context.Context, id int) (vehicle entities.Vehicle, err error) {
	repo.logger.Info("repository-get-vehicle-by-id")

	query := "SELECT * from public.vehicle WHERE id = $1"
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

		err = rows.Scan(&vehicle.ID, &vehicle.Name, &vehicle.Brand, &vehicle.Color, &vehicle.Year, &vehicle.Plate)
		if err != nil {
			repo.logger.Error("repository-failed:", err.Error())
			return
		}
	}

	return
}

func (repo *mysqlRepo) UpdateVehicle(ctx context.Context, vehicle entities.Vehicle, id int) (err error) {
	repo.logger.Info("repository-update-vehicle")

	query := "UPDATE public.vehicle SET name = $1, brand = $2, color = $3, year = $4, plate = $5 WHERE id = $6"

	stmt, err := repo.DB.PrepareContext(ctx, query)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	_, err = stmt.ExecContext(ctx, vehicle.Name, vehicle.Brand, vehicle.Color, vehicle.Year, vehicle.Plate, id)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	return nil
}

func (repo *mysqlRepo) DeleteVehicle(ctx context.Context, id int) (err error) {
	repo.logger.Info("repository-delete-vehicle")

	query := "DELETE FROM public.vehicle WHERE id = $1"

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
