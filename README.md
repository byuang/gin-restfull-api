# Description
A simple API Tags with CRUD and JWT auth

# Features
This API developed with Go, Gin-Gonic, Gorm, JWT, Postgresql, Docker
 
# Tech Used
 ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens) ![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white) ![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
      
# Getting Start:
Before you running the program, make sure you've run this command:
- `go get -u all`
- `.env with config db`
- `docker compose up -d`

### Run the program
`go run main.go`

The program will run on http://localhost:8888


### API Route List
| Method | URL                                        | Description           | Authorization           |
| ------ | ----------------------------------------   | --------------------- | ------------------------|
| POST   | localhost:8888/api/authentication/register | Register              |                         |
| POST   | localhost:8888/api/authentication/login    | Login                 |                         |
| GET    | localhost:8888/api/tags                    | Get All Tags          | Add Authorization token |
| POST   | localhost:8888/api/tags                    | Create Tags           | Add Authorization token |
| POST   | localhost:8888/api/tags/{id}               | Update Tags           | Add Authorization token |
| GET    | localhost:8888/api/tags/{id}               | Get Tags Details      | Add Authorization token |
| DELETE | localhost:8888/api/tags/{id}               | Delete Tags           | Add Authorization token |

 
<!-- </> with 💛 by readMD (https://readmd.itsvg.in) -->
