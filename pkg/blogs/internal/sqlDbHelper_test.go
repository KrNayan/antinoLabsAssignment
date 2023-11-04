package internal

import (
	"antinolabsassignment/pkg/common/models/blog"
	log "antinolabsassignment/pkg/common/utilities/logger"
	"testing"
)

func TestSqlDbHelper_Post(t *testing.T) {
	logger := log.New()
	sqlInstance, err := NewSqlDbHelper(logger, "antinolabs")
	if err != nil {
		t.Fatal(err)
	}
	if err = sqlInstance.Post(blog.BlogConfig{
		EmailId: "krnayan59@gmail.com",
		Blog:    "Hey, I am Nayan",
	}); err != nil {
		t.Fatal(err)
	}
	logger.Info("blog posted successfully!")
}

func TestSqlDbHelper_Get(t *testing.T) {
	logger := log.New()
	sqlInstance, err := NewSqlDbHelper(logger, "antinolabs")
	if err != nil {
		t.Fatal(err)
	}

	var resp = &blog.BlogConfig{}
	if resp, err = sqlInstance.Get(9); err != nil || resp == nil {
		t.Fatal(err)
	}
	logger.Info("blog retrieved successfully: ", *resp)
}

func TestSqlDbHelper_Update(t *testing.T) {
	logger := log.New()
	sqlInstance, err := NewSqlDbHelper(logger, "antinolabs")
	if err != nil {
		t.Fatal(err)
	}

	if err = sqlInstance.Update(blog.BlogConfig{
		BlogId: 1,
		Blog:   "Test Update",
	}); err != nil {
		t.Fatal(err)
	}
	logger.Info("blog updated successfully")
}

func TestSqlDbHelper_Delete(t *testing.T) {
	logger := log.New()
	sqlInstance, err := NewSqlDbHelper(logger, "antinolabs")
	if err != nil {
		t.Fatal(err)
	}

	if err = sqlInstance.Delete(1); err != nil {
		t.Fatal(err)
	}
	logger.Info("blog deleted successfully")
}
