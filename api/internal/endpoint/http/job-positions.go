package endpoint_http

import (
	"context"
	"net/http"

	"github.com/IndominusByte/danspro-be/api/internal/constant"
	jobpositionsentity "github.com/IndominusByte/danspro-be/api/internal/entity/job-positions"
	"github.com/creent-production/cdk-go/auth"
	"github.com/creent-production/cdk-go/parser"
	"github.com/creent-production/cdk-go/response"
	"github.com/creent-production/cdk-go/validation"
	"github.com/go-chi/chi/v5"
	"github.com/gomodule/redigo/redis"
)

type jobPositionsUsecaseIface interface {
	GetAllJobPosition(ctx context.Context, rw http.ResponseWriter, payload *jobpositionsentity.QueryParamAllJobPositionSchema)
	GetJobPositionById(ctx context.Context, rw http.ResponseWriter, jobpositionId string)
}

func AddJobPositions(r *chi.Mux, uc jobPositionsUsecaseIface, redisCli *redis.Pool) {
	r.Route("/job-positions", func(r chi.Router) {
		// protected route
		r.Group(func(r chi.Router) {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
					if err := auth.ValidateJWT(r.Context(), redisCli, "jwtRequired"); err != nil {
						response.WriteJSONResponse(rw, 401, nil, map[string]interface{}{
							constant.Header: err.Error(),
						})
						return
					}
					// Token is authenticated, pass it through
					next.ServeHTTP(rw, r)
				})
			})
			r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
				var p jobpositionsentity.QueryParamAllJobPositionSchema

				if err := validation.ParseRequest(&p, r.URL.Query()); err != nil {
					response.WriteJSONResponse(rw, 422, nil, map[string]interface{}{
						constant.Body: constant.FailedParseBody,
					})
					return
				}

				uc.GetAllJobPosition(r.Context(), rw, &p)
			})
			r.Get("/{job_position_id}", func(rw http.ResponseWriter, r *http.Request) {
				jobpositionId, _ := parser.ParsePathToStr("/job-positions/(.*)", r.URL.Path)

				uc.GetJobPositionById(r.Context(), rw, jobpositionId)
			})
		})
		// public route
	})
}
