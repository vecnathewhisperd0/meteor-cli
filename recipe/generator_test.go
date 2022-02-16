package recipe_test

import (
	"os"
	"path"
	"testing"

	"github.com/odpf/meteor/recipe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFromTemplate(t *testing.T) {
	t.Run("should output recipe files using template to output directory", func(t *testing.T) {
		templatePath := "./testdata/generator/template.yaml"
		outputDir := "./testdata/generator/temp"

		t.Run("when recipe has a name", func(t *testing.T) {
			data := []recipe.TemplateData{
				{
					FileName: "recipe-one",
					Data: map[string]interface{}{
						"name":   "recipe-1",
						"broker": "main-broker.com:9092",
						"owner":  "john@example.com",
					},
				},
				{
					FileName: "recipe-two",
					Data: map[string]interface{}{
						"name":   "recipe-2",
						"broker": "secondary-broker.com:9092",
						"owner":  "jane@example.com",
					},
				},
			}

			cleanDir(t, outputDir)
			defer cleanDir(t, outputDir)

			err := recipe.FromTemplate(recipe.TemplateConfig{
				TemplateFilePath: templatePath,
				OutputDirPath:    outputDir,
				Data:             data,
			})
			require.NoError(t, err)

			assertRecipeFile(t,
				"./testdata/generator/expected.yaml",
				path.Join(outputDir, data[0].FileName+".yaml"),
			)
			assertRecipeFile(t,
				"./testdata/generator/expected-2.yaml",
				path.Join(outputDir, data[1].FileName+".yaml"),
			)
		})

		t.Run("when recipe does not have a name", func(t *testing.T) {
			data := []recipe.TemplateData{
				{
					FileName: "recipe-three",
					Data: map[string]interface{}{
						"broker": "main-broker.com:9092",
						"owner":  "john@example.com",
					},
				},
			}

			cleanDir(t, outputDir)
			defer cleanDir(t, outputDir)

			err := recipe.FromTemplate(recipe.TemplateConfig{
				TemplateFilePath: templatePath,
				OutputDirPath:    outputDir,
				Data:             data,
			})
			require.NoError(t, err)

			assertRecipeFile(t,
				"./testdata/generator/expected-3.yaml",
				path.Join(outputDir, data[0].FileName+".yaml"),
			)
		})

	})
}

func cleanDir(t *testing.T, dirPath string) {
	err := os.RemoveAll(dirPath)
	require.NoError(t, err)
}

func assertRecipeFile(t *testing.T, expectedPath, actualPath string) {
	expectedBytes, err := os.ReadFile(expectedPath)
	require.NoError(t, err)
	actualBytes, err := os.ReadFile(actualPath)
	require.NoError(t, err)
	assert.Equal(t, string(expectedBytes), string(actualBytes))
}
