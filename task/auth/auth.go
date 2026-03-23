package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RequestValidate(accessToken string) (string, error, int) {
	var username string
	authorizationURL := "http://localhost:8080/validate"
	client := &http.Client{}

	// Создание HTTP-запроса для проверки токена
	authorizeReq, err := http.NewRequest("GET", authorizationURL, nil)
	if err != nil {
		return "", fmt.Errorf("Ошибка при создании запроса на валидацию токена: %w", err), http.StatusInternalServerError
	}

	// Добавление токена в заголовок запроса
	authorizeReq.Header.Set("Authorization", accessToken)

	// Отправка запроса
	resp, err := client.Do(authorizeReq)
	if err != nil {
		return "", fmt.Errorf("Ошибка при выполнении запроса: %w", err), http.StatusInternalServerError
	}
	defer resp.Body.Close()

	// Проверка ответа от сервиса авторизации
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Токен не валиден"), http.StatusUnauthorized
	}

	err = json.NewDecoder(resp.Body).Decode(&username)
	if err != nil {
		return "", fmt.Errorf("Ошибка в теле запроса: %w", err), http.StatusBadRequest
	}

	return username, nil, http.StatusOK
}

func RequestRefreshToken(refreshToken string) (string, error, int) {
	var accessToken string
	var refreshURL string = "http://localhost:8080/refreshtoken"
	var client *http.Client = &http.Client{}

	//Создание HTTP-запроса для обновления токена
	refreshReq, err := http.NewRequest("GET", refreshURL, nil)
	if err != nil {
		return "", fmt.Errorf("Ошибка при создании запроса на обновление токена: %w", err), http.StatusInternalServerError
	}

	//Добавление токена в заголовок запроса
	refreshReq.Header.Set("Authorization", refreshToken)

	//Отправка запроса
	resp, err := client.Do(refreshReq)
	if err != nil {
		return "", fmt.Errorf("Ошибка при выполнении запроса: %w", err), http.StatusInternalServerError
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Ошибка: %s, Ответ: %s", http.StatusText(resp.StatusCode), string(bodyBytes)), resp.StatusCode
	}

	err = json.NewDecoder(resp.Body).Decode(&accessToken)
	if err != nil {
		return "", fmt.Errorf("Ошибка в теле запроса: %w", err), http.StatusBadRequest
	}

	return accessToken, nil, http.StatusOK
}
