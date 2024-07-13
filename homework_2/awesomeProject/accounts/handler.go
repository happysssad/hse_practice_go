package accounts

import (
	"awesomeProject/accounts/dto"
	"awesomeProject/accounts/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account), //инициализируем аккаунт
		guard:    &sync.RWMutex{},
	}
}

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex //защищаем мапу с помощью мьютекса
}

func (h *Handler) CreateAccount(c echo.Context) error {
	var request dto.CreateAccountRequest // {"name": "alice", "amount": 50}
	if err := c.Bind(&request); err != nil {
		c.Logger().Error(err)

		return c.String(http.StatusBadRequest, "Invalid request")
	}

	if len(request.Name) == 0 {
		return c.String(http.StatusBadRequest, "Empty name")
	} //проверяем, что имя непустое

	h.guard.Lock()

	if _, ok := h.accounts[request.Name]; ok {
		h.guard.Unlock()

		return c.String(http.StatusForbidden, "Account already exists")
	}

	h.accounts[request.Name] = &models.Account{ //записываем аккаунт
		Name:   request.Name,
		Amount: request.Amount,
	}

	h.guard.Unlock()

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) GetAccount(c echo.Context) error {
	name := c.QueryParams().Get("name")

	h.guard.RLock() //lock на чтение

	account, ok := h.accounts[name]

	h.guard.RUnlock()

	if !ok {
		return c.String(http.StatusNotFound, "Account not found")
	}

	response := dto.GetAccountResponse{
		Name:   account.Name,
		Amount: account.Amount,
	}

	return c.JSON(http.StatusOK, response) //кидаем структуру, сам в json превратит
}

func (h *Handler) DeleteAccount(c echo.Context) error { // Удаляет аккаунт
	var req dto.DeleteAccountRequest

	if len(req.Name) == 0 {
		return c.String(http.StatusBadRequest, "Empty name")
	}

	h.guard.RLock()
	defer h.guard.RUnlock()

	if _, ok := h.accounts[req.Name]; !ok {
		return c.String(http.StatusNotFound, "Account not found")
	}

	delete(h.accounts, req.Name)

	return c.NoContent(http.StatusOK)
}

func (h *Handler) PatchAccount(c echo.Context) error { // Меняет баланс
	var req dto.PatchAccountRequest

	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid req")
	}

	if len(req.Name) == 0 {
		return c.String(http.StatusBadRequest, "empty name")
	}

	if req.Amount == 0 {
		return c.String(http.StatusBadRequest, "empty amount")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[req.Name]
	if !ok {
		return c.String(http.StatusNotFound, "account not found")
	}

	account.Amount += req.Amount

	return c.NoContent(http.StatusOK)
}

func (h *Handler) ChangeAccount(c echo.Context) error { // Меняет имя
	var req dto.ChangeAccountRequest

	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "Invalid req")
	}

	if len(req.Name) == 0 {
		return c.String(http.StatusBadRequest, "Empty name")
	}

	if len(req.NewName) == 0 {
		return c.String(http.StatusBadRequest, "Empty new name")
	}

	h.guard.Lock()
	defer h.guard.Unlock()

	account, ok := h.accounts[req.Name]
	if !ok {
		return c.String(http.StatusNotFound, "Account not found")
	}

	if _, ok := h.accounts[req.NewName]; ok {
		return c.String(http.StatusConflict, "Account with new name already exists")
	}

	account.Name = req.NewName
	delete(h.accounts, req.Name)
	h.accounts[req.NewName] = account

	return c.NoContent(http.StatusOK)
}

// Написать клиент консольный, который делает запросы
