# Day 6 - Pointers and Errors

What we learned today!!!!!!

## Pointers

1. Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

2. The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

## nil

1. Pointers can be nil
2. When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
3. Useful for when you want to describe a value that could be missing

## Errors

1. Errors are the way to signify failure when calling a function/method
2. By listening to our tests owe concluded that checking for a string in an errors would result in flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.

## Create new tupes for existing ones

1. useful for adding more domain specific meaning to values
2. Can let you implement interfaces
