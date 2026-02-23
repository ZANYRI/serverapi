package data

import (

)

type Request struct {
	Provider string `json:"provider"`
	Model 	string 	`json:"model"`
	Message string	`json:"message"`
}