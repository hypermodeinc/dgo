/*
 * SPDX-FileCopyrightText: © Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package dgo_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// LoginParams stores the information needed to perform a login request.
type LoginParams struct {
	Endpoint   string
	UserID     string
	Passwd     string
	Namespace  uint64
	RefreshJwt string
}

type GraphQLParams struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type HttpToken struct {
	UserId       string
	Password     string
	AccessJwt    string
	RefreshToken string
}

type GraphQLResponse struct {
	Data       json.RawMessage        `json:"data,omitempty"`
	Errors     GqlErrorList           `json:"errors,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

type GqlErrorList []*GqlError

type GqlError struct {
	Message    string                 `json:"message"`
	Path       []interface{}          `json:"path,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

func MakeGQLRequestHelper(t *testing.T, endpoint string, params *GraphQLParams,
	token *HttpToken) *GraphQLResponse {

	b, err := json.Marshal(params)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(b))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	if token.AccessJwt != "" {
		req.Header.Set("X-Dgraph-AccessToken", token.AccessJwt)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	require.NoError(t, err)

	var gqlResp GraphQLResponse
	err = json.Unmarshal(b, &gqlResp)
	require.NoError(t, err)

	return &gqlResp
}

func (errList GqlErrorList) Error() string {
	var buf bytes.Buffer
	for i, gqlErr := range errList {
		if i > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(gqlErr.Error())
	}
	return buf.String()
}

func (gqlErr *GqlError) Error() string {
	var buf bytes.Buffer
	if gqlErr == nil {
		return ""
	}

	buf.WriteString(gqlErr.Message)
	return buf.String()
}

func MakeGQLRequest(t *testing.T, endpoint string, params *GraphQLParams,
	token *HttpToken) *GraphQLResponse {
	resp := MakeGQLRequestHelper(t, endpoint, params, token)
	if len(resp.Errors) == 0 || !strings.Contains(resp.Errors.Error(), "Token is expired") {
		return resp
	}
	var err error
	token, err = HttpLogin(&LoginParams{
		Endpoint:   endpoint,
		UserID:     token.UserId,
		Passwd:     token.Password,
		RefreshJwt: token.RefreshToken,
	})
	require.NoError(t, err)
	return MakeGQLRequestHelper(t, endpoint, params, token)
}

// HttpLogin sends a HTTP request to the server
// and returns the access JWT and refresh JWT extracted from
// the HTTP response
func HttpLogin(params *LoginParams) (*HttpToken, error) {
	login := `mutation login($userId: String, $password: String, $namespace: Int, $refreshToken: String) {
		login(userId: $userId, password: $password, namespace: $namespace, refreshToken: $refreshToken) {
			response {
				accessJWT
				refreshJWT
			}
		}
	}`

	gqlParams := GraphQLParams{
		Query: login,
		Variables: map[string]interface{}{
			"userId":       params.UserID,
			"password":     params.Passwd,
			"namespace":    params.Namespace,
			"refreshToken": params.RefreshJwt,
		},
	}
	body, err := json.Marshal(gqlParams)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal body: %w", err)
	}

	req, err := http.NewRequest("POST", params.Endpoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("login through curl failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read from response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got non 200 response from the server with %s ", string(respBody))
	}
	var outputJson map[string]interface{}
	if err := json.Unmarshal(respBody, &outputJson); err != nil {
		var errOutputJson map[string]interface{}
		if err := json.Unmarshal(respBody, &errOutputJson); err == nil {
			if _, ok := errOutputJson["errors"]; ok {
				return nil, fmt.Errorf("response error: %v", string(respBody))
			}
		}
		return nil, fmt.Errorf("unable to unmarshal the output to get JWTs: %w", err)
	}

	data, found := outputJson["data"].(map[string]interface{})
	if !found {
		return nil, fmt.Errorf("data entry found in the output: %w", err)
	}

	l, found := data["login"].(map[string]interface{})
	if !found {
		return nil, fmt.Errorf("data entry found in the output: %w", err)
	}

	response, found := l["response"].(map[string]interface{})
	if !found {
		return nil, fmt.Errorf("data entry found in the output: %w", err)
	}

	newAccessJwt, found := response["accessJWT"].(string)
	if !found || newAccessJwt == "" {
		return nil, errors.New("no access JWT found in the output")
	}
	newRefreshJwt, found := response["refreshJWT"].(string)
	if !found || newRefreshJwt == "" {
		return nil, errors.New("no refresh JWT found in the output")
	}

	return &HttpToken{
		UserId:       params.UserID,
		Password:     params.Passwd,
		AccessJwt:    newAccessJwt,
		RefreshToken: newRefreshJwt,
	}, nil
}
