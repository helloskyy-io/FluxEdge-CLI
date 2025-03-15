# Edge CLI 🚀  
**A lightweight, high-performance CLI for interacting with the Edge Platform API.**  

![Go](https://img.shields.io/badge/Go-1.21-blue.svg)  
![License](https://img.shields.io/badge/license-MIT-green)  
![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)  

---

## 📌 Overview  
`edge-cli` is a **fast, lightweight, and modular command-line interface** built in **Go** to manage Edge platform deployments, automate resource provisioning, and interact with the API efficiently.  

### Key Features  
✅ Fetch available machines (`get-machines`)  
✅ Rent machines dynamically (`rent-machine`)  
✅ Deploy workloads from Git repositories (`deploy`)  
✅ Monitor active deployments (`monitor`)  
✅ Track account balance & alerts (`check-balance`)  
✅ Designed for **automation**, **performance**, and **scalability**  

---

## 📌 Installation  
### 1️⃣ Prerequisites  
- Install **Go 1.21+** ([Download](https://go.dev/dl/))  
- Git installed (`sudo apt install git` or `brew install git`)  

### 2️⃣ Clone the Repository  
```bash
git clone https://github.com/yourusername/edge-cli.git
cd edge-cli
```

### 3️⃣ Build the CLI
```
go build -o edgeapi
```

### 4️⃣ Run the CLI
```
./edgeapi
```
✅ You should see:
```
Welcome to Edge CLI! Use --help to see available commands.
```

## 📌 Usage

### Available Commands
```
./edgeapi get-machines          # List available machines
./edgeapi rent-machine --id 1234 # Rent a machine
./edgeapi deploy --repo github.com/myrepo # Deploy workload
./edgeapi check-balance         # Check balance & alerts
```

### Example API Call (Get Machines)
```
./edgeapi get-machines
```

✅ Response:
```
[
  {"id": "1234", "name": "Server 1"},
  {"id": "5678", "name": "Server 2"}
]
```
