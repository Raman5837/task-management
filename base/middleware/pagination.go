package middleware

import (
	"math"
	"net/url"
	"strconv"

	"github.com/Raman5837/task-management/base/configuration"
	"github.com/gofiber/fiber/v2"
)

var Logger = configuration.GetLogger()

const (
	DEFAULT_MIN_PAGESIZE = 25
	DEFAULT_MAX_PAGESIZE = 100
)

type OffsetPaginationRequestOptions struct {
	Limit  int
	Offset int
}

type OffsetPaginationResponseOptions struct {
	PageSize   int `json:"page_size"`
	TotalCount int `json:"total_count"`

	PrevUrl string `json:"prev_url"`
	NextUrl string `json:"next_url"`
}

// Pagination middleware based on limit-offset approach
func OffsetBasedPaginationMiddleware() fiber.Handler {

	return func(context *fiber.Ctx) error {

		invalidLimitNumberMessage := "Invalid limit number"
		invalidOffsetNumberMessage := "Invalid offset number"

		requestedLimit := context.Query("limit", "")
		requestedOffset := context.Query("offset", "")
		nonPaginatedRequest := requestedLimit == "" && requestedOffset == ""

		Logger.Debug(
			"Current request has requested limit of %v and offset of %v", requestedLimit, requestedOffset,
		)

		if nonPaginatedRequest {
			// Setting Pagination Options In The Context Locals
			context.Locals("paginationOptions", OffsetPaginationRequestOptions{Limit: 25, Offset: 0})
			Logger.Debug("This was non-paginated API request, so added default pagination options")

		} else {

			limit, parsingError := strconv.Atoi(requestedLimit)
			if parsingError != nil || limit < DEFAULT_MIN_PAGESIZE || limit > DEFAULT_MAX_PAGESIZE {
				return context.Status(fiber.StatusBadRequest).JSON(invalidLimitNumberMessage)
			}

			offset, parsingError := strconv.Atoi(requestedOffset)
			if parsingError != nil || limit < DEFAULT_MIN_PAGESIZE || limit > DEFAULT_MAX_PAGESIZE {
				return context.Status(fiber.StatusBadRequest).JSON(invalidOffsetNumberMessage)
			}

			paginationOptions := OffsetPaginationRequestOptions{Limit: limit, Offset: offset}

			// Setting pagination options in the context locals
			context.Locals("paginationOptions", paginationOptions)
			Logger.Debug("Context for limit-offset based pagination options set in the API")
		}

		// Execute next middleware or handler function
		return context.Next()

	}

}

// Middleware to send API response in limit-offset based paginated fashion
func LimitOffsetBasedPaginatedResponse(context *fiber.Ctx, totalCount int, data any) error {

	paginatedRequest := context.Locals("paginationOptions").(OffsetPaginationRequestOptions)

	currentLimit := paginatedRequest.Limit
	currentOffset := paginatedRequest.Offset
	Logger.Debug("Current limit and offset are:- %v, %v", currentLimit, currentOffset)

	nextOffset := currentOffset + currentLimit
	prevOffset := currentOffset - currentLimit

	if prevOffset < 0 {
		prevOffset = 0
		Logger.Debug("Prev and Next offset after modification are:- %v, %v", prevOffset, nextOffset)

	} else {
		Logger.Debug("Prev and Next offset are:- %v, %v", prevOffset, nextOffset)
	}

	// Build next and prev URLs
	nextURL := buildPaginationURL(context.OriginalURL(), currentLimit, nextOffset)
	prevURL := buildPaginationURL(context.OriginalURL(), currentLimit, prevOffset)

	pagination := OffsetPaginationResponseOptions{
		PageSize:   currentLimit,
		TotalCount: totalCount,
		NextUrl:    nextURL,
		PrevUrl:    prevURL,
	}

	_ = pagination.PageSize
	currentPage := float64(currentOffset) / float64(currentLimit)
	totalPages := float64(pagination.TotalCount) / float64(currentLimit)

	return context.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"error":        nil,
			"data":         data,
			"previous":     pagination.PrevUrl,
			"next":         pagination.NextUrl,
			"count":        pagination.TotalCount,
			"total_pages":  math.Ceil(totalPages),
			"current_page": math.Ceil(currentPage) + 1,
		},
	)

}

// Helper function to build next and prev URL
func buildPaginationURL(originalUrl string, limit int, offset int) string {

	url, err := url.Parse(originalUrl)
	if err != nil {
		Logger.Error(err, "Something went wrong while parsing pagination URL")
		return originalUrl
	}

	queryParam := url.Query()
	queryParam.Set("limit", strconv.Itoa(limit))
	queryParam.Set("offset", strconv.Itoa(offset))

	url.RawQuery = queryParam.Encode()

	return url.String()
}
