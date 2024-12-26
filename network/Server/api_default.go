package Server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type Info struct {
	Name    []string
	Surname []string
	Age     []int
}

func Hello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		formTemplate := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Заполните форму</title>
			</head>
			<body>
				<h1>Введите информацию</h1>
				<form action="/hello" method="POST">
					<label for="name">Имя:</label>
					<input type="text" id="name" name="name" required>
					
					<label for="surname">Фамилия:</label>
					<input type="text" id="surname" name="surname" required>
			
					<label for="age">Возраст:</label>
					<input type="number" id="age" name="age" required>
			
					<button type="submit">Отправить</button>
				</form>
			</body>
			</html>
			`

		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		// Отправим HTML в ответ
		tmpl := template.New("form")
		tmpl, err := tmpl.Parse(formTemplate)
		if err != nil {
			http.Error(w, "Ошибка шаблона", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)

	case "POST":
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Некорректные данные формы", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		surname := r.FormValue("surname")
		age := r.FormValue("age")

		responseMessage := fmt.Sprintf("Привет, %s %s! Ваш возраст: %s", name, surname, age)
		response := Response{
			Message: responseMessage,
		}
		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
	}
}
