import socket

# Créer un socket UDP
sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

# Associer le socket à l'adresse et au port
sock.bind(("0.0.0.0", 9090))

print("Le port 9090 est ouvert en UDP. En attente de données...")

while True:
    try:
        # Recevoir des données (1024 octets max)
        data, addr = sock.recvfrom(1024)
        print(f"Message reçu de {addr}: {data.decode()}")
    except KeyboardInterrupt:
        print("\nFermeture du serveur UDP.")
        sock.close()
        break
