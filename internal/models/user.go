package models

import "time"

type CreateUserRequest struct {
    Name string `json:"name" validate:"required,min=1,max=100"`
    DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UpdateUserRequest struct {
    Name string `json:"name" validate:"required,min=1,max=100"`
    DOB  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UserResponse struct {
    ID   int32     `json:"id"`
    Name string    `json:"name"`
    DOB  string    `json:"dob"`
    Age  *int      `json:"age,omitempty"`
}

func CalculateAge(dob time.Time) int {
    now := time.Now()
    age := now.Year() - dob.Year()
    
    if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
        age--
    }
    
    return age
}