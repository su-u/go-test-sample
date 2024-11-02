package calc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	actual := Sum(1, 2)
	expected := 3

	assert.Equal(t, expected, actual)
}

func TestSum2(t *testing.T) {
	t.Run("正の数値", func(t *testing.T) {
		actual := Sum(1, 2)
		expected := 3

		assert.Equal(t, expected, actual)
	})
	t.Run("負の数値", func(t *testing.T) {
		actual := Sum(-1, -2)
		expected := -3

		assert.Equal(t, expected, actual)
	})
}

// 並列化していない場合
func TestSum3(t *testing.T) {
	t.Run("正の数値", func(t *testing.T) {
		actual := Sum(1, 2)
		expected := 3

		time.Sleep(1 * time.Second)

		assert.Equal(t, expected, actual)
	})
	t.Run("負の数値", func(t *testing.T) {
		actual := Sum(-1, -2)
		expected := -3

		time.Sleep(1 * time.Second)

		assert.Equal(t, expected, actual)
	})
}

// 並列化してテストを実行
func TestSum4(t *testing.T) {
	t.Run("正の数値", func(t *testing.T) {
		t.Parallel()
		actual := Sum(1, 2)
		expected := 3

		time.Sleep(1 * time.Second)

		assert.Equal(t, expected, actual)
	})
	t.Run("負の数値", func(t *testing.T) {
		t.Parallel()
		actual := Sum(-1, -2)
		expected := -3

		time.Sleep(1 * time.Second)

		assert.Equal(t, expected, actual)
	})
}

func TestSum5(t *testing.T) {
	actual := Sum(1, 2)
	expected := 4 // 本来は3である

	assert.Equal(t, expected, actual)
}
