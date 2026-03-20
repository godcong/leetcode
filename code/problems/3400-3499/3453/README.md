# 分割正方形 I

## Problem Info

- **Number:** 3453
- **Daily Question Date:** 2026-03-21
- **Difficulty:** Medium

## Problem Link

- [LeetCode](https://leetcode-cn.com/problems/separate-squares-i/)

## Description

给你一个二维整数数组 `squares` ，其中 `squares[i] = [xi, yi, li]` 表示一个与 x 轴平行的正方形的左下角坐标和正方形的边长。

找到一个**最小的** y 坐标，它对应一条水平线，该线需要满足它以上正方形的总面积 **等于** 该线以下正方形的总面积。

答案如果与实际答案的误差在 `10-5` 以内，将视为正确答案。

**注意**：正方形 **可能会** 重叠。重叠区域应该被 **多次计数** 。

**示例 1：**

**输入：** squares = \[\[0,0,1],\[2,2,1]]

**输出：** 1.00000

**解释：**

![](https://pic.leetcode.cn/1739609465-UaFzhk-4062example1drawio.png)

任何在 `y = 1` 和 `y = 2` 之间的水平线都会有 1 平方单位的面积在其上方，1 平方单位的面积在其下方。最小的 y 坐标是 1。

**示例 2：**

**输入：** squares = \[\[0,0,2],\[1,1,1]]

**输出：** 1.16667

**解释：**

![](https://pic.leetcode.cn/1739609527-TWqefZ-4062example2drawio.png)

面积如下：

- 线下的面积：`7/6 * 2 (红色) + 1/6 (蓝色) = 15/6 = 2.5`。
- 线上的面积：`5/6 * 2 (红色) + 5/6 (蓝色) = 15/6 = 2.5`。

由于线以上和线以下的面积相等，输出为 `7/6 = 1.16667`。

**提示：**

- `1 <= squares.length <= 5 * 104`
- `squares[i] = [xi, yi, li]`
- `squares[i].length == 3`
- `0 <= xi, yi <= 109`
- `1 <= li <= 109`
- 所有正方形的总面积不超过 `1012`。