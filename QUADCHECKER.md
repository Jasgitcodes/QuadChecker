# 🧩 QuadChecker

A Go-based ASCII pattern recognition system that generates and identifies rectangle patterns produced by different Quad functions (QuadA–QuadE).

---

## 📌 Overview

QuadChecker is designed to:

- Generate ASCII rectangles using predefined rule sets (QuadA–QuadE)
- Analyze input ASCII shapes
- Detect which Quad function generated a given pattern
- Demonstrate rule-based grid traversal and string matching in Go

---

## 🏗️ Project Structure
QUADCHECKER/
│
├── main.go # Core detector (pattern recognition engine)
├── main-test.go # CLI tester for generating quad outputs
├── quadA.go # QuadA pattern generator
├── quadB.go # QuadB pattern generator
├── quadC.go # QuadC pattern generator
├── quadD.go # QuadD pattern generator
├── quadE.go # QuadE pattern generator
│
├── go.mod
├── go.sum
│
├── quadChecker/ # compiled binary (optional)
├── quadA/
├── quadB/
├── quadC/
├── quadD/
├── quadE/


---

## ⚙️ How It Works

### 1. Pattern Generation
Each Quad function generates a rectangle using `(x, y)` dimensions with unique rules for:
- Corners  
- Borders  
- Inner spaces  

---

### 2. Pattern Detection (`main.go`)
- Reads ASCII input from standard input
- Extracts width (x) and height (y)
- Compares input against all Quad generators
- Outputs matching Quad types

---

## 🧠 Core Concept

All quads follow this model:

```

(row, col) → character output

Each Quad differs only in:

Corner rules
Border rules
Symbol mapping

▶️ How to Run

Generate a quad:

go run main-test.go 5 3

Detect a quad:
echo "o--o
|  |
o--o" | go run main.go

🎯 Learning Outcomes
Grid-based algorithm design
Rule-based pattern generation
String manipulation in Go
Input/output stream handling
Modular code architecture

