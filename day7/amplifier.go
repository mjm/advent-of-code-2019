package day7

type Amplifier struct {
	Template *VM
	VM       *VM
}

func (amp *Amplifier) Init(phaseSetting int) {
	amp.VM = amp.Template.Clone()
	amp.VM.Input <- phaseSetting
}

func (amp *Amplifier) Run() {
	amp.VM.Execute()
}

func (amp *Amplifier) PipeTo(other *Amplifier) {
	amp.VM.PipeTo(other.VM)
}

func (amp *Amplifier) Send(n int) {
	amp.VM.Input <- n
}

func (amp *Amplifier) Output() int {
	return <-amp.VM.Output
}

type AmplifierArray struct {
	Amps []*Amplifier
}

func NewAmplifierArray(template *VM, size int) *AmplifierArray {
	amps := make([]*Amplifier, size)
	for i := range amps {
		amps[i] = &Amplifier{Template: template}
	}

	return &AmplifierArray{Amps: amps}
}

func (a *AmplifierArray) Run(phaseSettings []int) int {
	if len(phaseSettings) != len(a.Amps) {
		panic("amplifier array and phase settings count mismatch")
	}

	for i, setting := range phaseSettings {
		a.Amps[i].Init(setting)
	}
	go a.Amps[0].Send(0)

	for i := 0; i < len(a.Amps)-1; i++ {
		go a.Amps[i].PipeTo(a.Amps[i+1])
	}

	for _, amp := range a.Amps {
		go amp.Run()
	}

	return a.Amps[len(a.Amps)-1].Output()
}

func (a *AmplifierArray) RunWithFeedback(phaseSettings []int) int {
	if len(phaseSettings) != len(a.Amps) {
		panic("amplifier array and phase settings count mismatch")
	}

	firstAmp := a.Amps[0]
	lastAmp := a.Amps[len(a.Amps)-1]

	for i, setting := range phaseSettings {
		a.Amps[i].Init(setting)
	}
	go firstAmp.Send(0)

	for i := 0; i < len(a.Amps)-1; i++ {
		go a.Amps[i].PipeTo(a.Amps[i+1])
	}

	done := make(chan int)
	go func() {
		var lastVal int
		for val := range lastAmp.VM.Output {
			lastVal = val
			firstAmp.VM.Input <- val
		}
		done <- lastVal
	}()

	for _, amp := range a.Amps {
		go amp.Run()
	}

	return <-done
}

func (a *AmplifierArray) HighestSignal(phaseSettings []int) (int, []int) {
	var highestSequence []int
	var highestOutput int

	for nums := range AllPermutations(phaseSettings) {
		out := a.Run(nums)
		if out > highestOutput {
			highestOutput = out
			highestSequence = nums
		}
	}

	return highestOutput, highestSequence
}

func (a *AmplifierArray) HighestFeedbackSignal(phaseSettings []int) (int, []int) {
	var highestSequence []int
	var highestOutput int

	for nums := range AllPermutations(phaseSettings) {
		out := a.RunWithFeedback(nums)
		if out > highestOutput {
			highestOutput = out
			highestSequence = nums
		}
	}

	return highestOutput, highestSequence
}
