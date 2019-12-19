# Purpose
This repo aims to provide all the functionality related to algorithms in go. Sorting is the first step in the process of making something vast. Currently, WIP. Contributors are most welcome.  
fun fact: gorithims stands for Go + Algorithms

# Installation
```
go get -d github.com/raunakjodhawat/gorithims
```
   
# Sorting implementation
Implements following sorting functions
1. Insertion Sort => insertion()
2. Selection Sort => selection()
3. Merge Sort => merge()
4. Heap Sort => heap()
5. Quick Sort => quick()
6. Bubble Sort => bubble()
7. Bucket Sort => bucket()
8. Radix Sort => radix()

# Functions
## Sorting Functions
- Insertion()
- Selection()
- Merge()
- Heap()
- Quick()
- Bubble()
- Bucket()
- Radix()

###### Optional Parameters
```
    toPrint:
    for functions with input parameter as (toPrint ...bool)
    toPrint is a boolean value, that is optional
        if true is provided, it will print to console
        if nothing (or false) is provided, it won't print
        regardless it sends the actual raw data in multi-dimensional string array
```
## Complexity functions
1. GetAllComplexity()
```
    import "github.com/raunakjodhawat/gorithims/src/Utility"
    Utility.GetAllComplexity(toPrint ...bool)
    
    /*  
        Returns and print a 2d string array of list of
        Time and Space Complexity of sorting algorithims 
        defined for this repo
    */
```
2. GetAllTimeComplexity()
```
    import "github.com/raunakjodhawat/gorithims/src/Utility"
    Utility.GetAllTimeComplexity(toPrint ...bool)
    
    /*  
        Returns and print a 2d string array of list of all
        Time Complexity of sorting algorithims 
        defined for this repo
    */
```
3. GetAllSpaceComplexity()
```
    import "github.com/raunakjodhawat/gorithims/src/Utility"
    Utility.GetAllTimeComplexity(toPrint ...bool)
    
    /*  
        Returns and print a 2d string array of list of all
        Space Complexity of sorting algorithims 
        defined for this repo
    */
```
4. GetTimeComplexityByKey()
5. GetSpaceComplexityByKey()
6. GetAllComplexityByKey()
7. GetAllFunc()
#### Other functions
1. GetCredits()
```
    import "github.com/raunakjodhawat/gorithims/src/Utility"
    Utility.GetCredits(toPrint ...bool)
    
    /*  
        Returns and print a 2d string array of list of References
        used in creation of this project
    */
```

## External Dependency
1. fmt
2. time