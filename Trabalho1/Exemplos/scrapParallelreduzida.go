func scrapParallel(url string, rchan chan Result) {
	defer close(rchan)
	resp, err := http.Get(url)

	//Trecho removido para simplificação

	var r Result
	
  //Trecho removido para simplificação

	rchan <- r
}
