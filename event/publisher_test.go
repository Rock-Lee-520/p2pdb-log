package event

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func randInt(min int, max int) byte {

	rand.Seed(time.Now().UnixNano())

	return byte(min + rand.Intn(max-min))

}

func randUpString(l int) []byte {

	var result bytes.Buffer

	var temp byte

	for i := 0; i < l; {

		if randInt(65, 91) != temp {

			temp = randInt(65, 91)

			result.WriteByte(temp)

			i++

		}

	}

	return result.Bytes()

}

func TestPublishAsyncEvent(t *testing.T) {
	data := randUpString(19)
	PublishSyncEvent(StoreSqlInsertEvent, data)
}
