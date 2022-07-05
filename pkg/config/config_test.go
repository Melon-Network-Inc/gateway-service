package config

import (
	"io/ioutil"
	"os"
	"path"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// initDirs makes directories for testing
func initDirs(dir, name string) func() {
	var (
		fullName = name + ".yaml"
		cleanup  = true
		clean    func()
	)

	if dir != "" {
		root, err := ioutil.TempDir("", "")

		clean = func() {
			os.Chdir("..")
			os.RemoveAll(root)
		}

		defer func() {
			if cleanup {
				clean()
			}
		}()

		Expect(err).NotTo(HaveOccurred())

		err = os.Chdir(root)
		Expect(err).NotTo(HaveOccurred())

		err = os.MkdirAll(dir, 0750)
		Expect(err).NotTo(HaveOccurred())
	} else {
		defaultDir := path.Join("../../")
		err := os.Chdir(defaultDir)
		Expect(err).NotTo(HaveOccurred())

		clean = func() {
			if name != "" {
				os.Remove(path.Join(defaultDir, fullName))
			}
		}

		defer func() {
			if cleanup {
				clean()
			}
		}()
	}

	if name != "" {
		err := ioutil.WriteFile(path.Join(dir, fullName), content, 0640)
		Expect(err).NotTo(HaveOccurred())
	}

	cleanup = false
	return clean
}

var _ = Describe("Config", func() {
	AfterEach(func() {
		os.Setenv(`MELON_CONFIG_PATH`, "")
		os.Setenv(`MELON_CONFIG_NAME`, "")
		os.Setenv(`MELON_CONFIG_TYPE`, "")
	})

	It("should not return nil when using default values", func() {
		config := New()
		Expect(config).NotTo(BeNil())
	})

	It("should return Configuration type with all fields filled", func() {
		config := New()
		Expect(config).NotTo(BeNil())

		Expect(config.Redis).NotTo(BeNil())
	})

	It(`should return valid config when MELON_CONFIG_PATH is set and
			config file exists in that path`, func() {
		dir := `//`
		os.Setenv(`MELON_CONFIG_PATH`, dir)
		cleanDirs := initDirs(dir, "")
		defer cleanDirs()
	})

	It("should return nil when config file does not exists", func() {
		os.Setenv(`MELON_CONFIG_NAME`, "improbable")

		config := New()
		Expect(config).To(BeNil())
	})

	It("should return nil when MELON_CONFIG_TYPE is set but is not valid", func() {
		configTypes := []string{"dev", "prod", "testing", "mmm", "."}
		for _, configType := range configTypes {
			os.Setenv(`MELON_CONFIG_TYPE`, configType)

			config := New()
			Expect(config).To(BeNil())
		}
	})
})
