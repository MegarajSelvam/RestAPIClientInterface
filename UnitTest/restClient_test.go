package UnitTest

import (
	"RestAPIClientInterface/Common"
	"RestAPIClientInterface/MockRestAPIClient"
	"RestAPIClientInterface/RestAPIClient"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebCallV1(t *testing.T) {
	MockClient := new(MockRestAPIClient.MockHttpClient)
	MockClient.On("Do", Common.GetRequest()).Return(Common.GetResponse())
	RestAPIClient.Client = MockClient
	resp, err := RestAPIClient.WebCallV1(Common.GetRequest())
	assert.Equal(t, resp.StatusCode, 200)
	assert.Nil(t, err)
	require.JSONEq(t, Common.ResponseString, string(Common.ReadContent(resp.Body)))
}
