![Logo](/profile/frame_002.jpg)

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

## 📌 Installation (for users)
### **1️⃣ Download the Binary**
To install `edge-cli` without compiling from source, download the latest release:

```bash
curl -L -o edgeapi https://github.com/yourusername/edge-cli/releases/latest/download/edgeapi
chmod +x edgeapi
```

### 2️⃣ Add edgeapi to Your PATH
✅ Temporary (for the current session)
```bash
export PATH=$PATH:$(pwd)
```
✅ Permanent (for future sessions)
```bash
echo 'export PATH=$PATH:$HOME/edge-cli' >> ~/.bashrc && source ~/.bashrc
```
3️⃣ Set Your API Key
Before running any commands, export your API key: (get this key from the FluxEdge UI > Account > APIkeys)
```bash
export API_KEY="your-api-key-here"
```
(For Windows: set API_KEY=your-api-key-here)

### 📌 Usage

### Available Commands
```bash
./edgeapi get-machines          # List available machines
./edgeapi rent-machine --id 1234 # Rent a machine
./edgeapi deploy --repo github.com/myrepo # Deploy workload
./edgeapi check-balance         # Check balance & alerts
```

### Example API Call (Get Machines)
```bash
./edgeapi get-machines
```
✅ Response:
```json
[
  {"id": "1234", "name": "Server 1"},
  {"id": "5678", "name": "Server 2"}
]
```


## 📌 Installation (For Contributors)

### 1️⃣ Fork & Clone the Repo
To contribute, start by forking the repo and cloning it:
```bash
git clone https://github.com/yourusername/edge-cli.git
cd edge-cli
```

### 2️⃣ Prerequisites  
- Install **Go 1.24+** ([Download](https://go.dev/dl/))  
- Git installed (`sudo apt install git` or `brew install git`)  
- Go dependencies (`go mod tidy`)

### 3️⃣ Build the CLI
```
go build -o edgeapi
```

### 4️⃣ Run the CLI
```
./edgeapi --help
```
✅ You should see:
```
Welcome to Edge CLI! Use --help to see available commands.


Usage
edgeapi [command]

Available Commands

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
### 5️⃣ Run Tests
To validate your changes, run the test suite:
```bash
go test ./tests/
```


### 📌 Contributing Guidelines
✅ Follow Go best practices
✅ Use feature branches (feature/add-new-command)
✅ Write tests for new features (tests/)
✅ Ensure API keys are NEVER hardcoded
✅ Submit a pull request with a clear description

### 📌 Security Warning
- NEVER hardcode your API key inside the code.
- Use a .env file or environment variables to store your API key.
- The .env file is ignored via .gitignore and should never be committed.


## 📌 License

This project is licensed under Business Source License (BSL 1.1)

```
Business Source License 1.1

Licensed Work: Edge CLI (edge-cli)

Licensor: InFLux LLC

Change Date: March 16, 2035 (on this date, this software will be relicensed under the [Apache 2.0/MIT License]).

Licensed Grant:
- You may use, modify, and distribute this software for any purpose EXCEPT providing a substantially similar or competitive service to [Edge CLI, FluxEdge, or other InFLux LLC products].
- After the Change Date, this software will be available under [Apache 2.0/MIT License].

Additional Use Grant:
- Non-commercial and personal use is fully permitted.
- Contributions are welcome and licensed under the same terms.
- Internal commercial use within your own company is permitted.
- You may create integrations, plugins, and extensions.

Restrictions:
- You **may not** offer the Licensed Work as a **cloud service, SaaS, or managed service** that competes with [Edge CLI, FluxEdge, or other InFLux LLC products].
- You **may not** resell, sublicense, or commercially distribute the Licensed Work **without explicit written permission from InFLux LLC**.
- You **may not** remove or alter this license notice.

Disclaimer:
This software is provided "as is," without warranties or guarantees. The Licensor (InFLux LLC) is not responsible for any damages resulting from the use of this software.
```