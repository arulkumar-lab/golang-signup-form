	package main

	import (
		"log"
		"net/http"
		"html/template"
		)

	type PageVariables struct {
		PageTitle string
		FormElements []InputTextField
	}

	type InputTextField struct {
		Name string
		PlaceHolder string
		className string
		Text string
	}

	func signupHandler(w http.ResponseWriter, r *http.Request){
		Title := "GoLang Sample application"
		MyFormElements := []InputTextField{
			{
				Name:      "firstName",
				PlaceHolder:     "Enter First Name",
				className: "input-text",
				Text:      "First Name",
			},{
				Name:      "lastName",
				PlaceHolder:     "Enter Last Name",
				className: "input-text",
				Text:      "Last Name",
			},{
				Name:      "email",
				PlaceHolder:     "Enter Email",
				className: "input-text",
				Text:      "Email",
			},
			{
				Name:      "Address",
				PlaceHolder:     "Enter Address",
				className: "input-text",
				Text:      "Address",
			},
	}
		HomePageVariables := PageVariables{
			PageTitle : Title,
			FormElements: MyFormElements,
		}

		t, err := template.ParseFiles("index.html")
		if err != nil{
			log.Print("Template parsing error...", err)
		}

		err = t.Execute(w, HomePageVariables)
		if err != nil{
			log.Print("Template Executing error", err)
		}
	}

	func successHandler(w http.ResponseWriter, r *http.Request)  {
		r.ParseForm()

		//requestParams := r.Form.Get("firstName")

		Title := "Signup success!"
		MyPageVariables := PageVariables{
			PageTitle: Title,
		}
		t, err := template.ParseFiles("./template/success.html")
		if err != nil {
			log.Print("Template parsing error: ", err)
		}
		err = t.Execute(w, MyPageVariables)
		if err != nil {
			log.Print("Template loading error", err)
		}
	}

	func main() {
		http.HandleFunc("/", signupHandler)
		http.HandleFunc("/selected", successHandler)
		http.ListenAndServe("3300", nil)
		log.Fatal(http.ListenAndServe(":3300", nil))

	}