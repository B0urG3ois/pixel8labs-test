package github

func New(githubUC GithubUsecase) *Handler {
	return &Handler{
		githubUC: githubUC,
	}
}
