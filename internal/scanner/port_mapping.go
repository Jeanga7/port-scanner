package scanner

// PortMapping contient les correspondances entre les numéros de port et les services associés
var PortMapping = map[int]string{
    21:    "FTP",
    22:    "SSH",
    23:    "Telnet",
    25:    "SMTP",
    53:    "DNS",
    80:    "HTTP",
    110:   "POP3",
    115:   "SFTP",
    135:   "PRC",
    139:   "NetBIOS",
    143:   "IMAP",
    194:   "CRI",
    443:   "SSL",
    445:   "SMB",
    1433:  "MSSQL",
    3306:  "MySQL",
    3389:  "Remote Desktop",
    5632:  "PCAnywhere",
    5900:  "VNC",
    25565: "Minecraft",
}