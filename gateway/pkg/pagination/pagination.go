package pagination

import "strconv"

type PageParams struct {
	Page     int
	PageSize int
}

func GetPageParams(c interface{}) PageParams {
	params := PageParams{
		Page:     1,
		PageSize: 20,
	}

	if page := GetIntParam(c, "page"); page > 0 {
		params.Page = page
	}

	if pageSize := GetIntParam(c, "page_size"); pageSize > 0 {
		if pageSize > 100 {
			pageSize = 100
		}
		params.PageSize = pageSize
	}

	return params
}

func GetIntParam(c interface{}, key string) int {
	switch v := c.(type) {
	case map[string]interface{}:
		if val, ok := v[key]; ok {
			switch v := val.(type) {
			case float64:
				return int(v)
			case string:
				if i, err := strconv.Atoi(v); err == nil {
					return i
				}
			}
		}
	}
	return 0
}

func GetStringParam(c interface{}, key string) string {
	switch v := c.(type) {
	case map[string]interface{}:
		if val, ok := v[key]; ok {
			switch v := val.(type) {
			case string:
				return v
			}
		}
	}
	return ""
}

func GetInt64Param(c interface{}, key string) int64 {
	switch v := c.(type) {
	case map[string]interface{}:
		if val, ok := v[key]; ok {
			switch v := val.(type) {
			case float64:
				return int64(v)
			case string:
				if i, err := strconv.ParseInt(v, 10, 64); err == nil {
					return i
				}
			}
		}
	}
	return 0
}

func GetBoolParam(c interface{}, key string) bool {
	switch v := c.(type) {
	case map[string]interface{}:
		if val, ok := v[key]; ok {
			switch v := val.(type) {
			case bool:
				return v
			case string:
				return v == "true" || v == "1"
			case float64:
				return v == 1
			}
		}
	}
	return false
}

func Offset(params PageParams) int {
	return (params.Page - 1) * params.PageSize
}

func TotalPages(total int64, pageSize int) int64 {
	pages := total / int64(pageSize)
	if total%int64(pageSize) > 0 {
		pages++
	}
	return pages
}
