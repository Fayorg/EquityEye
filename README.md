# EquityEye

**EquityEye** is an open-source component of the [Ledgerly](https://github.com/fayorg/ledgerly) project. It provides market data access and aggregation for Ledgerly and can also be self-hosted as a standalone service.

This component is designed for individuals and families who need market data for personal use. **Commercial or industrial use requires explicit permission from the author.**

---

## Features

- Real-time (or near real-time as needed) and historical market data retrieval
- API endpoints for seamless integration with applications
- Periodic data aggregation and storage for historical analysis
- Self-hosting capability for standalone usage
- Support for multiple financial data providers and multiples keys per provider for redundancy and throughput
- Open-source and easy to extend

---

## Table of Contents

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [License](#license)
- [Contributing](#contributing)
- [Contact](#contact)

---

## Getting Started

These instructions will guide you through setting up EquityEye on your local machine for development or self-hosting.

### Prerequisites

- **Programming Language**: Go (1.22.2 or newer)
- **Cache**: Redis (recommended) or any compatible cache server
- **Database**: Timescale (recommended) or any SQL-compatible database
- **Dependencies**: Docker (optional, for containerized deployment)

### Clone the Repository

```bash
git clone https://github.com/Fayorg/EquityEye
cd equityeye
```

## License

EquityEye is licensed under the CC BY-NC-SA 4.0 license.
You are free to use and modify this software for personal or family purposes. Commercial use requires prior written permission from the author.

For more details, see the [LICENSE](LICENSE.md) file.

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository.
2. Create a new feature branch (git checkout -b feature/your-feature).
3. Commit your changes (git commit -m 'Add your feature').
4. Push to your branch (git push origin feature/your-feature).
5. Open a pull request.

## Contact

If you have questions, feedback, or commercial licensing inquiries, please contact:

- Author: [Elie "Fayorg" Baier](mailto:elie.baier@epfl.ch)
- GitHub: [Fayorg](https://github.com/fayorg)