package main

func getMax(arr []int, l int, r int) int {

	if l == r {
		return arr[l]
	}

	mid := (l + r) / 2
	x := getMax(arr, l, mid)
	y := getMax(arr, mid+1, r)
	if x > y {
		return x
	}

	return y
}

func main() {

	//_, ok := aa[9]
	//fmt.Println(ok)
}

var aa = map[int]bool{
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	//99999:true,
	182: true,
	961: true,
	217: true,
	9:   true,
}