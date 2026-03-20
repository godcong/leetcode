# LeetCode Project Directory Structure Specification

## Overview

This document defines the directory structure and organization standards for the LeetCode solution repository.

## 1. Root Directory Structure

```
leetcode/
├── code/                          # All problem solutions
│   ├── problems/                  # Regular problem set
│   ├── contest/                   # Competition problems
│   └── interview/                 # Interview questions
├── common/                        # Common definitions
├── cmd/                           # Command-line tools
├── docs/                          # Documentation
└── README.md                      # Main index
```

## 2. Code Directory Organization

### 2.1 Problems (Regular Problem Set)

**Structure:**
```
code/problems/
├── 0000-0099/                    # Grouped by hundreds
│   ├── 0001/                     # Individual problem directory
│   │   ├── two-sum.go
│   │   ├── two-sum_test.go
│   │   └── README.md
│   ├── 0002/
│   └── ...
├── 0100-0199/
└── ...
```

**Rules:**
- Group problems in ranges of 100 (0000-0099, 0100-0199, etc.)
- Each problem has its own directory named with the problem number
- Include solution file, test file, and README in each problem directory

### 2.2 Contest (Competition Problems)

**Structure:**
```
code/contest/
├── lcp/                          # LCP competition series
│   ├── LCP_06/
│   │   ├── min-count.go
│   │   ├── min-count_test.go
│   │   └── README.md
│   ├── LCP_19/
│   └── ...
└── weekly/                       # Weekly contest series
```

**Rules:**
- Organize by competition series (LCP, weekly, etc.)
- Each competition problem has independent directory
- Use format: `{series}_{number}` for directory names

### 2.3 Interview (Interview Questions)

**Structure:**
```
code/interview/
├── programmer/                   # Programmer Interview Gold
│   ├── Interview_01_02/
│   │   ├── check-permutation-lcci.go
│   │   ├── check-permutation-lcci_test.go
│   │   └── README.md
│   └── ...
└── sword-offer/                  # Sword Refers Offer
    ├── offer-v1/                # First Edition (Classic)
    │   ├── 0000-0099/
    │   │   ├── 010-fibonacci/
    │   │   │   ├── fibonacci.go
    │   │   │   ├── fibonacci_test.go
    │   │   │   └── README.md
    │   │   └── ...
    │   └── ...
    └── offer-v2/                # Second Edition (LeetCode ID)
        ├── 0000-0099/
        │   ├── 029-stream-median/
        │   │   ├── stream-median.go
        │   │   ├── stream-median_test.go
        │   │   └── README.md
        │   └── ...
        └── 0100-0199/
```

**Rules:**
- Separate `programmer` and `sword-offer` series
- Sword Offer divided into V1 (classic) and V2 (LeetCode ID version)
- Each edition grouped by problem number ranges
- Problem directories use format: `{number}-{keyword}`

## 3. Key Principles

### 3.1 Directory Isolation
- **Each problem must have its own directory**
- Prevents file naming conflicts
- Enables individual README per problem
- Facilitates easy navigation

### 3.2 Hierarchical Organization
```
Level 1: Category (problems, contest, interview)
Level 2: Series/Subcategory (lcp, programmer, sword-offer)
Level 3: Version/Edition (offer-v1, offer-v2)
Level 4: Range Grouping (0000-0099, 0100-0199)
Level 5: Individual Problem ({number}/ or {number}-{keyword}/)
```

### 3.3 English-Only Naming
- All directory names must be in English
- No Chinese characters or special symbols
- Use hyphens (-) for word separation
- Use underscores (_) for series prefixes

## 4. Special Directories

### 4.1 Common Definitions
```
common/
├── common.go          # Data structure definitions
├── common_func.go     # Utility functions
└── utils/             # Additional utilities
```

### 4.2 Command Tools
```
cmd/
├── gen/              # Problem generator
├── readme/           # README generator
└── query/            # Query tool
```

## 5. Migration Guidelines

When reorganizing existing files:

1. **Identify category** - Determine which section the problem belongs to
2. **Check range** - Find appropriate range group (e.g., 0000-0099)
3. **Create directory** - Make problem-specific directory
4. **Move files** - Relocate solution, test, and README
5. **Update package** - Ensure package name matches new location
6. **Verify imports** - Update import paths if needed

## 6. Examples

### Good Structure ✅
```
code/problems/0000-0099/0001/two-sum.go
code/interview/sword-offer/offer-v1/0000-0099/010-fibonacci/fibonacci.go
code/contest/lcp/LCP_06/min-count.go
```

### Bad Structure ❌
```
code/0001.two-sum.go                    # No isolation
code/problems/all-problems/             # Mixed together
code/interview/sword-offer/files/       # No categorization
```

## 7. Maintenance

- Regular audits to ensure compliance
- Automated checks in CI/CD pipeline
- Document any exceptions or special cases
- Keep this specification up to date

## Related Documents

- [Naming Convention](NAMING_CONVENTION.md)
- [Development Guide](DEVELOPMENT_GUIDE.md)
- [README Template](README_TEMPLATE.md)
