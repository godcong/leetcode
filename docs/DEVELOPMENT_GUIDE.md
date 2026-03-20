# LeetCode Development Guide

## Overview

This guide provides development standards and best practices for contributing to the LeetCode solution repository.

## 1. Project Structure

```
leetcode/
├── code/                      # Solution code
│   ├── problems/              # Regular problems
│   ├── contest/               # Competition problems
│   └── interview/             # Interview questions
├── common/                    # Common definitions
├── cmd/                       # CLI tools
├── docs/                      # Documentation
├── go.mod                     # Go module definition
└── README.md                  # Main index
```

## 2. Development Environment

### 2.1 Requirements

- Go 1.14 or higher
- Git for version control
- Code editor with Go support (VS Code, GoLand, etc.)

### 2.2 Setup

```bash
# Clone repository
git clone <repository-url>
cd leetcode

# Install dependencies
go mod download

# Verify setup
go test ./common/...
```

## 3. Using Common Definitions

### 3.1 Import Common Package

```go
import . "github.com/godcong/leetcode/common"
// or
import . "leetcode/common"
```

### 3.2 Available Data Structures

#### ListNode (Linked List)
```go
type ListNode struct {
    Val  int
    Next *ListNode
}

// Helper functions
func CreateList(arr []int) *ListNode
func PrintList(head *ListNode) string
```

#### TreeNode (Binary Tree)
```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Helper functions
func CreateTree(arr []interface{}) *TreeNode
func PrintTree(root *TreeNode) string
```

#### Node (N-ary Tree)
```go
type Node struct {
    Val      int
    Children []*Node
}
```

### 3.3 Example Usage

```go
package _0001

import . "github.com/godcong/leetcode/common"

func twoSum(head *ListNode) []int {
    // Use common ListNode
    result := []int{}
    for head != nil {
        result = append(result, head.Val)
        head = head.Next
    }
    return result
}
```

## 4. Writing Solutions

### 4.1 Basic Structure

```go
package _0001

// Problem: Two Sum
// Link: https://leetcode-cn.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
    // Your solution here
    m := make(map[int]int)
    for i, num := range nums {
        if j, ok := m[target-num]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}
```

### 4.2 Comments

Use English comments:

```go
// Good: Clear English comment
// Initialize hash map to store value -> index
m := make(map[int]int)

// Avoid: Chinese comments
// 初始化哈希表
```

### 4.3 Code Style

**Naming:**
- Functions: camelCase (LeetCode standard)
- Variables: camelCase
- Constants: PascalCase
- Types: PascalCase

**Formatting:**
- Use `gofmt` or `goimports`
- No trailing whitespace
- Proper indentation (tabs)

## 5. Writing Tests

### 5.1 Test File Structure

```go
package _0001

import (
    "testing"
    "github.com/godcong/leetcode/common"
)

func TestTwoSum(t *testing.T) {
    tests := []struct {
        nums   []int
        target int
        want   []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
        {[]int{3, 2, 4}, 6, []int{1, 2}},
        {[]int{3, 3}, 6, []int{0, 1}},
    }

    for _, tt := range tests {
        t.Run("", func(t *testing.T) {
            got := twoSum(tt.nums, tt.target)
            if !equal(got, tt.want) {
                t.Errorf("twoSum(%v, %d) = %v, want %v", 
                    tt.nums, tt.target, got, tt.want)
            }
        })
    }
}

func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}
```

### 5.2 Test Naming

- Use `Test{FunctionName}` pattern
- One test function per problem function
- Use table-driven tests for multiple cases

### 5.3 Running Tests

```bash
# Run all tests
go test ./...

# Run specific problem tests
go test ./code/problems/0000-0099/0001/...

# Run with verbose output
go test -v ./code/problems/0000-0099/0001/...

# Run with coverage
go test -cover ./...
```

## 6. Writing README

### 6.1 README Template

```markdown
# 0001. Two Sum

## Problem Link
- [LeetCode](https://leetcode-cn.com/problems/two-sum/)

## Description
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

**Example 1:**
```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
```

## Approach

### Hash Map

Use a hash map to store the value -> index mapping.

**Time Complexity:** O(n)  
**Space Complexity:** O(n)

## Code

```go
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        if j, ok := m[target-num]; ok {
            return []int{j, i}
        }
        m[num] = i
    }
    return nil
}
```

## Related Problems
- [0015. 3Sum]()
- [0018. 4Sum]()
```

### 6.2 README Guidelines

- Write in English (or bilingual EN/ZH)
- Include time and space complexity
- Provide at least one approach
- Link to related problems

## 7. Using CLI Tools

### 7.1 Generate New Problem

```bash
# Generate problem template
go run cmd/gen/main.go -type problem -id 0001

# Output:
# Created: code/problems/0000-0099/0001/two-sum.go
# Created: code/problems/0000-0099/0001/two-sum_test.go
# Created: code/problems/0000-0099/0001/README.md
```

### 7.2 Generate README Index

```bash
# Generate main README.md
go run cmd/readme/main.go

# Update statistics
go run cmd/readme/main.go --stats
```

## 8. Git Workflow

### 8.1 Commit Message Format

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types:**
- `feat`: New feature/solution
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting)
- `refactor`: Code refactoring
- `test`: Adding tests
- `chore`: Build/config changes

**Examples:**
```
feat(problems): add solution for problem 0001

- Add two-sum solution using hash map
- Add comprehensive test cases
- Add README with explanation

Closes #123
```

```
fix(sword-offer): correct package name in Offer_010

Package name should match directory structure

Fixes #456
```

### 8.2 Branch Strategy

```bash
# Main branch
main

# Feature branches
feature/add-problem-0001
fix/package-names
docs/update-readme
```

### 8.3 Pull Request Process

1. Create feature branch
2. Make changes and commit
3. Run tests locally
4. Push to remote
5. Create pull request
6. Code review
7. Merge after approval

## 9. Quality Assurance

### 9.1 Pre-commit Checklist

- [ ] Code compiles (`go build`)
- [ ] Tests pass (`go test`)
- [ ] Code formatted (`gofmt`)
- [ ] No linting errors (`golint`)
- [ ] Package names correct
- [ ] Import paths correct
- [ ] README updated (if needed)

### 9.2 Continuous Integration

CI pipeline runs:
- Build verification
- All tests
- Code coverage check
- Linting checks

### 9.3 Code Review Guidelines

Reviewers check:
- Correctness of solution
- Code quality and style
- Test coverage
- Documentation completeness
- Naming conventions

## 10. Troubleshooting

### Common Issues

#### Package Name Mismatch

**Error:**
```
package name _0002 doesn't match directory 0001
```

**Solution:**
```bash
# Ensure package name matches directory
# Directory: code/problems/0000-0099/0001/
# Package: package _0001
```

#### Import Path Errors

**Error:**
```
cannot find package "leetcode/common"
```

**Solution:**
```bash
# Use correct import path
import . "github.com/godcong/leetcode/common"

# Or update go.mod module path
```

#### Test Failures

**Debug:**
```bash
# Run with verbose output
go test -v ./path/to/problem/...

# Run specific test
go test -v -run TestFunctionName ./path/to/problem/...
```

## 11. Best Practices

### 11.1 Solution Quality

- **Multiple approaches**: Provide at least 2 different solutions if possible
- **Optimal solution**: Highlight the most efficient approach
- **Edge cases**: Handle all edge cases in tests
- **Documentation**: Explain trade-offs between approaches

### 11.2 Code Organization

- **Single responsibility**: One function per file (mostly)
- **Clear naming**: Self-documenting variable names
- **Comments**: Explain why, not what
- **Structure**: Group related problems together

### 11.3 Testing

- **Comprehensive**: Cover all edge cases
- **Isolated**: Each test independent
- **Repeatable**: Same result every time
- **Fast**: Quick execution

## 12. Performance Tips

### 12.1 Time Complexity Analysis

Common complexities (best to worst):
```
O(1) < O(log n) < O(n) < O(n log n) < O(n²) < O(2^n) < O(n!)
```

### 12.2 Space Optimization

- Reuse input arrays when possible
- Use in-place algorithms
- Avoid unnecessary data structures
- Consider trade-offs (time vs space)

## Related Documents

- [Directory Structure](DIRECTORY_STRUCTURE.md)
- [Naming Convention](NAMING_CONVENTION.md)
- [Code Style](CODE_STYLE.md)
