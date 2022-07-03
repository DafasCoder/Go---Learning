# Day 4 - Arrays and Slices

What we learned today!!

1. Arrays are declare like 'name []int'

   - Arrays can be intialize like this 'name [5]int{"Jack", "John"}'
   - Arrays have a fixed sized

2. Leaving the array type empty creates a slice and intialize the capicity make
   it an array
3. for loop can also be made like this 'for \_, value := range values'

   - This allows to count the length of the array
   - It capture the value similar to a for each
   - Leaving \_ allows us to ignore the index

4. 'reflect.DeepEqual' is another way of comparing values but it isn't type-safe
   so it will compile. Best option would be two sorting to check each indiviual
   value

   - Require the import for "reflect" package

5. Slices

   - How they have a fixed capacity but you can create new slices from old ones using append
   - So by doing 'number[1:]' we take 1 to the end
   - Slice can be any size

6. Using "test cover" command allows to see which function are cover by your test suite
