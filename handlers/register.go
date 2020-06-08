package handlers

import (
	"fmt"
	"net/http"
)

func Register(appEnvironment *UrzaEnvironment) http.Handler {
	return UrzaApp{appEnvironment, register}
}

func register(ue *UrzaEnvironment, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("FUNC REGISTER")
	return nil
}
