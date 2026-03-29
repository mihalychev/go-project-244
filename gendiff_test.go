package code

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	json    = "json"
	plain   = "plain"
	stylish = "stylish"

	flatJsonFile1   = "testdata/fixtures/flat_json/file1.json"
	flatJsonFile2   = "testdata/fixtures/flat_json/file2.json"
	flatYamlFile1   = "testdata/fixtures/flat_yaml/file1.yaml"
	flatYamlFile2   = "testdata/fixtures/flat_yaml/file2.yml"
	nestedJsonFile1 = "testdata/fixtures/nested_json/file1.json"
	nestedJsonFile2 = "testdata/fixtures/nested_json/file2.json"
	nestedYamlFile1 = "testdata/fixtures/nested_yaml/file1.yaml"
	nestedYamlFile2 = "testdata/fixtures/nested_yaml/file2.yml"

	flatStylishExpected   = "testdata/fixtures/expected/stylish/flat_expected.txt"
	nestedStylishExpected = "testdata/fixtures/expected/stylish/nested_expected.txt"
	nestedPlainExpected   = "testdata/fixtures/expected/plain/nested_expected.txt"
	nestedJsonExpected    = "testdata/fixtures/expected/json/nested_expected.json"
)

func TestGenDiff(t *testing.T) {
	cases := []struct {
		name         string
		format       string
		filepath1    string
		filepath2    string
		expectedPath string
	}{
		{"Stylish Flat JSON", stylish, flatJsonFile1, flatJsonFile2, flatStylishExpected},
		{"Stylish Flat YAML", stylish, flatYamlFile1, flatYamlFile2, flatStylishExpected},
		{"Stylish Nested JSON", stylish, nestedJsonFile1, nestedJsonFile2, nestedStylishExpected},
		{"Stylish Nested YAML", stylish, nestedYamlFile1, nestedYamlFile2, nestedStylishExpected},

		{"Plain Nested JSON", plain, nestedJsonFile1, nestedJsonFile2, nestedPlainExpected},
		{"Plain Nested YAML", plain, nestedYamlFile1, nestedYamlFile2, nestedPlainExpected},

		{"JSON Nested JSON", json, nestedJsonFile1, nestedJsonFile2, nestedJsonExpected},
		{"JSON Nested YAML", json, nestedYamlFile1, nestedYamlFile2, nestedJsonExpected},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			expectedBytes, err := os.ReadFile(tc.expectedPath)
			require.NoError(t, err)

			expected := strings.TrimSpace(string(expectedBytes))

			res, err := GenDiff(tc.filepath1, tc.filepath2, tc.format)
			require.NoError(t, err)

			assert.Equal(t, expected, strings.TrimSpace(res))
		})
	}
}

func TestGenDiff_FileDoesNotExist(t *testing.T) {
	incorrectPath := "incorrect/path"

	_, err1 := GenDiff(incorrectPath, flatJsonFile2, "")
	_, err2 := GenDiff(flatJsonFile1, incorrectPath, "")

	assert.Error(t, err1)
	assert.ErrorIs(t, err1, os.ErrNotExist)

	assert.Error(t, err2)
	assert.ErrorIs(t, err2, os.ErrNotExist)
}

func TestGenDiff_UnsupportedFileExtension(t *testing.T) {
	_, err := GenDiff(flatStylishExpected, flatJsonFile2, "")
	require.Error(t, err)
	assert.EqualError(t, err, "unsupported file extension: .txt")
}

func TestGenDiff_UnsupportedFormat(t *testing.T) {
	_, err := GenDiff(flatJsonFile1, flatJsonFile2, "unsupported")
	require.Error(t, err)
	assert.EqualError(t, err, "unsupported format: unsupported")
}
