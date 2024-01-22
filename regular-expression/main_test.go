package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEnglishName(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid English name", "John Doe", true},
		{"Invalid English name", "John Doe 123", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, validateEnglishName(tc.input))
		})
	}
}

func TestValidatePhoneNumber(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid phone number", "0123456789", true},
		{"Invalid if phone number exceeded limitation", "01234567890", false},
		{"Invalid if phone number contained special character", "+012345678", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, validatePhoneNumber(tc.input))
		})
	}
}

func TestValidateEmail(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid email", "johndoe@mail.com", true},
		{"Invalid email if not contain @", "johndoemail.com", false},
		{"Invalid email if not contain domain", "johndoe@.com", false},
		{"Invalid email if not contain top level domain", "johndoe@mail", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, validateEmail(tc.input))
		})
	}
}

func TestExtractStreetName(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"Valid street name", "12, Ratchada Road, Bangkok, Thailand 12345", "Ratchada"},
		{"Invalid street name if not contain road suffix", "12, Ratchada, Bangkok, Thailand 12345", ""},
		{"Invalid street name if not contain comma", "12 Ratchada Road Bangkok Thailand 12345", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, extractStreetName(tc.input))
		})
	}
}
