# LeetCode Naming Convention Specification

## Overview

This document defines the naming conventions for files, directories, and packages in the LeetCode solution repository.

## 1. File Naming Patterns

### Core Principle
**Directory provides the index, filename only needs function name**

✅ Good:
```
code/problems/0000-0099/0001/two-sum.go
code/problems/0000-0099/0001/two-sum_test.go
```

❌ Bad:
```
code/problems/0000-0099/0001/0001.two-sum.go  # Redundant index
code/problems/0000-0099/0001/0001.two-sum_test.go
```

### 1.1 Problems (Regular Problem Set)

**Pattern:** `{function-name}.go`

**Examples:**
```
two-sum.go
two-sum_test.go
find-min.go
find-min_test.go
```

**Rules:**
- `function-name`: LeetCode official function name in kebab-case
- Test files use `_test.go` suffix
- **No index prefix needed** (directory already has it)

### 1.2 Contest Problems (LCP Series)

**Pattern:** `{function-name}.go`

**Examples:**
```
min-count.go
min-count_test.go
minimum-operations.go
minimum-operations_test.go
```

**Rules:**
- Function name in kebab-case
- **No series prefix or index needed** (directory already identifies it)

### 1.3 Interview Programmer Series

**Pattern:** `{function-name}.go`

**Examples:**
```
deleteNode.go
deleteNode_test.go
check-permutation-lcci.go
check-permutation-lcci_test.go
one-away-lcci.go
```

### 1.4 Sword Refers Offer Series

#### Offer V1 (First Edition - Classic)

**Directory:** `offer-v1/{range}/{num}-{keyword}/`
**File Pattern:** `{function-name}.go`

**Examples:**
```
offer-v1/0000-0099/010-fibonacci/fibonacci.go
offer-v1/0000-0099/010-fibonacci/fibonacci_test.go
offer-v1/0000-0099/022-find-kth-node/find-kth-node.go
offer-v1/0000-0099/022-find-kth-node/find-kth-node_test.go
```

**Keyword Mapping:**
```
010 → fibonacci (斐波那契数列)
020 → is-number (表示数值的字符串)
022 → find-kth-node (链表中倒数第 k 个节点)
042 → max-sub-array (连续子数组的最大和)
047 → max-value-of-gifts (礼物的最大价值)
052 → first-common-node (两个链表的第一个公共节点)
058 → reverse-words (翻转单词顺序)
```

#### Offer V2 (Second Edition - LeetCode ID)

**Directory:** `offer-v2/{range}/{num}-{leetcod-id}/`
**File Pattern:** `{function-name}.go` or `{leetcod-id}.go`

**Examples:**
```
offer-v2/0000-0099/029-stream-median/stream-median.go
offer-v2/0000-0099/029-stream-median/stream-median_test.go
offer-v2/0000-0099/041-max-sliding-window/max-sliding-window.go
offer-v2/0000-0099/041-max-sliding-window/max-sliding-window_test.go
```

**Alternative (using LeetCode IDs):**
```
offer-v2/0000-0099/029-stream-median/4ueAj6.go
offer-v2/0000-0099/029-stream-median/4ueAj6_test.go
offer-v2/0000-0099/041-max-sliding-window/qIsx9U.go
offer-v2/0000-0099/041-max-sliding-window/qIsx9U_test.go
```

## 2. Package Naming Rules

### 2.1 General Principles

- Package names must match **directory name** (not full path)
- Use underscore prefix to avoid conflicts with standard libraries
- Keep package names concise and meaningful

### 2.2 Package Name Patterns

| Category | Directory | Package Pattern | Example |
|----------|-----------|-----------------|---------|
| **Problems** | `0001/` | `_{index}` | `package _0001` |
| **Contest LCP** | `LCP_06/` | `_{prefix}_{index}` | `package _LCP_06` |
| **Interview Programmer** | `Interview_02_03/` | `_{type}_{index}` | `package _Interview_02_03` |
| **Sword Offer V1** | `10-fibonacci/` | `_{version}_{num}` | `package _Offer_V1_010` |
| **Sword Offer V2** | `029-stream-median/` | `_{version}_{num}` | `package _Offer_V2_029` |

### 2.3 Examples by Category

#### Problems
```go
// code/problems/0000-0099/0001/two-sum.go
package _0001

func twoSum(nums []int, target int) []int {
    // Implementation
}
```

#### Contest
```go
// code/contest/lcp/LCP_06/min-count.go
package _LCP_06

func minCount(coins []int) int {
    // Implementation
}
```

#### Interview Programmer
```go
// code/interview/programmer/Interview_02_03/deleteNode.go
package _Interview_02_03

import . "github.com/godcong/leetcode/common"

func deleteNode(node *ListNode) {
    // Implementation
}
```

#### Sword Offer V1
```go
// code/interview/sword-offer/offer-v1/0000-0099/010-fibonacci/fibonacci.go
package _Offer_V1_010

func fib(n int) int {
    // Implementation
}
```

#### Sword Offer V2
```go
// code/interview/sword-offer/offer-v2/0000-0099/029-stream-median/stream-median.go
package _Offer_V2_029

func streamMedian() /* return type */ {
    // Implementation
}
```

## 3. Directory Naming Rules

### 3.1 Range Directories

**Pattern:** `{start}-{end}` (4-digit numbers)

**Examples:**
```
0000-0099
0100-0199
1000-1099
```

**Rules:**
- Always use 4-digit numbers
- Inclusive ranges
- Zero-padded

### 3.2 Problem Directories

**Pattern A (Problems):** `{index}`
**Pattern B (Sword Offer):** `{num}-{keyword}`

**Examples:**
```
0001                    # Problems
0153                    # Problems  
Interview_01_02        # Interview Programmer
10-fibonacci           # Sword Offer V1
029-stream-median      # Sword Offer V2
```

### 3.3 Series Directories

Use descriptive English names:
```
problems                # Regular problem set
contest                 # Competition problems
interview               # Interview questions
lcp                     # LCP competition series
programmer              # Programmer interview gold
sword-offer             # Sword refers offer
offer-v1                # Sword offer first edition
offer-v2                # Sword offer second edition
```

## 4. Special Naming Cases

### 4.1 Chinese to English Translation

For Sword Offer V1, translate Chinese problem names to English keywords:

| Chinese | English Keyword |
|---------|----------------|
| 斐波那契数列 | fibonacci |
| 表示数值的字符串 | is-number |
| 链表中倒数第 k 个节点 | find-kth-node |
| 连续子数组的最大和 | max-sub-array |
| 礼物的最大价值 | max-value-of-gifts |
| 两个链表的第一个公共节点 | first-common-node |
| 翻转单词顺序 | reverse-words |

### 4.2 LeetCode ID Usage

For Sword Offer V2, can use LeetCode problem IDs:

```
4ueAj6.go              # Using LeetCode ID directly
stream-median.go       # Using descriptive keyword
```

Recommend using descriptive keywords for better readability.

### 4.3 Duplicate Function Names

When multiple problems have same function name:

```
// Problem 1: Remove Duplicates from Sorted Array
0026.remove-duplicates.go
package _0026

// Problem 2: Remove Duplicates from Sorted List
0083.remove-duplicates.go
package _0083
```

Directory isolation prevents conflicts.

## 5. Import Path Conventions

### 5.1 Common Definitions

```go
import . "github.com/godcong/leetcode/common"
// or
import . "leetcode/common"
```

### 5.2 Relative Imports (if needed)

```go
import . "../../common"
```

### 5.3 Standard Library Imports

Always use standard Go imports:
```go
import (
    "math"
    "sort"
    "strings"
)
```

## 6. README File Naming

Each problem directory must include:

```
README.md              # Problem description and solution
```

**Optional:**
```
SOLUTION.md           # Detailed solution explanation
NOTES.md              # Notes and tips
COMPLEXITY.md         # Complexity analysis
```

## 7. Test File Naming

**Pattern:** `{base-name}_test.go`

**Examples:**
```
two-sum_test.go
fibonacci_test.go
min-count_test.go
```

**Rules:**
- Same base name as solution file
- Use `_test.go` suffix (Go convention)
- Same package as solution file

## 8. Common Mistakes to Avoid

### ❌ Bad Examples

```go
// Wrong: No directory isolation
code/0001.two-sum.go

// Wrong: Inconsistent naming
code/problems/two_sum.go
code/problems/TwoSum.go

// Wrong: Package doesn't match directory
// Directory: 0001/
package leetcode  // Should be: package _0001

// Wrong: Chinese characters in names
code/两数之和.go  // Should use English
```

### ✅ Good Examples

```go
// Correct: Proper isolation
code/problems/0000-0099/0001/two-sum.go

// Correct: Consistent kebab-case
code/problems/0000-0099/0001/two-sum.go
code/problems/0000-0099/0001/two-sum_test.go

// Correct: Package matches directory
// Directory: 0001/
package _0001

// Correct: English only
code/problems/0000-0099/0001/two-sum.go
```

## 9. Summary Table

| Component | Format | Example |
|-----------|--------|---------|
| **Problem File** | `{function}.go` | `two-sum.go` |
| **Test File** | `{function}_test.go` | `two-sum_test.go` |
| **Directory** | `{index}` or `{num}-{keyword}` | `0001/`, `10-fibonacci/` |
| **Package** | Match directory name | `_0001`, `_Offer_V1_010` |
| **Range** | `{start}-{end}` | `0000-0099` |

**Key Principle:** Directory provides context, filename only needs function name!

## Related Documents

- [Directory Structure](DIRECTORY_STRUCTURE.md)
- [Development Guide](DEVELOPMENT_GUIDE.md)
- [Code Style](CODE_STYLE.md)
