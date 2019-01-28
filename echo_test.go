package echoer

import (
	"github.com/pupimvictor/do-echo-svr/models"
	"testing"
)

func TestEcho(t *testing.T){
	tests := []struct{
		name string
		givenMsg string
		expectedResp *models.Echo
	}{
		{
			name: "test1",
			givenMsg: "msg1",
			expectedResp: &models.Echo{
				Echo: "msg1",
			},
		},{
			name: "test2",
			givenMsg: "msg1 msg1",
			expectedResp: &models.Echo{
				Echo: "msg1 msg1",
			},
		},{
			name: "testEmptyMsg",
			givenMsg: "",
			expectedResp: &models.Echo{
				Echo: "",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(ts *testing.T){
			msg := &models.Message{Msg: &test.givenMsg}
			echoResp := Echo(msg)
			if echoResp.Echo != test.expectedResp.Echo {
				t.Errorf("expected %+v got %+v", test.expectedResp, echoResp)
			}
		})
	}
}

