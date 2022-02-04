package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"goji.io"
	"goji.io/pat"

	"github.com/google/uuid"
)

type UploadData struct {
	FullName      string    `json:"full_name"`
	Email         string    `json:"email"`
	SignaturePath string    `json:"signature_path"`
	Questions     string    `json:"questions"`
	PhotoPath     string    `json:"photo_path"`
	SubmittedAt   time.Time `json:"submitted_at"`
}

func savePhoto(id string, r *http.Request, field string) (string, error) {
	file, _, err := r.FormFile(field)
	if err != nil {
		log.Print("0")
		return "", err
	}

	defer file.Close()

	var snapshot []byte
	_, err = file.Read(snapshot)
	if err != nil {
		log.Print("1")
		return "", err
	}

	filepath := fmt.Sprintf("uploads/%s-%s.png", id, field)
	os.MkdirAll("uploads/", os.ModePerm)

	dst, err := os.Create(filepath)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filepath, err
}

func uploadPhoto(w http.ResponseWriter, r *http.Request) {
	uid := uuid.New()
	submitterId := strings.Replace(uid.String(), "-", "", -1)
	signaturePath, err := savePhoto(submitterId, r, "signature")
	if err != nil {
		http.Error(w, fmt.Sprintf("Err: %s", err), http.StatusBadRequest)
		return
	}

	photoPath, err := savePhoto(submitterId, r, "photo")
	if err != nil {
		http.Error(w, fmt.Sprintf("Err: %s", err), http.StatusBadRequest)
		return
	}

	uploadInformation := UploadData{
		FullName:      r.FormValue("fullName"),
		Email:         r.FormValue("email"),
		Questions:     r.FormValue("questions"),
		PhotoPath:     photoPath,
		SignaturePath: signaturePath,
		SubmittedAt:   time.Now(),
	}

	os.MkdirAll("forms/", os.ModePerm)
	jsonStr, err := json.Marshal(uploadInformation)
	if err != nil {
		http.Error(w, "err", http.StatusBadRequest)
	}
	os.WriteFile(fmt.Sprintf("forms/%s.json", submitterId), jsonStr, 0644)

	fmt.Fprint(w, "Uploaded!")
}

func main() {
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	mux := goji.NewMux()

	fs := http.FileServer(http.Dir(workdir + "/dist/"))
	mux.Handle(pat.Get("/*"), fs)

	mux.HandleFunc(pat.Post("/api/upload"), uploadPhoto)

	http.ListenAndServe("localhost:8000", mux)
}
