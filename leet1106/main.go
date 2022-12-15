package main

func main() {
	// parseBoolExpr("&(t,f)")
}

// // !33 |124 &38 ,44 (40 )41 f102 t116
// func parseBoolExpr(expression string) bool {
// 	arr1 := []int{}
// 	arr2 := []int{}
// 	res := false
// 	for i := 0; i < len(expression); i++ {
// 		temp := int(expression[i])
// 		if temp == 33 || temp == 124 || temp == 38 {
// 			arr1 = append(arr1, temp)
// 		}
// 		if temp == 102 || temp == 116 || temp == 40 {
// 			arr2 = append(arr2, temp)
// 		}
// 		if temp == 41 {
// 			fu := arr1[len(arr1)-1]
// 			if fu == 124 {
// 				te := false
// 				t := arr2[len(arr2)-1]
// 				for t != 40 {
// 					te |= (t == 116)
// 					t = arr2[len(arr1)-1]
// 					arr2 = arr2[:len(arr2)-1]
// 				}
// 				if te {
// 					arr2 = append(arr2, 116)
// 				} else {
// 					arr2 = append(arr2, 102)
// 				}
// 			}
// 		}
// 	}
// 	return res
// }
