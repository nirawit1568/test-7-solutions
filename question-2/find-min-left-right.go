package question2

import "fmt"

func FindMinLeftRight(input string) string {

	result := make([]int, len(input)+1)

	// ฟังชั่นสำหรับเพิ่มตัวเลขให้ตรงตามเงื่อนไข
	changeNumber := func() {

		for i, v := range input {

			left, right := result[i], result[i+1]

			switch v {
			case 'L':
				// ถ้าเป็นค่าเริ่มต้น(เป็น 0 ทุก index) หรือ ค่าซ้ายไม่เท่ากับขวาให้อัพเดทตามเงื่อนไข
				if left == 0 || left == right {
					result[i] = right + 1
				}

			case 'R':
				// ถ้าเป็นค่าเริ่มต้น(เป็น 0 ทุก index) หรือ ค่าขวาไม่เท่ากับซ้ายให้อัพเดทตามเงื่อนไข
				if right == 0 || right == left {
					result[i+1] = left + 1
				}

			case '=':
				// ถ้าค่าไม่เท่ากันระหว่างขวากับซ้าย ให้อัพเดทเลย
				if right != left {
					if i == 0 { // ถ้า "=" อยู่ตัวแรกสุดให้เลือกเอาค่าจากขวามาใส่ซ้าย
						result[i] = right
					} else { // ถ้า "=" อยู่ตัวสุดท้ายให้เลือกเอาค่าจากซ้ายมาใส่ขวาา
						result[i+1] = left
					}
				}
			}
		}
	}

	// ฟังชั่นสำหรับเช็คว่าผลลัพธ์เรียงลำดับได้ถูกต้องแล้วหรือไม่
	checkNumber := func() bool {

		var isCorrect bool

		for i, v := range input {

			left, right := result[i], result[i+1]

			if v == 'L' && !(left > right) {
				isCorrect = false
				break

			} else if v == 'R' && !(right > left) {
				isCorrect = false
				break

			} else if v == '=' && !(left == right) {
				isCorrect = false
				break
			}

			if i == len(input)-1 {
				isCorrect = true
			}
		}

		return isCorrect
	}

	for !checkNumber() {
		changeNumber()
	}

	// เรียงผลลัพธ์
	output := ""
	for _, num := range result {
		output += fmt.Sprintf("%d", num)
	}

	return output
}
