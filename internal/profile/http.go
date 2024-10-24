package profile

import (
	"codeathon.runwayclub.dev/domain"
	"codeathon.runwayclub.dev/internal/endpoint"
	"github.com/gin-gonic/gin"
)

func Api(service ProfileService) {
	r := endpoint.GetEngine()
	r.GET("/profiles", func(c *gin.Context) {
		list, err := service.List(c, &domain.ListOpts{Offset: 0, Limit: 10})
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, list)
	})

	r.GET("/profile/:id", func(c *gin.Context) {
		id := c.Param("id")
		profile, err := service.GetById(c, id)
		if err != nil {
			c.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, profile)
	})

	r.PUT("/profile", func(c *gin.Context) {
		profile := &domain.Profile{}
		err := c.BindJSON(profile)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		err = service.Update(c, profile)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, profile)
	})

	r.DELETE("/profile/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := service.Delete(c, id)
		if err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "deleted",
		})
	})
}
