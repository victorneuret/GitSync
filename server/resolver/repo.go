package resolver

import (
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/victorneuret/GitSync/database"
	"github.com/victorneuret/GitSync/models"
	"context"
)

func (r *queryResolver) GetAllRepos(ctx context.Context) ([]models.Repo, error) {
	var repos []database.Repo
	database.DB.Find(&repos)

	var gqlRepo []models.Repo
	for i := 0; i < len(repos); i++ {
		gqlRepo = append(gqlRepo, models.Repo{
			Name: repos[i].Name,
			Private: repos[i].Private,
			GithubURL: repos[i].GithubURL,
			Owner: repos[i].Owner,
			Updater: repos[i].Updater,
		})
	}
	return gqlRepo, nil
}

func (r *queryResolver) GetRepo(ctx context.Context, name string) (*models.Repo, error) {
	var repo database.Repo
	if database.DB.Where(&database.Repo{Name: name}).First(&repo).RecordNotFound() {
		return nil, gqlerror.Errorf("Repository " + name + " does not exist")
	}

	return &models.Repo{
		Name: repo.Name,
		Private: repo.Private,
		GithubURL: repo.GithubURL,
		Owner: repo.Owner,
		Updater: repo.Updater,
	}, nil
}

func (r *MutationResolverType) CreateRepo(ctx context.Context, input models.NewRepo) (*models.Repo, error) {
	if !database.DB.Where(&database.Repo{Name: input.Name}).First(&database.Repo{}).RecordNotFound() {
		return nil, gqlerror.Errorf("Repository " + input.Name + " already exist")
	}

	user := database.Repo{
		Name: input.Name,
		Private: input.Private,
		Owner: input.Owner,
	}
	database.DB.Create(&user)

	gqlRepo := database.Repo{}
	database.DB.Where(&database.Repo{Name: input.Name}).First(&gqlRepo)
	return &models.Repo{
		Name: gqlRepo.Name,
		Private: gqlRepo.Private,
		Owner: gqlRepo.Owner,
	}, nil
}

func (r *MutationResolverType) UpdateRepo(ctx context.Context, name string, input models.ModifRepo) (*models.Repo, error) {
	var repo database.Repo
	if database.DB.Where(&database.Repo{Name: name}).First(&repo).RecordNotFound() {
		return nil, gqlerror.Errorf("Repository " + name + " does not exist")
	}

	repo.Private = input.Private
	if input.Updater != "" {
		repo.Updater = input.Updater
	}
	database.DB.Save(&repo)

	return &models.Repo{
		Name: repo.Name,
		Private: repo.Private,
		GithubURL: repo.GithubURL,
		Owner: repo.Owner,
		Updater: repo.Updater,
	}, nil
}