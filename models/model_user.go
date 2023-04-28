/*
 * ECE 461 - Spring 2023 - Project 2
 *
 * API for ECE 461/Spring 2023/Project 2: A Trustworthy Module Registry
 *
 * API version: 3.0.2
 * Contact: davisjam@purdue.edu
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type User struct {
	Name string `json:"name" bson:"name"`
	// Is this user an admin?
	IsAdmin bool `json:"isAdmin" bson:"isadmin"`
}