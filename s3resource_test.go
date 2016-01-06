package s3resource_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/concourse/s3-resource"
	"io/ioutil"

	"testing"
)

func TestS3Resource(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "S3Resource Suite")
}

var _ = Describe("S3Resource", func() {

	Describe("S3client", func() {
		Context("when AWS keys are set", func() {
			It("should not fail", func() {
				_, err := s3resource.NewS3Client(
					ioutil.Discard,
					"AKAIABCD123",
					"OIHGSD%768^SD&S(S^DASD",
					"eu-west-1",
					"",
					"",
				)
				Ω(err).ShouldNot(HaveOccurred())
			})
		})

		Context("when AWS keys are set and credentialsSource is profile", func() {
			It("should fail", func() {
				_, err := s3resource.NewS3Client(
					ioutil.Discard,
					"AKAIABCD123",
					"OIHGSD%768^SD&S(S^DASD",
					"eu-west-1",
					"",
					"profile",
				)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("when an unknown credentails_source is set", func() {
			It("should fail", func() {
				_, err := s3resource.NewS3Client(
					ioutil.Discard,
					"AKAIABCD123",
					"OIHGSD%768^SD&S(S^DASD",
					"eu-west-1",
					"",
					"hello world",
				)
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("when credentials_source is set to static and access_key_id"+
			"or secret_access_key are omitted", func() {
			It("should fail", func() {
				_, err := s3resource.NewS3Client(
					ioutil.Discard,
					"",
					"OIHGSD%768^SD&S(S^DASD",
					"eu-west-1",
					"",
					"static",
				)
				Ω(err).Should(HaveOccurred())
			})
		})

	})
})
