package generator

import (
	"crypto/rand"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGeneratePassword_ValidLength(t *testing.T) {
	const MaxPasswordLengthToTest = 40

	// --- TEST VALID LENGTHS ---
	for length := MinPasswordLength; length <= MaxPasswordLengthToTest; length++ {
		password, entropy, err := GeneratePassword(length)
		require.NoError(t, err)
		assert.NotEmpty(t, password)
		assert.Len(t, password, length)
		assert.Greater(t, entropy, 0.0)
	}
}

func TestGeneratePassword_InvalidLength(t *testing.T) {
	// --- TEST INVALID LENGTHS ---
	password, entropy, err := GeneratePassword(MinPasswordLength - 1)
	require.Error(t, err)
	assert.Empty(t, password)
	assert.Equal(t, entropy, 0.0)
}

func TestGeneratePassword_EmptyPassword(t *testing.T) {
	// --- TEST EMPTY LENGTH ---
	emptyLength := 0
	password, entropy, err := GeneratePassword(emptyLength)
	require.Error(t, err)
	assert.Empty(t, password)
	assert.Equal(t, entropy, 0.0)
}

func TestGeneratePassword_ErrorOnRandIntFailure(t *testing.T) {
	// --- TEST ERROR ON RAND INT FAILURE ---
	originalRand := rand.Reader
	defer func() { rand.Reader = originalRand }()
	rand.Reader = &mockRandReader{}

	password, entropy, err := GeneratePassword(MinPasswordLength)
	require.Error(t, err)
	assert.Empty(t, password)
	assert.Equal(t, entropy, 0.0)
}

type mockRandReader struct{}

func (m *mockRandReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("mock error")
}

func TestCalculateEntropy(t *testing.T) {
	// --- TEST CALCULATE ENTROPY ---
	charPool := lowercase + uppercase + numbers + specialChars
	charPoolLen := float64(len(charPool))

	tests := []struct {
		password string
		expected float64
	}{
		{"", 0.0},                                       // Empty password
		{"a", math.Log2(charPoolLen)},                   // One character, entropy of a single character from a pool of charPoolLen
		{"ab", math.Log2(charPoolLen) * 2},              // Two characters
		{"abcdEFGH1234!@", math.Log2(charPoolLen) * 14}, // Complex password
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			entropy := calculateEntropy(tt.password, len(charPool))
			assert.InDelta(t, tt.expected, entropy, 0.01)
		})
	}
}

func TestPasswordUniqueness(t *testing.T) {
	// --- TEST PASSWORD UNIQUENESS ---
	passwords := make(map[string]struct{})
	numPasswords := 100
	for i := 0; i < numPasswords; i++ {
		password, _, err := GeneratePassword(MinPasswordLength)
		require.NoError(t, err)
		_, exists := passwords[password]
		assert.False(t, exists, "Password is not unique")
		passwords[password] = struct{}{}
	}
}

func TestGeneratePassword_CharacterPool(t *testing.T) {
	// --- TEST CHARACTER POOL ---
	length := MinPasswordLength
	password, _, err := GeneratePassword(length)
	require.NoError(t, err)
	assert.Len(t, password, length)

	// --- ENSURE PASSWORD CONTAINS ONLY ALLOWED CHARACTERS ---
	charPool := lowercase + uppercase + numbers + specialChars
	for _, char := range password {
		assert.Contains(t, charPool, string(char))
	}
}
