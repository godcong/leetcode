package old

/*
874. 模拟行走机器人
机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：

-2：向左转 90 度
-1：向右转 90 度
1 <= x <= 9：向前移动 x 个单位长度
在网格上有一些格子被视为障碍物。

第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])

机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。

返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。



示例 1：

输入: commands = [4,-1,3], obstacles = []
输出: 25
解释: 机器人将会到达 (3, 4)
示例 2：

输入: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
输出: 65
解释: 机器人在左转走到 (1, 8) 之前将被困在 (1, 4) 处


提示：

0 <= commands.length <= 10000
0 <= obstacles.length <= 10000
-30000 <= obstacle[i][0] <= 30000
-30000 <= obstacle[i][1] <= 30000
答案保证小于 2 ^ 31
*/
const (
	robotSimN = iota //北
	robotSimS        //东
	robotSimE        //南
	robotSimW        //西
)

func robotSim(commands []int, obstacles [][]int) int {
	current := [2]int{0, 0}
	direction := robotSimN
	max := -1
	obstaclesMap := make(map[[2]int]bool, len(obstacles))
	for i := range obstacles {
		obstaclesMap[[2]int{obstacles[i][0], obstacles[i][1]}] = true
	}
	var robotSimCoordinate func(d int, n int)
	robotSimCoordinate = func(d int, n int) {
		if n == 0 {
			return
		}
		switch d {
		case robotSimN:
			if obstaclesMap[[2]int{current[0], current[1] + 1}] {
				return
			}
			current[1]++
		case robotSimW:
			if obstaclesMap[[2]int{current[0] - 1, current[1]}] {
				return
			}
			current[0]--
		case robotSimE:
			if obstaclesMap[[2]int{current[0], current[1] - 1}] {
				return
			}
			current[1]--
		case robotSimS:
			if obstaclesMap[[2]int{current[0] + 1, current[1]}] {
				return
			}
			current[0]++
		}
		robotSimCoordinate(d, n-1)
	}
	for i := range commands {
		if commands[i] < 0 {
			direction = robotSimDirection(direction, commands[i])
		} else {
			robotSimCoordinate(direction, commands[i])
			max = robotSimDirectionMax(max, current[0]*current[0]+current[1]*current[1])
		}
	}
	return max
}

func robotSimDirectionMax(old, new int) int {
	if old > new {
		return old
	}
	return new
}

func robotSimDirection(d int, n int) int {
	if n == -2 {
		return (d + 3) % 4
	}
	return (d + 1) % 4
}
