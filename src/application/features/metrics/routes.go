package metrics

import "github.com/gin-gonic/gin"

func Routes(r *gin.RouterGroup) {
	m := r.Group("/metrics")
	{
		m.GET("/", Index)
	}
}
