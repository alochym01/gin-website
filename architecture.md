# Clean Architecture
1. Folder Structure
    ```bash
    .
    ├── cmd
    │   └── web
    │       └── main.go
    ├── config
    │   ├── driver.go
    │   ├── mysql.go
    │   └── sqlite.go
    ├── Docker-MySQL.md
    ├── go.mod
    ├── go.sum
    ├── handler
    │   └── album.go
    ├── Makefile
    ├── models
    │   ├── albums.go
    │   └── prepareDB.go
    ├── README.md
    ├── repository
    │   └── albumrepo.go
    └── router
       └── router.go
    ```
2. `cmd/web/main.go` is entry point to run go application and define:
    1. Initial Database connection
        ```go
        // sqlite
        config.DB = config.SqliteConn("foo.db")
        defer config.DB.Close()
        config.DB.Ping()
        models.PreparesqliteDB(config.DB)
        ```
    3. Run application
        ```go
        router := router.Router()
        router.Run()
        ```
3. `config/driver.go` is contained global pointer database connection `*sql.DB`
4. `config/sqlite.go` and `config/mysql.go` are define function `SqliteConn(f string) *sql.DB` return pointer database connection
5. `repository/albumrepo.go` defines an `AlbumRepo interface` with methods:
    ```go
    type AlbumRepo interface {
        Get() ([]models.Album, error)
        GetByID(id string) (models.Album, error)
        Create(title string, artist string, price float64) error
        Update(title string, artist string, price float64, id string) error
        Delete(id string) error
    }
    ```
6. `models/albums.go` defines:
   1. An Album table
        ```go
        type Album struct {
            ID     int     `json:"id"`
            Title  string  `json:"title"`
            Artist string  `json:"artist"`
            Price  float64 `json:"price"`
        }
        ```
   2. A struct for binding an client request body for an album
        ```go
        type RequestAlbum struct {
            Title  string  `json:"title"`
            Artist string  `json:"artist"`
            Price  float64 `json:"price"`
        }
        ```
   3. Using global pointer database connection `*sql.DB` from `config/driver.go`
   4. Implementation all method of `AlbumRepo interface`
        ```go
        func (al Album) Get() ([]models.Album, error) {}
        func (al Album) GetByID(id string) (models.Album, error) {}
        func (al Album) Create(title string, artist string, price float64) error {}
        func (al Album) Update(title string, artist string, price float64, id string) error {}
        func (al Album) Delete(id string) error {}
        ```
7. `handler/album.go` defines:
   1. An AlbumHandler struct with albumRepo is repository Album interface
        ```go
        type AlbumHandler struct {
            albumRepo repository.AlbumRepo
        }
        ```
   2. Define NewAlbumHandler function to return AlbumHandler object
        ```go
        func NewAlbumHandler() *AlbumHandler {
            return &AlbumHandler{
                albumRepo: models.Album{},
            }
        }
        ```
   3. Implementation all CRUD method on AlbumHandler and query database using `AlbumHandler.albumRepo` attribute
        ```go
        func (al AlbumHandler) Index(c *gin.Context) {}
        func (al AlbumHandler) Show(c *gin.Context) {}
        func (al AlbumHandler) Create(c *gin.Context) {}
        func (al AlbumHandler) Update(c *gin.Context) {}
        func (al AlbumHandler) Delete(c *gin.Context) {}
        ```
8. `router/router.go` defines:
   1. Create a *gin.Engine
        ```go
        router := gin.New()
        ```
   2. Create an AlbumHandler
        ```go
        albumH := handler.NewAlbumHandler()
        ```
   3. Doing route base on url
        ```go
        router.GET("/albums", albumH.Index)
        router.GET("/albums/:id", albumH.Show)
        router.POST("/albums", albumH.Create)
        router.PUT("/albums/:id", albumH.Update)
        router.DELETE("/albums/:id", albumH.Delete)
        ```
