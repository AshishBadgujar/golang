package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func problemPuller(fileName string) ([]problem, error) {
	if fObj, err := os.Open(fileName); err == nil {
		csvR := csv.NewReader(fObj)
		if cLines, err := csvR.ReadAll(); err == nil {
			return problemParser(cLines), nil
		} else {
			return nil, fmt.Errorf("error in reading data in csv %s", err.Error())
		}
	} else {
		return nil, fmt.Errorf("error in opening  %s file; %s", fileName, err.Error())
	}
}

func problemParser(lines [][]string) []problem {
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{
			q: lines[i][0],
			a: lines[i][1],
		}
	}
	return r
}

func main() {
	fName := flag.String("f", "quiz.csv", "path of csv file")
	timer := flag.Int("t", 30, "timer of the quiz")
	flag.Parse()

	problems, err := problemPuller(*fName)
	if err != nil {
		fmt.Sprintf("Something went wrong: %s", err.Error())
		os.Exit(1)
	}
	correctAns := 0

	tObj := time.NewTimer(time.Duration(*timer) * time.Second) //timer obj
	ansC := make(chan string)

problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problem %d: %s=", i+1, p.q)
		go func() {
			fmt.Scanf("%s\n", &answer)
			ansC <- answer
		}()
		select {
		//timer expire
		case <-tObj.C:
			fmt.Println()
			break problemLoop
			//answers
		case iAns := <-ansC:
			if iAns == p.a {
				correctAns++
			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}
	}
	fmt.Printf("Your result is %d out of %d\n", correctAns, len(problems))
	fmt.Println("Press enter to exit")
	<-ansC //close program
}
