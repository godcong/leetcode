# 外观数列

## Problem Info

- **Number:** 38
- **Daily Question Date:** 2026-03-21
- **Difficulty:** Medium

## Problem Link

- [LeetCode](https://leetcode-cn.com/problems/count-and-say/)

## Description

「外观数列」是一个数位字符串序列，由递归公式定义：

- `countAndSay(1) = "1"`
- `countAndSay(n)` 是 `countAndSay(n-1)` 的行程长度编码。

<!--THE END-->

[行程长度编码](https://baike.baidu.com/item/%E8%A1%8C%E7%A8%8B%E9%95%BF%E5%BA%A6%E7%BC%96%E7%A0%81/2931940)（RLE）是一种字符串压缩方法，其工作原理是通过将连续相同字符（重复两次或更多次）替换为字符重复次数（运行长度）和字符的串联。例如，要压缩字符串 `"3322251"` ，我们将 `"33"` 用 `"23"` 替换，将 `"222"` 用 `"32"` 替换，将 `"5"` 用 `"15"` 替换并将 `"1"` 用 `"11"` 替换。因此压缩后字符串变为 `"23321511"`。

给定一个整数 `n` ，返回 **外观数列** 的第 `n` 个元素。

**示例 1：**

**输入：**n = 4

**输出：**"1211"

**解释：**

countAndSay(1) = "1"

countAndSay(2) = "1" 的行程长度编码 = "11"

countAndSay(3) = "11" 的行程长度编码 = "21"

countAndSay(4) = "21" 的行程长度编码 = "1211"

**示例 2：**

**输入：**n = 1

**输出：**"1"

**解释：**

这是基本情况。

**提示：**

- `1 <= n <= 30`

**进阶：**你能迭代解决该问题吗？