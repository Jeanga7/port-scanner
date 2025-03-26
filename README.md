# 🛠️ **Active** - Port Scanner

## 📖 **Description du projet**

**Active** est un outil de scan de ports léger permettant de vérifier si des ports sont ouverts ou fermés sur un hôte donné. Ce programme utilise les protocoles **TCP** et **UDP** pour analyser les ports et retourne des informations détaillées sur leur statut (ouvert ou fermé). Le programme peut aussi identifier le service associé au port lorsqu'il est ouvert. 🧐

### **Pourquoi les scans de ports sont-ils importants ?**
Les scans de ports sont un élément crucial dans les tests de pénétration (pentesting). Ils permettent aux professionnels de la sécurité de découvrir les services et les vulnérabilités potentielles en explorant les ports d’un système. Un port ouvert peut permettre l’accès à un service, ce qui en fait une cible privilégiée pour les attaques.

## 🚀 **Fonctionnement**

Le programme scanne les ports spécifiés sur un hôte donné en utilisant les protocoles TCP ou UDP. Il affiche si le port est ouvert ou fermé et, si applicable, le service associé à ce port.

### **Exemples de ports courants**:
- Port **80**: HTTP
- Port **21**: FTP
- Port **22**: SSH
- Port **53**: DNS
- Port **443**: HTTPS
- ... et bien d'autres !

## ⚙️ **Installation et utilisation**

### 1. **Installation**

Clonez ce dépôt et accédez au répertoire du projet.

```bash
git clone https://github.com/votre-utilisateur/active.git
cd active
```

### 2. **Utilisation**

#### 📜 **Commandes disponibles**

Lancez le programme avec les options suivantes :

```bash
$> active --help
Usage: active [OPTIONS] [HOST] [PORT]
Options:
  -p               Range of ports to scan
  -u               UDP scan
  -t               TCP scan
  --help           Show this message and exit.
```

#### 🖥️ **Scanner un port avec TCP**

Pour scanner un port avec le protocole **TCP**, utilisez la commande suivante :

```bash
$> active -t 127.0.0.1 -p 80
Port 80 is open (HTTP)
```

#### 🌐 **Scanner un port avec UDP**

Pour scanner un port avec le protocole **UDP**, utilisez la commande suivante :

```bash
$> active -u 127.0.0.1 -p 8080
Port 8080 is open (HTTP)
```

#### 🔍 **Scanner une plage de ports**

Vous pouvez également scanner une plage de ports :

```bash
$> active -t 127.0.0.1 -p 80-82
Port 80 is open (HTTP)
Port 81 is open (HTTP)
Port 82 is closed
```

#### 🆘 **Aide**

Pour afficher l'aide et obtenir plus d'informations sur l'utilisation des commandes :

```bash
$> active --help
```

### 3. **Exemples d'exécution**

- **Scanner un port spécifique avec TCP** :

```bash
$> active -t 127.0.0.1 -p 80
Port 80 is open (HTTP)
```

- **Scanner un port spécifique avec UDP** :

```bash
$> active -u 127.0.0.1 -p 8080
Port 8080 is open (HTTP)
```

- **Scanner une plage de ports avec TCP** :

```bash
$> active -t 127.0.0.1 -p 80-83
Port 80 is open (HTTP)
Port 81 is open (HTTP)
Port 82 is closed
Port 83 is open (HTTP)
```

### 4. **Commandes utilitaires**

#### 🔍 **Voir les ports ouverts**
Pour afficher les ports ouverts sur votre machine, utilisez la commande suivante :

```bash
netstat -ano | findstr :80
```

#### 🌐 **Lancer un serveur web sur un port**
Pour tester un port en lançant un serveur web local, utilisez cette commande :

```bash
python -m http.server 80
```

#### Ouverture de port sur windows (firewall)
```bash 
netsh advfirewall firewall add rule name="Open TCP 8080" dir=in action=allow protocol=TCP localport=8080
netsh advfirewall firewall add rule name="Open UDP 9090" dir=in action=allow protocol=UDP localport=9090
```
#### Fermeture de port sur windows (firewall)
```bash 
netsh advfirewall firewall delete rule name="Open TCP 8080"
netsh advfirewall firewall delete rule name="Open UDP 9090"
```

---

## 📝 **Détails supplémentaires**

### **Services associés aux ports**
Voici quelques exemples de services associés à des ports communs :

- **Port 21**: FTP
- **Port 22**: SSH
- **Port 25**: SMTP
- **Port 80**: HTTP
- **Port 443**: HTTPS

Ces services sont souvent utilisés pour identifier les ports ouverts sur un hôte lors d’un scan.

---

## 🧪 **Tests**

### **Tests réalisés**

1. **Port 80**: Vérifiez si le port 80 est ouvert sur l'adresse locale :

```bash
$> active -t 127.0.0.1 -p 80
Port 80 is open (HTTP)
```

2. **UDP Port 8080**: Testez un port UDP en utilisant un serveur local avec le port 8080 :

```bash
$> active -u 127.0.0.1 -p 8080
Port 8080 is open (HTTP)
```

### **Bonus**: Affichage du service

Le programme affiche également le service associé à chaque port ouvert (HTTP, FTP, etc.).

---

## 🏁 **Conclusion**

Active est un outil simple mais puissant pour scanner les ports et identifier les services associés. Il est idéal pour les tests de pénétration, l'audit de sécurité, ou pour tout utilisateur cherchant à mieux comprendre les services qui tournent sur son réseau.

---

### 📚 **Références**
- [Wikipedia - Port Scanning](https://en.wikipedia.org/wiki/Port_scanning)
- [Nmap - Network Mapper](https://nmap.org/)

---

### 📄 **Exemples d'exécution**

```bash
$> active --help
Usage: active [OPTIONS] [HOST] [PORT]
Options:
  -p               Range of ports to scan
  -u               UDP scan
  -t               TCP scan
  --help           Show this message and exit.
```

