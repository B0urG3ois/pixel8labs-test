package github

func New(githubRepo GithubResource) *Usecase {
	return &Usecase{
		GithubRepo: githubRepo,
	}
}
