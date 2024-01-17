package lib_test

import (
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"goth.stack/lib"
)

func TestPublish(t *testing.T) {
	db, mock := redismock.NewClientMock()
	mock.ExpectPublish("mychannel", "mymessage").SetVal(1)

	err := lib.Publish(db, "mychannel", "mymessage")
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Then you can check the channel name in your test
func TestSubscribe(t *testing.T) {
	db, _ := redismock.NewClientMock()

	pubsub, channel := lib.Subscribe(db, "mychannel")
	assert.NotNil(t, pubsub)
	assert.Equal(t, "mychannel", channel)
}
