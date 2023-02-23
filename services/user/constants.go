package user

import "ps-tag-onboarding-go/services"

const (
	RouteUser = "/user"

	USER_AGE_MINIMUM_ERROR       services.ServiceErrorT = 1
	USER_EMAIL_FORMAT_ERROR      services.ServiceErrorT = 2
	USER_EMAIL_REQUIRED_ERROR    services.ServiceErrorT = 3
	USER_NAME_REQUIRED_ERROR     services.ServiceErrorT = 4
	USER_NAME_UNIQUE_ERROR       services.ServiceErrorT = 5
	USER_NOT_FOUND_ERROR         services.ServiceErrorT = 6
	USER_VALIDATION_FAILED_ERROR services.ServiceErrorT = 7
	USER_DATABASE_ERROR          services.ServiceErrorT = 8
)
