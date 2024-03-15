package usecase

import "context"

func (uc *usecase) Assign(ctx context.Context, vehicleID, truckerID int) (err error) {
	uc.logger.Info("usecase-assign-trucker")
	err = uc.repoAssign.Assign(ctx, vehicleID, truckerID)
	return
}
