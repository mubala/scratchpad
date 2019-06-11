package genericds;

import ( 
  "testing"
  "errors"
  "fmt"
)


func  IntComparator (x, y interface{}) (int , error){
   a,a_is_int := x.(int)
   b, b_is_int := y.(int)
   if a_is_int && b_is_int {
     return a - b, nil
   } else {
     return -1, errors.New("Wrong type used ")
   }
}


func Test_BinarySearch ( t *testing.T ) {
  comparator := IntComparator
  sortedIntArray := []interface {}{ 1, 3, 5, 7, 9, 11, 13}
  test_plan := []struct {
      name string 
      list []interface{}
      toFind  interface{}
      location int 
      errorFound error
  } {
    {
      name: "Find element in second half the array  ",
      toFind : 9,
      list: sortedIntArray,
      location : 4,
      errorFound: nil,
    },
    {
      name: "Find first element of the array  ",
      toFind : 1,
      list: sortedIntArray,
      location : 0,
      errorFound: nil,
    }, 
    {
      name: "Find mid element of the array  ",
      toFind : 7,
      list: sortedIntArray,
      location : 3,
      errorFound: nil,
    }, 
    {
      name: "Find last element of the array  ",
      toFind : 13,
      list: sortedIntArray,
      location : 6,
      errorFound: nil,
    }, 
    {
      name: "Find `non present` element in the array ",
      toFind : 2,
      list: sortedIntArray,
      location : -1,
      errorFound: nil,
    },
  { 
    name: "Find incompatible value ( type ) in the array",
    toFind : "Test",
    list: sortedIntArray,
    location : -1,
    errorFound: errors.New("Wrong type used "),
  }, 

  }

  for _,test := range test_plan {
    fmt.Println ( "Test " , test.name)
    t.Run(test.name, func ( t *testing.T){
       t.Log(  " Testing ", test.name )
       locationFound, err1 := BinarySearch( test.list, test.toFind, comparator)
       if  test.errorFound != nil && err1 == nil  {
          t.Errorf ( " Test %v failed  - expected error nil , %v", test.name,   test.errorFound ) 
       } else if   test.location != locationFound {
            t.Errorf(" `%v`   %v not found in the list  ( %v != %v)", test.name, test.toFind, locationFound, test.location)
          }
    })
  }
    
}