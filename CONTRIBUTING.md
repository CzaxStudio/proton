# Contributing to Proton

Thank you for your interest in contributing to Proton! We welcome bug reports, feature requests, and pull requests to help make this Go GUI library better.

---

## Getting Started & Prerequisites

Because Proton interacts directly with system graphics, you need to ensure your local development environment has the necessary dependencies installed before compiling.

### Dependencies

* **Go**: Version 1.21 or higher is recommended.
* **Linux Users**: Ensure you have X11/Wayland development packages installed (e.g., `libwayland-dev`, `libx11-dev` depending on your specific window manager engine).
* **Windows/macOS**: Standard Go installation should work out of the box.

### Local Setup

1. **Fork** the repository on GitHub.
2. **Clone** your fork locally:
   ```bash
   git clone [https://github.com/YOUR_USERNAME/proton.git](https://github.com/YOUR_USERNAME/proton.git)
   cd proton
