package iam

import "github.com/escaletech/go-escale/httpclient"

type IAM struct {
	URL              string
	token            string
	env              string
	timeoutInSeconds int
	HTTPClient       httpclient.HTTPClientInterface
}

type Permission struct {
	Rule    string   `json:"rule,omitempty"`
	Title   string   `json:"title,omitempty"`
	Actions []string `json:"actions,omitempty"`
}

type IamApplication struct {
	ID          string       `json:"id,omitempty"`
	Title       string       `json:"title,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
}
