package interfaces

import (
	"context"
	"gobrax-api/domain/entities"
)

type RepositoryTrucker interface {
	CreateTrucker(ctx context.Context, trucker entities.Trucker) error
	GetTrucker(ctx context.Context) (truckers []entities.Trucker, err error)
	GetTruckerByID(ctx context.Context, id int) (trucker entities.Trucker, err error)
	UpdateTrucker(ctx context.Context, trucker entities.Trucker, id int) (err error)
	DeleteTrucker(ctx context.Context, id int) (err error)
}

type RepositoryVehicle interface {
	CreateVehicle(ctx context.Context, vehicle entities.Vehicle) error
	GetVehicle(ctx context.Context) (vehicles []entities.Vehicle, err error)
	GetVehicleByID(ctx context.Context, id int) (vehicle entities.Vehicle, err error)
	UpdateVehicle(ctx context.Context, vehicle entities.Vehicle, id int) (err error)
	DeleteVehicle(ctx context.Context, id int) (err error)
}

type RepositoryAssign interface {
	Assign(ctx context.Context, vehicleID, truckerID int) error
}
