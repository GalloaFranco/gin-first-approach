package services

import (
	"github.com/GalloaFranco/gin-first-approach/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) CloseDB() {}

func (mock *MockRepository) Update(video entities.Video) {}

func (mock *MockRepository) Delete(video entities.Video) {}

func (mock *MockRepository) Save(video entities.Video) entities.Video {
	args := mock.Called()
	result := args.Get(0)
	return result.(entities.Video)
}

func (mock *MockRepository) FindAll() []entities.Video {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.Video)
}

func Test_Should_New_Function_Return_IVideoService_Type(t *testing.T) {
	result := New(nil)
	assert.Equal(t, reflect.TypeOf(New(nil)), reflect.TypeOf(result))
}

func Test_Should_FindAll_Method_Be_Called(t *testing.T) {
	mockRepo := new(MockRepository)
	mockRepo.On("FindAll").Return([]entities.Video{
		entities.Video{
			ID:          uint64(123),
			Title:       "FakeTitle",
			Description: "FakeDescription",
			URL:         "FakeUrl",
		},
	})

	sut := New(mockRepo)
	result := sut.FindAll()

	// AssertExpectations asserts that everything specified with On and Return was
	// in fact called as expected.  Calls may have occurred in any order.
	mockRepo.AssertExpectations(t)

	assert.Equal(t, "FakeTitle", result[0].Title)
	assert.Equal(t, "FakeDescription", result[0].Description)
	assert.Equal(t, "FakeUrl", result[0].URL)
	assert.Equal(t, uint64(123), result[0].ID)

}

func Test_Should_Save_Method_Be_Called(t *testing.T) {
	video := entities.Video{
		ID:          uint64(123),
		Title:       "FakeTitle",
		Description: "FakeDescription",
		URL:         "FakeUrl",
	}
	mockRepo := new(MockRepository)
	mockRepo.On("Save", mock.Anything).Return(video)

	sut := New(mockRepo)
	result := sut.Save(video)

	assert.Equal(t, "FakeTitle", result.Title)
	mockRepo.AssertExpectations(t)

	assert.Equal(t, "FakeTitle", result.Title)
	assert.Equal(t, "FakeDescription", result.Description)
	assert.Equal(t, "FakeUrl", result.URL)
	assert.Equal(t, uint64(123), result.ID)
}
