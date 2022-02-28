package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid // file name

	log.Printf("%scripts", vl)
	// open the video
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

// upload local files to server
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File cannot bigger than 500MB!")
		return
	}

	file, _, err := r.FormFile("file") // header omitted
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
	}
	filename := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+filename, data, 0666) // try not to use 777
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploading successfully")
}
