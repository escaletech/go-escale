package mysqldb_test

import (
	"strconv"
	"testing"

	"github.com/escaletech/go-escale/mysqldb"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type cases struct {
	description      string
	err              error
	expectedResponse bool
}

func TestIsRecordNotFoundError(t *testing.T) {
	testCases := []cases{
		{
			"Error type is \"ErrRecordNotFound\"",
			gorm.ErrRecordNotFound,
			true,
		},
		{
			"Error type is a generic one",
			*new(error),
			false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			t.Run("returns "+strconv.FormatBool(testCase.expectedResponse), func(t *testing.T) {

				resp := mysqldb.IsRecordNotFoundError(testCase.err)
				assert.Equal(t, testCase.expectedResponse, resp)
			})
		})
	}
}
