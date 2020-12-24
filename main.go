package main

import (
	"errors"
	"fmt"
)

func clockwisePrint(a [][]int) {
	var m, n, k, l int
	/*  k - starting row index
	    m - ending row index
	    l - starting column index
	    n - ending column index
	    i - iterator
	*/

	R := len(a)
	C := len(a[0])

	if R != C {
		err := errors.New("Invalid input. Number of columns and rows must be equal")
		fmt.Println(err)
		return
	}

	m = R
	n = C

	for k < m && l < n {
		/* Print the first row from the remaining rows */
		for i := l; i < n; i++ {
			fmt.Printf("%d ", a[k][i])
		}
		k++

		/* Print the last column from the remaining columns
		 */
		for i := k; i < m; i++ {
			fmt.Printf("%d ", a[i][n-1])
		}
		n--

		/* Print the last row from the remaining rows */
		if k < m {
			for i := n - 1; i >= l; i-- {
				fmt.Printf("%d ", a[m-1][i])
			}
			m--
		}

		/* Print the first column from the remaining columns
		 */
		if l < n {
			for i := m - 1; i >= k; i-- {
				fmt.Printf("%d ", a[i][l])
			}
			l++
		}
	}
}

func main() {

	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	clockwisePrint(a)
}
