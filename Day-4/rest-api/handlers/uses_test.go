package handlers

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"rest-api/middleware"
	"rest-api/models"
	"rest-api/models/mockmodels"
	"testing"
)

func TestSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// sample newUser that we would get after reading JSON body
	newUser := models.NewUser{
		Name:     "John Doe",
		Age:      30,
		Email:    "d@email.com",
		Password: "abc",
	}

	// User that we would get from CreateUser method of models package
	mockUser := models.User{
		Id:           "ab49a45c-ec2c-47a5-8675-9f072e2d9216",
		Email:        "d@email.com",
		Name:         "John Doe",
		Age:          30,
		PasswordHash: "2a$10$EimVQRw4YiKIoMqh3JMwOesA9ngPGZT.chFEmPSaHzYl.mlnhLr12",
	}

	tt := []struct {
		name             string
		body             []byte
		expectedStatus   int
		expectedResponse string
		mockStore        func(m *mockmodels.MockService)
	}{
		{
			name: "Ok",
			body: []byte(`{
   				 "name": "John Doe",
   				 "age": 30,
                 "email": "d@email.com",
                 "password": "abc"
			}`),
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"id":"ab49a45c-ec2c-47a5-8675-9f072e2d9216","email":"d@email.com","name":"John Doe","age":30,"password_hash":"2a$10$EimVQRw4YiKIoMqh3JMwOesA9ngPGZT.chFEmPSaHzYl.mlnhLr12"}`,
			mockStore: func(m *mockmodels.MockService) {
				m.EXPECT().CreateUser(gomock.Eq(newUser)).Return(mockUser, nil).Times(1)
			},
		},
	}
	//we need gin router to register the /signup endpoint
	router := gin.New()

	ctrl := gomock.NewController(t)

	//NewMockService would return the implementation of the interface
	mockService := mockmodels.NewMockService(ctrl)

	// setting the handler with the mock implementation of the interface
	h := handler{
		conn:     mockService,
		validate: validator.New(),
	}

	router.POST("/signup", h.Signup)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// calling the mock store field which is a function, to set the test expectations
			tc.mockStore(mockService)

			// Create a fake TraceID.
			traceID := "fake-trace-id"

			ctx := context.Background()
			// Insert the TraceId into the context.
			ctx = context.WithValue(ctx, middleware.TraceIdKey, traceID)

			// creating the request object with context
			req := httptest.NewRequestWithContext(ctx, http.MethodPost, "/signup", bytes.NewReader(tc.body))
			rec := httptest.NewRecorder()

			//making call to the handler function
			router.ServeHTTP(rec, req)
			require.Equal(t, tc.expectedStatus, rec.Code)
			require.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}

}
