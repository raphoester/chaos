# Chaos - Stop hardcoding test values !

## Why Chaos ?

Let's say that you want to write any kind of test.
You have to provide some input values to your function, and you have to check that the output is what you expect.

```go 
func TestSaveUser(t *testing.T) {
    repository := NewRepository()
    if err := repository.SaveUser("John", "Doe"); err != nil {
        t.Errorf("error while saving user: %v", err)
    }
}
```

Do you want to write "John" and "Doe" everywhere in your tests ?

- What if you have to test a function that takes 10 parameters ?
- What if you have to write 10 tests for this function ?
- What if some of your tests require you to call this function in a loop ?

Hardcoded values cause pain.

The alternative is to use a random value generator.

But random values are not reproductible.
This leads to flaky tests, and you can't reproduce the bug that your CI found.

## Chaos to the rescue

Chaos is a library that generates random values for you, and makes them reproducible.

```go
func TestSaveUser(t *testing.T) {
    repository := NewRepository()

	// generate a random string of length 10 with the seed "firstName"
    firstName := chaos.String(10, "firstName") 
    lastName := chaos.String(10, "lastName")
	
    if err := repository.SaveUser(firstName, lastName); err != nil {
        t.Errorf("Error while saving user: %v", err)
    }
}
```

Now you can run your tests with the same seed, and you will get the same values every time.

There are many other functions that allow you to generate different kinds of values: 
- Int
- Float
- Bool
- Duration...