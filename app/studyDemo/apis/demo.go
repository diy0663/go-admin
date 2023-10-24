package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
)

type Demo struct {
	api.Api
}

func (d *Demo) GetDemo(c *gin.Context) {
	d.MakeContext(c)
	d.OK(nil, "this is demo")
}
