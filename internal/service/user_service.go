package service

import (
    "context"
    "time"

    "github.com/mars-alien/user-api-go/internal/models"
    "github.com/mars-alien/user-api-go/internal/repository"
)

type UserService interface {
    CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error)
    GetUser(ctx context.Context, id int32) (*models.UserResponse, error)
    ListUsers(ctx context.Context, page, pageSize int) ([]models.UserResponse, error)
    UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (*models.UserResponse, error)
    DeleteUser(ctx context.Context, id int32) error
}

type userService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.UserResponse, error) {
    dob, err := time.Parse("2006-01-02", req.DOB)
    if err != nil {
        return nil, err
    }

    user, err := s.repo.Create(ctx, req.Name, dob)
    if err != nil {
        return nil, err
    }

    return &models.UserResponse{
        ID:   user.ID,
        Name: user.Name,
        DOB:  user.Dob.Format("2006-01-02"),
    }, nil
}

func (s *userService) GetUser(ctx context.Context, id int32) (*models.UserResponse, error) {
    user, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }

    age := models.CalculateAge(user.Dob)
    
    return &models.UserResponse{
        ID:   user.ID,
        Name: user.Name,
        DOB:  user.Dob.Format("2006-01-02"),
        Age:  &age,
    }, nil
}

func (s *userService) ListUsers(ctx context.Context, page, pageSize int) ([]models.UserResponse, error) {
    if page < 1 {
        page = 1
    }
    if pageSize < 1 || pageSize > 100 {
        pageSize = 10
    }

    offset := (page - 1) * pageSize
    
    users, err := s.repo.List(ctx, int32(pageSize), int32(offset))
    if err != nil {
        return nil, err
    }

    responses := make([]models.UserResponse, len(users))
    for i, user := range users {
        age := models.CalculateAge(user.Dob)
        responses[i] = models.UserResponse{
            ID:   user.ID,
            Name: user.Name,
            DOB:  user.Dob.Format("2006-01-02"),
            Age:  &age,
        }
    }

    return responses, nil
}

func (s *userService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (*models.UserResponse, error) {
    dob, err := time.Parse("2006-01-02", req.DOB)
    if err != nil {
        return nil, err
    }

    user, err := s.repo.Update(ctx, id, req.Name, dob)
    if err != nil {
        return nil, err
    }

    return &models.UserResponse{
        ID:   user.ID,
        Name: user.Name,
        DOB:  user.Dob.Format("2006-01-02"),
    }, nil
}

func (s *userService) DeleteUser(ctx context.Context, id int32) error {
    return s.repo.Delete(ctx, id)
}