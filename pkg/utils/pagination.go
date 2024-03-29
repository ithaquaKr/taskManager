package utils

import (
	"math"
	"strconv"
)

const (
	// TODO: make it configurable
	defaultPageSize = 10
)

// Paginate query params
type PaginationQuery struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
}

// Paginate query response
type PaginationResponse struct {
	TotalCount   int  `json:"total_count"`
	TotalPage    int  `json:"total_page"`
	HasMore      bool `json:"has_more"`
	NextPage     int  `json:"next_page"`
	PreviousPage int  `json:"previous_page"`
}

// Set page size
func (q *PaginationQuery) SetPageSize(sizeQuery string) error {
	if sizeQuery == "" {
		q.PageSize = defaultPageSize
		return nil
	}
	n, err := strconv.Atoi(sizeQuery)
	if err != nil {
		return err
	}
	q.PageSize = n

	return nil
}

// Set page number
func (q *PaginationQuery) SetPageNumber(pageQuery string) error {
	if pageQuery == "" {
		q.PageSize = 0
		return nil
	}
	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}
	q.Page = n

	return nil
}

// Get Page number
func (q *PaginationQuery) GetPageNumber() int {
	return q.Page
}

// Get Page size
func (q *PaginationQuery) GetPageSize() int {
	return q.PageSize
}

// Get offset
func (q *PaginationQuery) GetOffset() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.PageSize
}

// Get limit
func (q *PaginationQuery) GetLimit() int {
	return q.PageSize
}

// Get total number of pages
func GetTotalPages(totalItems, pageSize int) int {
	return int(math.Ceil(float64(totalItems) / float64(pageSize)))
}

// Check if there is a next page
func HasNextPage(currentPage, totalPages, pageSize int) bool {
	return currentPage < totalPages/pageSize
}

// Get next page
func GetNextPage(currentPage int) int {
	return currentPage + 1
}

// Get previous page
func GetPreviousPage(currentPage int) int {
	if currentPage == 0 {
		return 0
	}
	return currentPage - 1
}

// Get has more
func GetHasMore(currentPage, totalCount, pageSize int) bool {
	return currentPage < totalCount/pageSize
}
