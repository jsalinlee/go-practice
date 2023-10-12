// Package dog provides conversions from human measurements to dog measurements
package dog

// Years converts human time to dog time in respect to the average duration of one Earth solar orbit
func Years(n int) int {
	return n * 7
}
