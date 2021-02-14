package server_test

import (
	"io/ioutil"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Weiyuan-Lane/example-test-go/internal/server"
	mock_queryparams "github.com/Weiyuan-Lane/example-test-go/test/server/utils/queryparams"
)

var _ = Describe("Server", func() {
	Describe(".RenderPageParamsResponse", func() {
		var (
			mockCtrl             *gomock.Controller
			mockPageParamsEntity *mock_queryparams.MockPageParams
		)

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockPageParamsEntity = mock_queryparams.NewMockPageParams(mockCtrl)
		})

		AfterEach(func() {
			mockCtrl.Finish()
		})

		It("should render the page values in response", func() {
			recorder := httptest.NewRecorder()

			gomock.InOrder(
				mockPageParamsEntity.
					EXPECT().
					Page().
					Times(1).
					Return(7, nil),
				mockPageParamsEntity.
					EXPECT().
					PerPage().
					Times(1).
					Return(19, nil),
			)

			server.RenderPageParamsResponse(recorder, mockPageParamsEntity)
			response := recorder.Result()
			defer response.Body.Close()

			body, _ := ioutil.ReadAll(response.Body)
			strBody := string(body)

			{
				Expect(strBody).To(Equal("page: 7\nperPage: 19\n"))
			}
		})
	})
})
