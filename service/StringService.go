package service

// StringService interface
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// ServiceMiddleware is a chainable behavior modifier for StringService.
type ServiceMiddleware func(StringService) StringService
