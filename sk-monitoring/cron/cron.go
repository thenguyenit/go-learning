package cron

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/thenguyenit/sk-monitoring/message/slack"

	"github.com/thenguyenit/sk-monitoring/tail"
)

//Monitor will monitor the cron log
func Monitor() {
	//example: 2018-11-24 06:30:19 i-060717ce5ef67f4e6 cron-bptu4vkg5jfiw-bptu4vkg5jfiw: Cron Complete
	pattern := "(.*) i-060717ce5ef67f4e6 cron-bptu4vkg5jfiw-bptu4vkg5jfiw: (.*)"

	//minute
	delayTimeAllow, _ := strconv.ParseFloat(os.Getenv("CRON_DELAY_TIME_ALLOW"), 64)

	logFilePath := os.Getenv("LOG_PATH") + "/production/cron.log"
	flagFilePath := os.Getenv("LOG_PATH") + "/production/flag_send_message_to_slack_cron"
	endOfLine, err := tail.Get(logFilePath)
	if err == nil {
		fmt.Println(endOfLine)

		slack := slack.New()

		r, _ := regexp.Compile(pattern)
		output := r.FindStringSubmatch(endOfLine)
		if len(output) > 1 {
			latestRunDate, _ := time.Parse("2006-01-02 03:04:05", output[1])
			lastestMessage := output[1] + ": " + output[2]

			elapsed := time.Since(latestRunDate)
			minutesElapsed := strconv.FormatFloat(elapsed.Minutes(), 'f', 0, 64)
			fmt.Println("Elapsed: " + minutesElapsed + " minutes")

			if elapsed.Minutes() > delayTimeAllow {

				message := ":naichuoi: `Last heartbeat is older than " + minutesElapsed + " minutes`" + "\n" + "Last message: " + lastestMessage
				slack.ChatPostMessage(message)

				//Create flag send a notify message
				ioutil.WriteFile(flagFilePath, []byte(message), 0777)

			} else if _, err := os.Stat(flagFilePath); elapsed.Minutes() <= 1 && !os.IsNotExist(err) {
				message := ":tada:  Heart is beating :tada:" + "\n" + "Last message: " + lastestMessage
				slack.ChatPostMessage(message)
				//Remove flag send a notify message
				os.Remove(flagFilePath)

			}
		}
	} else {
		log.Fatal(err)
	}
}
