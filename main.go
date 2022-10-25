package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/alex-ant/gomath/gaussian-elimination"
	"github.com/alex-ant/gomath/rational"
	"github.com/enescakir/emoji"
	"github.com/maja42/goval"
)

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}
func gcd2(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a = b
		b = a % b
	}
	return a
}
func gcd(arg ...string) (res int) {
	res = 1
	for _, i := range arg {
		j, _ := strconv.Atoi(i)
		res = gcd2(res, j)
	}
	return
}

func gauss(num int, reader *bufio.Reader) {
	equations := make([][]rational.Rational, 0)
	for i := 0; i < num; i++ {
		temp, _ := reader.ReadString('\n')
		temp = strings.TrimSuffix(temp, "\n")
		x := strings.Fields(temp)
		equation := make([]rational.Rational, 0)
		for j := 0; j <= num; j++ {
			val, _ := strconv.Atoi(x[j])
			equation = append(equation, rational.New(int64(val), 1))
		}
		equations = append(equations, equation)
	}
	result, gaussErr := gaussian.SolveGaussian(equations, false)
	if gaussErr != nil {
		fmt.Fprintln(os.Stderr, gaussErr)
	}
	for _, i := range result {
		fmt.Print(i[0].GetNumerator(), " ")
	}
	fmt.Println()
}
func quadratic(reader *bufio.Reader) {
	var x1, x2 float64
	equation, _ := reader.ReadString('\n')
	equation = strings.TrimSuffix(equation, "\n")
	equationArr := strings.Fields(equation)
	a, _ := strconv.Atoi(equationArr[0])
	b, _ := strconv.Atoi(equationArr[1])
	c, _ := strconv.Atoi(equationArr[2])
	x1 = (float64(-b) + math.Sqrt(float64(b*b-4*c*a))) / float64(2*a)
	x2 = (float64(-b) - math.Sqrt(float64(b*b-4*c*a))) / float64(2*a)
	fmt.Println(x1, x2)
}

// x^2 + 4x - 4
// (x - 2)(x + 2)
func main() {
	reader := bufio.NewReader(os.Stdin)
	eval := goval.NewEvaluator()
	for {
		pandaEmoji := emoji.Panda
		emojisArr := []emoji.Emoji{
			emoji.FaceBlowingAKiss,
			emoji.FaceSavoringFood,
			emoji.FaceScreamingInFear,
			emoji.FaceWithHandOverMouth,
			emoji.FaceWithMedicalMask,
			emoji.FaceWithOpenMouth,
			emoji.FaceWithRaisedEyebrow,
			emoji.FaceWithRollingEyes,
			emoji.FaceWithSteamFromNose,
			emoji.FaceWithTearsOfJoy,
			emoji.FaceWithThermometer,
			emoji.TiredFace,
			emoji.FaceWithTongue,
			emoji.FaceWithoutMouth,
			emoji.AngryFace,
			emoji.ClownFace,
		}
		emojiPrompt := []emoji.Emoji{
			emojisArr[rand.Intn(len(emojisArr))],
			emojisArr[rand.Intn(len(emojisArr))],
			emojisArr[rand.Intn(len(emojisArr))],
			emojisArr[rand.Intn(len(emojisArr))],
			emojisArr[rand.Intn(len(emojisArr))],
		}
		for _, v := range emojiPrompt {
			fmt.Print(v + " ")
		}
		fmt.Printf("%v >> ", pandaEmoji)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		cmd = strings.TrimSuffix(cmd, "\n")
		cmdArr := strings.Fields(cmd)
		fmt.Print(string("\033[32m"))
		switch cmdArr[0] {
		case "calc":
			result, err := eval.Evaluate(cmdArr[1], nil, nil)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Println(result)
		case "cd":
			if len(cmdArr) == 1 {
				os.Chdir("/home")
			} else {
				os.Chdir(cmdArr[1])
			}
		case "mkdir":
			os.MkdirAll(cmdArr[1], 0750)
		case "chmod":
			fs, _ := strconv.Atoi((cmdArr[2]))
			os.Chmod(cmdArr[1], os.FileMode(fs))
		case "chown":
			os.Chown(cmdArr[1], os.Getegid(), os.Geteuid())
		case "touch":
			os.Create(cmdArr[1])
		case "prime":
			num, _ := strconv.Atoi(cmdArr[1])
			if isPrime(num) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		case "gcd":
			fmt.Println(gcd(cmdArr[1:]...))
		case "gauss":
			num, _ := strconv.Atoi(cmdArr[1])
			gauss(num, reader)

		case "quad":
			quadratic(reader)

		case "exit":
			os.Exit(0)
		default:
			execCmd := exec.Command(cmdArr[0], cmdArr[1:]...)
			execCmd.Stderr = os.Stderr
			execCmd.Stdout = os.Stdout
			execCmd.Run()
		}

	}

}
