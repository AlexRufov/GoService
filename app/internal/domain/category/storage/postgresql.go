package postgresql

import (
	"GoService/app/pkg/client/postgresql"
	"GoService/app/pkg/logging"
)

type storage struct {
	client postgresql.Client
	logger *logging.Logger
}
