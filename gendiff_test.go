package code

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const stylish = "stylish"

func TestGendiff(t *testing.T) {
	cases := []struct {
		name         string
		format       string
		filepath1    string
		filepath2    string
		expectedPath string
	}{
		{"Flat JSON", stylish, "testdata/fixtures/flat_json/file1.json", "testdata/fixtures/flat_json/file2.json", "testdata/fixtures/flat_expected.txt"},
		{"Flat YAML", stylish, "testdata/fixtures/flat_yaml/file1.yaml", "testdata/fixtures/flat_yaml/file2.yml", "testdata/fixtures/flat_expected.txt"},
		{"Nested JSON", stylish, "testdata/fixtures/nested_json/file1.json", "testdata/fixtures/nested_json/file2.json", "testdata/fixtures/nested_expected.txt"},
		{"Nested YAML", stylish, "testdata/fixtures/nested_yaml/file1.yaml", "testdata/fixtures/nested_yaml/file2.yml", "testdata/fixtures/nested_expected.txt"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			expectedBytes, err := os.ReadFile(tc.expectedPath)
			require.NoError(t, err)

			expected := strings.TrimSpace(string(expectedBytes))

			res, err := Gendiff(tc.format, tc.filepath1, tc.filepath2)
			require.NoError(t, err)

			assert.Equal(t, expected, strings.TrimSpace(res))
		})
	}
}

func TestGendiff_FileDoesNotExist(t *testing.T) {
	incorrectPath := "incorrect/path"

	_, err1 := Gendiff("", incorrectPath, "testdata/fixtures/flat_json/file2.json")
	_, err2 := Gendiff("", "testdata/fixtures/flat_json/file1.json", incorrectPath)

	assert.Error(t, err1)
	assert.ErrorIs(t, err1, os.ErrNotExist)

	assert.Error(t, err2)
	assert.ErrorIs(t, err2, os.ErrNotExist)
}

func TestGendiff_UnsupportedFileExtension(t *testing.T) {
	_, err := Gendiff("", "testdata/fixtures/flat_expected.txt", "testdata/fixtures/flat_json/file2.json")
	require.Error(t, err)
	assert.EqualError(t, err, "unsupported file extension: .txt")
}

func TestGendiff_UnsupportedFormat(t *testing.T) {
	_, err := Gendiff("unsupported", "testdata/fixtures/flat_json/file1.json", "testdata/fixtures/flat_json/file2.json")
	require.Error(t, err)
	assert.EqualError(t, err, "unsupported format: unsupported")
}
