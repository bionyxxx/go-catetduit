package user

import "golang.org/x/crypto/bcrypt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUserByID(id uint) (*UserResponse, error) {
	userData, err := s.repo.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	userResp := &UserResponse{
		ID:        &userData.ID,
		Name:      userData.Name,
		Phone:     userData.Phone,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt.Unix(),
		UpdatedAt: userData.UpdatedAt.Unix(),
	}

	return userResp, nil
}

// ChangePassword
func (s *Service) ChangePassword(userID uint, newPassword string) error {
	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Update password in the repository
	err = s.repo.ChangePassword(userID, string(hashedPassword))

	return err
}

// check old password
func (s *Service) CheckOldPassword(userID uint, oldPassword string) (bool, error) {
	userData, err := s.repo.GetUserByID(userID)

	if err != nil {
		return false, err
	}

	// Compare old password
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(oldPassword))
	if err != nil {
		return false, nil
	}

	return true, nil
}
