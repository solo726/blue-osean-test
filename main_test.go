package main

import (
	"testing"
	"net/http/httptest"
	"context"
	"time"
	"strings"
)

//func TestHandleHi_Recorder(t *testing.T) {
//	rw := httptest.NewRecorder()
//	req := httptest.NewRequest("GET", "/", nil)
//	handleHi(rw, req)
//	if !strings.Contains(rw.Body.String(), "visitor number") {
//		t.Errorf("Unexpected output: %s", rw.Body)
//	}
//	t.Log(rw.Body.String())
//}

func TestHandleHi_Recorder_Parallel(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	for i := 0; i < 2; i++ {
		go func() {
			rw := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", nil)
			handleHi(rw, req)
			select {
			case <-ctx.Done():
				t.Error("request timeout...")
			default:
				if !strings.Contains(rw.Body.String(), "visitor number") {
					t.Errorf("Unexpected output: %s", rw.Body)
				}
			}
			t.Log(rw.Body.String())
		}()
	}
	time.Sleep(2 * time.Second)
}

func BenchmarkHi(b *testing.B) {
	b.ReportAllocs()

	req := httptest.NewRequest("GET", "/", nil)
	for i := 0; i < b.N; i++ {
		rw := httptest.NewRecorder()
		handleHi(rw, req)
	}

}
