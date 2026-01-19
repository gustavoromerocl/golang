package db

import (
	"context"
	"log"
	"strconv"

	"github.com/gustavoromerocl/social/cmd/internal/store"
)

func Seed(store *store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	log.Println("Creating", len(users), "users...")
	for i := range users {
		err := store.Users.Create(ctx, &users[i])
		if err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}
	log.Println("Users created successfully")

	posts := generatePosts(200, users)
	log.Println("Creating", len(posts), "posts...")
	for i := range posts {
		err := store.Posts.Create(ctx, &posts[i])
		if err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}
	log.Println("Posts created successfully")

	comments := generateComments(300, users, posts)
	log.Println("Creating", len(comments), "comments...")
	for i := range comments {
		_, err := store.Comments.Create(ctx, &comments[i])
		if err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}
	log.Println("Seeding complete")
}

func generateUsers(n int) []store.User {
	users := make([]store.User, n)
	for i := 0; i < n; i++ {
		users[i] = store.User{
			Username: "user" + strconv.Itoa(i+1),
			Email:    "user" + strconv.Itoa(i+1) + "@example.com",
			Password: "password123",
		}
	}
	return users
}

func generatePosts(n int, users []store.User) []store.Post {
	posts := make([]store.Post, n)
	for i := 0; i < n; i++ {
		posts[i] = store.Post{
			Title:   "Post Title " + strconv.Itoa(i+1),
			Content: "This is the content of post number " + strconv.Itoa(i+1),
			UserID:  users[i%len(users)].ID,
			Tags:    []string{"tag1", "tag2"},
		}
	}
	return posts
}

func generateComments(n int, users []store.User, posts []store.Post) []store.Comment {
	comments := make([]store.Comment, n)
	for i := 0; i < n; i++ {
		comments[i] = store.Comment{
			PostID:  posts[i%len(posts)].ID,
			UserID:  users[i%len(users)].ID,
			Content: "This is a comment number " + strconv.Itoa(i+1),
		}
	}
	return comments
}
