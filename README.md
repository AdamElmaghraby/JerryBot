# JerryBot

JerryBot is a sophisticated Discord bot that utilizes OpenAI's GPT-3.5-turbo through the `go-openai` library, designed to enhance interactions within Discord servers. With its advanced conversational abilities and flexible message routing, JerryBot offers an engaging and customizable chat experience.

## Features

- **AI-Powered Chat**: Integrates with GPT-3.5-turbo to provide dynamic and contextually relevant responses.
- **Customizable Command Handling**: Uses a flexible routing system to handle various commands and interactions.
- **Moderation**: Includes profanity filtering and moderation to maintain a respectful chat environment.

## Configuration

JerryBot’s configuration is managed through a `config.json` file, located in the `config` directory. The configuration file includes:

- **Token**: Your Discord bot token.
- **BotPrefix**: The prefix used to invoke bot commands.
- **OpenApiKey**: Your OpenAI API key.

### Sample `config.json`

```json
{
    "Token": "YOUR_DISCORD_BOT_TOKEN",
    "BotPrefix": "!j ",
    "OpenApiKey": "YOUR_OPENAI_API_KEY"
}
```

## Components

### `config.go`

Handles reading and managing configuration from `config.json`.

### `mux.go`

Provides a message routing system for Discord commands. It supports fuzzy matching and context-aware command handling.

### `gpt.go`

Implements the integration with OpenAI's GPT-3.5-turbo and includes profanity filtering to ensure content moderation.

### `main.go`

Initializes and runs the bot, setting up command routes and handling bot lifecycle events.

## Getting Started

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/AdamElmaghraby/JerryBot.git
   ```

2. **Install Dependencies**:
   ```bash
   cd JerryBot
   go mod download
   ```

3. **Configure the Bot**:
   - Edit `config/config.json` with your Discord bot token, bot prefix, and OpenAI API key.

4. **Run the Bot**:
   ```bash
   go run main.go
   ```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to customize further based on specific features or additional documentation!
