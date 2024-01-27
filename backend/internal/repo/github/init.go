package github

func New(githubSvc githubResource, db dbTX) *Repo {
	return &Repo{
		githubSvc: githubSvc,
		db:        db,
	}
}
