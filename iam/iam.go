package iam

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/escaletech/go-escale/httpclient"
)

func New(iamURL, iamToken, enviroment string, timeoutInSeconds int, httpClient httpclient.HTTPClientInterface) *IAM {
	return &IAM{
		URL:              iamURL,
		token:            iamToken,
		env:              enviroment,
		timeoutInSeconds: timeoutInSeconds,
		HTTPClient:       httpClient,
	}
}

func (i *IAM) RegisterApplicationFromPermissions(applicationName, applicationTitle string, permissions []Permission) error {
	iamApplication := IamApplication{
		ID:          applicationName,
		Title:       applicationTitle,
		Permissions: permissions,
	}
	return i.doRegistration(&iamApplication)
}

func (i *IAM) RegisterApplicationFromRoutes(applicationName, applicationTitle string, routes map[string][]string) error {
	iamApplication := IamApplication{
		ID:          applicationName,
		Title:       applicationTitle,
		Permissions: convertRoutersToPermissions(routes),
	}
	return i.doRegistration(&iamApplication)
}

func (i *IAM) doRegistration(application *IamApplication) error {

	headers := map[string]string{
		"Authorization": "Basic " + i.token,
	}

	body, _ := json.Marshal(application)

	resp, err := i.HTTPClient.DoRequest(httpclient.Request{
		Method:  "PUT",
		URL:     i.URL + "/services/" + application.ID,
		Headers: headers,
		Body:    bytes.NewBuffer(body),
		Config: httpclient.Config{
			TimeoutInSeconds: i.timeoutInSeconds,
		},
	})
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	respBody, _ := io.ReadAll(resp.Body)
	return errors.New(string(respBody))
}

func convertRoutersToPermissions(routes map[string][]string) []Permission {
	permissions := make([]Permission, 0)
	for url, methods := range routes {
		if len(methods) == 0 {
			continue
		}
		permission, err := generateRoutePermission(url, methods)
		if err != nil {
			continue
		}
		permissions = append(permissions, permission)
	}
	return permissions
}

func generateRoutePermission(url string, methods []string) (Permission, error) {
	rule := replaceUrlDash(replacePathParams(removeFirstDash(url)))
	return Permission{
		Rule:    rule,
		Title:   generateNameFromRule(rule),
		Actions: getActionsFromMethods(methods),
	}, nil
}

func getActionsFromMethods(methods []string) []string {
	actionsMap := make(map[string]struct{})
	for _, method := range methods {
		if method == "GET" {
			actionsMap["read"] = struct{}{}
		} else {
			actionsMap["write"] = struct{}{}
		}
	}
	actions := make([]string, 0)
	for k := range actionsMap {
		actions = append(actions, k)
	}
	return actions
}

func generateNameFromRule(rule string) string {
	clearPathparams := strings.Replace(rule, ":*", "", -1)
	splited := strings.Split(clearPathparams, ":")
	if len(splited) <= 1 {
		return "Basic Route"
	}
	splited = splited[1:]
	if len(splited) > 4 {
		splited = splited[len(splited)-4:]
	}
	return strings.Join(splited, " ")

}

func replacePathParams(url string) string {
	r, _ := regexp.Compile("{(.*?)}")
	return r.ReplaceAllString(url, "*")
}

func replaceUrlDash(url string) string {
	return strings.Replace(url, "/", ":", -1)
}

func removeFirstDash(url string) string {
	if len(url) > 0 && url[0] == '/' {
		return url[1:]
	}
	return url
}
