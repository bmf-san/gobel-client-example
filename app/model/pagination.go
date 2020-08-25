package model

import (
	"net/http"
	"strconv"
)

// A Pagination represents the singular of pagination.
type Pagination struct {
	PaginationCount     int        `json:"pagination_count"`
	PaginationPageCount []struct{} `json:"pagination_pagecount"`
	PaginationPage      int        `json:"pagination_page"`
	PaginationLimit     int        `json:"pagination_limit"`
}

// Convert converts response header to pagination struct.
func (p *Pagination) Convert(header http.Header) error {
	paginationCount, err := strconv.Atoi(header.Get("Pagination-Count"))
	if err != nil {
		return err
	}

	paginationPageCount, err := strconv.Atoi(header.Get("Pagination-PageCount"))
	if err != nil {
		return err
	}

	paginationPage, err := strconv.Atoi(header.Get("Pagination-Page"))
	if err != nil {
		return err
	}

	paginationLimit, err := strconv.Atoi(header.Get("Pagination-Limit"))
	if err != nil {
		return err
	}

	p.PaginationCount = paginationCount
	// Since range starts from 0 on the template side, 1 is added.
	p.PaginationPageCount = make([]struct{}, paginationPageCount+1)
	p.PaginationPage = paginationPage
	p.PaginationLimit = paginationLimit

	return nil
}
