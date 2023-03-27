package handlers

import (
	"cd-handler/files"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Files struct {
}

type Request struct {
	Secret string `json:"secret"`
	Entry  string `json:"entry"`
	Files  []File `json:"files"`
}

func RegisterSecretHandler(secret string, run func() error) {
	secretHandler := func(writer http.ResponseWriter, r *http.Request) {
		formSecret := r.FormValue("secret")

		_, header, err := r.FormFile("entry")
		if err != nil {
			http.Error(writer, "bad entry", http.StatusBadRequest)
			return
		}
		err = files.FileWriter(*header)
		if err != nil {
			http.Error(writer, "bad entry file", http.StatusBadRequest)
			return
		}

		if r.Method != "POST" {
			http.Error(writer, "", http.StatusMethodNotAllowed)
			return
		}

		if formSecret != secret {
			http.Error(writer, "bad secret", http.StatusUnauthorized)
			return
		}

		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			http.Error(writer, "bad files", http.StatusBadRequest)
			return
		}

		fhs := r.MultipartForm.File["myFiles"]

		for _, fh := range fhs {
			err = files.FileWriter(*fh)
			if err != nil {
				http.Error(writer, "bad file "+fh.Filename, http.StatusBadRequest)
				return
			}
		}

		err = run()
		if err != nil {
			fmt.Println(err)
			http.Error(writer, "bad entry exec", http.StatusBadRequest)
			return
		}

		writer.Header().Set("Content-Type", "application/json")

		responseJson, err := json.Marshal(Response{Status: "ok"})
		if err != nil {
			http.Error(writer, "json error", http.StatusInternalServerError)
			return
		}

		writer.Write(responseJson)
	}

	http.HandleFunc("/secret", secretHandler)
}
