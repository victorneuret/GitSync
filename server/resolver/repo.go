package resolver

import (
	"context"
	"github.com/vektah/gqlparser/gqlerror"
	"github.com/victorneuret/GitSync/database"
	"github.com/victorneuret/GitSync/models"
	"github.com/victorneuret/GitSync/app"
)


// QUERY

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
			UpdateOnMaster: repos[i].UpdateOnMaster,
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
		UpdateOnMaster: repo.UpdateOnMaster,
	}, nil
}

func (r *queryResolver) GetRepoFromOwner(ctx context.Context, owner string) ([]models.Repo, error) {
	var repos []database.Repo
	if database.DB.Find(&repos).Where(&database.Repo{Owner: owner}).RecordNotFound() {
		return nil, gqlerror.Errorf("No repositorys from " + owner + " found")
	}

	var gqlRepo []models.Repo
	for i := 0; i < len(repos); i++ {
		gqlRepo = append(gqlRepo, models.Repo{
			Name: repos[i].Name,
			Private: repos[i].Private,
			GithubURL: repos[i].GithubURL,
			Owner: repos[i].Owner,
			Updater: repos[i].Updater,
			UpdateOnMaster: repos[i].UpdateOnMaster,
		})
	}
	return gqlRepo, nil
}



// MUTATION

func (r *MutationResolverType) CreateRepo(ctx context.Context, input models.NewRepo, token string) (*models.Repo, error) {
	if !database.DB.Where(&database.Repo{Name: input.Name, Owner: input.Owner}).First(&database.Repo{}).RecordNotFound() {
		return nil, gqlerror.Errorf("Repository " + input.Name + " already exist")
	}

	var user database.User
	if database.DB.Where(&database.User{Login: input.Owner}).First(&user).RecordNotFound() {
		return nil, gqlerror.Errorf("User " + input.Owner + " does not exist")
	}
	if user.BlihUsername == "" || user.BlihToken == "" {
		return nil, gqlerror.Errorf("User " + input.Owner + " have missing blih informations")
	}


	if !app.CreateGitHubRepo(input.Name, input.Private, token) {
		return nil, gqlerror.Errorf("Can't create github repo " + input.Name)
	}
	if !app.CreateBlihRepo(input.Name, user.BlihUsername, user.BlihToken) {
		return nil, gqlerror.Errorf("Can't create blih repo " + input.Name)
	}
	if !app.SetMirror(input.Name, user.BlihUsername, user.Login) {
		return nil, gqlerror.Errorf("Mirror of " + input.Name + " failed")
	}
	if !app.CreateGitHubHook(input.Name, user.Login, user.Token) {
		return nil, gqlerror.Errorf("Github webhook creation of " + input.Name + " failed")
	}



	repo := database.Repo{
		Name: input.Name,
		Private: input.Private,
		GithubURL: "https://github.com/" + input.Owner + "/" + input.Name,
		Owner: input.Owner,
		Updater: input.Owner,
		UpdateOnMaster: input.UpdateOnMaster,
	}
	database.DB.Create(&repo)

	gqlRepo := database.Repo{}
	database.DB.Where(&database.Repo{Name: input.Name}).First(&gqlRepo)
	return &models.Repo{
		Name: gqlRepo.Name,
		Private: gqlRepo.Private,
		GithubURL: gqlRepo.GithubURL,
		Owner: gqlRepo.Owner,
		Updater: gqlRepo.Updater,
		UpdateOnMaster: gqlRepo.UpdateOnMaster,
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
	repo.UpdateOnMaster = input.UpdateOnMaster
	database.DB.Save(&repo)

	return &models.Repo{
		Name: repo.Name,
		Private: repo.Private,
		GithubURL: repo.GithubURL,
		Owner: repo.Owner,
		Updater: repo.Updater,
		UpdateOnMaster: repo.UpdateOnMaster,
	}, nil
}