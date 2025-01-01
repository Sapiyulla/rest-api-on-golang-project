package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Text struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

type TextModel struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

var Texts = []Text{
	Text{Id: 0, Title: "Martini", Year: 2003},
	Text{Id: 1, Title: "Kolan", Year: 2006},
	Text{Id: 2, Title: "Jackob", Year: 1988},
}

func getInfo(c *gin.Context) {
	c.IndentedJSON(200, Texts)
}

func postText(c *gin.Context) {
	var textM TextModel
	var text Text

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&textM); err != nil {
		return
	}
	text.Id = len(Texts)
	text.Title = textM.Title
	text.Year = textM.Year
	// Add the new album to the slice.
	Texts = append(Texts, text)
	c.IndentedJSON(201, Texts)
}

func getInfoForId(c *gin.Context) {
	id := c.Param("id")
	for _, item := range Texts {
		if strconv.Itoa(item.Id) == id {
			c.JSON(200, item)
			return
		}
	}
	c.JSON(404, gin.H{
		"status": "error",
	})
}

func deleteInfo(c *gin.Context) {
	Texts = make([]Text, len(Texts), cap(Texts))
}

func deleteInfoForId(c *gin.Context) {
	p := c.Param("id")
	if len(Texts) == 0 {
		c.JSON(400, gin.H{
			"status": "error",
		})
		return
	}
	for index, item := range Texts {
		if strconv.Itoa(item.Id) == p {
			_own := Texts[index+1:]
			for i, item := range _own {
				_own[i].Id = item.Id - 1
			}
			Texts = Texts[:index]
			Texts = append(Texts, _own...)
			c.JSON(200, gin.H{
				"status": "deleted",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"status": "error",
	})
}

func putInfoForId(c *gin.Context) {
	p := c.Param("id")
	var text Text
	for index, item := range Texts {
		if strconv.Itoa(item.Id) == p {
			id := item.Id
			err := c.BindJSON(&text)
			if err != nil {
				c.JSON(400, gin.H{
					"status": "error",
				})
				return
			}
			Texts[index].Title = text.Title
			Texts[index].Year = text.Year
			Texts[index].Id = id
			c.JSON(202, gin.H{
				"status": "accepted",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"status": "error",
	})
}

func main() {
	r := gin.Default()

	r.GET("/text", getInfo)
	r.POST("/text", postText)
	r.GET("/text/:id", getInfoForId)
	r.DELETE("/text", deleteInfo)
	r.DELETE("/text/:id", deleteInfoForId)
	r.PUT("/text/:id", putInfoForId)

	r.Run(":9000")
}
