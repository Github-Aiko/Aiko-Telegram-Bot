package bot

import "testing"

func TestSendMessage(t *testing.T) {

	bot, _ := New("5856609274:AAFn7XeOaaGuKc-VmYQo3ywPQ5X2eF9esNk", "http://amiss:921920@54.175.181.68:12345")

	err := bot.SendMessage("-1001658662143", "hello")
	if err != nil {
		t.Log(err)
	}

}
