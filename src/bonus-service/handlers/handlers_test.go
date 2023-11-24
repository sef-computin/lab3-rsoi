package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lab2/src/bonus-service/dbhandler/mock_dbhandler"
	"lab2/src/bonus-service/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreatePrivilege(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockBonusDB, record *models.Privilege)
	tests := []struct {
		name                 string
		input                *models.Privilege
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			username: "TestMax",
			input: &models.Privilege{
				Username: "TestMax",
				Balance:  0,
			},

			mockBehaviour: func(r *mock_dbhandler.MockBonusDB, input *models.Privilege) {
				r.EXPECT().CreatePrivilege(input)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			privilege := &models.Privilege{
				Username: test.username,
				Balance:  0,
			}

			dbhandler := mock_dbhandler.NewMockBonusDB(c)
			test.mockBehaviour(dbhandler, test.input)

			handler := BonusHandler{DBHandler: dbhandler}

			r := gin.New()
			r.POST("/api/v1/bonus/privilege", handler.CreatePrivilegeHandler)

			jsonValue, _ := json.Marshal(privilege)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/bonus/privilege", bytes.NewBuffer(jsonValue))
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}

func TestUpdatePrivilege(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockBonusDB, record *models.Privilege)
	tests := []struct {
		name                 string
		input                *models.Privilege
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			username: "TestMax",
			input: &models.Privilege{
				Username: "TestMax",
				Balance:  0,
			},

			mockBehaviour: func(r *mock_dbhandler.MockBonusDB, input *models.Privilege) {
				r.EXPECT().UpdatePrivilege(input)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			privilege := &models.Privilege{
				Username: test.username,
				Balance:  0,
			}

			dbhandler := mock_dbhandler.NewMockBonusDB(c)
			test.mockBehaviour(dbhandler, test.input)

			handler := BonusHandler{DBHandler: dbhandler}

			r := gin.New()
			r.POST("/api/v1/bonus/privilege", handler.UpdatePrivilegeHandler)

			jsonValue, _ := json.Marshal(privilege)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/bonus/privilege", bytes.NewBuffer(jsonValue))
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}

func TestGetPrivilegeByUsername(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockBonusDB, username string, output *models.Privilege)
	tests := []struct {
		name                 string
		input                models.Privilege
		output               models.Privilege
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			username: "TestMax",
			input: models.Privilege{
				Username: "TestMax",
				Balance:  0,
			},

			mockBehaviour: func(r *mock_dbhandler.MockBonusDB, username string, output *models.Privilege) {
				r.EXPECT().GetPrvilegeByUsername(username).Return(output, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			var privilege models.Privilege

			dbhandler := mock_dbhandler.NewMockBonusDB(c)
			test.mockBehaviour(dbhandler, test.username, &privilege)

			handler := BonusHandler{DBHandler: dbhandler}

			r := gin.New()
			r.GET("/api/v1/bonus/:username", handler.GetPrivilegeByUsernameHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/bonus/%v", test.username), nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}

func TestGetHistoryById(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockBonusDB, privilegeId string, output []*models.PrivilegeHistory)
	tests := []struct {
		name                 string
		input                models.Privilege
		output               models.PrivilegeHistory
		privilegeId          string
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "Ok",
			username:    "TestMax",
			privilegeId: "1",
			input: models.Privilege{
				Username: "TestMax",
				Balance:  0,
			},

			mockBehaviour: func(r *mock_dbhandler.MockBonusDB, privilegeId string, output []*models.PrivilegeHistory) {
				r.EXPECT().GetHistoryById(privilegeId).Return(output, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			var privilege_history []*models.PrivilegeHistory
			dbhandler := mock_dbhandler.NewMockBonusDB(c)
			test.mockBehaviour(dbhandler, test.privilegeId, privilege_history)

			handler := BonusHandler{DBHandler: dbhandler}

			r := gin.New()
			r.GET("/api/v1/bonus/:privilegeId", handler.GetHistoryByIdHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/bonus/%v", test.privilegeId), nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}
