package echoer

import (
	"github.com/pupimvictor/do-echo-svr/models"
)

func Echo (in *models.Message)  (*models.Echo){
	if in.Msg != nil {
		return &models.Echo{Echo: *in.Msg}
	}
	return &models.Echo{Echo: ""}
}