package nativeapi

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/deluan/rest"
	"github.com/navidrome/navidrome/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type mockMaintenance struct {
	pruneCalled bool
	pruneErr    error
}

func (m *mockMaintenance) DeleteMissingFiles(context.Context, []string) error {
	return nil
}

func (m *mockMaintenance) DeleteAllMissingFiles(context.Context) error {
	return nil
}

func (m *mockMaintenance) PruneMissing(context.Context) error {
	m.pruneCalled = true
	return m.pruneErr
}

var _ core.Maintenance = (*mockMaintenance)(nil)

var _ = Describe("Missing files API", func() {
	Describe("pruneMissingFiles", func() {
		It("calls maintenance and returns a JSON response", func() {
			maintenance := &mockMaintenance{}
			req := httptest.NewRequest(http.MethodPost, "/missing/prune", nil)
			resp := httptest.NewRecorder()

			pruneMissingFiles(maintenance).ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusOK))
			Expect(resp.Body.String()).To(Equal(`{"pruned":true}`))
			Expect(maintenance.pruneCalled).To(BeTrue())
		})

		It("returns forbidden when pruning is not permitted", func() {
			maintenance := &mockMaintenance{pruneErr: rest.ErrPermissionDenied}
			req := httptest.NewRequest(http.MethodPost, "/missing/prune", nil)
			resp := httptest.NewRecorder()

			pruneMissingFiles(maintenance).ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusForbidden))
			Expect(maintenance.pruneCalled).To(BeTrue())
		})

		It("returns an internal server error when pruning fails", func() {
			maintenance := &mockMaintenance{pruneErr: errors.New("db failed")}
			req := httptest.NewRequest(http.MethodPost, "/missing/prune", nil)
			resp := httptest.NewRecorder()

			pruneMissingFiles(maintenance).ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(http.StatusInternalServerError))
			Expect(maintenance.pruneCalled).To(BeTrue())
		})
	})
})
