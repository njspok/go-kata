package big_numbers

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNumberFromInt(t *testing.T) {
	require.Equal(t, Number("123"), NumberFromUint(123))
	require.Equal(t, Number("0"), NumberFromUint(0))
	require.Equal(t, Number("9"), NumberFromUint(9))
}

func Test(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		require.Equal(t, Number("0"), Sum("0", "0"))
		require.Equal(t, Number("2"), Sum("1", "1"))
		require.Equal(t, Number("4"), Sum("2", "2"))
		require.Equal(t, Number("10"), Sum("1", "9"))
		require.Equal(t, Number("18"), Sum("9", "9"))
		require.Equal(t, Number("19"), Sum("10", "9"))
		require.Equal(t, Number("120"), Sum("111", "9"))
		require.Equal(t, Number("353646"), Sum("121323", "232323"))
		require.Equal(t, Number("91002328220491911630239667963"), Sum("63829983432984289347293874", "90938498237058927340892374089"))
	})
	t.Run("replace", func(t *testing.T) {
		rand.Seed(time.Now().Unix())
		aug := uint(rand.Intn(1000000))
		add := uint(rand.Intn(1000000))

		for i := 0; i < 1000000; i++ {
			require.Equalf(t,
				NumberFromUint(aug+add),
				Sum(NumberFromUint(aug), NumberFromUint(add)),
				fmt.Sprintf("%v,%v", aug, add),
			)
		}
	})
}
