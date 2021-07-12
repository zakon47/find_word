package find_word_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/zakon47/find_word"
	"testing"
)

var (
	list1 = []string{"НАШЕ", "ВАШЕ", "AUTO", "SEX", "47"}
	list2 = []string{"НАше", "ваШЕ", "AutO", "Sex", "47"}
	list3 = []string{"555", ""}
)

func TestHasPrefix(t *testing.T) {
	data := []struct {
		testName string
		str      string
		list     []string
		out1     string
		out2     string
		err      error
	}{
		{testName: "test0", str: "", list: list1, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test1", str: "СЛОВОНАШЕ", list: list1, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test2", str: "НАШЕСЛОВО", list: list1, out1: "НАШЕ", out2: "СЛОВО", err: nil},
		{testName: "test3", str: "НашеСЛОВО", list: list1, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test4", str: "НашеСЛОВО", list: list3, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test5", str: "47НашеСЛОВО", list: list1, out1: "47", out2: "НашеСЛОВО", err: nil},
		{testName: "test6", str: "47", list: list1, out1: "47", out2: "", err: nil},
	}
	for _, test := range data {
		t.Run(test.testName, func(t *testing.T) {
			w1, w2, err := find_word.HasPrefix(test.str, test.list)
			assert.Equal(t, test.out1, w1)
			assert.Equal(t, test.out2, w2)
			assert.Equal(t, test.err, err)
		})
	}
}
func TestHasPrefixWithoutCase(t *testing.T) {
	data := []struct {
		testName string
		str      string
		list     []string
		out1     string
		out2     string
		err      error
	}{
		{testName: "test0", str: "", list: list1, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test1", str: "СЛОВОНАШЕ", list: list2, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test2", str: "НАШЕСЛОВО", list: list2, out1: "НАШЕ", out2: "СЛОВО", err: nil},
		{testName: "test3", str: "НашеСЛОВО", list: list2, out1: "Наше", out2: "СЛОВО", err: nil},
		{testName: "test4", str: "НашеСЛОВО", list: list3, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test5", str: "47НашеСЛОВО", list: list2, out1: "47", out2: "НашеСЛОВО", err: nil},
		{testName: "test6", str: "47", list: list1, out1: "47", out2: "", err: nil},
	}
	for _, test := range data {
		t.Run(test.testName, func(t *testing.T) {
			w1, w2, err := find_word.HasPrefixWithoutCase(test.str, test.list)
			assert.Equal(t, test.out1, w1)
			assert.Equal(t, test.out2, w2)
			assert.Equal(t, test.err, err)
		})
	}
}
func TestHasSuffix(t *testing.T) {
	data := []struct {
		testName string
		str      string
		list     []string
		out1     string
		out2     string
		err      error
	}{
		{testName: "test0", str: "", list: list1, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test1", str: "СЛОВОНАШЕ", list: list1, out1: "НАШЕ", out2: "СЛОВО", err: nil},
		{testName: "test2", str: "СЛОВОНаше", list: list1, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test3", str: "НАШЕСЛОВО", list: list1, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test4", str: "НашеСЛОВО", list: list3, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test5", str: "НашеСЛОВО47", list: list1, out1: "47", out2: "НашеСЛОВО", err: nil},
		{testName: "test6", str: "47", list: list1, out1: "47", out2: "", err: nil},
	}
	for _, test := range data {
		t.Run(test.testName, func(t *testing.T) {
			w1, w2, err := find_word.HasSuffix(test.str, test.list)
			assert.Equal(t, test.out1, w1)
			assert.Equal(t, test.out2, w2)
			assert.Equal(t, test.err, err)
		})
	}
}
func TestHasSuffixWithoutCase(t *testing.T) {
	data := []struct {
		testName string
		str      string
		list     []string
		out1     string
		out2     string
		err      error
	}{
		{testName: "test0", str: "", list: list1, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test1", str: "СЛОВОНАШЕ", list: list2, out1: "НАШЕ", out2: "СЛОВО", err: nil},
		{testName: "test2", str: "НАШЕСЛОВО", list: list2, out1: "", out2: "", err: find_word.ErrFindWord},
		{testName: "test3", str: "СЛОВОНаше", list: list2, out1: "Наше", out2: "СЛОВО", err: nil},
		{testName: "test4", str: "x47", list: list2, out1: "47", out2: "x", err: nil},
		{testName: "test5", str: "НашеСЛОВО47", list: list1, out1: "47", out2: "НашеСЛОВО", err: nil},
		{testName: "test6", str: "47", list: list1, out1: "47", out2: "", err: nil},
	}
	for _, test := range data {
		t.Run(test.testName, func(t *testing.T) {
			w1, w2, err := find_word.HasSuffixWithoutCase(test.str, test.list)
			assert.Equal(t, test.out1, w1)
			assert.Equal(t, test.out2, w2)
			assert.Equal(t, test.err, err)
		})
	}
}
