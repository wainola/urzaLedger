package handlers

import (
	"encoding/json"
	"net/http"
	"urza/utils"
)

type PostIncomeBody struct {
	Income    string
	Amount    int
	Source    string
	CreatedAt string
}

func Incomes(appEnvironment *UrzaEnvironment) http.Handler {
	return UrzaApp{appEnvironment, postIncome}
}

func PostIncomeValidation(p PostIncomeBody) bool {
	if p.Income == "" {
		return false
	} else if p.Amount == 0 {
		return false
	} else if p.Source == "" {
		return false
	} else {
		return true
	}
}

func postIncome(ue *UrzaEnvironment, w http.ResponseWriter, r *http.Request) error {
	body := r.Body

	_, userId := utils.ExtractUrlToProcess(r.URL.Path, "post", "incomes")

	bodyToValidate := PostIncomeBody{}

	err := json.NewDecoder(body).Decode(&bodyToValidate)

	if err != nil {
		panic(err)
	}

	validation := PostIncomeValidation(bodyToValidate)

	if !validation {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		responseToSend := Response{http.StatusUnprocessableEntity, "Unprocessable Entity"}
		json.NewEncoder(w).Encode(responseToSend)
	}

	result := ue.DB.CreateIncome(bodyToValidate.Income, bodyToValidate.Amount, bodyToValidate.Source, bodyToValidate.CreatedAt, userId)

	if result {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		res := Response{201, "Created"}
		json.NewEncoder(w).Encode(res)
	}

	return nil
}

func GetIncomes(appEnvironment *UrzaEnvironment) http.Handler {
	return UrzaApp{appEnvironment, getIncomes}
}

func getIncomes(ue *UrzaEnvironment, w http.ResponseWriter, r *http.Request) error {
	_, userId := utils.ExtractUrlToProcess(r.URL.Path, "post", "incomes")

	result := ue.DB.GetIncomes(userId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

	return nil
}
