# 同位字符串连接的最小长度

## Problem Info

- **Number:** 3138
- **Daily Question Date:** 2026-03-21
- **Difficulty:** Medium

## Problem Link

- [LeetCode](https://leetcode-cn.com/problems/minimum-length-of-anagram-concatenation/)

## Description

给你一个字符串 `s` ，它由某个字符串 `t` 和若干 `t` 的 **同位字符串** 连接而成。

请你返回字符串 `t` 的 **最小** 可能长度。

**同位字符串** 指的是重新排列一个字符串的字母得到的另外一个字符串。例如，"aab"，"aba" 和 "baa" 是 "aab" 的同位字符串。

**示例 1：**

**输入：**s = "abba"

**输出：**2

**解释：**

一个可能的字符串 `t` 为 `"ba"` 。

**示例 2：**

**输入：**s = "cdef"

**输出：**4

**解释：**

一个可能的字符串 `t` 为 `"cdef"` ，注意 `t` 可能等于 `s` 。

**示例 3：**

**输入：**s = "abcbcacabbaccba"

**输出：**3

**提示：**

- `1 <= s.length <= 105`
- `s` 只包含小写英文字母。