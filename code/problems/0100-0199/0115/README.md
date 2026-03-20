# 不同的子序列

## Problem Info

- **Number:** 115
- **Daily Question Date:** 2026-03-21
- **Difficulty:** Hard

## Problem Link

- [LeetCode](https://leetcode-cn.com/problems/distinct-subsequences/)

## Description

给你两个字符串 `s` 和 `t` ，统计并返回在 `s` 的 **子序列** 中 `t` 出现的个数。

测试用例保证结果在 32 位有符号整数范围内。

**示例 1：**

```
输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
rabbbit
rabbbit
rabbbit
```

**示例 2：**

```
输入：s = "babgbag", t = "bag"
输出：5
解释：
如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。 
babgbag
babgbag
babgbag
babgbag
babgbag
```

**提示：**

- `1 <= s.length, t.length <= 1000`
- `s` 和 `t` 由英文字母组成