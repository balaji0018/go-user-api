package service

import (
	"context"
	db "go-user-api/db/sqlc"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

// userrespoanse is what we send back to fronted
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age"`
}

type UserService struct {
	queries *db.Queries
}

// //Newuserservice connects this service to the database
func NewUserService(conn *pgxpool.Pool) *UserService {
	return &UserService{
		queries: db.New(conn),
	}
}

// Calculateage figures out age from date
func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	//If birthday hasnt happened yet this year subrtract 1
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// CreateUser handles the logic to add a new user
func (s *UserService) CreateUser(ctx context.Context, name string, dobString string) (UserResponse, error) {
	parsedDob, err := time.Parse("2006-01-02", dobString)
	if err != nil {
		return UserResponse{}, err
	}

	//saving to database using code SQLC genertaed
	user, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  pgtype.Date{Time: parsedDob, Valid: true},
	})
	if err != nil {
		return UserResponse{}, err
	}

	//Return the result
	return UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Time.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob.Time),
	}, nil
}

// Getuser fetches a user and calculate thier age
func (s *UserService) GetUser(ctx context.Context, id int32) (UserResponse, error) {
	user, err := s.queries.GetUser(ctx, id)
	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:   user.ID,
		Name: user.Name,
		Dob:  user.Dob.Time.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob.Time), //Dynamic Calculation
	}, nil
}
func (s *UserService) GetAllUsers(ctx context.Context) ([]UserResponse, error) {
	// Use SQLC generated ListUsers
	users, err := s.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []UserResponse
	for _, u := range users {
		result = append(result, UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Time.Format("2006-01-02"),
			Age:  CalculateAge(u.Dob.Time), // dynamic age calculation
		})
	}

	return result, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.queries.DeleteUser(ctx, id)
}
