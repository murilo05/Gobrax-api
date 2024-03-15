package repository

import "context"

func (repo *mysqlRepo) Assign(ctx context.Context, vehicleID, truckerID int) (err error) {
	repo.logger.Info("repository-assign")

	query := `INSERT INTO public.assign (vehicle_id, trucker_id) 
    VALUES ($1, $2)`
	_, err = repo.DB.Exec(query, vehicleID, truckerID)
	if err != nil {
		repo.logger.Error("repository-failed:", err.Error())
		return
	}

	return nil
}
