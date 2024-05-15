package conn

import (
	"errors"
	"io"

	"github.com/surrealdb/surrealdb.go/pkg/model"
)

type Connection interface {
	Connect(url string, exitChannel chan error) (Connection, error)
	Send(method string, params []interface{}) (interface{}, error)
	Close() error
	LiveNotifications(id string) (chan model.Notification, error)
}

var (
	ErrConnectionClosed       = errors.New("connection closed")
	ErrConnectionNotConnected = errors.New("connection not connected")
	ErrTimeout                = errors.New("timeout")
	ErrClosedPipe             = io.ErrClosedPipe
)
