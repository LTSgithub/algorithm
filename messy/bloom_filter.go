package main

import (
	"fmt"
	"hash/crc32"
	"math"
)

/*
公式一：确定空间
	m = n * ln(p) / ln2 / ln2

	m：bit数组的大小
	n：样本量
	p：预期失误率

	例如：样本量n为100亿，失误率p为万分之一，那么需要的内存大概是16G

公式二：确定hash函数的个数
	k = ln2 * m / n = 0.7 * m / n

	k：hash函数的个数
	m：如果m是小数，向上取整
	n：样本量

公式三：当真实的 m，k 都向上取整之后的失误率
	p = (1 - e(-n * k / m)(幂))(k)（幂）
*/

type Bitmap struct {
	Words  []uint32 `json:"words"`
	Length int      `json:"length"`
}

func (b *Bitmap) Has(num int) bool {
	word, bit := num/32, uint(num%32)
	return word < len(b.Words) && (b.Words[word]&(1<<bit)) != 0
}

func (b *Bitmap) Add(num int) {

	if num > math.MaxUint32 {
		return
	}

	word, bit := num/32, uint(num%32)
	for word >= len(b.Words) {
		b.Words = append(b.Words, 0)
	}

	if b.Words[word]&(1<<bit) == 0 {
		b.Words[word] |= 1 << bit
		b.Length++
	}
}

func (b *Bitmap) Delete(num int) {

	// 如果没有，返回
	if !b.Has(num) {
		return
	}

	if num > math.MaxUint32 {
		return
	}

	word, bit := num/32, uint(num%32)
	for word >= len(b.Words) {
		b.Words = append(b.Words, 0)
	}

	if b.Words[word]&(1<<bit) != 0 {
		b.Words[word] &= ^(1 << bit)
		b.Length--
	}
}

func (b *Bitmap) Len() int {

	return b.Length
}

func (b *Bitmap) NumList() []uint32 {

	buf := make([]uint32, 0)

	for i, v := range b.Words {
		if v == 0 {
			continue
		}
		for j := uint(0); j < 32; j++ {
			if v&(1<<j) != 0 {
				buf = append(buf, uint32(32*i)+uint32(j))
			}
		}
	}

	return buf
}

type BloomFilter struct {
	Bitmap
}

func (m *BloomFilter) hashUint(ss string) int {
	ret := int(crc32.ChecksumIEEE([]byte(ss)))

	if ret >= 0 {
		return ret
	}

	return -ret
}

func (m *BloomFilter) Insert(ss string) {
	ret := m.hashUint(ss)
	m.Add(ret)
}

func (m *BloomFilter) IsExist(ss string) bool {

	return m.Has(m.hashUint(ss))
}

func main() {

	ss := "aaaaaa"

	bl := BloomFilter{}

	bl.Insert(ss)

	fmt.Println(bl.IsExist(ss))

}
