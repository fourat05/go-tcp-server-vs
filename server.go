package main

import (
    "fmt"
    "net"
	"os"
	"bufio"
	"strings"
	"time"
)

const (
    HOST = "localhost"
    PORT = "8888"
	PROTOCOL = "tcp"
	QUIT_CMD = "QUIT"
	EHLO_CMD = "EHLO"
	DATE_CMD = "DATE"
)

// Flag pour marquer si le serveur avait reçu la commande EHLO
var EHLO_FIRED = false

// Handles incoming requests.
func handleRequest(c net.Conn) {
	fmt.Printf("Connexion établie %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		
		temp := strings.Split(strings.TrimSpace(string(netData)), " ")
		verb := temp[0]
		arg := ""
		if len(temp) > 1 {
			arg = temp[1]
		}
		result := ""
		switch verb {
			case QUIT_CMD:{
				result = "221 Bye"
				break
			}
			case EHLO_CMD:{
				EHLO_FIRED = true;
				result = "250 Pleased to meet you " + arg
			}
			case DATE_CMD:{
				if EHLO_FIRED {
					result = time.Now().Format("2006-01-02T15:04:05")
				} else {
					result = "550 Bad state"
				}
			}
			default:
				result = "Saisie irronée !"
		}
		result = result + "\n"
		c.Write([]byte(result))
	}
	c.Close()
}


func main() {
	server := HOST+":"+PORT
    // Démarrer l'écoute
    l, err := net.Listen(PROTOCOL, server)
    if err != nil {
        fmt.Println("Erreur lors du lancement du serveur:", err.Error())
        os.Exit(1)
    }

	fmt.Println("Serveur démarré", server)
	
	for {
        // Attente de messages
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Erreur lors de la lecture du message:", err.Error())
            os.Exit(1)
        }
        // Nouvelle goroutine pour accépter des connexions simultanées.
        go handleRequest(conn)
	}
	
}