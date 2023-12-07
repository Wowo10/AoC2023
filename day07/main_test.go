package main_test

import (
	"testing"

	sut "day07"

	"github.com/stretchr/testify/assert"
)

func TestFive(t *testing.T) {
	c := sut.CheckCategory("22222")

	assert.Equal(t, c, sut.Five)
}

func TestFour(t *testing.T) {
	c := sut.CheckCategory("22221")

	assert.Equal(t, c, sut.Four)
}

func TestFour_2(t *testing.T) {
	c := sut.CheckCategory("22122")

	assert.Equal(t, c, sut.Four)
}

func TestFull(t *testing.T) {
	c := sut.CheckCategory("22211")

	assert.Equal(t, c, sut.Full)
}

func TestFull_2(t *testing.T) {
	c := sut.CheckCategory("21212")

	assert.Equal(t, c, sut.Full)
}

func TestThree(t *testing.T) {
	c := sut.CheckCategory("22213")

	assert.Equal(t, c, sut.Three)
}

func TestTwoPairs(t *testing.T) {
	c := sut.CheckCategory("22113")

	assert.Equal(t, c, sut.TwoPairs)
}

func TestPair(t *testing.T) {
	c := sut.CheckCategory("22143")

	assert.Equal(t, c, sut.Pair)
}

func TestJoker(t *testing.T) {
	c := sut.CheckCategory("2222J")

	assert.Equal(t, c, sut.Five)
}

func TestJoker2(t *testing.T) {
	c := sut.CheckCategory("222JJ")

	assert.Equal(t, c, sut.Five)
}

func TestJoker3(t *testing.T) {
	c := sut.CheckCategory("22JJJ")

	assert.Equal(t, c, sut.Five)
}

func TestJoker4(t *testing.T) {
	c := sut.CheckCategory("2JJJJ")

	assert.Equal(t, c, sut.Five)
}

func TestJoker5(t *testing.T) {
	c := sut.CheckCategory("JJJJJ")

	assert.Equal(t, c, sut.Five)
}

func TestJoker6(t *testing.T) {
	c := sut.CheckCategory("2221J")

	assert.Equal(t, c, sut.Four)
}

func TestJoker7(t *testing.T) {
	c := sut.CheckCategory("2211J")

	assert.Equal(t, c, sut.Full)
}

func TestJoker8(t *testing.T) {
	c := sut.CheckCategory("221JJ")

	assert.Equal(t, c, sut.Four)
}

func TestJoker9(t *testing.T) {
	c := sut.CheckCategory("21JJJ")

	assert.Equal(t, c, sut.Four)
}

func TestNone(t *testing.T) {
	c := sut.CheckCategory("21345")

	assert.Equal(t, c, sut.None)
}

func TestCMP_1(t *testing.T) {
	h1 := sut.CreateHand("21345 1")
	h2 := sut.CreateHand("22222 1")

	assert.Negative(t, sut.CmpHands(h1, h2))
}

func TestCMP_2(t *testing.T) {
	h1 := sut.CreateHand("22222 1")
	h2 := sut.CreateHand("21345 1")

	assert.Positive(t, sut.CmpHands(h1, h2))
}

func TestCMP_3(t *testing.T) {
	h1 := sut.CreateHand("33332 1")
	h2 := sut.CreateHand("2AAAA 1")

	assert.Positive(t, sut.CmpHands(h1, h2))
}

func TestCMP_4(t *testing.T) {
	h1 := sut.CreateHand("77888 1")
	h2 := sut.CreateHand("77788 1")

	assert.Positive(t, sut.CmpHands(h1, h2))
}

func TestCMP_5(t *testing.T) {
	h1 := sut.CreateHand("AAAAA 1")
	h2 := sut.CreateHand("22222 1")

	assert.Positive(t, sut.CmpHands(h1, h2))
}
