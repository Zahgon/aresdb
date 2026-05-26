package utils

import (
	"math/rand"

	"time"

	"github.com/uber/aresdb/client"
	"go.uber.org/zap"
)

var (
	random        = rand.New(rand.NewSource(0))
	unitToSeconds = map[string]time.Duration{
		"m": time.Minute,
		"h": time.Hour,
		"d": 24 * time.Hour,
	}

	logger = zap.NewExample().Sugar()
)

type TableCreationResp struct {
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func parseAndGenerateRandomTime(timeStr string) uint32 { _ = "STUB: not implemented"; return 0 }

func PanicIfErr(err error) { _ = "STUB: not implemented"; return }

func logError(msg string) { _ = "STUB: not implemented"; return }

func ingestDataForArrayTestTable(connector client.Connector, tableName string, dataFilePath string) {
	_ = "STUB: not implemented"
	return
}

func generateArrayTableColValue(val string, arraySize int) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func generateArrayValue(valType string, arraySize int) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func generateArrayItemValue(valType string, itemNo int) string {
	_ = "STUB: not implemented"
	return ""
}

func MakeQuery(host string, port int, queryName, queryType, queryPath string) string {
	_ = "STUB: not implemented"
	return ""
}

func CreateTable(host string, port int, tableName string, tableSchemaPath string) {
	_ = "STUB: not implemented"
	return
}

// TODO(lucafuji): need a better error code mapping here to tell it's a table exists error.

func IngestDataForTable(host string, port int, tableName string, dataPath string) {
	_ = "STUB: not implemented"
	return
}

// TODO for some hard coded table for array test now
