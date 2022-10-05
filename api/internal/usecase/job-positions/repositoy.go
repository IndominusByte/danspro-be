package jobpositions

import (
	"context"

	authentity "github.com/IndominusByte/danspro-be/api/internal/entity/auth"
	jobpositionsentity "github.com/IndominusByte/danspro-be/api/internal/entity/job-positions"
)

type jobPositionsRepo interface {
	GetJobPositionById(ctx context.Context, jobpositionId string) (*jobpositionsentity.JobPosition, error)
	GetAllJobPositionPaginate(ctx context.Context,
		payload *jobpositionsentity.QueryParamAllJobPositionSchema) (*jobpositionsentity.JobPositionPaginate, error)
}

type authRepo interface {
	GetUserById(ctx context.Context, userId int) (*authentity.User, error)
}
