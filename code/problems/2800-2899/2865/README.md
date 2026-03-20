# 美丽塔 I

## Problem Info

- **Number:** 2865
- **Daily Question Date:** 2026-03-21
- **Difficulty:** Medium

## Problem Link

- [LeetCode](https://leetcode-cn.com/problems/beautiful-towers-i/)

## Description

给定一个包含 `n` 个整数的数组 `heights` 表示 `n` 座连续的塔中砖块的数量。你的任务是移除一些砖块来形成一个 **山脉状** 的塔排列。在这种布置中，塔高度先是非递减，有一个或多个连续塔达到最大峰值，然后非递增排列。

返回满足山脉状塔排列的方案中，**高度和的最大值** 。

**示例 1：**

```
输入：maxHeights = [5,3,4,1,1]
输出：13
解释：我们移除一些砖块来形成 heights = [5,3,3,1,1]，峰值位于下标 0。
```

**示例 2：**

```
输入：maxHeights = [6,5,3,9,2,7]
输出：22
解释：我们移除一些砖块来形成 heights = [3,3,3,9,2,2]，峰值位于下标 3。
```

**示例 3：**

```
输入：maxHeights = [3,2,5,5,2,3]
输出：18
解释：我们移除一些砖块来形成 heights = [2,2,5,5,2,2]，峰值位于下标 2 或 3。
```

**提示：**

- `1 <= n == heights.length <= 103`
- `1 <= heights[i] <= 109`