# Go Examples Project - AI Coding Instructions

This is a Go learning repository structured as an interactive command-line tool inspired by "Go by Example". It contains algorithm problems, Go language samples, and comprehensive tests.

## Project Architecture

### Core Structure
- **`main.go`**: CLI router using `flag.NewFlagSet()` pattern for subcommands (e.g., `defer`, `hello`, `fibonacci`)
- **`samples/`**: Educational Go examples for language features (types, functions, data structures)
- **`problem/`**: LeetCode-style algorithm implementations
- **`testing/`**: Comprehensive test suites for both samples and problems

### CLI Pattern
The main.go uses a sophisticated flag-based routing system:
```go
switch os.Args[1] {
case "hello":
    helloCmd.Parse(os.Args[2:])
    samples.Hello(*helloName)
}
```
When adding new commands, follow this pattern: define flags, create FlagSet, add switch case.

## Package Organization

### Relative Import Convention
Uses relative imports throughout: `import problem "../problem"` and `import samples "./samples"`
- Problem solutions in `package problem`
- Learning samples in `package samples` 
- Tests in `package testing` or `package testing_test`

### No Go Modules
Project runs without go.mod - uses relative imports and GOPATH-style structure. Tests run with `go test` from individual directories, not with module-aware commands.

## Testing Patterns

### Algorithm Problem Tests
Follow this comprehensive table-driven pattern:
```go
func TestTwoSum(t *testing.T) {
    cases := []struct {
        nums   []int
        target int
        want   []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
        // Multiple comprehensive test cases
    }
    for _, c := range cases {
        got := problem.TwoSum(c.nums, c.target)
        if !equal(got, c.want) {
            t.Errorf("TwoSum(%v, %d) = %v; want %v", c.nums, c.target, got, c.want)
        }
    }
}
```

### Helper Functions in Tests
Create specialized helper functions for complex data structures:
- `createList()` and `listToSlice()` for LinkedList problems
- Custom `equal()` functions that handle order-independent comparisons
- Type aliases like `type ListNode = problem.ListNode` for cleaner test code

### Sample Function Tests
Use simpler structure for learning examples:
```go
cases := []struct {
    in, want int
}{
    {0, 0}, {1, 1}, {2, 1},
}
```

## Code Conventions

### Function Naming
- Algorithm functions: `TwoSum`, `MergeTwoList` (PascalCase, LeetCode style)
- Sample functions: `Hello`, `Fibonacci`, `DeferSamples` (PascalCase with descriptive suffixes)
- Test helpers: `createList`, `listToSlice` (camelCase)

### Comments
- Include LeetCode URLs: `//https://leetcode.com/problems/two-sum/`
- Descriptive function comments: `// Fibonacci Sequence is the series of numbers: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, ...`
- Inline comments for algorithm explanations

### Data Structure Patterns
- **ListNode**: Use consistent field naming (`Val int`, `Next *ListNode` for problems; `Value int`, `Next *ListNode` for samples)
- **Maps for lookups**: `numIdx := make(map[int]int) // Key: Value Value: Index`
- **Multiple assignment**: `node.Next, l1 = l1, l1.Next`

## Running the Project

### CLI Usage
Run examples with: `go run main.go <command> [flags]`
Examples:
- `go run main.go hello -name="World"`
- `go run main.go fibonacci -number=10`
- `go run main.go func -contain="apple" -prefix="a" -min-len=5`

### Testing
Tests must be run from the testing directory: `cd testing && go test`
No module-aware testing - project lacks go.mod intentionally.

### Development Workflow
1. Implement algorithm in `problem/` with LeetCode URL comment
2. Create comprehensive test in `testing/` with extensive edge cases
3. Add CLI integration in `main.go` if it's a sample (not for pure algorithm problems)
4. Use helper functions for complex test setup/teardown

## Key Patterns to Maintain

- **Relative imports**: Always use `"../problem"` style imports
- **Table-driven tests**: Extensive test cases covering edge cases
- **CLI flag patterns**: Follow the established FlagSet pattern in main.go
- **No external dependencies**: Pure Go standard library only
- **Educational focus**: Code should be clear and well-commented for learning