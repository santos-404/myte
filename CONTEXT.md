# Myte programming language — Project context

Welcome to the internal documentation and development context for the **Myte** programming language. This file contains essential information on what Myte is, what technologies and processes are involved, and how development is organized.

> 🚧 Myte is under heavy construction. Nothing is guaranteed to work, and many components are placeholders.

---

## Contents

- [1. General information](#1-general-information)  
- [2. Project goals](#2-project-goals)  
- [3. Technologies And tools](#3-technologies--tools)  
- [4. Building and running](#4-building-and-running)  
- [5. Contribution Guidelines](#5-contribution-guidelines)  
- [6. License and Copyright](#6-license-and-copyright)

---

## 1. General information

- **Project Name**: Myte
- **Status**: Pre-alpha / Experimental
- **Purpose**: To explore compiler construction, language design, and create a modern language from scratch.
- **Language type**: Interpreted progamming language. It follows an imperative and procedural paradigm.
- **Main Repository**: [github.com/santos-404/myte](https://github.com/santos-404/myte)

---

## 2. Project goals

The primary goals of Myte are:

- Learn and apply language design principles.
- Build a working interpreter from scratch.
- Explore typing systems, parsing techniques, and runtime design.
- Create something usable, expressive, and maybe fun!

> Note: This is both an educational and experimental project.

---

## 3. Technologies and tools

Currently used (or planned) technologies:

- **Language**: Go, using only the Go standard library — no third-party dependencies.
- **Lexer and parser**: Hand-written recursive descent parser and a custom lexer.
- **AST Construction**: Manually built abstract syntax tree (AST) structures.
- **Evaluation**: Tree-walk interpreter pattern.
- **Testing**: Basic test suite.
- **Docs**: Markdown-based internal documentation for now.

---

## 4. Building and running

> These instructions will evolve once the language has a buildable structure.

---

## 5. Contribution guidelines

Everyone is welcome to contribute ideas, issues, code, or just follow along.

If you’d like to help, feel free. You can refer to [CONTRIBUTING.md](./CONTRIBUTING.md) for details.

---

## 6. License and Copyright

This project is licensed under the MIT License.

> Copyright (c) 2025 Javier Santos

See the [LICENSE](./LICENSE) file for full terms.

This project does not include any GPL-licensed code and is safe for use in proprietary environments.
