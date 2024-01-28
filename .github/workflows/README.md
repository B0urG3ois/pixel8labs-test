# GitHub Action Workflow

### Code Analysis

Code analysis is performed by the linter in GitHub Actions workflows, and it's set up to operate on both `PUSH` events and `PULL REQUESTS` targeting the `main` or `master` branch. This configuration ensures that crucial branches consistently adhere to established standards.

### Pipeline

It's intended as a GitHub Actions workflow for deploying the backend to a hosting server, specifically Heroku.
So if there is a merge from `PULL REQUEST`, this Github Action will run and deploy it into Heroku.