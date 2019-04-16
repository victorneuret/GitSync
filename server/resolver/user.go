package resolver

import (
	"context"
	"github.com/vektah/gqlparser/gqlerror"

	"github.com/victorneuret/GitSync/database"
	"github.com/victorneuret/GitSync/models"
)

func (r *MutationResolverType) CreateUser(ctx context.Context, input models.NewUser) (*models.User, error) {
	if !database.DB.Where(&database.User{Login: input.Login}).First(&database.User{}).RecordNotFound() {
		return nil, gqlerror.Errorf("User " + input.Login + " already exist")
	}

	user := database.User{
		Name: input.Name,
		Login: input.Login,
		Email: input.Email,
		AvatarURL: input.AvatarURL,
		Token: input.Token,
	}
	database.DB.Create(&user)

	gqlUser := database.User{}
	database.DB.Where(&database.User{Login: input.Login}).First(&gqlUser)
	return &models.User{
		ID: int (gqlUser.ID),
		Name: gqlUser.Name,
		Login: gqlUser.Login,
		Email: gqlUser.Email,
		AvatarURL: gqlUser.AvatarURL,
	}, nil
}

func (r *MutationResolverType) UpdateUser(ctx context.Context, login string, input models.NewUser) (*models.User, error) {
	var user database.User
	if database.DB.Where(&database.User{Login: login}).First(&user).RecordNotFound() {
		return nil, gqlerror.Errorf("User " + login + " does not exist")
	}

	if input.Name != "" {
		user.Name = input.Name
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.AvatarURL != "" {
		user.AvatarURL = input.AvatarURL
	}
	if input.Token != "" {
		user.Token = input.Token
	}
	database.DB.Save(&user)

	return &models.User{
		ID: int (user.ID),
		Name: user.Name,
		Login: user.Login,
		Email: user.Email,
		AvatarURL: user.AvatarURL,
	}, nil
}

func (r *queryResolver) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []database.User
	database.DB.Find(&users)

	var gqlUsers []models.User
	for i := 0; i < len(users); i++ {
		gqlUsers = append(gqlUsers, models.User{
			ID: int (users[i].ID),
			Name: users[i].Name,
			Login: users[i].Login,
			Email: users[i].Email,
			AvatarURL: users[i].AvatarURL,
		})
	}
	return gqlUsers, nil
}

func (r *queryResolver) GetUser(ctx context.Context, login string) (*models.User, error) {
	var user database.User
	if database.DB.Where(&database.User{Login: login}).First(&user).RecordNotFound() {
		return nil, gqlerror.Errorf("User " + login + " does not exist")
	}

	return &models.User{
		ID: int (user.ID),
		Name: user.Name,
		Login: user.Login,
		Email: user.Email,
		AvatarURL: user.AvatarURL,
	}, nil
}