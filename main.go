package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	"github.com/enescakir/emoji"
	"github.com/maja42/goval"
)

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
		case "exit":
			os.Exit(0)
		case "cd":
			os.Chdir(cmdArr[1])
		default:
			execCmd := exec.Command(cmdArr[0], cmdArr[1:]...)
			execCmd.Stderr = os.Stderr
			execCmd.Stdout = os.Stdout
			execCmd.Run()
		}

	}

}
