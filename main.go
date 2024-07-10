package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"paulobraga.com/study/lib"
)

type Album struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Id     int     `json:"id"`
}

type App struct {
	DB     *sql.DB
	Routes *gin.Engine
	conn   string
}

func (a *App) getAlbums(c *gin.Context) {
	db := a.DB

	if db == nil {
		lib.LogService("Error: Error fetching albums database - database not initialized", c)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error fetching albums"})
		return
	}

	sqlStatement := `SELECT title, artist, price, id FROM albums`

	rows, err := db.Query(sqlStatement)
	if err != nil {

		lib.LogService("Error: Error fetching albums - "+err.Error(), c)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Error fetching albums"})
		return
	}

	var albums []Album

	for rows.Next() {
		var album Album
		err = rows.Scan(&album.Title, &album.Artist, &album.Price, &album.Id)
		if err != nil {
			lib.LogService("Error: Error fetching albums - "+err.Error(), c)
			return
		}
		albums = append(albums, album)
	}
	lib.LogService("Info: Albums fetched successfully", c)

	if len(albums) == 0 {
		lib.LogService("Error: No albums found", c)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "No albums found"})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func (a *App) getAlbumById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	db := a.DB

	var row *sql.Row

	sqlStatement := `SELECT * FROM albums WHERE id = $1`
	row = db.QueryRow(sqlStatement, id)

	var album Album
	err := row.Scan(&album.Title, &album.Artist, &album.Price, &album.Id)
	if err != nil {
		lib.LogService("Error: Error fetching album by id - "+err.Error(), c)
		return
	}
	lib.LogService("Info: Album fetched successfully", c)
	c.IndentedJSON(http.StatusOK, album)
}

func (a *App) createAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		lib.LogService("Error: Error creating album - "+err.Error(), c)
		return
	}

	db := a.DB

	sqlStatement := `INSERT INTO albums (title, artist, price) VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(sqlStatement, newAlbum.Title, newAlbum.Artist, newAlbum.Price).Scan(&newAlbum.Id)
	if err != nil {
		lib.LogService("Error: Error creating album - "+err.Error(), c)
		return
	}
	lib.LogService("Info: Album created successfully", c)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func (a *App) deleteAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		lib.LogService("Error: Error deleting album - "+err.Error(), c)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	db := a.DB

	sqlStatement := `DELETE FROM albums WHERE id = $1`

	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		lib.LogService("Error: Error deleting album - "+err.Error(), c)
		return
	}
	lib.LogService("Info: Album deleted successfully", c)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
}

func (a *App) updateAlbum(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		lib.LogService("Error: Error updating album - "+err.Error(), c)
		return
	}

	db := a.DB

	var updatedAlbum Album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		lib.LogService("Error: Error updating album - "+err.Error(), c)
		return
	}

	sqlStatement := `UPDATE albums SET title = $1, artist = $2, price = $3 WHERE id = $4`

	_, err = db.Exec(sqlStatement, updatedAlbum.Title, updatedAlbum.Artist, updatedAlbum.Price, id)
	if err != nil {
		lib.LogService("Error: Error updating album - "+err.Error(), c)
		return
	}
	lib.LogService("Info: Album updated successfully", c)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "album updated"})
}

type X bool

func (c X) Int(a, b *int) int {
	if c {
		return *a
	}
	return *b
}

func main() {

	app := App{}
	// Create connection pool
	app.CreateConnection("albums")
	// Migrate the schema
	app.Migrate()

	// Start the server in release mode
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/albums", app.getAlbums)
	router.GET("/albums/:id", app.getAlbumById)
	router.POST("/albums", app.createAlbum)
	router.PUT("/albums/:id", app.updateAlbum)
	router.DELETE("/albums/:id", app.deleteAlbum)
	router.Run(":8080")

}
