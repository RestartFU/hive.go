package hive

type Time struct {
	time time
}

func (t Time) String() string { return string(t.time) }

type time string

func TimeAll() Time     { return Time{time: "all"} }
func TimeMonthly() Time { return Time{time: "monthly"} }
