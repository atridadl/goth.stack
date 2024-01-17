package lib

func SendSSE(channel string, message string) error {
	// Create a channel to receive an error from the goroutine
	errCh := make(chan error, 1)

	// Use a goroutine to send the message asynchronously
	go func() {
		err := Publish(RedisClient, channel, message)
		errCh <- err // Send the error to the channel
	}()

	// Wait for the goroutine to finish and check for errors
	err := <-errCh
	if err != nil {
		return err
	}

	return nil
}
