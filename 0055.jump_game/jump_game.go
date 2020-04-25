package jump_game

func canJump(nums []int) bool {
	jump := 0
	for i, n := range nums {
		if i > jump {
			return false
		}
		if n+i > jump {
			jump = n + i
		}
	}
	return true
}
