package usecase

import (
	"context"
	"gobrax-api/domain/entities"
)

func (uc *usecase) CreateVehicle(ctx context.Context, vehicle entities.Vehicle) (err error) {
	uc.logger.Info("usecase-create-vehicle")
	err = uc.repoVehicle.CreateVehicle(ctx, vehicle)
	return
}

func (uc *usecase) GetVehicle(ctx context.Context) (vehicles []entities.Vehicle, err error) {
	uc.logger.Info("usecase-get-vehicle")
	vehicles, err = uc.repoVehicle.GetVehicle(ctx)
	return
}

func (uc *usecase) GetVehicleByID(ctx context.Context, id int) (vehicle entities.Vehicle, err error) {
	uc.logger.Info("usecase-get-vehicle-by-id")
	vehicle, err = uc.repoVehicle.GetVehicleByID(ctx, id)
	return
}

func (uc *usecase) UpdateVehicle(ctx context.Context, vehicle entities.Vehicle, id int) (err error) {
	uc.logger.Info("update-vehicle")
	err = uc.repoVehicle.UpdateVehicle(ctx, vehicle, id)
	return
}

func (uc *usecase) DeleteVehicle(ctx context.Context, id int) (err error) {
	uc.logger.Info("delete-vehicle")
	err = uc.repoVehicle.DeleteVehicle(ctx, id)
	return
}
