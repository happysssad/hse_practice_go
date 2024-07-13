package main

import (
	"awesomeProject/accounts/dto"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	Amount  int
	NewName string
}

func main() {
	portVal := flag.Int("port", 8080, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command to execute")
	nameVal := flag.String("name", "", "name of account")
	amountVal := flag.Int("amount", 0, "amount of account")
	newNameVal := flag.String("new", "", "new name of account")

	flag.Parse()

	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *amountVal,
		NewName: *newNameVal,
	}

	if err := do(cmd); err != nil {
		panic(err)
	}
}

func (c *Command) Do() error {
	switch c.Cmd {
	case "create":
		return c.create()
	default:
		return fmt.Errorf("unknown command: %s", c.Cmd)
	}
}

func (c *Command) create() error {
	panic("implement me")
}

func do(cmd Command) error {
	switch cmd.Cmd {
	case "create":
		if err := create(cmd); err != nil {
			return fmt.Errorf("create account failed: %w", err)
		}

		return nil
	case "get":
		if err := get(cmd); err != nil {
			return fmt.Errorf("get account failed: %w", err)
		}

		return nil
	case "delete":
		if err := deleteCli(cmd); err != nil {
			return fmt.Errorf("delete account failed %w", err)
		}

		return nil
	case "patch":
		if err := patch(cmd); err != nil {
			return fmt.Errorf("patch account failed %w", err)
		}

		return nil
	case "change":
		if err := change(cmd); err != nil {
			return fmt.Errorf("change account failed %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unknown command %s", cmd.Cmd)
	}
}

func change(cmd Command) error {
	req := dto.ChangeAccountRequest{
		Name:    cmd.Name,
		NewName: cmd.NewName,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err2 := http.Post(
		fmt.Sprintf("http://%s:%d/account/change", cmd.Host, cmd.Port), "application/json",
		bytes.NewBuffer(data))

	if err2 != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, _ := io.ReadAll(resp.Body)
	return fmt.Errorf("http post failed: %s", string(body))
}

func patch(cmd Command) error {
	request := dto.PatchAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("Json marshal failed: %w", err)
	}

	resp, err2 := http.Post(
		fmt.Sprintf("http://%s:%d/account/patch", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)

	if err2 != nil {
		return fmt.Errorf("Server patch failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, _ := io.ReadAll(resp.Body)

	return fmt.Errorf("resp error: %s", string(body))
}

func deleteCli(cmd Command) error {
	req := dto.DeleteAccountRequest{
		Name: cmd.Name,
	}

	_, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("JSON marshal failed: %w", err)
	}

	r, err := http.NewRequest(http.MethodDelete,
		fmt.Sprintf("http://%s:%d/account/delete?name=%s", cmd.Host, cmd.Port, cmd.Name),
		nil)
	if err != nil {
		return fmt.Errorf("http delete creation failed: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("http delete creation failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error: %s", string(body))
}

func get(cmd Command) error {
	resp, err := http.Get(
		fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}

		return fmt.Errorf("resp error %s", string(body))
	}

	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}

	fmt.Printf("response account name: %s and amount: %d", response.Name, response.Amount)

	return nil
}

func create(cmd Command) error {
	request := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("json marshal failed: %w", err)
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}

	defer func() { //всегда после того, как сделали запрос, надо дефер, чтобы избежать утечку ресурсов
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}

	return fmt.Errorf("resp error %s", string(body))
}
