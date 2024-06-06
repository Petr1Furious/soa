package swagger

import (
	"bytes"
	"net/http/httptest"
	"testing"
)

func TestGenerateSessionID(t *testing.T) {
	for i := 0; i < 1000; i++ {
		id1, err1 := generateSessionID()
		if err1 != nil {
			t.Fatalf("generateSessionID() error = %v", err1)
		}

		id2, err2 := generateSessionID()
		if err2 != nil {
			t.Fatalf("generateSessionID() error = %v", err2)
		}

		if id1 == id2 {
			t.Fatalf("generateSessionID() should generate unique IDs, got %v and %v", id1, id2)
		}
	}
}

func TestReadBody(t *testing.T) {
	body := []byte(`{"key":"value"}`)
	req := httptest.NewRequest("POST", "http://example.com", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	result, ok := ReadBody(w, req)

	if !ok {
		t.Errorf("ReadBody returned false, expected true")
	}
	if !bytes.Equal(result, body) {
		t.Errorf("ReadBody returned %v, expected %v", result, body)
	}
}
