package bowling

type Status string

const (
	OpenedStatus Status = "opened"
	StrikeStatus Status = "strike"
	SpareStatus  Status = "spare"

	FirstRollStatus  Status = "first"
	SecondRollStatus Status = "second"
	ThirdRollStatus  Status = "third"
	FinalStatus      Status = "final"
)
