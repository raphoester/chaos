# Chaos - Stop hardcoding test values !
 
**WARNING: This library is still in its early development phase (semver 0.x). Retrocompatibility is not ensured.**

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
- What if some of your tests require you to call this function multiple times ?

Hardcoded values cause pain.

"Why not use random values ?" - you may say

But random values are not deterministic.

Using random values leads to what we call "flaky" tests.

With flaky tests, it will happen that you simply can't reproduce that bug that your CI found.

## Chaos to the rescue

Chaos is a library that generates "random" (deterministic) values for you, and makes them reproducible.

```go
func TestSaveUser(t *testing.T) {
    repository := NewRepository()

    // generate a random string of length 10 with the default seed
    firstName := chaos.String(10) 
    lastName := chaos.String(10)
	
    if err := repository.SaveUser(firstName, lastName); err != nil {
        t.Errorf("Error while saving user: %v", err)
    }
}
```

Unlike your random seed (typically the current time), the hardcoded seed doesn't change between executions.
That means you will get the same values every single time.

## Available Functions

There are many functions that allow you to generate different kinds of values: 
- Int
- Float
- Bool
- Duration
- UUID
- Time
- Slice items...

For detailed usage of each function, please refer to the source code and test files.

## Contributing

If you'd like to contribute to Chaos, please follow these steps:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Running Tests

To run the tests for Chaos:

1. Make sure you have Go installed on your system.
2. Clone the repository:
   ```bash
   git clone https://github.com/your-username/chaos.git
   cd chaos
   ```
3. Run the tests using the `go test` command:
   ```bash
   go test -v ./...
   ```

This will run all the tests in the package and display the results.
