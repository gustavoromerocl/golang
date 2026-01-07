package main

import (
	"net/http"

	"github.com/gustavoromerocl/social/cmd/internal/store"
)

type createPostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags,omitempty"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload createPostPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}
	userId := 1 // In a real implementation, this would come from authentication context

	post := store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		UserID:  int64(userId),
		Tags:    payload.Tags,
	}
	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, &post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

}
