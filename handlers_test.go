package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// A MASTER_KEY valida sem tocar no banco (DB nil: se consultasse, o teste entraria em pânico).
func TestValidateAceitaMasterKey(t *testing.T) {
	app := &App{MasterKey: "master-de-teste"}
	req := httptest.NewRequest(http.MethodGet, "/validate", nil)
	req.Header.Set("Authorization", "Bearer master-de-teste")
	rec := httptest.NewRecorder()

	app.validateKeyHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, esperado 200", rec.Code)
	}
}

func TestValidateSemChaveRetorna401(t *testing.T) {
	app := &App{MasterKey: "master-de-teste"}
	rec := httptest.NewRecorder()

	app.validateKeyHandler(rec, httptest.NewRequest(http.MethodGet, "/validate", nil))

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, esperado 401", rec.Code)
	}
}

// MasterKey vazia nunca pode validar uma chave vazia ou qualquer outra.
func TestValidateMasterKeyVaziaNaoValida(t *testing.T) {
	app := &App{MasterKey: ""}
	req := httptest.NewRequest(http.MethodGet, "/validate", nil)
	req.Header.Set("Authorization", "Bearer ")
	rec := httptest.NewRecorder()

	app.validateKeyHandler(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, esperado 401", rec.Code)
	}
}
