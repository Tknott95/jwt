package main

import (
	"fmt"
	"github.com/tknott95/JwtPlaytime"
	"html/template"
	"log"
	"net/http"
)

const loginTemplate = `
<h1>Enter your username and password</h1>
<form action="/" method="POST">
	<input type="text" name="user" required>

	<label for="password">Password</label>
	<input type="password" name="password" required>

	<input type="submit" value="Submit">
</form>
`

func authHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		t, _ := template.New("login").Parse(loginTemplate)
		t.Execute(w, nil)
	case "POST":
		user := r.FormValue("user")
		pass := r.FormValue("password")
		err := users.AuthenticateUser(user, pass)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users.SetSession(w, user)
		w.Write([]byte("Signed in successfully"))
	}
}

func restrictedHandler(w http.ResponseWriter, r *http.Request) {
	user := users.GetSession(w, r)
	w.Write([]byte(user))
}

func oauthRestrictedHandler(w http.ResponseWriter, r *http.Request) {
	user, err := users.VerifyToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Write([]byte(user))
}

func sanitizeInputExample(str string) {
	fmt.Println("JS: ", template.JSEscapeString(str))
	fmt.Println("HTML: ", template.HTMLEscapeString(str))
}

func main() {
	sanitizeInputExample("<script>alert(\"Hi!\");</sciprt>")

	username, password := "tex", "hola123"

	err := users.NewUser(username, password)
	if err != nil {
		fmt.Printf("User already exists: %s\n", err.Error())
	} else {
		fmt.Printf("Succesfully created and authenticated user \033[32m%s\033[0m\n", username)
	}

	http.HandleFunc("/", authHandler)
	http.HandleFunc("/auth/gplus/authorize", users.AuthURLHandler)
	http.HandleFunc("/auth/gplus/callback", users.CallbackURLHandler)
	http.HandleFunc("/oauth", oauthRestrictedHandler)
	http.HandleFunc("/restricted", restrictedHandler)

	log.Fatal(http.ListenAndServeTLS(":3000", "server.pem", "server.key", nil))
}
