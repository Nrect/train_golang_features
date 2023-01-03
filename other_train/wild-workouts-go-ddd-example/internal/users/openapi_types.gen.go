// Package main provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package main

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// User defines model for User.
type User struct {
	Balance     int    `json:"balance"`
	DisplayName string `json:"displayName"`
	Role        string `json:"role"`
}