package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnifyNewline(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		actual := UnifyNewline("あ\n\nあ")
		expect := "あ\nあ"
		assert.Equal(t, expect, actual)
	})
	t.Run("non-effect", func(t *testing.T) {
		actual := UnifyNewline("あ\nあ")
		expect := "あ\nあ"
		assert.Equal(t, expect, actual)
	})
	t.Run("multiple", func(t *testing.T) {
		actual := UnifyNewline("あ\n\nい\n\nう")
		expect := "あ\nい\nう"
		assert.Equal(t, expect, actual)
	})
}

func TestWrapText(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		actual := WrapText("あいうえおかきくけこ", 3)
		expect := []string{"あいう", "えおか", "きくけ", "こ"}
		assert.Equal(t, expect, actual)
	})
	t.Run("less than point", func(t *testing.T) {
		actual := WrapText("あ", 3)
		expect := []string{"あ"}
		assert.Equal(t, expect, actual)
	})
	t.Run("just size", func(t *testing.T) {
		actual := WrapText("いろはにほへ", 3)
		expect := []string{"いろは", "にほへ"}
		assert.Equal(t, expect, actual)
	})
}

func TestFmtText(t *testing.T) {
	t.Run("standard", func(t *testing.T) {
		actual := FmtText("色は匂へど 散りぬるを 我が世誰そ 常ならむ 有為の奥山 今日越えて 浅き夢見じ 酔ひもせず")
		expect := []string{"色は匂へど 散りぬるを 我が世誰そ 常なら", "む 有為の奥山 今日越えて 浅き夢見じ 酔", "ひもせず"}
		assert.Equal(t, expect, actual)
	})
	t.Run("over", func(t *testing.T) {
		actual := FmtText(strings.Repeat("ブルータスお前もか", 19))
		expect := []string{"ブルータスお前もかブルータスお前もかブルー", "タスお前もかブルータスお前もかブルータスお", "前もかブルータスお前もかブルータスお前もか", "ブルータスお前もかブルータスお前もかブルー", "タスお前もかブルータスお前もかブルータスお", "前もかブルータスお前もかブルータスお前もか", "ブルータスお前もかブルータスお前もかブルー", "タスお前もかブルータスお前もかブルータスお", "前"}
		assert.Equal(t, expect, actual)
	})
	t.Run("over with newlines", func(t *testing.T) {
		actual := FmtText("あ\nい\nう\nえ\nお\nか\nき\nく\nけけけ\nこ\nさ\nし")
		expect := []string{"あ", "い", "う", "え", "お", "か", "き", "く", "け"}
		assert.Equal(t, expect, actual)
	})
}
