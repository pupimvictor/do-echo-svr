package restapi

import (
	"encoding/json"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/pupimvictor/do-echo-svr/models"
	"github.com/pupimvictor/do-echo-svr/restapi/operations"
	"github.com/pupimvictor/do-echo-svr/restapi/operations/echo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoEchoHandler(t *testing.T) {
	tests := []struct {
		name         string
		givenMsg     string
		expectedResp *models.Echo
	}{
		{
			name:     "test1",
			givenMsg: "msg1",
			expectedResp: &models.Echo{
				Echo: "msg1",
			},
		}, {
			name:     "test2",
			givenMsg: "msg1 msg1",
			expectedResp: &models.Echo{
				Echo: "msg1 msg1",
			},
		}, {
			name:     "testEmptyMsg",
			givenMsg: "",
			expectedResp: &models.Echo{
				Echo: "",
			},
		},
	}

	spec, err := loads.Spec("../swagger.yml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	api := operations.NewEchoerAPI(spec)
	server := NewServer(api)
	server.ConfigureAPI()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "", nil)
			if err != nil {
				t.Fatal(err)
			}

			msg := test.givenMsg
			params := echo.EchoParams{
				HTTPRequest: req,
				Body: &models.Message{
					Msg: &msg,
				},
			}

			r := api.EchoEchoHandler.Handle(params)
			w := httptest.NewRecorder()
			r.WriteResponse(w, runtime.JSONProducer())
			if w.Code != 200 {
				t.Errorf("status code %d", w.Code)
			}
			var echoResp models.Echo
			json.Unmarshal(w.Body.Bytes(), &echoResp)
			if echoResp.Echo != test.expectedResp.Echo {
				t.Errorf("expected %+v got %+v", test.expectedResp, echoResp)
			}
		})
	}
}
