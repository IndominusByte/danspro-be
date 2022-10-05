package jobpositions

import (
	"context"
	"fmt"

	jobpositionsentity "github.com/IndominusByte/danspro-be/api/internal/entity/job-positions"
	"github.com/creent-production/cdk-go/pagination"
	"github.com/jmoiron/sqlx"
)

type RepoJobPositions struct {
	db      *sqlx.DB
	queries map[string]string
	execs   map[string]string
}

var queries = map[string]string{
	"getJobPositionByDynamic": `SELECT id, kind, url, company, company_url, location, title, description, how_to_apply, company_logo, created_at FROM account.job_positions`,
}
var execs = map[string]string{}

func New(db *sqlx.DB) (*RepoJobPositions, error) {
	rp := &RepoJobPositions{
		db:      db,
		queries: queries,
		execs:   execs,
	}

	err := rp.Validate()
	if err != nil {
		return nil, err
	}

	return rp, nil
}

// Validate will validate sql query to db
func (r *RepoJobPositions) Validate() error {
	for _, q := range r.queries {
		_, err := r.db.PrepareNamedContext(context.Background(), q)
		if err != nil {
			return err
		}
	}

	for _, e := range r.execs {
		_, err := r.db.PrepareNamedContext(context.Background(), e)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *RepoJobPositions) GetJobPositionById(ctx context.Context, jobpositionId string) (*jobpositionsentity.JobPosition, error) {
	var t jobpositionsentity.JobPosition
	stmt, _ := r.db.PrepareNamedContext(ctx, r.queries["getJobPositionByDynamic"]+" WHERE id = :id")

	return &t, stmt.GetContext(ctx, &t, jobpositionsentity.JobPosition{Id: jobpositionId})
}

func (r *RepoJobPositions) GetAllJobPositionPaginate(ctx context.Context,
	payload *jobpositionsentity.QueryParamAllJobPositionSchema) (*jobpositionsentity.JobPositionPaginate, error) {

	var results jobpositionsentity.JobPositionPaginate

	query := r.queries["getJobPositionByDynamic"] + " WHERE 1=1"
	if len(payload.Description) > 0 {
		query += ` AND lower(description) LIKE '%'|| lower(:description) ||'%'`
	}
	if len(payload.Location) > 0 {
		query += ` AND lower(location) LIKE '%'|| lower(:location) ||'%'`
	}
	if payload.FullTime.Valid {
		query += ` AND kind = 'Full Time'`
	}

	query += ` ORDER BY id DESC`

	// pagination
	var count struct{ Total int }
	stmt_count, _ := r.db.PrepareNamedContext(ctx, fmt.Sprintf("SELECT count(*) AS total FROM (%s) AS anon_1", query))
	err := stmt_count.GetContext(ctx, &count, payload)
	if err != nil {
		return &results, err
	}
	payload.Offset = (payload.Page - 1) * payload.PerPage

	// results
	query += ` LIMIT :per_page OFFSET :offset`
	stmt, _ := r.db.PrepareNamedContext(ctx, query)
	err = stmt.SelectContext(ctx, &results.Data, payload)
	if err != nil {
		return &results, err
	}

	paginate := pagination.Paginate{Page: payload.Page, PerPage: payload.PerPage, Total: count.Total}
	results.Total = paginate.Total
	results.NextNum = paginate.NextNum()
	results.PrevNum = paginate.PrevNum()
	results.Page = paginate.Page
	results.IterPages = paginate.IterPages()

	return &results, nil
}
