package usecase

import (
	"context"
	"gobrax-api/domain/entities"
)

func (uc *usecase) CreateTrucker(ctx context.Context, trucker entities.Trucker) (err error) {
	uc.logger.Info("usecase-create-trucker")
	err = uc.repoTrucker.CreateTrucker(ctx, trucker)
	return err
}

func (uc *usecase) GetTrucker(ctx context.Context) (truckers []entities.Trucker, err error) {
	uc.logger.Info("usecase-get-trucker")
	truckers, err = uc.repoTrucker.GetTrucker(ctx)
	return truckers, err
}

func (uc *usecase) GetTruckerByID(ctx context.Context, id int) (trucker entities.Trucker, err error) {
	uc.logger.Info("usecase-get-trucker-by-id")
	trucker, err = uc.repoTrucker.GetTruckerByID(ctx, id)
	return trucker, err
}

func (uc *usecase) UpdateTrucker(ctx context.Context, trucker entities.Trucker, id int) (err error) {
	uc.logger.Info("usecase-update-trucker")
	err = uc.repoTrucker.UpdateTrucker(ctx, trucker, id)
	return err
}

func (uc *usecase) DeleteTrucker(ctx context.Context, id int) (err error) {
	uc.logger.Info("usecase-delete-trucker")
	err = uc.repoTrucker.DeleteTrucker(ctx, id)
	return err
}
