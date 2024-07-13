package dto

type CreateAccountRequest struct { //создаем аккаунт пользователя
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type PatchAccountRequest struct { //меняем имя аккаунта
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ChangeAccountRequest struct { //начислить деньги аккаунту
	Name    string `json:"name"`
	NewName string `json:"new name"`
}

type DeleteAccountRequest struct { //удалить аккаунт
	Name string `json:"name"`
}
