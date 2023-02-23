package user

import "ps-tag-onboarding-go/services"

func (s *UserService) processUserGet(userID string) (*User, error) {
	user, err := s.UserGet(userID)
	if err != nil {
		return nil, services.ServiceErrorReply{
			ErrorCode:    USER_NOT_FOUND_ERROR,
			ErrorContext: err.Error(),
		}
	}

	return user, nil
}

func (s *UserService) processUserSet(user *User) error {
	err := s.UserSet(user)
	if err != nil {
		return services.ServiceErrorReply{
			ErrorCode:    USER_DATABASE_ERROR,
			ErrorContext: err.Error(),
		}
	}

	return nil
}
