package resolver

import (
	"context"
	"github.com/vektah/gqlparser/gqlerror"

	"github.com/victorneuret/GitSync/models"
	"github.com/victorneuret/GitSync/database"
)

func (r *MutationResolverType) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	if !database.DB.Where(&database.User{Name: input.Name}).First(&database.User{}).RecordNotFound() {
		return nil, gqlerror.Errorf("User " + input.Name + " already exist")
	}

	user := database.User{
		Name: input.Name,
	}
	database.DB.Create(&user)

	gqlUser := database.User{}
	database.DB.Where(&database.User{Name: input.Name}).First(&gqlUser)
	return &models.User{ID: int (gqlUser.ID), Name: gqlUser.Name}, nil
}

func (r *MutationResolverType) UpdateUser(ctx context.Context, name string, input models.NewUser) (*models.User, error) {
	var user database.User
	if database.DB.Where(&database.User{Name: name}).First(&user).RecordNotFound() {
		return nil, gqlerror.Errorf("User " + name + " does not exist")
	}

	user.Name = input.Name
	database.DB.Save(&user)

	return &models.User{ ID: int (user.ID), Name: user.Name}, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	var users []database.User
	database.DB.Find(&users)

	var gqlUsers []models.User
	for i := 0; i < len(users); i++ {
		gqlUsers = append(gqlUsers, models.User{
			ID: int (users[i].ID),
			Name: users[i].Name,
		})
	}
	return gqlUsers, nil
}

func (r *queryResolver) GetUser(ctx context.Context, name string) (*models.User, error) {
	var user database.User
	if database.DB.Where(&database.User{Name: name}).First(&user).RecordNotFound() {
		return nil, gqlerror.Errorf("User " + name + " does not exist")
	}

	return &models.User{ ID: int (user.ID), Name: user.Name}, nil
}