package calc

import (
	"errors"
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

// アサーションサンプル
func TestSum6(t *testing.T) {
	// 2つのオブジェクトが等価であること
	assert.Equal(t, 3, 3)

	// 2つのオブジェクトが等価ではないこと
	assert.NotEqual(t, 3, 4)

	// オブジェクトがゼロ値、もしくは長さが0のスライスであること
	assert.Empty(t, 0)
	assert.Empty(t, "")
	assert.Empty(t, []int{})

	// エラーであること
	err := errors.New("error")
	assert.Error(t, err)

	// エラーではないこと(nilであること)
	err = nil
	assert.NoError(t, err)

	// 1番目の要素が、2番目の要素以上であること
	assert.GreaterOrEqual(t, 2, 1)

	// オブジェクトが特定の長さを持っていること
	assert.Len(t, []int{1, 2}, 2)

	// nilであること
	err = nil
	assert.Nil(t, err)

	// nilではないこと
	v := 0
	assert.NotNil(t, v)

	// ゼロ値であること
	v = 0
	assert.Zero(t, v)
}
