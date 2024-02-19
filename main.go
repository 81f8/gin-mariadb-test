package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Country struct {
	CountryID     int     `db:"country_id" json:"countryID"`
	CountryName   string  `db:"name" json:"countryName"`
	Area          float64 `db:"area" json:"area"`
	Language      string  `db:"lang" json:"language"`
	ContinentName string  `db:"continentName" json:"continentName"`
}

func main() {
	r := gin.Default()

	db, err := sqlx.Open("mysql", "root:salam@tcp(localhost:3306)/nation")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r.GET("/doc", func(c *gin.Context) {
		htmlBytes, err := os.ReadFile("./doc.html")
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "Failed to read HTML file")
			return
		}

		htmlString := string(htmlBytes)
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlString))
	})

	r.GET("/countries", func(c *gin.Context) {
		var countries []Country
		err := db.Select(&countries, `
			SELECT c.country_id, c.name, c.area, l.language AS lang, ct.name AS continentName
			FROM countries c
			JOIN country_languages cl ON c.country_id = cl.country_id
			JOIN languages l ON cl.language_id = l.language_id
			JOIN regions r ON c.region_id = r.region_id
			JOIN continents ct ON r.continent_id = ct.continent_id
		`)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch countries"})
			return
		}

		c.JSON(http.StatusOK, countries)
	})

	r.GET("/countries/:id", func(c *gin.Context) {
		var country Country
		id := c.Param("id")
		err := db.Get(&country, `
			SELECT c.country_id, c.name, c.area, l.language AS lang, ct.name AS continentName
			FROM countries c
			JOIN country_languages cl ON c.country_id = cl.country_id
			JOIN languages l ON cl.language_id = l.language_id
			JOIN regions r ON c.region_id = r.region_id
			JOIN continents ct ON r.continent_id = ct.continent_id
			WHERE c.country_id = ?
		`, id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
				return
			}
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch country"})
			return
		}

		c.JSON(http.StatusOK, country)
	})

	r.POST("/countries", func(c *gin.Context) {
		var country Country
		if err := c.ShouldBindJSON(&country); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		result, err := db.Exec(`
			INSERT INTO countries (name, area)
			VALUES (?, ?)
		`, country.CountryName, country.Area)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create country"})
			return
		}

		countryID, _ := result.LastInsertId()
		country.CountryID = int(countryID)

		c.JSON(http.StatusCreated, country)
	})

	r.PUT("/countries/:id", func(c *gin.Context) {
		id := c.Param("id")
		var country Country
		if err := c.ShouldBindJSON(&country); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		_, err := db.Exec(`
			UPDATE countries SET name = ?, area = ?
			WHERE country_id = ?
		`, country.CountryName, country.Area, id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update country"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Country updated successfully"})
	})

	r.DELETE("/countries/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec(`
			DELETE FROM countries WHERE country_id = ?
		`, id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete country"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Country deleted successfully"})
	})

	r.Run(":8080")
}
