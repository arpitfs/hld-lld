package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var Data = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
}

type Response struct {
	Data       []string
	Page       int
	PageSize   int
	TotalItems int
	TotalPages int
}

func pagination(data []string, page, pageSize int) ([]string, int, int) {
	totalItems := len(data)
	totalPages := totalItems / pageSize

	if page < 1 {
		page = 1
	} else if page > totalPages {
		page = totalPages
	}

	start := (page - 1) * pageSize
	end := start + pageSize

	if start > totalItems {
		start = totalItems
	}

	if end > totalItems {
		end = totalItems
	}

	return data[start:end], page, totalPages
}

func getPaginatedData(w http.ResponseWriter, r *http.Request) {
	pageParam := r.URL.Query().Get("page")
	pageSizeParam := r.URL.Query().Get("page_size")
	page, err := strconv.Atoi(pageParam)
	if page < 1 || err != nil {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil || pageSize < 1 {
		pageSize = 5
	}

	paginatedData, currentPage, totalPages := pagination(Data, page, pageSize)
	response := Response{
		Data:       paginatedData,
		Page:       currentPage,
		PageSize:   pageSize,
		TotalItems: len(Data),
		TotalPages: totalPages,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getCursorPaginatedData(w http.ResponseWriter, r *http.Request) {
	cursor := r.URL.Query().Get("cursor")
	limitParam := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitParam)
	if limit < 1 || err != nil {
		limit = 5
	}

	cursorPaginatedData, nextCursor, hasMore := CursorPagination(Data, cursor, limit)
	response := CursorResponse{
		Data:       cursorPaginatedData,
		HasMore:    hasMore,
		NextCursor: nextCursor,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/pagination", getPaginatedData)
	http.HandleFunc("/cursorpagination", getCursorPaginatedData)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
