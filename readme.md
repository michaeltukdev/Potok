> [!WARNING]
> THIS IS STILL IN ACTIVE DEVELOPMENT. A LOT OF THE FUNCTIONALITY DOES NOT WORK.

# Potok

**Potok** is a secure, cross-platform, self-hosted CLI tool for automatically syncing and backing up [Obsidian](https://obsidian.md/) vaults (or technically anything) with end-to-end encryption. Potok is free, open-source, and designed for privacy and flexibility.

## Features

- **End-to-End Encryption (E2EE):** Vaults are encrypted locally before uploading to server; only you hold the keys.
- **Self-Hosted:** As a fan of self-hosting, you run your own Potok server—no third-party cloud required.
- **Multiple Vaults:** You can easily sync and manage multiple vaults independently of one another.
- **Automatic Sync:** Detects changes within vaults and syncs them automatically
- **Cross-Platform:** Works on Windows and Linux (Untested on MacOS - would appreciate if someone could do this for me).
- **CLI-Based:** For simplicity, we currently only support CLI through commands (potentially a GUI in the future)
- **Secure Key Storage:** Your encryption keys and API keys never leave your operating system, they're stored within your OS keyring.
- **No Vendor Lock-in:** Free and open-source.

## Getting Started

## Usage

### **Configure the Client**

### **List Your Vaults**

### **Sync a Vault (WIP)**

## Commands

| Command            | Description                                 |
|--------------------|---------------------------------------------|
| `set-api-url`      | Set the Potok server URL                    |
| `set-api-key`      | Set your API key (stored securely)          |
| `list-vaults`      | List all your vaults on the server          |
| `sync`             | Sync your vaults (automatic, coming soon)   |
| `upload`           | Encrypt and upload a new vault (planned)    |
| `download`         | Download and decrypt a vault (planned)      |

## Configuration

Potok stores your configuration in `~/.potok/config.json`:

```json
{
  "api_url": "http://localhost:8080",
}
```

API keys and vault passwords are stored securely in your OS keyring.


## Security

- All data is encrypted locally before upload.
- The server never sees your passwords or unencrypted data.
- API keys and passwords are never displayed after being set.

## Roadmap

- [ ] Automatic file watching and sync
- [ ] File-level sync and conflict handling
- [ ] Web dashboard for server admin
- [ ] Cross-platform installer
- [ ] More documentation and examples

## Contributors

Thanks to everyone who has contributed to Potok!

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/michaeltukdev">
        <img src="https://avatars.githubusercontent.com/u/66504185" width="80px;" alt=""/>
        <br />
        <sub><b>Michael Tuk</b></sub>
      </a>
      <br />Creator & Maintainer
    </td>
    <!-- You could be here!-->
  </tr>
</table>

Contributions are welcome! Please open issues or pull requests.

## Acknowledgements

- [Obsidian](https://obsidian.md/)
- [spf13/cobra](https://github.com/spf13/cobra)
- [zalando/go-keyring](https://github.com/zalando/go-keyring)
- [fsnotify](https://github.com/fsnotify/fsnotify)



## Running with Docker
You can run the Potok server and its database with a single command using Docker.

Make sure you have Docker and Docker Compose installed.

From the root of the project, run the following command:

<b>docker compose up --build</b>

The server will be available at http://localhost:8080.
