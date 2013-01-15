package rosalind

import "fmt"

type consM struct {
	A      []int
	C      []int
	G      []int
	T      []int
	length int
	cap    int
}

func NewConsM(length int) *consM {
	if length < 1 {
		panic("Size of consM needs to be >0")
	}
	c := &consM{
		A:      make([]int, length, length),
		C:      make([]int, length, length),
		G:      make([]int, length, length),
		T:      make([]int, length, length),
		length: length,
		cap:    length,
	}
	return c
}

func (c *consM) extend(l int) {
	newA := make([]int, 2*c.cap, 2*c.cap)
	newC := make([]int, 2*c.cap, 2*c.cap)
	newG := make([]int, 2*c.cap, 2*c.cap)
	newT := make([]int, 2*c.cap, 2*c.cap)
	copy(c.A, newA)
	copy(c.C, newC)
	copy(c.G, newG)
	copy(c.T, newT)
	c.A = newA
	c.C = newC
	c.G = newG
	c.T = newT
	c.length = l
	c.cap = 2 * c.cap
}

func (c *consM) AddString(s string) {
	for len(s) > c.cap {
		c.extend(len(s))
	}
	for i, v := range s {
		switch {
		case v == 'A' || v == 'a':
			c.A[i]++
		case v == 'C' || v == 'c':
			c.C[i]++
		case v == 'G' || v == 'g':
			c.G[i]++
		case v == 'T' || v == 't':
			c.T[i]++
		}
	}
}

func (c *consM) AddBytes(b []byte) {
	for len(b) > c.cap {
		c.extend(len(b))
	}
	for i, v := range b {
		switch {
		case v == 'A' || v == 'a':
			c.A[i]++
		case v == 'C' || v == 'c':
			c.C[i]++
		case v == 'G' || v == 'g':
			c.G[i]++
		case v == 'T' || v == 't':
			c.T[i]++
		}
	}
}

func (c *consM) max(i int) byte {
	var m int
	var l byte
	a := [4]int{c.A[i], c.C[i], c.G[i], c.T[i]}
	ls := [4]byte{'A', 'C', 'G', 'T'}
	for i, v := range a {
		if v > m {
			m = v
			l = ls[i]
		}
	}
	return l
}

func (c *consM) Consensus() []byte {
	output := make([]byte, c.length, c.length)
	for i := 0; i < c.length; i++ {
		output[i] = c.max(i)
	}
	return output
}

func (c *consM) _s(ia []int) string {
	var s string
	for _, v := range ia[:c.length] {
		s += fmt.Sprintf("%d ", v)
	}

	return s

}
func (c *consM) String() string {
	return fmt.Sprintf("A: %s\nC: %s\nG: %s\nT: %s", c._s(c.A), c._s(c.C), c._s(c.G), c._s(c.T))
}
