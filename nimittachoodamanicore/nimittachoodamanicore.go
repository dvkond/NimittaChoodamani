package nimittachoodamanicore

import "math/rand"

func About() string {
	return "The Divine Oracle Answers to your questions. Translated by Sri Raghavendra Rao."
}

func Help() string {
	return "Just think of a DEFINITE question (Eg: Will I get a promotion this year? When will I get married? Can I change job now?), think of God, offer a prayer, give yourself a few seconds of introspection and run NimittaChoodamani (without any params!)."
}

func RollTheDice() (string, string, string) {
	//=ЕСЛИ(СЛУЧМЕЖДУ(1;100)<25; 1; ЕСЛИ(СЛУЧМЕЖДУ(1;100)<50; 2; ЕСЛИ(СЛУЧМЕЖДУ(1;100)<75; 10; 100)))
	dice1 := compareLess(0.25, 0.5, 0.75)
	//=ЕСЛИ(СЛУЧМЕЖДУ(1;100)<35; 1; ЕСЛИ(СЛУЧМЕЖДУ(1;100)<70; 2; ЕСЛИ(СЛУЧМЕЖДУ(1;100)<95; 10; 100)))
	dice2 := compareLess(0.35, 0.7, 0.95)
	//=ЕСЛИ(СЛУЧМЕЖДУ(1;100)>75; 1; ЕСЛИ(СЛУЧМЕЖДУ(1;100)>50; 2; ЕСЛИ(СЛУЧМЕЖДУ(1;100)>25; 10; 100)))
	dice3 := compareMore(0.75, 0.5, 0.25)
	return dice1, dice2, dice3
}

func compareLess(threshold1, threshold2, threshold3 float64) string {
	if rand.Float64() < threshold1 {
		return "1"
	} else {
		if rand.Float64() < threshold2 {
			return "2"
		} else {
			if rand.Float64() < threshold3 {
				return "10"
			} else {
				return "100"
			}
		}
	}
}

func compareMore(threshold1, threshold2, threshold3 float64) string {
	if rand.Float64() > threshold1 {
		return "1"
	} else {
		if rand.Float64() > threshold2 {
			return "2"
		} else {
			if rand.Float64() > threshold3 {
				return "10"
			} else {
				return "100"
			}
		}
	}
}

func Answer(dice1, dice2, dice3 string) string {
	return getAnswer(dice1 + dice2 + dice3)
}
