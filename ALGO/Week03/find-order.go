package main

func main()  {

}

func findOrder(numCourses int, prerequisites [][]int) []int {
	// 入度表，记录本课程先休课程（每个节点的前驱）
	inDegree := make([]int, numCourses)
	// 出边表，记录本课程是哪些课程的先休
	outEdge := make([][]int, numCourses)
	// 结果集
	res := []int{}

	// 开启遍历
	for i := 0; i< len(prerequisites); i++ {
		// 更新入度
		inDegree[prerequisites[i][0]]++
		// 更新出边
		outEdge[prerequisites[i][1]] = append(outEdge[prerequisites[i][1]], prerequisites[i][0])
	}

	// 队列用于记录入度为0的课程
	queue := []int{}

	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	if len(queue) == 0 {
		return res
	}

	for len(queue) > 0 {
		// 取出队头
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)

		// 处理队头的出边，把入度为0的入队
		for i:= 0; i < len(outEdge[node]); i++ {
			out := outEdge[node][i]
			inDegree[out]--
			if inDegree[out] == 0 {
				queue = append(queue, out)
			}
		}
	}

	// 检查是否所有课程都被安排，如果没有，说明有环，无法安排，返回空
	if len(res) == numCourses {
		return res
	}

	return nil
}