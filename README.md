# Encrypte-Quantum-Ledger-Database
Encrypte Quantum Ledger Database (ELDB) is a fully managed ledger database service provided by EncrypteID. It is designed to provide a transparent, immutable, and cryptographically verifiable transaction log owned by a central trusted authority. Inspired From Amazon Quantum Ledger Database

## How They Work 
This Architect is an ordered sequence of data changes. Each change in this architect is linked with the previous chain through cryptographic means. It attaches a so-called digest to each change, which is iteratively computed over all changes until this point.
![alt text](assets/image.png)