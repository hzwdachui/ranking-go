package serviceranklist

import (
	"testing"

	"example.com/rankingSystem/logger"
)

func TestMain(m *testing.M) {
	logger.InitLogger("debug")
	m.Run()
}
