package main

import (
	"fmt"
	"strconv"
)

type CursorResponse struct {
	Data       []string
	NextCursor string
	HasMore    bool
}

func CursorPagination(data []string, cursor string, limit int) ([]string, string, bool) {
	start := 0
	if cursor != "" {
		start, _ = strconv.Atoi(cursor)
	}
	end := start + limit
	if end > len(data) {
		end = len(data)
	}

	hasMore := end < len(data)
	nextCursor := ""
	if hasMore {
		nextCursor = strconv.Itoa(end)
	}
	fmt.Println(end, nextCursor, hasMore)

	return data[start:end], nextCursor, hasMore
}
