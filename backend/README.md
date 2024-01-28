# Backend

This backend is a critical component that handles authentication through GitHub OAuth, and integrating with GitHub API.

## Tech Stack

### Golang 1.21
Go, or Golang, is renowned for its efficient concurrent programming with lightweight goroutines and channels, fostering the development of scalable applications. The language's simplicity, readable syntax, and static typing contribute to code reliability. Go's strong standard library, cross-platform support, and built-in testing framework enhance productivity, while its performance benefits from quick compilation to machine code. The community support and tooling, exemplified by commands like `go fmt` and `go vet`, further underscore Go's appeal for building robust, scalable, and concurrent applications.

### Router Chi V5 for Routing
Chi v5 is a handy router for Go that makes web development easier. It keeps your code clean with an easy-to-understand style and supports extra features through middleware. You can structure your web applications well with its features like parameterized routing and route groups. Chi is lightweight, so you only use the parts you need, making your code easy to manage. It quickly directs incoming web requests to the right places, ensuring your application runs smoothly. Plus, there's an active community for support and resources. For the latest details, check out the official Chi documentation.

### GitHub OAuth and API Integration
To enhance our interaction with the GitHub API, we've developed wrappers using Go. These wrappers simplify the process of making requests to GitHub, offering a neat and straightforward interface for accessing a variety of information, including user data and repositories. This abstraction allows for efficient communication with the GitHub platform while maintaining a clean and organized code structure.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/B0urG3ois/pixel8labs-test.git

2. Install Golang 1.21:
   ```bash
   https://go.dev/doc/install

3. Run application:
   ```bash
   cd /backend
   go run cmd/app/main.go

4. Application Start:
   ```bash
   Application start on port 8000
   http://localhost:8000/api/*

## Documentation

- Postman Collection
  ```bash
  https://api.postman.com/collections/12053329-8d8ac234-9f8d-4305-a8b7-3566f438328d?access_key=PMAT-01HN79YPPTZBJ7T8FS6BCBS3WY

- Endpoints Collection
    - `/api/login`: Initiate the OAuth flow to link with a GitHub account.
    - `/api/logout`: Log the user out.
    - `/api/callback`: Manage code exchange against the front-callback layer's code.
    - `/api/v1/user`: Retrieve the user information details.
    - `/api/v1/repositories`: Retrieve the user's six most recent repositories.

- Backend Public URL
  ```bash
  https://pixel8labs-test-be31bc5e9b1e.herokuapp.com

## Database Uses

Using a `Database` for counting total visitors has its advantages, providing a persistent and reliable storage solution for such data. However, it's essential to note that relying solely on caching can introduce faster retrieval times but may lack persistence. 
To strike a balance between `efficiency` and `durability`, a `best practice` at least for me to involves a combined approach of using both a database and caching. This strategy optimizes performance by leveraging the speed of caching while ensuring data integrity and durability through the use of a database. This hybrid approach represents a best practice in achieving an optimal balance between performance and robust data management.

----------------------------------------------------------------
Please feel free to examine the code and make any required modifications or improvements to customize it for your particular needs.


