package q77486

type Person struct {
	parent string
	amount int
	result int
}

func (p *Person) cal(amount int) {
	bbing := amount / 10
	amount = amount - bbing
	p.setResult(amount)
	if bbing == 0 {
		return
	}
	if p.parent != "" {
		person[p.parent].cal(bbing)
	}
}
func (p *Person) setResult(cash int) {
	p.result += cash
}

var person map[string]*Person

func solution(enroll []string, referral []string, seller []string, amount []int) []int {
	person = make(map[string]*Person)
	p := &Person{
		parent: "",
		amount: 0,
		result: 0,
	}
	person["center"] = p
	for i, s := range enroll {
		prnt := referral[i]
		if referral[i] == "-" {
			prnt = "center"
		}
		person[s] = &Person{
			parent: prnt,
			amount: 0,
			result: 0,
		}
	}
	for i, s := range seller {
		person[s].cal(amount[i] * 100)
	}
	var answer []int
	for _, s := range enroll {
		answer = append(answer, person[s].result)
	}
	return answer
}
