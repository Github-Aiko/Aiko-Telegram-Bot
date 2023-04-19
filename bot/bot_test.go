package bot

import "testing"

func TestSendMessage(t *testing.T) {

	bot, _ := New("5856609274:AAFn7XeOaaGuKc-VmYQo3ywPQ5X2eF9esNk")

	err := bot.SendMessage("-1001658662143", "Aiko ....")
	if err != nil {
		t.Log(err)
	}

}
