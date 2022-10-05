package jobpositions

import (
	"context"
	"net/http"

	"github.com/IndominusByte/danspro-be/api/internal/constant"
	jobpositionsentity "github.com/IndominusByte/danspro-be/api/internal/entity/job-positions"
	"github.com/creent-production/cdk-go/response"
	"github.com/creent-production/cdk-go/validation"
)

type JobPositionsUsecase struct {
	jobPositionsRepo jobPositionsRepo
	authRepo         authRepo
}

func NewJobPositionsUsecase(jobPositionRepo jobPositionsRepo, authRepo authRepo) *JobPositionsUsecase {
	return &JobPositionsUsecase{
		jobPositionsRepo: jobPositionRepo,
		authRepo:         authRepo,
	}
}

func (uc *JobPositionsUsecase) GetJobPositionById(ctx context.Context, rw http.ResponseWriter, jobpositionId string) {
	jobposition, err := uc.jobPositionsRepo.GetJobPositionById(ctx, jobpositionId)
	if err != nil {
		response.WriteJSONResponse(rw, 404, nil, map[string]interface{}{
			constant.App: "job position not found.",
		})
		return
	}

	response.WriteJSONResponse(rw, 200, jobposition, nil)

}

func (uc *JobPositionsUsecase) GetAllJobPosition(ctx context.Context, rw http.ResponseWriter,
	payload *jobpositionsentity.QueryParamAllJobPositionSchema) {

	if err := validation.StructValidate(payload); err != nil {
		response.WriteJSONResponse(rw, 422, nil, err)
		return
	}

	results, _ := uc.jobPositionsRepo.GetAllJobPositionPaginate(ctx, payload)

	response.WriteJSONResponse(rw, 200, results, nil)
}
