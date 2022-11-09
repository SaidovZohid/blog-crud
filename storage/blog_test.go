package storage

import (
	"testing"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createBlog(t *testing.T) *Blog {
	var blog *Blog
	blog, err := dbManager.Create(&Blog{
		Title: faker.Sentence(),
		Description: faker.Sentence(),
		Author: faker.Name(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, blog)
	return blog
}

func deleteBlog(id int, t *testing.T) {
	err := dbManager.DeleteBlog(id)
	require.NoError(t, err)
}

func TestGetBlog(t *testing.T) {
	b := createBlog(t)
	blog, err := dbManager.GetBlog(b.Id)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	deleteBlog(b.Id, t)
}

func TestCreateBlog(t *testing.T) {
	b := createBlog(t)
	deleteBlog(b.Id, t)
}

func TestUpdateBlog(t *testing.T) {
	b := createBlog(t)
	b.Author = faker.Name()
	b.Title = faker.Sentence()
	b.Description = faker.Sentence()
	blog, err := dbManager.UpdateBlog(b)
	require.NoError(t, err)
	require.NotEmpty(t, blog)

	deleteBlog(b.Id, t)
}

func TestDeleteBlog(t *testing.T) {
	b := createBlog(t)
	deleteBlog(b.Id, t)
}


func TestGetAllBlogs(t *testing.T) {
	b := createBlog(t)

	blogs, err := dbManager.GetAll(&GetBlogsQueryParam{
		Limit: 10,
		Page: 1,
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(blogs), 1)

	deleteBlog(b.Id, t)
}