package utility_test

import (
	"encoding/hex"
	"testing"

	"github.com/CityOfZion/neo-go-sdk/utility"
	"github.com/stretchr/testify/assert"
)

func TestBase58(t *testing.T) {
	t.Run(".Decode()", func(t *testing.T) {
		t.Run("HappyCase", func(t *testing.T) {
			testCases := []struct {
				description string
				in          string
				out         string
			}{
				{
					description: "Empty",
					in:          "",
					out:         "",
				},
				{
					description: "SingleDecimal",
					in:          "1",
					out:         "00",
				},
				{
					description: "FullString",
					in:          "16UwLL9Risc3QfPqBUvKofHmBQ7wMtjvM",
					out:         "00010966776006953d5567439e5e39f86a0d273beed61967f6",
				},
			}

			base58 := utility.NewBase58()

			for _, testCase := range testCases {
				t.Run(testCase.description, func(t *testing.T) {
					result, err := base58.Decode(testCase.in)
					assert.Nil(t, err)
					assert.Equal(t, testCase.out, hex.EncodeToString(result))
				})
			}
		})

		t.Run("SadCase", func(t *testing.T) {
			testCases := []struct {
				description string
				in          string
			}{
				{
					description: "InvalidChar",
					in:          "0",
				},
			}

			base58 := utility.NewBase58()

			for _, testCase := range testCases {
				t.Run(testCase.description, func(t *testing.T) {
					_, err := base58.Decode(testCase.in)
					assert.NotNil(t, err)
				})
			}
		})
	})
}
