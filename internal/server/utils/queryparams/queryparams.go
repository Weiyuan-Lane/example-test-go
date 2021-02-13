package queryparams

import (
	"net/url"
	"strconv"
)

const (
	pageQueryKey    = "page"
	perPageQueryKey = "per_page"
)

type PageParams interface {
	Page() (int, error)
	PerPage() (int, error)
}

func New(queryValues url.Values) PageParams {
	return QueryParams{
		queryValues: queryValues,
	}
}

type QueryParams struct {
	queryValues url.Values
}

func (q QueryParams) Page() (int, error) {
	strVal := q.queryValues.Get(pageQueryKey)
	if strVal == "" {
		return 0, nil
	}

	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		return 0, err
	}

	return intVal, nil
}

func (q QueryParams) PerPage() (int, error) {
	strVal := q.queryValues.Get(perPageQueryKey)
	if strVal == "" {
		return 10, nil
	}

	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		return 0, err
	}

	return intVal, nil
}
