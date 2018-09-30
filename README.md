# go-simple-api-rest

Simple API REST desarrollada en Go con PostgreSQL.
Incluye CRUD de Notas y de Usuarios, aplica JWT.


## Requiere
github.com/lib/pq
github.com/gorilla/mux


## Endpoints
login/ -> POST: name, password -> TOKEN


api/notes -> GET -> []Note

api/notes/{id} -> GET -> Note

api/notes -> POST: user_id, title, description

api/notes/{id} -> PUT: title, description

api/notes/{id} -> DELETE


api/users -> GET -> []User

api/users/{id} -> GET -> User

api/users -> POST: name, password, confirm_password

api/users/{id} -> PUT: name, password, confirm_password

api/users/{id} -> DELETE


## Migrations
```bash
# Unix
./go-notes-apirest -migrate

# Windows
go-notes-apirest -migrate
```