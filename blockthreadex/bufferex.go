package blockthreadex

func addEmailsToQueue(emails []string) chan string {
	ln := len(emails)

	ch := make(chan string, ln)

	for i := 0; i < ln; i++ {
		ch <- emails[i]
	}

	return ch
}

func TestAddEmails() chan string {
	es := []string{
		"hey",
		"hellow",
		"good",
	}
	return addEmailsToQueue(es)
}
