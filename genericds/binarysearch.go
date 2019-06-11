package genericds;

// The `comparator` is used to compare the `toFind` with the elements in `sortedArray`
// Return `index` of the `toFind` in `sortedArray` if found
// Returns -1 if not found 
func BinarySearch ( 
       sortedArray []interface{}, 
       toFind interface{}, 
       comparator func ( interface{}, interface{}) (int, error) ,
    )  (int, error) {
    start, end  := 0, len(sortedArray)
    for start <= end {
       mid := (end - start  ) / 2 + start 
       comparisionResult, typeError := comparator(toFind, sortedArray[mid])
       if typeError != nil {
         return -1, typeError
       }
       switch {
       case comparisionResult < 0 :
         end = mid-1
       case comparisionResult > 0 :
         start = mid+1
       default:
         return mid, nil  
       }
    }
    return -1, nil 
}