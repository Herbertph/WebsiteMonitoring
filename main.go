package main

import (
	"fmt"
	"net/http"
	"time"
)

// Função que faz uma requisição HTTP a uma URL e verifica o status
func checkWebsite(url string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Failed to reach %s. Error: %s\n", url, err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == 200 {
        fmt.Printf("Site %s is up! Status Code: %d\n", url, resp.StatusCode)
    } else {
        fmt.Printf("Site %s might be down. Status Code: %d\n", url, resp.StatusCode)
    }
}

func main() {
    fmt.Println("Site Monitoring Service Started...")

    // Testando a verificação de um site
    url := "http://herbertph.com"
    checkWebsite(url)

    // Adicionando um delay de 5 segundos antes de terminar
    time.Sleep(5 * time.Second)
}
