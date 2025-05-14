package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var i, n, j, idx, compalier int
	fmt.Scan(&n)
	for i  =0; i< n; i++{
		fmt.Scan(&A[i])
	}
	for i =0; i < n; i++{
		idx = i
		for j =i +1; j < n; j++{
			if A[i] > A[j]{
				idx = j
			}
		}
		compalier = A[idx]
		A[idx] = A[i]
		A[i] = compalier
	}
	for i =0; i < n; i++{
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}