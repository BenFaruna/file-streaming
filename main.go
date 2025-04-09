package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ServerResponse struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	OK     bool   `json:"ok"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/stream-upload", StreamUpload)
	mux.HandleFunc("/file-upload", FileUpload)

	mux.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func FileUpload(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)
	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("file")
	if err != nil {
		writer.Encode(ServerResponse{Status: http.StatusBadRequest, Msg: "file upload failed"})
	}
	fmt.Println(r.Header.Get("Content-Type"))

	filename := header.Filename
	temp, _ := os.CreateTemp(".", fmt.Sprintf("*-%s", filename))
	defer temp.Close()

	_, err = io.Copy(temp, file)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		writer.Encode(ServerResponse{
			Status: http.StatusInternalServerError,
			Msg:    "error getting file content",
			OK:     false,
		})
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := writer.Encode(ServerResponse{
		Status: http.StatusCreated,
		Msg:    "file uploaded",
		OK:     true,
	}); err != nil {
		return
	}
}

func StreamUpload(w http.ResponseWriter, r *http.Request) {
	writer := json.NewEncoder(w)
	defer r.Body.Close()
	temp, _ := os.CreateTemp(".", "*.tmp")
	defer temp.Close()

	_, err := io.Copy(temp, r.Body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		writer.Encode(ServerResponse{
			Status: http.StatusInternalServerError,
			Msg:    "error getting file content",
			OK:     false,
		})
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := writer.Encode(ServerResponse{
		Status: http.StatusCreated,
		Msg:    "file uploaded",
		OK:     true,
	}); err != nil {
		return
	}
}
