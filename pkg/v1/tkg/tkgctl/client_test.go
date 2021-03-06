// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package tkgctl

import (
	"io/ioutil"
	"os"
	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	fakeproviders "github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/fakes/providers"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/providerinterface"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/tkgconfigreaderwriter"
	"github.com/vmware-tanzu/tanzu-framework/pkg/v1/tkg/tkgconfigupdater"
)

var _ = Describe("ensureBoMandProvidersPrerequisite", func() {
	var (
		err                    error
		providerGetter         providerinterface.ProviderInterface
		tkgConfigUpdaterClient tkgconfigupdater.Client
		tkgConfigReaderWriter  tkgconfigreaderwriter.TKGConfigReaderWriter
	)

	BeforeEach(func() {
		testingDir, err = os.MkdirTemp("", "test")
		err = os.MkdirAll(testingDir, 0o700)
		Expect(err).ToNot(HaveOccurred())
		prepareConfiDir(testingDir)
		providerGetter = fakeproviders.FakeProviderGetter()
		tkgConfigReaderWriter, err = tkgconfigreaderwriter.NewReaderWriterFromConfigFile("../fakes/config/config.yaml", "../fakes/config/config.yaml")
		Expect(err).ToNot(HaveOccurred())
		tkgConfigUpdaterClient = tkgconfigupdater.New(testingDir, providerGetter, tkgConfigReaderWriter)
	})

	Context("When two goroutines try to modify the file under configDir", func() {
		It("should not return errors", func() {
			errs := make(chan error, 2)
			defer close(errs)
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				err := ensureBoMandProvidersPrerequisite(testingDir, tkgConfigUpdaterClient, true)
				errs <- err
			}()

			go func() {
				defer wg.Done()
				err := ensureBoMandProvidersPrerequisite(testingDir, tkgConfigUpdaterClient, true)
				errs <- err
			}()
			wg.Wait()
			var err1, err2 error
			err1 = <-errs
			err2 = <-errs

			Expect(err1).ToNot(HaveOccurred())
			Expect(err2).ToNot(HaveOccurred())
		})
	})
})

var _ = Describe("Unit test for New", func() {
	var (
		err       error
		options   Options
		configDir string
	)
	JustBeforeEach(func() {
		configDir, _ = ioutil.TempDir("", "cluster_client_test")
		prepareConfiDir(configDir)
		options = Options{
			ConfigDir:      configDir,
			ProviderGetter: fakeproviders.FakeProviderGetter(),
		}
		_, err = New(options)
	})

	Context("Create tkgctl client with all clients", func() {
		It("should create the tkg client", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	AfterEach(func() {
		os.Remove(configDir)
	})
})
