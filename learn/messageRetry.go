package main

func getMessageRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	pl := len(primary)
	sl := len(secondary) + pl
	tl := len(tertiary) + sl
	return [3]string{primary, secondary, tertiary}, [3]int{pl, sl, tl}
}
