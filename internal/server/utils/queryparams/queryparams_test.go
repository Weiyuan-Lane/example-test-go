package queryparams_test

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Weiyuan-Lane/example-test-go/internal/server/utils/queryparams"
)

var _ = Describe("Queryparams", func() {
	Describe("#Page", func() {
		It("should return the default page value if not in query params", func() {
			params := url.Values{}
			queryParams := queryparams.New(params)

			val, err := queryParams.Page()

			{
				Expect(val).To(Equal(0))
				Expect(err).To(BeNil())
			}
		})

		It("should return an error if the page value is not a number", func() {
			params := url.Values{
				"page": []string{"not-a-number"},
			}
			queryParams := queryparams.New(params)

			val, err := queryParams.Page()

			{
				Expect(val).To(Equal(0))
				Expect(err).ToNot(BeNil())
			}
		})

		It("should return the converted page value", func() {
			params := url.Values{
				"page": []string{"17"},
			}
			queryParams := queryparams.New(params)

			val, err := queryParams.Page()

			{
				Expect(val).To(Equal(17))
				Expect(err).To(BeNil())
			}
		})
	})

	Describe("#PerPage", func() {
		It("should return the default page value if not in query params", func() {
			params := url.Values{}
			queryParams := queryparams.New(params)
			val, err := queryParams.PerPage()

			{
				Expect(val).To(Equal(10))
				Expect(err).To(BeNil())
			}
		})

		It("should return an error if the page value is not a number", func() {
			params := url.Values{
				"per_page": []string{"not-a-number"},
			}
			queryParams := queryparams.New(params)
			val, err := queryParams.PerPage()

			{
				Expect(val).To(Equal(0))
				Expect(err).ToNot(BeNil())
			}
		})

		It("should return the converted page value", func() {
			params := url.Values{
				"per_page": []string{"17"},
			}
			queryParams := queryparams.New(params)
			val, err := queryParams.PerPage()

			{
				Expect(val).To(Equal(17))
				Expect(err).To(BeNil())
			}
		})
	})
})
