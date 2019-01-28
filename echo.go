package echoer

import (
	"github.com/pupimvictor/do-echo-svr/models"
)

//Echo receives a Message and return its content in a Echo object
//if the message is nil, it will return a empty echo
func Echo (in *models.Message)  (*models.Echo){
	if in.Msg != nil {
		return &models.Echo{Echo: *in.Msg}
	}
	return &models.Echo{Echo: ""}
}