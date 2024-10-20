package question1

func max(value1, value2 int) int {
	if value1 > value2 {
		return value1
	} else {
		return value2
	}
}

func FindMaxTreeSum(input [][]int) int {

	/*
		วิธีคิดคือต้องคิดจากข้างล่างขึ้นมา โดยเริ่มจากแถวก่อนล่างสุด
		โดยลองบวกค่าที่เชื่อมกับตัวเอง 2 ค่าด้านล่าง
		แล้วเลือกผลบวกของค่าที่เยอะที่สุด แล้วนำค่านั้นมาแทนตำแหน่งตัวเอง
		ทำแบบนี้ขึ้นมาเรื่อยๆ จะได้เส้นที่มีค่าผลรวมมากที่สุด
	*/

	for i := len(input) - 2; i >= 0; i-- {
		for j := 0; j < len(input[i]); j++ {
			input[i][j] += max(input[i+1][j], input[i+1][j+1])
		}
	}

	return input[0][0]
}
