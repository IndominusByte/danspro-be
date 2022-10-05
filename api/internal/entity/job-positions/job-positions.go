package jobpositions

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type QueryParamAllJobPositionSchema struct {
	Page        int       `schema:"page" validate:"required,gte=1"`
	PerPage     int       `schema:"per_page" validate:"required,gte=1" db:"per_page"`
	Description string    `schema:"description" validate:"omitempty" db:"description"`
	Location    string    `schema:"location" validate:"omitempty" db:"location"`
	FullTime    null.Bool `schema:"full_time" validate:"omitempty"`
	Offset      int       `schema:"-" db:"offset"`
}

type JobPosition struct {
	Id          string    `json:"id" db:"id"`
	Kind        string    `json:"kind" db:"kind"`
	Url         string    `json:"url" db:"url"`
	Company     string    `json:"company" db:"company"`
	CompanyUrl  string    `json:"company_url" db:"company_url"`
	Location    string    `json:"location" db:"location"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	HowToApply  string    `json:"how_to_apply" db:"how_to_apply"`
	CompanyLogo string    `json:"company_logo" db:"company_logo"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type JobPositionPaginate struct {
	Data      []JobPosition `json:"data"`
	Total     int           `json:"total"`
	NextNum   null.Int      `json:"next_num"`
	PrevNum   null.Int      `json:"prev_num"`
	Page      int           `json:"page"`
	IterPages []null.Int    `json:"iter_pages"`
}
