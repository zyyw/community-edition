// Copyright 2021 VMware Tanzu Community Edition contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package harbor_test

import (
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/vmware-tanzu/community-edition/addons/packages/test/pkg/repo"
	"github.com/vmware-tanzu/community-edition/addons/packages/test/pkg/ytt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Harbor Ytt Templates", func() {
	var (
		values string
		output string
		err    error

		configDir = filepath.Join(repo.RootDir(), "addons/packages/harbor/2.4.0/bundle/config")

		ValuesFromFile = func(filename string) string {
			data, err := ioutil.ReadFile(filepath.Join(repo.RootDir(), "addons/packages/harbor/2.4.0/test/unittest/fixtures/values", filename))
			Expect(err).NotTo(HaveOccurred())

			return string(data)
		}

		ExpectOutputEqualToFile = func(filename string) {
			data, err := ioutil.ReadFile(filepath.Join(repo.RootDir(), "addons/packages/harbor/2.4.0/test/unittest/fixtures/expected", filename))
			Expect(err).NotTo(HaveOccurred())

			Expect(output).To(BeEquivalentTo(string(data)))
		}
	)

	BeforeEach(func() {
		values = ""
	})

	JustBeforeEach(func() {
		var filePaths []string

		for _, p := range []string{"upstream/**/*.yaml", "overlay/*.yaml", "*.yaml", "*.star"} {
			matches, err := filepath.Glob(filepath.Join(configDir, p))
			Expect(err).NotTo(HaveOccurred())
			filePaths = append(filePaths, matches...)
		}

		filePaths = append(filePaths,
			filepath.Join(repo.RootDir(), "addons/packages/harbor/2.4.0/test/unittest/fixtures/values/default.yaml"),
		)
		output, err = ytt.RenderYTTTemplate(ytt.CommandOptions{}, filePaths, strings.NewReader(values))
	})

	Context("default", func() {
		BeforeEach(func() {
			values = ""
		})

		It("renders with a default configuration", func() {
			Expect(err).NotTo(HaveOccurred())
			ExpectOutputEqualToFile("default.yaml")
		})
	})

	Context("existing pvc for registry", func() {
		BeforeEach(func() {
			values = ValuesFromFile("registry-existing-pvc.yaml")
		})

		It("renders with a existing pvc for registry configuration", func() {
			Expect(err).NotTo(HaveOccurred())
			ExpectOutputEqualToFile("registry-existing-pvc.yaml")
		})
	})

	Context("azure storage for registry", func() {
		BeforeEach(func() {
			values = ValuesFromFile("registry-azure-storage.yaml")
		})

		It("renders with a azure storage configuration", func() {
			Expect(err).NotTo(HaveOccurred())
			ExpectOutputEqualToFile("registry-azure-storage.yaml")
		})
	})

	Context("s3 storage for registry", func() {
		BeforeEach(func() {
			values = ValuesFromFile("registry-s3-storage.yaml")
		})

		It("renders with a s3 storage configuration", func() {
			Expect(err).NotTo(HaveOccurred())
			ExpectOutputEqualToFile("registry-s3-storage.yaml")
		})
	})

	Context("configuring tlsCertificateSecretName", func() {
		BeforeEach(func() {
			values = ValuesFromFile("tls-certificate-secret-name.yaml")
		})

		It("renders with a tlsCertificateSecretName configuration", func() {
			Expect(err).NotTo(HaveOccurred())
			ExpectOutputEqualToFile("tls-certificate-secret-name.yaml")
		})
	})

	Context("gcs storage for registry", func() {
		BeforeEach(func() {
			values = ValuesFromFile("registry-gcs-storage.yaml")
		})

		It("renders with a gcs storage configuration", func() {
			Expect(err).NotTo(HaveOccurred())
			ExpectOutputEqualToFile("registry-gcs-storage.yaml")
		})
	})

	Context("configuring timeoutPolicy for HTTPProxy", func() {
		BeforeEach(func() {
			values = ValuesFromFile("httpproxy-timeout.yaml")
		})

		It("renders with a HTTPProxy timeoutPolicy configuration", func() {
			Expect(err).NotTo(HaveOccurred())
			ExpectOutputEqualToFile("httpproxy-timeout.yaml")
		})
	})
})
